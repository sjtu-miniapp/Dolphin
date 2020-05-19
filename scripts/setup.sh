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
#./server_rest_main -cfg=cfg.yaml
      ;;
    redis)
      docker pull nihplod/redis
      docker run -dit --name redis -p 6379:6379 nihplod/redis redis-server
      ;;
    sql)
      docker pull nihplod/mysql
      docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${SQL_PASSWORD} -d nihplod/mysql
      ;;
    etcd)
      /etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380  --initial-cluster my-etcd-1=http://0.0.0.0:2380
      ;;
    *)
      echo "Invalid option!"
      ;;
esac

