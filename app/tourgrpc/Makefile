PROTOC_VERSION = 3.14.0
PROTOC_LINUX_ZIP = protoc-$(PROTOC_VERSION)-linux-x86_64.zip

install-pb-tools:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_LINUX_ZIP)
	unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local bin/protoc
	unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_LINUX_ZIP)

genproto:
	protoc --proto_path=internal/proto \
				 --go_out=internal/pb \
				 --go-grpc_out=internal/pb \
				 internal/proto/*.proto

cleanpb:
	rm internal/pb/*.go

runserver:
	go run cmd/server/main.go

runclient: 
	go run cmd/client/main.go

.PHONY:
	install-pb-tools genproto cleanpb runserver runclient