#!/bin/bash

PASS=0
FAIL=0

source ./test/testlib.sh

printf "\nTesting datafun...\n"


#initailization
rm -rf /tmp/test-datafun
mkdir -p /tmp/test-datafun


##########
# TESTS
##########

test_module_start "aggregate"
test/aggregate/mean.test.sh
#test/aggregate/mode.test.sh
test/aggregate/product.test.sh
test/aggregate/sum.test.sh

#filter
#test/reduce/gt.test.sh
#test/reduce/gte.test.sh
#test/reduce/lt.test.sh
#test/reduce/lte.test.sh


test_module_start "map"
test/map/abs.test.sh
test/map/inv.test.sh
test/map/scale.test.sh

#search
#test/search/max.test.sh
#test/search/min.test.sh

test_module_start "util"
test/util/splitlines.test.sh


#any cleanup
test_all_done 
