#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "quotient" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
1.4404077884922451e-11
EOF)

assert_empty $OUTPUT



### TEST: csv vertical input

#setup data for next test
#DATA="1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16"
#TEST_FILE="$(pwd)/test/resources/test_mean.txt"
#printf $DATA > $TEST_FILE

test_start "quotient" "" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.0017094017094017094,0.0023809523809523807,0.0025974025974025974,0.0026041666666666665
EOF)

assert_empty $OUTPUT


### TEST: csv horizontal input

test_start "quotient" "-x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.041666666666666664
0.014880952380952382
0.006818181818181818
0.003869047619047619
EOF)

assert_empty $OUTPUT


### TEST: csv chunked vertical input

test_start "quotient" "-c 2" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.2,0.3333333333333333,0.42857142857142855,0.5
0.6923076923076923,0.7142857142857143,0.7333333333333333,0.75
EOF)

assert_empty $OUTPUT



### TEST: csv chunked vertical input with odd number of chunks

test_start "quotient" "-c 3" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.022222222222222223,0.03333333333333333,0.03896103896103896,0.041666666666666664
13,14,15,16
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input

test_start "quotient" "-c 2 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.5,0.75
0.8333333333333334,0.875
0.9,0.9166666666666666
0.9285714285714286,0.9375
EOF)

assert_empty $OUTPUT



### TEST: csv chunked horizontal input with odd number of chunks

test_start "quotient" "-c 3 -x" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
0.16666666666666666,4
0.11904761904761905,8
0.08181818181818182,12
0.06190476190476191,16
EOF)

assert_empty $OUTPUT


#rm $TEST_FILE




