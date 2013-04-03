#!/bin/bash

source ./test/testlib.sh

test_file_start 


#setup data for next test
DATA="0\n10\n20"
TEST_FILE="$(pwd)/test/resources/test-interplin.txt"
printf $DATA > $TEST_FILE

test_start "interplin" "-n 2" "test-interplin.txt"

OUTPUT=$(diff $RET -<<EOF
0
3.3333333333333335
6.666666666666667
10
13.333333333333334
16.666666666666668
20
EOF)

assert_empty $OUTPUT

rm $TEST_FILE




