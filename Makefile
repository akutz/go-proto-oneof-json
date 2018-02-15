example.pb.go: example.proto
	protoc -I "$(<D)" --go_out=plugins=grpc:"$(@D)" "$<"
