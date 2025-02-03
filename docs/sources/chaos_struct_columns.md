---
title: "Source: chaos_struct_columns - Collect logs of all possible struct columns in DuckDB"
description: "Chaos struct columns captures all the struct columns while collecting logs from different sources."
---

The `chaos_struct_columns` table captures all the struct columns while collecting logs from different sources.

## Examples

### Basic info

```sql
select
  simple_struct,
  array_struct,
  nested_struct
from
  chaos_struct_columns;
```
