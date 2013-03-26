#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "gte" "-n -5.678" "data.txt"

OUTPUT=$(diff $RET -<<EOF
45
23.0
10442
0.42424
2
45.0
-5.678
6
EOF)

assert_empty $OUTPUT



test_start "gte" "-n 2" "data.txt"

OUTPUT=$(diff $RET -<<EOF
45
23.0
10442
2
45.0
6
EOF)

assert_empty $OUTPUT



test_start "gte" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
45
23.0
10442
0.42424
2
45.0
6
EOF)

assert_empty $OUTPUT




