.PHONY: proto
proto: ## Регенерация протоколов pb
	if not exist api\pb mkdir api\pb
	protoc --go_out=api/pb --go-grpc_out=api/pb --proto_path=api/proto api/proto/products.proto

update:
	@echo "Update"
	go get github.com/Nariett/arox-pkg@main
