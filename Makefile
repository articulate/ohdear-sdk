GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

test: fmtcheck
	ginkgo ./ohdear/
