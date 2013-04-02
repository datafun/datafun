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


#setup data for next test
DATA="1,2,-3,4\n-5,6,7,8\n9,10,11,-12"
TEST_FILE="$(pwd)/test/resources/test-abs.txt"
printf $DATA > $TEST_FILE



test_start "abs" "" "test-abs.txt"

OUTPUT=$(diff $RET -<<EOF
1,2,3,4
5,6,7,8
9,10,11,12
EOF)

assert_empty $OUTPUT

#cleanup
rm -f $TEST_FILE