
clean:
	rm todo/todo.pb.go
	rm -f todo-client/node_modules

gen:
	@protoc -I todo/ todo/todo.proto --go_out=plugins=grpc:todo
	@mkdir -p todo-client/node_modules
	@protoc --proto_path=todo \
	--js_out=import_style=commonjs,binary:todo-client/node_modules/ \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:todo-client/node_modules/ \
	todo/todo.proto
