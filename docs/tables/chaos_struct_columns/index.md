---
title: "Tailpipe Table: chaos_struct_columns - Query Chaos struct columns"
description: "Chaos date time captures all the struct columns while collecting logs from different sources."
---

# Table: chaos_struct_columns - Query Chaos date time

The `chaos_struct_columns` table captures all the struct columns while collecting logs from different sources.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `chaos_struct_columns` ([examples](https://hub.tailpipe.io/plugins/turbot/chaos/tables/chaos_struct_columns#example-configurations)):

```sh
vi ~/.tailpipe/config/chaos.tpc
```

```hcl
partition "chaos_struct_columns" "struct_columns" {
  source "chaos_struct_columns" {
    row_count = 100
  }
}
```

## Collect

[Collect](https://tailpipe.io/docs/manage/collection) logs for all `chaos_struct_columns` partitions:

```sh
tailpipe collect chaos_struct_columns
```

Or for a single partition:

```sh
tailpipe collect chaos_struct_columns.struct_columns
```

## Query

### Basic info

```sql
select
  simple_struct,
  array_struct,
  nested_struct
from
  chaos_struct_columns;
```