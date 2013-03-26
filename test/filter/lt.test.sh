#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "lt" "-n -6" "data.txt"

OUTPUT=$(diff $RET -<<EOF
-10000.34
EOF)

assert_empty $OUTPUT



test_start "lt" "-n 2" "data.txt"

OUTPUT=$(diff $RET -<<EOF
-10000.34
0.42424
-5.678
EOF)

assert_empty $OUTPUT



test_start "lt" "-n 23.43" "data.txt"

OUTPUT=$(diff $RET -<<EOF
23.0
-10000.34
0.42424
2
-5.678
6
EOF)

assert_empty $OUTPUT




