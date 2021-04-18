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

start_adminer() {
    docker kill datacenter_adminer
    docker rm datacenter_adminer
    docker run --name datacenter_adminer \
        -p 9434:8080 \
        --link mysql \
        -e 'MYSQL_PASSWORD=admin' \
        -d adminer 
}

##### 启动mysql 服务
start_mysql() {
  docker kill mysql
  docker rm mysql
  docker run --name mysql -d -p 3306:3306 -e 'MYSQL_ROOT_PASSWORD=admin' -v /home/dev/mysql8:/var/lib/mysql mysql:8.0.21
}

##### 启动redis 服务
start_redis() {
  docker stop redis
  docker rm redis
  echo "Start Redis Service..."
  docker run --name redis -d \
    --publish 6379:6379 \
    --env 'REDIS_PASSWORD=admin' \
    --volume /home/dev/redis:/var/lib/redis \
    sameersbn/redis:latest
}

### 启动 etcd 服务
start_etcd() {
  rm -rf /home/dev/etcd && mkdir -p /home/dev/etcd && \
    docker stop etcd && \
    docker rm etcd || true && \
    docker run -d \
    -p 2379:2379 \
    -p 2380:2380 \
    --mount type=bind,source=/home/dev/etcd,destination=/etcd-data \
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
    -v /home/dev/elasticsearch:/usr/share/elasticsearch/data \
    spencezhou/elasticsearch
}

install_prometheus() {
    wget https://github.com/prometheus/prometheus/releases/download/v2.26.0/prometheus-2.26.0.linux-amd64.tar.gz
    tar -C /usr/local/ -xvf prometheus-2.26.0.linux-amd64.tar.gz
    ln -sv /usr/local/prometheus-2.26.0.linux-amd64/ /usr/local/Prometheus
    rm prometheus-2.26.0.linux-amd64.tar.gz

}

install_grafana() {
    sudo apt-get install -y adduser libfontconfig1
    wget https://dl.grafana.com/oss/release/grafana_7.5.4_amd64.deb
    sudo dpkg -i grafana_7.5.4_amd64.deb
    rm grafana_7.5.4_amd64.deb
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
    green " 6. grafana"
    green " 7. all"
    green " 8. stop all container"
    yellow " 0. Exit"
    echo
    read -p "输入数字:" num
    case "$num" in
    1)
        start_mysql
        start_adminer
        ;;
    2)
        start_redis
        ;;
    3)
        start_etcd
        ;;
    4)
        start_es
        ;;
    5)
        start_mysql
        start_redis
        start_etcd
        start_es
        ;;
    6)
        install_prometheus
        ;;
    7)
        install_grafana
        ;;
    8)
        docker stop es
        docker stop etcd
        docker stop mysql
        docker stop redis
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