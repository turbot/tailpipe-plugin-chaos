---
title: "Tailpipe Table: chaos_date_time - Query Chaos date time"
description: "Chaos date time captures all the time range possibilities while collecting logs from different sources."
---

# Table: chaos_date_time - Query Chaos date time

The `chaos_date_time` table captures all the time range possibilities while collecting logs from different sources.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `chaos_date_time` ([examples](https://hub.tailpipe.io/plugins/turbot/chaos/tables/chaos_date_time#example-configurations)):

```sh
vi ~/.tailpipe/config/chaos.tpc
```

```hcl
partition "chaos_date_time" "chaos_date_time_range" {
  source "chaos_date_time" {
    row_count = 100
  }
}
```

## Collect

[Collect](https://tailpipe.io/docs/manage/collection) logs for all `chaos_date_time` partitions:

```sh
tailpipe collect chaos_date_time
```

Or for a single partition:

```sh
tailpipe collect chaos_date_time.chaos_date_time_range
```

## Query

### Basic info

```sql
select
  timestamp
from
  chaos_date_time;
```