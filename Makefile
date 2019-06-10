DICTS := $(basename $(notdir $(wildcard dict/*.yml)))

.PHONY: install
install:
	GO111MODULE=on go install

.PHONY: examples
examples:
	rm -rf examples
	mkdir -p examples
	for dict in $(DICTS); do \
	  for i in `seq 100`; \
	    do pipotron $$dict >> examples/$$dict.txt; \
	  done; \
	done

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
