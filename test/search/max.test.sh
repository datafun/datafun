#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "max" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
10442.000000
EOF)

assert_empty $OUTPUT



test_start "max" "-x" "data.txt"

OUTPUT=$(diff $RET -<<EOF
3,10442.000000
EOF)

assert_empty $OUTPUT



test_start "max" "-s 100 -n 0.5" "data.txt"

OUTPUT=$(diff $RET -<<EOF
101.500000,10442.000000
EOF)

assert_empty $OUTPUT