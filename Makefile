GOROOT ?= $(shell go env GOROOT)
#GOROOT ?= /home/yamahata/graphene/devel/go/go
#GOROOT ?= /home/yamahata/graphene/devel/go/go-1.11.5/go
GO = $(GOROOT)/bin/go

all: $(target)

#go_executables = \
        channel \
        execing-process \
        getwd \
        goroutines-list \
        goroutines-stress \
        goroutine \
        goroutines \
        goroutines1 \
        helloworld \
        many-goroutines \
        recursion \
        signal \
        sleep \
        sleep-go \
        timers
go_executables = DNS-test

exec_target = \
        $(go_executables)

manifests = \
        $(addsuffix .manifest,$(executables)) \
        manifests

level = ../../
include ../../Makefile

%: %.go
        $(GO) build $^
