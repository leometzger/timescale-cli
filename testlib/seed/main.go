package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/leometzger/timescale-cli/internal/db"
)

func main() {
	info := db.NewConnectionInfo("localhost", 5432, "postgres", "postgres", "password")
	conn := db.Connect(info)
	defer conn.Close(context.Background())

	// all exits on error
	slog.Info("dropping all objects...")
	dropAllObjects(conn)

	slog.Info("creating metrics hypertable...")
	createMetricsHypertable(conn)

	slog.Info("inserting seed metrics...")
	insertMetrics(conn)

	slog.Info("creating continuous aggregations...")
	createContinuousAggregations(conn)
}

func createMetricsHypertable(conn db.PgxIface) {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE metrics(
				created timestamp with time zone default now() not null,
				type_id integer                                not null,
				value   double precision                       not null
		);
	`)
	if err != nil {
		slog.Error("error creating table: " + err.Error())
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), `SELECT create_hypertable('metrics', 'created')`)
	if err != nil {
		slog.Error("error creating hypertable: " + err.Error())
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), `SELECT set_chunk_time_interval('metrics', INTERVAL '1 second');`)
	if err != nil {
		slog.Error("error setting chunk time interval: " + err.Error())
		os.Exit(1)
	}
}

func dropAllObjects(conn db.PgxIface) {
	_, err := conn.Exec(context.Background(), `
		DROP TABLE IF EXISTS metrics CASCADE;
	`)
	if err != nil {
		slog.Error("error dropping objects: " + err.Error())
		os.Exit(1)
	}
}

func createContinuousAggregations(conn db.PgxIface) {
	_, err := conn.Exec(context.Background(), `
		CREATE MATERIALIZED VIEW metrics_by_day WITH (timescaledb.continuous) AS
		SELECT 
			time_bucket(interval '1 day', created) AS bucket,
			type_id,
			count(*),
			sum(value)
		FROM metrics
		GROUP BY bucket, type_id;
	`)

	if err != nil {
		slog.Error("error creating continuous aggregates for testing porpose", "cause", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), `
		CREATE MATERIALIZED VIEW metrics_by_hour WITH (timescaledb.continuous) AS
		SELECT 
			time_bucket(interval '1 hour', created) AS bucket,
			type_id,
			count(*),
			sum(value)
		FROM metrics
		GROUP BY bucket, type_id;
	`)
	if err != nil {
		slog.Error("error creating continuous aggregates for testing porpose", "cause", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), `
		CREATE MATERIALIZED VIEW metrics_by_month WITH (timescaledb.continuous) AS
		SELECT 
			time_bucket(interval '1 month', bucket) AS bucket,
			type_id,
			count(*) as count,
			sum(sum)
		FROM metrics_by_hour
		GROUP BY 1, type_id;
	`)
	if err != nil {
		slog.Error("error creating monthly hierarquical continuous aggregates for testing porpose", "cause", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), `
		CREATE MATERIALIZED VIEW metrics_by_year WITH (timescaledb.continuous) as
		SELECT time_bucket('1 year'::interval, metrics_by_month.bucket) AS bucket,
			metrics_by_month.type_id,
			count(*) AS count,
			sum(metrics_by_month.sum) AS sum
		FROM metrics_by_month
		GROUP BY (time_bucket('1 year'::interval, metrics_by_month.bucket)), metrics_by_month.type_id;
	`)
	if err != nil {
		slog.Error("error creating yearly hierarquical continuous aggregates for testing porpose", "cause", err)
		os.Exit(1)
	}
}

// inserts some energy metrics into timescale
// it gets the data from the following link
// https://docs.timescale.com/tutorials/latest/energy-data/
func insertMetrics(conn db.PgxIface) {
	_, err := conn.Exec(context.Background(), `
		INSERT INTO metrics (created, type_id, value) 
		VALUES
			('2023-05-31 21:48:05.123016+00',22,0.573),
			('2023-05-31 21:48:05.123006+00',23,0.429),
			('2023-05-31 21:48:05.122433+00',12,1),
			('2023-05-31 21:48:05.12229+00',2,135),
			('2023-05-31 21:48:05.12229+00',1,297),
			('2023-05-31 21:48:05.122252+00',13,1.83),
			('2023-05-31 21:48:05.122244+00',11,1.84),
			('2023-05-31 21:48:05.122243+00',3,185),
			('2023-05-31 21:48:05.122171+00',21,0.684),
			('2023-05-31 21:48:04.126252+00',11,1.84),
			('2023-05-31 21:48:04.126241+00',21,0.685),
			('2023-05-31 21:48:04.126226+00',12,1),
			('2023-05-31 21:48:04.126225+00',2,138),
			('2023-05-31 21:48:04.126217+00',23,0.43),
			('2023-05-31 21:48:04.126209+00',13,1.83),
			('2023-05-31 21:48:04.12619+00',3,186),
			('2023-05-31 21:48:04.126158+00',22,0.582),
			('2023-05-31 21:48:04.117475+00',1,297),
			('2023-05-31 21:48:03.829933+00',4,621),
			('2023-05-31 21:48:03.12619+00',23,0.443),
			('2023-05-31 21:48:03.126186+00',13,1.86),
			('2023-05-31 21:48:03.126169+00',12,1),
			('2023-05-31 21:48:03.126094+00',11,1.84),
			('2023-05-31 21:48:03.126091+00',21,0.68),
			('2023-05-31 21:48:03.126075+00',22,0.572),
			('2023-05-31 21:48:03.126062+00',3,194),
			('2023-05-31 21:48:03.126051+00',2,135),
			('2023-05-31 21:48:03.117405+00',1,294),
			('2023-05-31 21:48:03.029592+00',4,620),
			('2023-05-31 21:48:02.226052+00',4,631),
			('2023-05-31 21:48:02.122441+00',21,0.68),
			('2023-05-31 21:48:02.122366+00',22,0.573),
			('2023-05-31 21:48:02.122366+00',13,1.83),
			('2023-05-31 21:48:02.122322+00',23,0.428),
			('2023-05-31 21:48:02.114168+00',12,1),
			('2023-05-31 21:48:02.114046+00',3,184),
			('2023-05-31 21:48:02.114046+00',11,1.83),
			('2023-05-31 21:48:02.113832+00',1,293),
			('2023-05-31 21:48:02.113751+00',2,135),
			('2023-05-31 21:48:01.117963+00',23,0.417),
			('2023-05-31 21:48:01.1177+00',21,0.685),
			('2023-05-31 21:48:01.117667+00',13,1.8),
			('2023-05-31 21:48:01.117635+00',11,1.85),
			('2023-05-31 21:48:01.117631+00',3,177),
			('2023-05-31 21:48:01.117596+00',12,1),
			('2023-05-31 21:48:01.117589+00',22,0.573),
			('2023-05-31 21:48:01.117589+00',1,297),
			('2023-05-31 21:48:01.117589+00',2,135),
			('2023-05-31 21:48:01.014066+00',4,627),
			('2023-05-31 21:48:00.208349+00',4,631),
			('2023-05-31 21:48:00.130577+00',13,1.82),
			('2023-05-31 21:48:00.130515+00',22,0.572),
			('2023-05-31 21:48:00.130485+00',23,0.427),
			('2023-05-31 21:48:00.130485+00',11,1.88),
			('2023-05-31 21:48:00.130485+00',21,0.696),
			('2023-05-31 21:48:00.130485+00',3,183),
			('2023-05-31 21:48:00.130483+00',2,135),
			('2023-05-31 21:48:00.130457+00',12,1),
			('2023-05-31 21:48:00.121824+00',1,308),
			('2023-05-31 21:47:59.126588+00',13,1.86),
			('2023-05-31 21:47:59.126563+00',22,0.577),
			('2023-05-31 21:47:59.126528+00',2,138),
			('2023-05-31 21:47:59.126509+00',11,1.89),
			('2023-05-31 21:47:59.12649+00',3,195),
			('2023-05-31 21:47:59.126484+00',12,1.01),
			('2023-05-31 21:47:59.126479+00',23,0.435),
			('2023-05-31 21:47:59.126472+00',21,0.699),
			('2023-05-31 21:47:59.117475+00',1,311),
			('2023-05-31 21:47:58.869795+00',4,630),
			('2023-05-31 21:47:58.115915+00',22,0.576),
			('2023-05-31 21:47:58.113164+00',11,1.84),
			('2023-05-31 21:47:58.113127+00',12,1),
			('2023-05-31 21:47:58.113099+00',2,135),
			('2023-05-31 21:47:58.113086+00',23,0.417),
			('2023-05-31 21:47:58.11308+00',3,177),
			('2023-05-31 21:47:58.11308+00',21,0.684),
			('2023-05-31 21:47:58.11305+00',13,1.8),
			('2023-05-31 21:47:58.104407+00',1,294),
			('2023-05-31 21:47:58.057049+00',4,609),
			('2023-05-31 21:47:57.257023+00',4,617),
			('2023-05-31 21:47:57.113802+00',3,186),
			('2023-05-31 21:47:57.113451+00',23,0.431);
	`)
	if err != nil {
		slog.Error("Error seeding the test")
		os.Exit(1)
	}
}
