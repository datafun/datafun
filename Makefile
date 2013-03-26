
PREFIX=/usr/local
GO=go #/usr/local/go/bin/go
ARGS=build
CMD=mkdir -p ./bin/$(GOOS)_$(GOARCH) && $(GO) $(ARGS)



all: aggregate 


### aggregate

aggregate: mean sum

mean:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/aggregate/$@.go


sum:
	$(CMD) -o ./bin/$(GOOS)_$(GOARCH)/$@ src/aggregate/$@.go


### map

map: abs

abs:
	$(CC) $(SRC) src/map/$@.c $(CARGS) $(INCLUDE) -o ./bin/$@





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

