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
    green " 1. start all servers"
    green " 2. stop all container"
    green " 3. prometheus"
    green " 4. grafana"
    green " 5. mysql client"
    yellow " 0. Exit"
    echo
    read -p "输入数字:" num
    case "$num" in
    1)
        docker-compose up -d
        ;;
    2)
        docker-compose down
        ;;
    3)
        install_prometheus
        ;;
    4)
        install_grafana
        ;;
    5)
        docker exec -it mysql mysql -uroot -padmin
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