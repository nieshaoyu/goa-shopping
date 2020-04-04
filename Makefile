
# workdir info
PACKAGE=goa-shopping
PREFIX=$(shell pwd)
CMD_PACKAGE=${PACKAGE}/cmd
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/shop


# which golint
GOLINT=$(shell which golangci-lint || echo '')


.PONY: lint test
default: lint build


lint:
	@echo "+ $@"
	@$(if $(GOLINT), , \
		$(error Please install golint: `go get -u github.com/golangci/golangci-lint/cmd/golangci-lint`))
	golangci-lint run --deadline=10m --disable-all -E errcheck ./...

build:
	@echo "+ build"
	go build -o ${OUTPUT_FILE} $(CMD_PACKAGE)

clean:
	@echo "+ $@"
	@rm -r "${OUTPUT_DIR}"

dep:
	go get goa.design/goa/v3 && go get goa.design/goa/v3/... && go get github.com/fatih/gomodifytags


gen-apidoc:
	@redoc-cli bundle gen/http/openapi.json -o gen/apidoc.html

GOA_DESIGN_PATH=design
GOA_GEN_OUTPUT=.
goa:
	echo "+ $@"
ifeq ($(GOA_GEN_OUTPUT), .)
	goa gen $(PACKAGE)/$(GOA_DESIGN_PATH)
else
	goa gen $(PACKAGE)/$(GOA_DESIGN_PATH) -o $(GOA_GEN_OUTPUT)
endif
	@gomodifytags -file gen/shop/service.go -line=0,100000 -add-tags json -w -override
	@make gen-apidoc