#!/bin/bash

source ./test/testlib.sh

test_file_start 


test_start "inv" "" "clean-data.txt"

OUTPUT=$(diff $RET -<<EOF
1,0.5,0.3333333333333333,0.25
0.2,0.16666666666666666,0.14285714285714285,0.125
0.1111111111111111,0.1,0.09090909090909091,0.08333333333333333
0.07692307692307693,0.07142857142857142,0.06666666666666667,0.0625
EOF)

assert_empty $OUTPUT
