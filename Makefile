proto_gen_go:
	protoc -I=protobufs/ --go_out=services/api-gateway/pb/ --go_opt=paths=source_relative --go-grpc_out=services/api-gateway/pb/ --go-grpc_opt=paths=source_relative protobufs/*.proto
