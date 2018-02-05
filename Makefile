# Stole this example from: https://github.com/vincentbernat/hellogopher/tree/feature/dep

BINARYNAME = tracker2jira
PACKAGE  = github.com/king-jam/tracker2jira
DATE    ?= $(shell date +%FT%T%z)
COMMITHASH ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
GOPATH   = $(CURDIR)/.gopath~
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))

GO      = go
GODOC   = godoc
GOFMT   = gofmt
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

LDFLAGS = -ldflags '-X $(PACKAGE)/rest/handlers/version.BuildDate=$(DATE) \
		-X $(PACKAGE)/rest/handlers/version.CommitHash=$(COMMITHASH)' \

SWAGGER_SERVER_DIR =

SWAGGER_CLIENT_DIR =



.PHONY: all
all: fmt vendor | $(BASE) ; $(info $(M) building executable…) @ ## Build program binary
	$Q cd $(BASE) && $(GO) build \
		-tags release \
		$(LDFLAGS) \
		-o bin/$(BINARYNAME) main.go

$(BASE): ; $(info $(M) setting GOPATH…)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# Tools

GODEP = $(BIN)/dep
$(BIN)/dep: | $(BASE) ; $(info $(M) building go dep...)
	$Q go get github.com/golang/dep/cmd/dep

GOLINT = $(BIN)/golint
$(BIN)/golint: | $(BASE) ; $(info $(M) building golint...)
	$Q go get github.com/golang/lint/golint

GOCOVMERGE = $(BIN)/gocovmerge
$(BIN)/gocovmerge: | $(BASE) ; $(info $(M) building gocovmerge...)
	$Q go get github.com/wadey/gocovmerge

GOCOV = $(BIN)/gocov
$(BIN)/gocov: | $(BASE) ; $(info $(M) building gocov...)
	$Q go get github.com/axw/gocov/...

GOCOVXML = $(BIN)/gocov-xml
$(bIN)/gocov-xml: | $(BASE) ; $(info $(M) building gocov-xml...)
	$Q go get github.com/AlekSi/gocov-xml

GO2XUNIT = $(BIN)/go2xunit
$(BIN)/go2xunit: | $(BASE) ; $(info $(M) building go2xunit...)
	$Q go get github.com/tebeka/go2xunit

SWAGGER = $(BIN)/swagger
$(BIN)/swagger: | $(BASE) ; $(info $(M) building swagger...)
	$Q go get github.com/go-swagger/go-swagger/cmd/swagger

GO_BINDATA_ASSETFS = $(BIN)/go-bindata-assetfs
$(BIN)/go-bindata-assetfs: | $(BASE) ; $(info $(M) building go-bindata-assetfs...)
	$Q go get github.com/jteeuwen/go-bindata/...
	$Q go get github.com/elazarl/go-bindata-assetfs/...

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt lint vendor | $(BASE) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q cd $(BASE) && $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

test-xml: fmt lint vendor | $(BASE) $(GO2XUNIT) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests with xUnit output
	$Q cd $(BASE) && 2>&1 $(GO) test -timeout 20s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

COVERAGE_MODE = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOVMERGE) $(GOCOV) $(GOCOVXML)
test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: fmt lint vendor test-coverage-tools | $(BASE) ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)/coverage
	$Q cd $(BASE) && for pkg in $(TESTPKGS); do \
		$(GO) test \
			-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $$pkg | \
					grep '^$(PACKAGE)/' | grep -v '^$(PACKAGE)/vendor/' | \
					tr '\n' ',')$$pkg \
			-covermode=$(COVERAGE_MODE) \
			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
	 done
	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: lint
lint: vendor | $(BASE) $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

# Dependency management

vendor: Gopkg.toml Gopkg.lock | $(BASE) $(GODEP) ; $(info $(M) retrieving dependencies…)
	$Q cd $(BASE) && $(GODEP) ensure
	@ln -nsf . vendor/src
	@touch $@
.PHONY: vendor-update
vendor-update: | $(BASE) $(GODEP)
ifeq "$(origin PKG)" "command line"
	$(info $(M) updating $(PKG) dependency…)
	$Q cd $(BASE) && $(GODEP) ensure -update $(PKG)
else
	$(info $(M) updating all dependencies…)
	$Q cd $(BASE) && $(GODEP) ensure -update
endif
	@ln -nsf . vendor/src
	@touch vendor

# Code Generation

.PHONY: swagger-server
swagger-server: $(BASE) $(SWAGGER) ; $(info $(M) generating swagger server...) @ ## Generates server
	$Q cd $(BASE)/rest && $(SWAGGER) generate server --flag-strategy=pflag --exclude-main -A t2j -s server -f swagger.yaml

.PHONY: swagger-client
swagger-client: $(BASE) $(SWAGGER) ; $(info $(M) generating swagger client...) @ ## Generates client
	$Q cd $(BASE)/rest && $(SWAGGER) generate client -A t2j -f swagger.yaml 2> /dev/null

# UI controls
.PHONY: ui
ui: $(BASE) $(GO_BINDATA_ASSETFS) ; $(info $(M) generating ui assets...) @ ## Generates UI assets
	$Q cd $(BASE)/ui && $(GO_BINDATA_ASSETFS) -pkg server dist/ && mv bindata_assetfs.go $(BASE)/rest/server/

.PHONY: swagger-ui
swagger-ui: $(BASE) $(GO_BINDATA_ASSETFS) ; $(info $(M) generating swagger ui assets...) @ ## Generates Swagger UI assets
	$Q cd $(BASE) && git clone https://github.com/swagger-api/swagger-ui.git
	$Q cd $(BASE)/swagger-ui && $(GO_BINDATA_ASSETFS) -pkg server dist/ && mv bindata_assetfs.go $(BASE)/rest/server/

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(GOPATH)
	@rm -rf bin
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
