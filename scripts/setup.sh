#!/bin/sh
# run in server
# python3.7 -m pip install mysql-connector
# encryption
# python3.7 -m pip install cffi
# python3.7 -m pip install cryptography
# python3.7 -m pip install python-dotenv

# export .env files
ENV_FILE=../.env
export $(egrep -v '^#' $ENV_FILE | xargs)
# TODO: DOCKER SWARM
case "$1" in
    local)
      # source setup.sh local
      ;;
    sql)
      docker pull nihplod/mysql
      docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${SQL_PASSWORD} -d nihplod/mysql
      ;;
    *)
      echo "Invalid option!"
      ;;
esac
