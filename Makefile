.PHONY: proto
proto:
	protoc -I ./protobuf-spec  ./protobuf-spec/*.proto --go_out=plugins=grpc:protobuf-spec
