DICTS := $(basename $(notdir $(wildcard dict/*.yml)))
GORELEASER_GITHUB_TOKEN ?=
GITHUB_TOKEN ?= $(GORELEASER_GITHUB_TOKEN)

.PHONY: install
install: packr
	GO111MODULE=on go install


.PHONY: dev
dev: clean
	GO111MODULE=on go install
	cd web; browser-sync

.PHONY: examples
examples:
	rm -rf examples
	mkdir -p examples
	for dict in $(DICTS); do \
	  for i in `seq 100`; \
	    do pipotron $$dict >> examples/$$dict.txt; \
	  done; \
	done

.PHONY: clean
clean:
	git clean -fxd

.PHONY: packr
packr:
	GO111MODULE=off go get github.com/gobuffalo/packr/v2/packr2
	packr2

.PHONY: docker
docker:
	docker build -t moul/pipotron .

.PHONY: functions
functions: packr
	mkdir -p functions-build
	GO111MODULE=on go install
	go build -o functions-build/pipotron ./functions/pipotron.go

.PHONY: goreleaser
goreleaser: packr
	GORELEASER_GITHUB_TOKEN=$(GORELEASER_GITHUB_TOKEN) GITHUB_TOKEN=$(GITHUB_TOKEN) goreleaser --rm-dist

.PHONY: goreleaser-dry-run
goreleaser-dry-run:
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY: netlify-dev
netlify-dev:
	netlify dev -c 'make dev'
