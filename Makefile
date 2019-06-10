DICTS := $(wildcard dict/*.yml)

.PHONY: install
install:
	go install

.PHONY: examples
examples: $(DICTS:.yml=.example.txt)

.PHONY: clean
clean:
	rm -rf dict/*.example.txt

%.example.txt: %.yml
	rm -f $@
	for i in `seq 10`; do pipotron $< >> $@; done
