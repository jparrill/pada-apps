# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BASE_FOLDER ?= $(shell bash -c 'pwd')
CMD_FOLDER = "${BASE_FOLDER}/cmd"
BIN_FOLDER = "${BASE_FOLDER}/bin"

all: build
build: 
	mkdir -p ${BIN_FOLDER}
	@for file in $(shell ls ${CMD_FOLDER}) ; do \
		$(GOBUILD) -o ${BIN_FOLDER}/$${file} -v -i "${CMD_FOLDER}/$${file}/main.go" ; \
	done
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BIN_FOLDER)
