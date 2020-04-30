#!/bin/sh
# run in server
# export .env files
ENV_FILE=../.env
# have to be sourced to run correctly
export $(egrep -v '^#' $ENV_FILE | xargs)
# TODO: DOCKER SWARM
case "$1" in
    local)
      # source setup.sh local
      ;;
    server)
      echo 1
#      ../backend/build/......?
#     docker
      ;;
    sql)
      docker pull nihplod/mysql
      docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${SQL_PASSWORD} -d nihplod/mysql
      ;;
    *)
      echo "Invalid option!"
      ;;
esac
