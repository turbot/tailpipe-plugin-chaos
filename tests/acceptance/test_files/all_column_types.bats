load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "verify single row count" {
  cat << EOF > ~/.tailpipe/config/chaos_all_col_types.tpc
partition "chaos_all_columns" "chaos_all_column_types" {
  source "chaos_all_columns" {
    row_count = 1
  }
}
EOF

  # tailpipe collect
  tailpipe collect chaos_all_columns.chaos_all_column_types

  # run tailpipe query and verify the row counts
  run tailpipe query "select count(*) as count from chaos_all_columns;" --output csv
  echo $output

  assert_equal "$output" "count
1"

  # remove the config file
  rm -rf ~/.tailpipe/config/chaos_all_col_types.tpc
}

@test "verify high row count" {
  cat << EOF > ~/.tailpipe/config/chaos_high_row_count.tpc
partition "chaos_all_columns" "chaos_high_row_count" {
  source "chaos_all_columns" {
    row_count = 100000
  }
}
EOF

  # tailpipe collect
  tailpipe collect chaos_all_columns.chaos_high_row_count

  # run tailpipe query and verify the row counts
  run tailpipe query "select count(*) as count from chaos_all_columns;" --output csv
  echo $output

  assert_equal "$output" "count
100000"

  # remove the config file
  rm -rf ~/.tailpipe/config/chaos_high_row_count.tpc
}

@test "verify very high row count" {
  skip "enable while testing"
  cat << EOF > ~/.tailpipe/config/chaos_very_high_row_count.tpc
partition "chaos_all_columns" "chaos_very_high_row_count" {
  source "chaos_all_columns" {
    row_count = 10000000
  }
}
EOF

  # tailpipe collect
  tailpipe collect chaos_all_columns.chaos_very_high_row_count

  # run tailpipe query and verify the row counts
  run tailpipe query "select count(*) as count from chaos_all_columns;" --output csv
  echo $output

  assert_equal "$output" "count
10000000"

  # remove the config file
  rm -rf ~/.tailpipe/config/chaos_very_high_row_count.tpc
}



function teardown() {
  rm -rf ~/.tailpipe/data
}