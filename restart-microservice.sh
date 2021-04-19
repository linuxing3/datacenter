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

logcmd() {
    eval $1 | tee -ai /var/atrandys.log
}


# 设置服务的目录
GatewayPath=$(pwd) #网关服务
userPath=$(pwd)/user/rpc #用户服务
commonPath=$(pwd)/common/rpc #公共服务
votesPath=$(pwd)/votes/rpc #投票服务
moviePath=$(pwd)/movies/rpc #movie服务
bookAddPath=$(pwd)/bookstore/rpc/adder #adder服务
bookCheckPath=$(pwd)/bookstore/rpc/checker #checker服务

# 基本配置文件名
configPath=/etc/rpc.yaml #配置文件
GateWayCnf=/etc/datacenter-api.yaml

# 设置服务可执行文件名称
UserRpc=user-server #定义网关服务
CommonRpc=common-server #定义公共服务
VotesRpc=votes-server #定义投票服务
MovieRpc=movie-server #定义movie服务
BookAdderRpc=bookadder-server #定义adder服务
BookCheckerRpc=bookchecker-server #定义checker服务

# 设置GatewayApi可执行文件名称
GateWayApi=datacenter-server #定义网关服务

# 普通Rpc服务启动函数
RpcServer(){
    mydir=$1
    myserver=$2
    mycnf=$3
    cd ${mydir}
    go build -o ${myserver} $mydir/rpc.go
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup ${mydir}/${myserver} -f ${mydir}${mycnf} &
}

# 附加Rpc服务启动函数
RpcServerPlus(){
    mydir=$1/$2/rpc
    myserver=${2}-server
    mycnf=/etc/${2}.yaml
    cd ${mydir}
    go build -o ${myserver} $mydir/$2.go
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup ${mydir}/${myserver} -f ${mydir}${mycnf} &
}

# 附加Rpc服务启动函数
RpcServerExtra(){
    # go run $GOROOT/src/go-zero-admin/rpc/oms/oms.go -f $GOROOT/src/go-zero-admin/rpc/oms/etc/oms.yaml 
    srcdir=$GOROOT/src/go-zero-admin/rpc/$1
    exedir=$GOPATH/src/github.com/linuxing3/datacenter
    myserver=${1}-server
    mycnf=${1}/rpc/etc/${1}.yaml
    cd ${srcdir}
    go build -o ${exedir}/${myserver} ${srcdir}/$1.go
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup ${exedir}/${myserver} -f ${exedir}/${mycnf} &
}

# 启动Api服务
StartServer(){
    mydir=$1
    myserver=$2
    mycnf=$3
    cd ${mydir}
    go build -o ${myserver} $mydir/datacenter.go
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup ${mydir}/${myserver} -f ${mydir}${mycnf} &
}

Logging() {
    tail -F $1/rpc/nohup.out
}

# 普通Rpc服务启动函数
StopAllServer(){
    kill -9 $(ps -ef|grep "${UserRpc}"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "${CommonRpc}"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "${VotesRpc}"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "${MovieRpc}"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "${BookAdderRpc}"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "${BookCheckerRpc}"|awk '{print $2}') >/dev/null 2>&1
    
    kill -9 $(ps -ef|grep "${GatewayApi}"|awk '{print $2}') >/dev/null 2>&1

    kill -9 $(ps -ef|grep "oms-server"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "pms-server"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "ums-server"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "sms-server"|awk '{print $2}') >/dev/null 2>&1
    kill -9 $(ps -ef|grep "sys-server"|awk '{print $2}') >/dev/null 2>&1
}

StartAllServer() {
    # StopAllServer
    RpcServer ${commonPath} ${CommonRpc} ${configPath}
    RpcServer ${userPath} ${UserRpc} ${configPath}
    RpcServer ${votesPath} ${VotesRpc} ${configPath}
    RpcServer ${moviePath} ${MovieRpc} ${configPath}
    RpcServer ${bookAddPath} ${BookAdderRpc} ${configPath}
    RpcServer ${bookCheckPath} ${BookCheckerRpc} ${configPath}

    RpcServerPlus ${GatewayPath} search
    RpcServerPlus ${GatewayPath} questions

    RpcServerExtra oms
    RpcServerExtra pms
    RpcServerExtra ums
    RpcServerExtra sms
    RpcServerExtra sys

    StartServer ${GatewayPath} ${GateWayApi} ${GateWayCnf}
}


function start_menu() {
    clear
    green " ===================================="
    green " Datacenter/Go-zero-admin微服务控制中心"
    green " 系统：debian10+"
    green " ===================================="
    echo
    red " 1. go-zero-admin服务"
    red " 2. 公共服务"
    red " 3. 用户服务"
    red " 4. 投票服务"
    red " 5. 电影书籍服务"
    red " 6. 搜索服务"
    red " 7. 问答抽奖服务"
    red " 8. 网关入口"
    red " 9. 启动全部服务"
    red " 10. 停止全部服务"
    red " 11. prometheus"
    red " 12. grafana"
    yellow " 0. Exit"
    echo
    read -p "输入数字:   " num
    case "$num" in
    1)
        RpcServerExtra oms
        RpcServerExtra pms
        RpcServerExtra ums
        RpcServerExtra sms
        RpcServerExtra sys
        ;;
    2)
        RpcServer ${commonPath} ${CommonRpc} ${configPath}
        ;;
    3)
        RpcServer ${userPath} ${UserRpc} ${configPath}
        ;;
    4)
        RpcServer ${votesPath} ${VotesRpc} ${configPath}
        ;;
    5)
        RpcServer ${moviePath} ${MovieRpc} ${configPath}
        RpcServer ${bookAddPath} ${BookAdderRpc} ${configPath}
        RpcServer ${bookCheckPath} ${BookCheckerRpc} ${configPath}
        ;;
    6)
        RpcServerPlus ${GatewayPath} search
        ;;
    7)
        RpcServerPlus ${GatewayPath} questions
        ;;
    8)
        StartServer ${GatewayPath} ${GateWayApi} ${GateWayCnf}
        ;;
    9)
        StartAllServer
        ;;
    10)
        StopAllServer
        ;;
    11)
        /usr/local/Prometheus/prometheus --config.file=config/prometheus/config.yml
        ;;
    12)
        /etc/init.d/grafana start
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