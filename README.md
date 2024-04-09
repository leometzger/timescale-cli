<p align="center">
  <img src="./assets/illustration.svg" width="35%" />
</p>

# TimescaleDB CLI

TimescaleDB CLI is a command-line tool designed to simplify some operations within TimescaleDB instances.
This unofficial tool is built to make the developers life easier. It does that by providing a terminal-based
interface for common operations and automation tasks.

## Usage

It is pretty simple to use it.

### Configuring

Add the configuration to the config file

```sh
tsctl config create --env staging --host db.timescale.staging --database tsdb
```

List configurated environments (TODO)

```sh
tsctl config ls
```

Check if the configuration is valid and connect the database

```sh
tsctl config check
```

### Aggregation

List aggregations from specified timescale instance

```sh
tsctl aggregation ls

tsctl aggregation ls --hypertable metrics

tsctl aggregation ls --view-name %hourly

tsctl aggregation ls --hypertable metrics --view-name %hourly
```

Refreshes hypertables from start to end using filters for view or hypertable.

```sh
tsctl aggregation refresh --env staging --start 2023-01-01 --end 2023-02-01

tsctl aggregation refresh --start 2023-01-01 --end 2023-02-01 --hypertable metrics

tsctl aggregation refresh --start 2023-01-01 --end 2023-02-01 --view-name %hourly
```

### Hypertable

List the hypertables, giving the main information about it.

```sh
tsctl hypertable ls

tsctl hypertable ls --name %hourly
```

Compress chunks manually from a start to end. (TODO)

```sh
tsctl hypertable compress --from 2023-01-01 --to 2024-01-01

tsctl hypertable compress --name %_metrics --from 2023-01-01 --to 2024-01-01
```

Give detailed information about the hypertable, like how many chunks it has. (TODO)

```sh
tsctl hypertable inspect metrics
```
