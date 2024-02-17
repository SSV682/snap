generate:
	@cd protos && protoc -I proto proto/analyzer/analyzer.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
	@cd protos && protoc -I proto proto/solver/solver.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative