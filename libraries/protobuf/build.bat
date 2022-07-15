set MODULE="wartech-studio.com/monster-reacher/libraries/protobuf"
protoc *.proto --go_out=. --go_opt=module=%MODULE% --go-grpc_out=. --go-grpc_opt=module=%MODULE%
protoc *.proto --gotag_out=:. --gotag_opt=module=%MODULE%
protoc gateway.proto --js_out=import_style=commonjs,binary:../../../monster-reacher-web/src/protoc --grpc-web_out=import_style=typescript,mode=grpcweb:../../../monster-reacher-web/src/protoc
protoc data_schema.proto --js_out=import_style=commonjs,binary:../../../monster-reacher-web/src/protoc --grpc-web_out=import_style=typescript,mode=grpcweb:../../../monster-reacher-web/src/protoc