---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/tailpipe.svg"
brand_color: "#a42a2d"
display_name: "Chaos (Tailpipe)"
name: "chaos"
description: "Tailpipe plugin for testing Tailpipe in weird and wonderful ways."
---

# Chaos

The Chaos plugin is used internally to test Tailpipe features and functionality in weird and wonderful ways.

## Installation

Download and install the latest Chaos plugin:

```bash
$ tailpipe plugin install chaos
Installing plugin chaos...
$
```

Create a partition and collect the logs:

```bash
$ cat << EOF > ~/.tailpipe/config/chaos_all_col_types.tpc
partition "chaos_all_columns" "all_column_types" {
  source "chaos_all_columns" {
    row_count = 100
  }
}
EOF
```

Run:

```bash
$ tailpipe collect chaos_all_columns.all_column_types

Collection complete.

Artifacts discovered: 0. Artifacts downloaded: 0. Artifacts extracted: 0. Rows enriched: 100. Rows converted: 100. Errors: 0.

200 files did not need compaction.


Timing (may overlap):
 - enrich:     102µs (active: 33.371µs)
 - convert:    5.029949084s (active: 48.756125ms)
 - total time: 5.030756s
```