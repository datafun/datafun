#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "sum" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
557.4062399999999
EOF)

assert_empty $OUTPUT



### TEST: csv vertical input

#setup data for next test
#DATA="1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16"
#TEST_FILE="$(pwd)/test/resources/test_mean.txt"
#printf $DATA > $TEST_FILE

test_start "sum" "" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
28,32,36,40
EOF)

assert_empty $OUTPUT


### TEST: csv horizontal input

test_start "sum" "-x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
10
26
42
58
EOF)

assert_empty $OUTPUT


### TEST: csv chunked vertical input

test_start "sum" "-c 2" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
6,8,10,12
22,24,26,28
EOF)

assert_empty $OUTPUT



### TEST: csv chunked vertical input with odd number of chunks

test_start "sum" "-c 3" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
15,18,21,24
13,14,15,16
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input

test_start "sum" "-c 2 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
3,7
11,15
19,23
27,31
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input with odd number of chunks

test_start "sum" "-c 3 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
6,4
18,8
30,12
42,16
EOF)

assert_empty $OUTPUT


#rm $TEST_FILE




