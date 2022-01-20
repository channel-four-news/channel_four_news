GO=go
SRC_DIR=src
SRC=$(wildcard $(SRC_DIR)/*.go)
PROG=channel_four_news

.PHONY: run clean

default: $(PROG)

run:
	$(GO) run $(SRC)

$(PROG): $(SRC)
	$(GO) build -o $(PROG) $(SRC)

clean:
	-rm $(PROG)

