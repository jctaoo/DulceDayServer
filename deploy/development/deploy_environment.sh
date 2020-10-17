docker run --name dulceday_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
docker run --name dylceday_redis -p 6379:6379 -d redis
docker run -d --name rabbitmq -h rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
