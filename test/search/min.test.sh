#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "min" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
-10000.340000
EOF)

assert_empty $OUTPUT



test_start "min" "-x" "data.txt"

OUTPUT=$(diff $RET -<<EOF
2,-10000.340000
EOF)

assert_empty $OUTPUT



test_start "min" "-s 200 -n 0.5" "data.txt"

OUTPUT=$(diff $RET -<<EOF
201.000000,-10000.340000
EOF)

assert_empty $OUTPUT