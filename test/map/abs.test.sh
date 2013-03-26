#!/bin/bash

source ./test/testlib.sh

test_file_start 


test_start "abs" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
45
23
10000.34
10442
0.42424
2
45
5.678
6
EOF)

assert_empty $OUTPUT
