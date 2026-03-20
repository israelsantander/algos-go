.PHONY: help test test-pkg test-sorting test-searching test-lists test-graphs test-linear test-recursion bench bench-pkg bench-sorting bench-searching bench-lists bench-graphs bench-linear bench-recursion

GO ?= go
PKG ?= ./...
BENCH ?= .
COUNT ?= 1
TESTFLAGS ?= -v
KNOWN_TARGETS := help test test-pkg test-sorting test-searching test-lists test-graphs test-linear test-recursion bench bench-pkg bench-sorting bench-searching bench-lists bench-graphs bench-linear bench-recursion
EXTRA_ARGS := $(filter-out $(KNOWN_TARGETS),$(MAKECMDGOALS))

help:
	@printf '%s\n' \
		'Available targets:' \
		'  make test                      Run all tests (verbose by default).' \
		'  make test-pkg PKG=./sorting    Run tests for one package.' \
		'  make test-sorting              Run sorting tests.' \
		'  make test-searching            Run searching tests.' \
		'  make test-lists                Run lists tests.' \
		'  make test-graphs               Run graphs tests.' \
		'  make test-linear               Run linear tests.' \
		'  make test-recursion            Run recursion tests.' \
		'  make bench                     Run all benchmarks with benchmem.' \
		'  make bench-pkg PKG=./sorting   Run benchmarks for one package.' \
		'  make bench-sorting             Run sorting benchmarks.' \
		'  make bench-searching           Run searching benchmarks.' \
		'  make bench-lists               Run lists benchmarks.' \
		'  make bench-graphs              Run graphs benchmarks.' \
		'  make bench-linear              Run linear benchmarks.' \
		'  make bench-recursion           Run recursion benchmarks.' \
		'' \
		'Variables:' \
		'  TESTFLAGS=-v (default for test targets)' \
		'  PKG=./sorting BENCH=Quick COUNT=5' \
		'' \
		'Passthrough args:' \
		'  make test-sorting -- -run TestReverse' \
		'  make test -- -count=1 -failfast'

test:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./...

test-pkg:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) $(PKG)

test-sorting:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./sorting

test-searching:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./searching

test-lists:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./lists

test-graphs:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./graphs

test-linear:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./linear

test-recursion:
	$(GO) test $(TESTFLAGS) $(EXTRA_ARGS) ./recursion

bench:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./...

bench-pkg:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) $(PKG)

bench-sorting:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./sorting

bench-searching:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./searching

bench-lists:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./lists

bench-graphs:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./graphs

bench-linear:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./linear

bench-recursion:
	$(GO) test -run='^$$' -bench='$(BENCH)' -benchmem -count=$(COUNT) ./recursion

%:
	@:
