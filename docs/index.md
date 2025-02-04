---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/tailpipe.svg"
brand_color: "#F18701"
display_name: "Chaos (Tailpipe)"
name: "chaos"
description: "Tailpipe plugin for testing Tailpipe in weird and wonderful ways."
---

# Chaos + Tailpipe

The Chaos plugin is used internally to test Tailpipe features and functionality in weird and wonderful ways.

[Tailpipe](https://tailpipe.io) is an open-source CLI tool that allows you to collect logs and query them with SQL.

The [Chaos Plugin for Tailpipe](https://hub.tailpipe.io/plugins/turbot/chaos) allows you to collect and query all column types table, all numeric column type table and more to test your plugins.

- **[Get started →](https://hub.tailpipe.io/plugins/turbot/chaos)**
- Documentation: [Table definitions & examples](https://hub.tailpipe.io/plugins/turbot/chaos/tables)
- Community: [Join #tailpipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/tailpipe-plugin-chaos/issues)

<img src="https://raw.githubusercontent.com/turbot/tailpipe-plugin-chaos/main/docs/images/chaos_all_columns_terminal.png" width="50%" type="thumbnail"/>

## Getting Started

Install Tailpipe from the [downloads](https://tailpipe.io/downloads) page:

```sh
# MacOS
brew install turbot/tap/tailpipe
```

```sh
# Linux or Windows (WSL)
sudo /bin/sh -c "$(curl -fsSL https://tailpipe.io/install/tailpipe.sh)"
```

Install the plugin:

```sh
tailpipe plugin install chaos
```

Configure table partition, and data source ([examples](https://hub.tailpipe.io/plugins/turbot/chaos/tables/chaos_all_columns#example-configurations)):

```sh
vi ~/.tailpipe/config/chaos.tpc
```

```hcl
partition "chaos_all_columns" "chaos_all_column_types" {
  source "chaos_all_columns" {
    row_count = 1
  }
}

partition "chaos_date_time" "date_time_inc" {
  source "chaos_date_time" {
    row_count = 100
  }
}

partition "chaos_struct_column" "struct_columns" {
  source "chaos_struct_columns" {
    row_count = 100
  }
}
```

Download, enrich, and save logs from your source ([examples](https://tailpipe.io/docs/reference/cli/collect)):

```sh
tailpipe collect chaos_all_columns
tailpipe collect chaos_date_time
tailpipe collect chaos_date_time
```

Enter interactive query mode:

```sh
tailpipe query
```

Run a query:

```sql
select
  smallint_column,
  float_column,
  boolean_column
from
  chaos_all_columns
limit 5;
```

```sh
+----------------------+-----------------------+------------------+
| smallint_column        | float_column          | boolean_column |
+----------------------+-----------------------+------------------+
| 5                      | 5.0                   | true           |
| 6                      | 6.0                   | false          |
| 7                      | 7.0                   | false          |
| 8                      | 8.0                   | true           |
| 9                      | 9.0                   | false          |
+----------------------+-----------------------+------------------+
```
