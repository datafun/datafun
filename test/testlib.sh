COLOR_CYAN="\x1B[36m"
COLOR_RED="\x1B[31m"
COLOR_GREEN="\x1B[32m"
COLOR_NORMAL="\x1B[39m"

function assert_empty() {
	if [ -z "$1" ]
	then
		printf "${COLOR_GREEN}passed${COLOR_NORMAL}\n"
		#PASS=$((PASS+1))
	else
		printf "${COLOR_RED}failed${COLOR_NORMAL}\n"
		printf "$1\n"
		#FAIL=`expr $FAIL + 1`
	fi
}

function test_file_start() {
	printf "  ${COLOR_CYAN}$(basename $0)${COLOR_NORMAL}\n"
}


function test_start() {
	local BIN=$1
	local ARGS=$2
	local INPUT=$3
	RET=0

	local OUT_FILE=/tmp/test-datafun/$BIN.test.out
	
	printf "    $BIN $ARGS ... "
	
	$(pwd)/bin/${GOOS}_$GOARCH/$BIN -i $(pwd)/test/resources/$INPUT -o $OUT_FILE $ARGS
	status=$?
	if [ $status -ne 0 ]; then
		echo "FAILED" > $OUT_FILE
	fi

	RET=$OUT_FILE
}

function test_bin() {
	echo "todo"
}



function test_all_done() {
	printf "\nFinished.\n"  #$PASS passed, $FAIL failed\n"
}

