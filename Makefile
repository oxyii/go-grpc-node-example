MODULE_NAME=github.com/oxyii/go-grpc-node-example

SOURCES=proto

NODE_DEPS=node_modules
PROTOC_GEN_TS_PATH=$(NODE_DEPS)/.bin/protoc-gen-ts

OUT_PB_TS=./proto
OUT_PB_GO=.

all: go ts

deps:
	go mod download && go install tool # go.sum must exists

go: deps
	protoc \
		-I ./$(SOURCES) \
		--go_out=$(OUT_PB_GO) --go_opt=paths=import --go_opt=module=$(MODULE_NAME) --go-grpc_out=$(OUT_PB_GO) --go-grpc_opt=paths=import --go-grpc_opt=module=$(MODULE_NAME) \
		\
		$$(find ./$(SOURCES) -iname "*.proto")

node:
	npm install --include=dev

ts: node
	mkdir -p $(OUT_PB_TS) && \
	protoc \
		-I ./$(SOURCES) \
		--plugin=protoc-gen-ts=$(PROTOC_GEN_TS_PATH) --ts_opt=target=node --ts_opt=no_namespace --ts_opt=unary_rpc_promise=true --ts_out=$(OUT_PB_TS) \
		\
		$$(find ./$(SOURCES) -iname "*.proto")

.PHONY: go ts