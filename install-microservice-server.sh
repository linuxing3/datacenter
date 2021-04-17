#!/bin/bash
## author by linuxing3
## Email linuxing3@qq.com
## wechat linuxing3


blue() {
    echo -e "\033[34m\033[01m$1\033[0m"
}
green() {
    echo -e "\033[32m\033[01m$1\033[0m"
}
red() {
    echo -e "\033[31m\033[01m$1\033[0m"
}
yellow() {
    echo -e "\033[33m\033[01m$1\033[0m"
}

##### 启动mysql 服务
start_mysql() {
  docker kill mysql
  docker rm mysql
  docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -v $(pwd)/mysql8:/var/lib/mysql -d mysql:8.0.21
}
##### 启动redis 服务
start_redis() {
  docker stop redis
  docker rm redis
  echo "Start Redis Service..."
  docker run --name redis -d \
    --publish 6379:6379 \
    --env 'REDIS_PASSWORD=admin' \
    --volume $(pwd)/redis:/var/lib/redis \
    sameersbn/redis:latest
}

### 启动 etcd 服务
star_etcd() {
  rm -rf $(pwd)/etcd && mkdir -p $(pwd)/etcd && \
    docker stop etcd && \
    docker rm etcd || true && \
    docker run -d \
    -p 2379:2379 \
    -p 2380:2380 \
    --mount type=bind,source=$(pwd)/etcd,destination=/etcd-data \
    --name etcd \
    quay.io/coreos/etcd \
    /usr/local/bin/etcd \
    --name s1 \
    --data-dir /etcd-data \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-peer-urls http://0.0.0.0:2380 \
    --initial-advertise-peer-urls http://0.0.0.0:2380 \
    --initial-cluster s1=http://0.0.0.0:2380 \
    --initial-cluster-token tkn \
    --initial-cluster-state new
}

#### 启动elasticsearch 服务
start_es() {
  docker stop es
  docker rm es
  docker run -d --name es \
    -p 9200:9200 -p 9300:9300 \
    -e "discovery.type=single-node" \
    -v $(pwd)/elasticsearch:/usr/share/elasticsearch/data \
    spencezhou/elasticsearch
}

install_prometheus() {
    wget https://github.com/prometheus/prometheus/releases/download/v2.26.0/prometheus-2.26.0.linux-amd64.tar.gz
    tar -C /usr/local/ -xvf prometheus-2.26.0.linux-amd64.tar.gz
    ln -sv /usr/local/prometheus-2.26.0.linux-amd64/ /usr/local/Prometheus
}

function start_menu() {
    clear
    green " ===================================="
    green " Datacenter docker 操控中心"
    green " 系统：debian10+gozero"
    green " ===================================="
    echo
    green " 1. mysql"
    green " 2. redis"
    green " 3. etcd"
    green " 4. elasticsearch"
    green " 5. prometheus"
    green " 6. all"
    yellow " 0. Exit"
    echo
    read -p "输入数字:" num
    case "$num" in
    1)
        start_mysql
        ;;
    2)
        start_redis
        ;;
    3)
        star_etcd
        ;;
    4)
        start_es
        ;;
    5)
        start_mysql
        start_redis
        star_etcd
        start_es
        ;;
    6)
        install_prometheus
        ;;
    0)
        exit 1
        ;;
    *)
        clear
        red "请输入正确选项"
        sleep 2s
        start_menu
        ;;
    esac
}

start_menu