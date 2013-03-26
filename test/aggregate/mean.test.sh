#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "mean" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
61.934026666666654
EOF)

assert_empty $OUTPUT



