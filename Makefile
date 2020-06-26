.PHONY: proto
proto:
	protoc -I ./protobuf-spec ./protobuf-spec/*.proto --go-grpc_out=protobuf-spec
