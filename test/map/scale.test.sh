#!/bin/bash

source ./test/testlib.sh

test_file_start 


test_start "scale" "-n 0.5" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.5,1,1.5,2
2.5,3,3.5,4
4.5,5,5.5,6
6.5,7,7.5,8
EOF)

assert_empty $OUTPUT

