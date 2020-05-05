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
      docker pull nihplod/mysql
      docker run -dit --name redis -p 6379:6379 nihplod/redis redis-server
      ;;
    sql)
      docker pull nihplod/mysql
      docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${SQL_PASSWORD} -d nihplod/mysql
      ;;
    etcd)
      /etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380  --initial-cluster my-etcd-1=http://0.0.0.0:2380
      ./service --registry=etcd --registry_address=xx.xx.xx.xx:2379
      ;;
    *)
      echo "Invalid option!"
      ;;
esac

etcd --name cd0 \
--initial-advertise-peer-urls http://192.168.3.45:2380 \
--listen-peer-urls http://192.168.3.45:2380 \
--listen-client-urls http://192.168.3.45:2379 \
--advertise-client-urls http://192.168.3.45:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster cd0=http://192.168.3.45:2380,cd1=http://192.168.3.45:2480,cd2=http://192.168.3.118:2380 \
--initial-cluster-state new
etcd --name cd1 \
--initial-advertise-peer-urls http://192.168.3.45:2480 \
--listen-peer-urls http://192.168.3.45:2480 \
--listen-client-urls http://192.168.3.45:2479 \
--advertise-client-urls http://192.168.3.45:2479 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster cd0=http://192.168.3.45:2380,cd1=http://192.168.3.45:2480,cd2=http://192.168.3.118:2380 \
--initial-cluster-state new
etcd --name cd2 \
--initial-advertise-peer-urls http://192.168.3.118:2380 \
--listen-peer-urls http://192.168.3.118:2380 \
--listen-client-urls http://192.168.3.118:2379 \
--advertise-client-urls http://192.168.3.118:2379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster cd0=http://192.168.3.45:2380,cd1=http://192.168.3.45:2480,cd2=http://192.168.3.118:2380 \
--initial-cluster-state new

etcdctl --endpoints="http://192.168.3.118:2379" endpoint health
go run main.go --selector=cache --registry=etcdv3 --registry_address=http://192.168.3.45:2479
go run main.go plugin.go --registry=etcd --registry_address=etcd1.foo.com:2379,etcd2.foo.com:2379,etcd3.foo.com:2379

etcd --name infra0 --initial-advertise-peer-urls http://10.0.1.10:2380 \
  --listen-peer-urls http://10.0.1.10:2380 \
  --listen-client-urls http://10.0.1.10:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://10.0.1.10:2379 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster infra0=http://10.0.1.10:2380,infra1=http://10.0.1.11:2380,infra2=http://10.0.1.12:2380 \
  --initial-cluster-state new

  $ etcd --name infra1 --initial-advertise-peer-urls http://10.0.1.11:2380 \
  --listen-peer-urls http://10.0.1.11:2380 \
  --listen-client-urls http://10.0.1.11:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://10.0.1.11:2379 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster infra0=http://10.0.1.10:2380,infra1=http://10.0.1.11:2380,infra2=http://10.0.1.12:2380 \
  --initial-cluster-state new

  etcd --name infra2 --initial-advertise-peer-urls http://10.0.1.12:2380 \
  --listen-peer-urls http://10.0.1.12:2380 \
  --listen-client-urls http://10.0.1.12:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://10.0.1.12:2379 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster infra0=http://10.0.1.10:2380,infra1=http://10.0.1.11:2380,infra2=http://10.0.1.12:2380 \
  --initial-cluster-state new


  etcd --name infra0 --initial-advertise-peer-urls http://localhost:2380 \
  --listen-peer-urls http://localhost:2380 \
  --listen-client-urls http://localhost:2379 \
  --advertise-client-urls http://localhost:2379 \
  --initial-cluster-token etcd-cluster-1 \
  --initial-cluster infra0=http://localhost:2380 \
  --initial-cluster-state new