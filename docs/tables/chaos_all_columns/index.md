---
title: "Tailpipe Table: chaos_all_columns - Query Chaos all columns"
description: "Chaos all columns captures all possible column types in DuckDB."
---

# Table: chaos_all_columns - Query Chaos all columns

The `chaos_all_columns` table captures all possible column types in DuckDB.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `chaos_all_columns` ([examples](https://hub.tailpipe.io/plugins/turbot/chaos/tables/chaos_all_columns#example-configurations)):

```sh
vi ~/.tailpipe/config/chaos.tpc
```

```hcl
partition "chaos_all_columns" "chaos_all_column_types" {
  source "chaos_all_columns" {
    row_count = 1
  }
}
```

## Collect

[Collect](https://tailpipe.io/docs/manage/collection) logs for all `chaos_all_columns` partitions:

```sh
tailpipe collect chaos_all_columns
```

Or for a single partition:

```sh
tailpipe collect chaos_all_columns.chaos_all_column_types
```

## Query

### Basic info

```sql
select
  boolean_column,
  date_time_column,
  ipaddress_column,
  json_column,
  long_string_column,
  epoch_column_seconds,
  string_to_array_column,
  inet_column,
  ltree_column
from
  chaos_all_columns;
```
