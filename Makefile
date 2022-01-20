GO=go
SRC_DIR=src
SRC=$(wildcard $(SRC_DIR)/*.go)
PROG=channel_four_news

.PHONY: run

run:
	$(GO) run $(SRC)

default: run

$(PROG): $(SRC)
	$(GO) build $(SRC)
