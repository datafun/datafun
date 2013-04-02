#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "product" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
1.4058518817922252e+14
EOF)

assert_empty $OUTPUT



### TEST: csv vertical input

#setup data for next test
#DATA="1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16"
#TEST_FILE="$(pwd)/test/resources/test_mean.txt"
#printf $DATA > $TEST_FILE

test_start "product" "" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
585,1680,3465,6144
EOF)

assert_empty $OUTPUT


### TEST: csv horizontal input

test_start "product" "-x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
24
1680
11880
43680
EOF)

assert_empty $OUTPUT


### TEST: csv chunked vertical input

test_start "product" "-c 2" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
5,12,21,32
117,140,165,192
EOF)

assert_empty $OUTPUT



### TEST: csv chunked vertical input with odd number of chunks

test_start "product" "-c 3" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
45,120,231,384
13,14,15,16
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input

test_start "product" "-c 2 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
2,12
30,56
90,132
182,240
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input with odd number of chunks

test_start "product" "-c 3 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
6,4
210,8
990,12
2730,16
EOF)

assert_empty $OUTPUT


#rm $TEST_FILE




