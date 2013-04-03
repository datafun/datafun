
PREFIX=/usr/local
GO=go #/usr/local/go/bin/go
ARGS=build
CMD=mkdir -p ./bin/$(GOOS)_$(GOARCH) && $(GO) $(ARGS)



all: aggregate filter map interpolate util


### aggregate
aggregate: mean product quotient sum
mean:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/aggregate/$@.go
product:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/aggregate/$@.go
quotient:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/aggregate/$@.go
sum:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/aggregate/$@.go



###filter
filter: nhead
nhead:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/filter/$@.go


### map
map: abs inv scale shift
abs:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/map/$@.go
inv:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/map/$@.go
scale:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/map/$@.go
shift:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/map/$@.go


### interpolate
interpolate: interplin
interplin:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/interpolate/$@.go


### util
util: splitlines
splitlines:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/util/$@.go



### UTILS

install:
	ls -tr bin/* | xargs -t -I {} install {} $(PREFIX)/bin/.

uninstall:
	ls bin | xargs -t -I {} rm $(PREFIX)/bin/{}

clean:
	rm -rf ./bin/*

test:
	./test/test.sh

.PHONY: clean install uninstall test

