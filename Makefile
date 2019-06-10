DICTS := $(wildcard dict/*.yml)

.PHONY: install
install:
	GO111MODULE=on go install

.PHONY: examples
examples: $(DICTS:.yml=.example.txt)

.PHONY: clean
clean:
	rm -rf dict/*.example.txt

%.example.txt: %.yml
	rm -f $@
	for i in `seq 100`; do pipotron $< >> $@; done
