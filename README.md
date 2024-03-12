# Intro / Idea

```sh
ts hypertable ls

ts aggregation ls
ts aggregation ls --hypertable metric
ts aggregation rm ca_metric_time_hourly
ts aggregation rm 123

ts aggregation inspect ca_metric_time
ts aggregation refresh
ts aggregation refresh --view-name '*hourly' --from 2023-01-01 --to 2023-06-01

ts aggregation compress --view-name '*daily*'
```
