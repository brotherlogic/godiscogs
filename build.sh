protoc -I=./ --go_out=plugins=grpc:./ godiscogs.proto
mv github.com/brotherlogic/godiscogs/* ./
