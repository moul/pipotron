DICTS := $(wildcard dict/*.yml)

.PHONY: install
install:
	GO111MODULE=on go install

.PHONY: examples
examples: $(DICTS:.yml=.example.txt)

.PHONY: packr
packr:
	go get github.com/gobuffalo/packr/packr
	packr

.PHONY: docker
docker:
	docker build -t ultreme/pipotron .

.PHONY: functions
functions: packr
	mkdir -p functions-build
	GO111MODULE=on go install
	go build -o functions-build/pipotron ./functions/pipotron.go

.PHONY: clean
clean:
	rm -rf dict/*.example.txt

%.example.txt: %.yml
	rm -f $@
	for i in `seq 100`; do pipotron $< >> $@; done
