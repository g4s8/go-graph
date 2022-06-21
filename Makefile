ifeq ($(V),1)
	Q = 
	V_BUILD = 
else
	Q = @
	V_BUILD = -v
endif

PACKAGES := $(shell go list ./...)
BUILD_FLAGS := $(V_BUILD)

.PHONY: all
all: build test test-race bench
	@echo "all done"

.PHONY: build
build:
	${Q}go build $(BUILD_FLAGS) $(PACKAGES)

.PHONY: test
test:
	${Q}go test $(BUILD_FLAGS) $(PACKAGES)

.PHONY: test-race
test-race:
	${Q}go test $(BUILD_FLAGS) -race $(PACKAGES)

.PHONY: bench
bench:
	${Q}go test $(BUILD_FLAGS) -bench=. $(PACKAGES)

.PHONY: vet
vet:
	${Q}go vet $(PACKAGES)
