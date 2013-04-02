#!/bin/bash

source ./test/testlib.sh

test_file_start 


test_start "shift" "-n -2" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
-1,0,1,2
3,4,5,6
7,8,9,10
11,12,13,14
EOF)

assert_empty $OUTPUT

