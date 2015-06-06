#!/bin/sh

TEST=$1
TEST="$(tr '[:lower:]' '[:upper:]' <<< ${TEST:0:1})${TEST:1}"

echo "## $TEST Results"
echo
echo "|Benchmark Name|Iterations|Per-Iteration|"
echo "|--|--|--|"
grep Benchmark \
  | tr '\t' '|'\
  | sed 's/^/|/' | sed 's/$/|/'
