#!/bin/bash

source ./test/testlib.sh

test_file_start 



test_start "sum" "" "data.txt"

OUTPUT=$(diff $RET -<<EOF
557.4062399999999
EOF)

assert_empty $OUTPUT



