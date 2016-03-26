nohup go run main.go >debug.main.info 2>&1 &
nohup go run morey/resource_server.go >debug.resource.server.info 2>&1 &
