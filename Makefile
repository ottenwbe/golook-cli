default: install

# build and install a release version of golook
release:
	@go install -ldflags "-s -w"

# build and install a version of golook
install:
	@go install

# unit tests
test:
	@ginkgo *

fmt:
	@for d in $$(go list ./... | grep -v vendor); do \
		go fmt $${d}; \
	done

vet:
	@for d in $$(go list ./... | grep -v vendor); do \
		go vet $${d};  \
	done

lint:
	@for d in $$(go list ./... | grep -v vendor); do \
		golint $${d};  \
	done

.PHONY: \
	release \
	install \
	test \
	fmt \
	vet \
	lint