DICTS := $(basename $(notdir $(wildcard dict/*.yml)))
GORELEASER_GITHUB_TOKEN ?=
GITHUB_TOKEN ?= $(GORELEASER_GITHUB_TOKEN)
DEPS := ./web/dicts.js ./dict/dict-packr.go

.PHONY: install
install: $(DEPS)
	GO111MODULE=on go install


.PHONY: dev
dev: clean
	GO111MODULE=on go install
	cd web; browser-sync

./web/dicts.js: $(wildcard ./dict/*.yml)
	cd dict; fs-bundler --format=js --callback=dicts *.yml > ../$@.tmp
	mv $@.tmp $@

./dict/dict-packr.go: $(wildcard ./dict/*.go) $(wildcard ./dict/*.yml)
	GO111MODULE=off go get github.com/gobuffalo/packr/v2/packr2
	rm -f ./dict/*~
	rm -f ./dict/.#*
	rm -f ./dict/*#
	packr2

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
	rm -f $(DEPS)

.PHONY: docker
docker:
	docker build -t moul/pipotron .

.PHONY: functions
functions: $(DEPS)
	mkdir -p functions-build
	GO111MODULE=on go install
	go build -o functions-build/pipotron ./functions/pipotron.go

.PHONY: goreleaser
goreleaser: $(DEPS)
	GORELEASER_GITHUB_TOKEN=$(GORELEASER_GITHUB_TOKEN) GITHUB_TOKEN=$(GITHUB_TOKEN) goreleaser --rm-dist

.PHONY: goreleaser-dry-run
goreleaser-dry-run:
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY: _netlify_prepare
_netlify_prepare:
	go get moul.io/fs-bundler

.PHONY: netlify-dev
netlify-dev:
	netlify dev -c 'make dev'

.PHONY: sam-local
sam-local: $(DEPS)
	@# pip install --user sam-aws-cli
	@echo ""
	@echo "Open: http://localhost:3000/index.html"
	@echo ""
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o functions-build/pipotron ./functions/pipotron.go
	sam local start-api --host=0.0.0.0 --static-dir=web
