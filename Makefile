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
		echo "building: cmd/$${file}/main.go" ; \
		$(GOBUILD) -o ${BIN_FOLDER}/$${file} -i "${CMD_FOLDER}/$${file}/main.go" ; \
	done

test: 
	$(GOTEST) -v  ./...

test_ci: 
	CI=true $(GOTEST) -v  ./...

cover:
	$(GOTEST) --cover ./...

clean: 
	$(GOCLEAN)
	rm -rf $(BIN_FOLDER)
