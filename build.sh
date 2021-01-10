#protoc -I=./ --go_out=plugins=grpc:./ godiscogs.proto
#protoc --proto_path ../../../ --go_out=plugins=grpc:./  -I=./ godiscogs.proto
/home/simon/.local/bin/protoc -I=./ --go_out ./ --go_opt=paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative --go-grpc_opt=require_unimplemented_servers=false godiscogs.proto
#mv github.com/brotherlogic/godiscogs/* ./
