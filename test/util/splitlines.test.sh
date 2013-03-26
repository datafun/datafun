#!/bin/bash

source ./test/testlib.sh

test_file_start 


BIN=splitlines
ARGS="-n 3 /tmp/test-datafun/sl_chunks.txt"

OUT_FILE=/tmp/test-datafun/$BIN.test.out
IN_FILE=$(pwd)/test/resources/data.txt
  
printf "    $BIN $ARGS ... "
  
$(pwd)/bin/${GOOS}_${GOARCH}/$BIN -i $IN_FILE -o $OUT_FILE -n 3 /tmp/test-datafun/sl_chunks.txt #> $OUT_FILE
status=$?
if [ $status -ne 0 ]; then
  echo "FAILED" > $OUT_FILE
fi

OUTPUT=$(diff $OUT_FILE -<<EOF
/tmp/test-datafun/0-sl_chunks.txt
/tmp/test-datafun/1-sl_chunks.txt
/tmp/test-datafun/2-sl_chunks.txt
EOF)

assert_empty $OUTPUT

##make sure dataset is actually there
cat /tmp/test-datafun/{0,1,2}-sl_chunks.txt > /tmp/test-datafun/sl_combined.txt
OUTPUT=$(diff /tmp/test-datafun/sl_combined.txt $IN_FILE)


printf "    diff combined.txt data.txt ... "
assert_empty $OUTPUT





