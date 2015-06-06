GOTESTFILES := $(wildcard *_test.go)
RESULTS = $(GOTESTFILES:_test.go=_results.md)

MARKDOWN=scripts/markdown.sh

%_test: %.go %_test.go
	go test -bench . $^

%_results.md: %.go %_test.go $(MARKDOWN)
	go test -bench . $(filter %.go,$^) | $(MARKDOWN) $* > $@

results: $(RESULTS)

clean:
	rm -rf *_results.md
