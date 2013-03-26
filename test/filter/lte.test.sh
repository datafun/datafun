#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "lte" "-n -5.678" "data.txt"

OUTPUT=$(diff $RET -<<EOF
-10000.34
-5.678
EOF)

assert_empty $OUTPUT



test_start "lte" "-n 2" "data.txt"

OUTPUT=$(diff $RET -<<EOF
-10000.34
0.42424
2
-5.678
EOF)

assert_empty $OUTPUT



test_start "lte" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
-10000.34
-5.678
EOF)

assert_empty $OUTPUT




