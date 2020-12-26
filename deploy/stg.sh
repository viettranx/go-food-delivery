#!/usr/bin/env bash

APP_NAME=food-delivery

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}
# docker rmi $(docker images -qa -f 'dangling=true')

docker run -d --name ${APP_NAME} \
  --network demo \
  -e VIRTUAL_HOST="food.custohub.com" \
  -e LETSENCRYPT_HOST="food.custohub.com" \
  -e LETSENCRYPT_EMAIL="mail@food.custohub.com" \
  -e DBConnStr="food-delivery:9dfbce690866fc19e5a718a54a9fe055@tcp(mysql:3306)/food-delivery?charset=utf8mb4&parseTime=true" \
  -p 8080:8080 \
  ${APP_NAME}