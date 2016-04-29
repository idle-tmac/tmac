kill -9  $(ps aux|grep main.go |grep -v grep|awk '{print $2}')
kill -9  $(ps aux|grep morey/resource_server.go|grep -v grep|awk '{print $2}')
kill -9 $(lsof -i :10086|grep -v COM|awk '{print $2}')
kill -9 $(lsof -i :8080|grep -v COM|awk '{print $2}')

