#! /usr/bin/bash
container_name="psql_container"
systemctl --user start docker-desktop

ck_container=$(docker ps -a | grep $container_name)

if [[ -z "$ck_container" ]] 
then
  docker run --name $container_name -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres
else 
  docker stop $container_name
  docker rm $container_name
  docker run --name $container_name -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres
fi

