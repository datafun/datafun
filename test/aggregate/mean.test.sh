#!/bin/bash

source ./test/testlib.sh

test_file_start 


### TEST: default input

test_start "mean" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
61.934026666666654
EOF)

assert_empty $OUTPUT


### TEST: csv vertical input

#setup data for next test
#DATA="1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16"
#TEST_FILE="$(pwd)/test/resources/test_mean.txt"
#printf $DATA > $TEST_FILE

test_start "mean" "" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
7,8,9,10
EOF)

assert_empty $OUTPUT


### TEST: csv horizontal input

test_start "mean" "-x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
2.5
6.5
10.5
14.5
EOF)

assert_empty $OUTPUT


### TEST: csv chunked vertical input

test_start "mean" "-c 2" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
3,4,5,6
11,12,13,14
EOF)

assert_empty $OUTPUT



### TEST: csv chunked vertical input with odd number of chunks

test_start "mean" "-c 3" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
5,6,7,8
13,14,15,16
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input

test_start "mean" "-c 2 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
1.5,3.5
5.5,7.5
9.5,11.5
13.5,15.5
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input with odd number of chunks

test_start "mean" "-c 3 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
2,4
6,8
10,12
14,16
EOF)

assert_empty $OUTPUT


#rm $TEST_FILE




