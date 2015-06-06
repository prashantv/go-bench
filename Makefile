GOTESTFILES := $(wildcard *_test.go)
RESULTS = $(GOTESTFILES:_test.go=_results.md)

%_test: %.go %_test.go
	go test -bench . $^

%_results.md: %.go %_test.go
	go test -bench . $^ | ./scripts/markdown.sh $* > $@

results: $(RESULTS)

clean_results:
	rm -rf *_results.md
