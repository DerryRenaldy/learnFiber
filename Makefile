# Docker Create Container Command
create_container:
	docker run -itd --name mysql-fiber-project -e MYSQL_ROOT_PASSWORD=root -v /home/derryrenaldy/Desktop/belajar/golang-fiber-demo/docker-volume-mysql:/var/lib/mysql -p 13306:3306 -h localhost mysql

# Docker Container Start
start_container:
	docker container start mysql-fiber-project

# Docker Delete Container Command
delete_container:
	docker container rm -f mysql-fiber-project

# Docker SSH
docker_ssh:
	docker exec -it mysql-fiber-project bash

# MYSQL Connect
mysql_connect:
	mysql -u root -h 127.0.0.1 -P 13306 -p