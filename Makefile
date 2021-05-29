gen:
	protoc -I=proto/ -I=proto/import/github.com/MobileStore-Grpc/product/proto/ \
	--go_out=. --go_opt=module=github.com/MobileStore-Grpc/search-product \
	--go-grpc_out=. --go-grpc_opt=module=github.com/MobileStore-Grpc/search-product \
	--grpc-gateway_out=. --grpc-gateway_opt=module=github.com/MobileStore-Grpc/search-product \
	--openapiv2_out=swagger \
	proto/*.proto proto/google/api/*.proto

clean:
	rm pb/*.go