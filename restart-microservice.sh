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
getewayPath=$(pwd) #网关服务
userPath=$(pwd)/user/rpc #用户服务
commonPath=$(pwd)/common/rpc #公共服务
votesPath=$(pwd)/votes/rpc #投票服务
moviePath=$(pwd)/movies/rpc #movie服务
bookAddPath=$(pwd)/bookstore/rpc/adder #adder服务
bookCheckPath=$(pwd)/bookstore/rpc/checker #checker服务

# 基本配置文件名
configPath=/etc/rpc.yaml #配置文件
gateWayCnf=/etc/datacenter-api.yaml

# 设置服务可执行文件名称
UserRpc=user-server #定义网关服务
CommonRpc=common-server #定义公共服务
VotesRpc=votes-server #定义投票服务
MovieRpc=movie-server #定义movie服务
BookAdderRpc=bookadder-server #定义adder服务
BookCheckerRpc=bookchecker-server #定义checker服务
geteWayApi=datacenter-server #定义网关服务

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

function start_menu() {
    clear
    green " ===================================="
    green " Datacenter 微服务 操控中心"
    green " 系统：debian10+gozero"
    green " ===================================="
    echo
    green " 1. 公共服务"
    green " 2. 用户服务"
    green " 3. 投票服务"
    green " 4. 电影服务"
    green " 5. 书籍服务"
    green " 6. 搜索服务"
    green " 7. 问答抽奖服务"
    green " 8. 网关入口"
    green " 9. prometheus"
    green " 10. 启动全部服务"
    yellow " 0. Exit"
    echo
    read -p "输入数字:" num
    case "$num" in
    1)
        RpcServer ${commonPath} ${CommonRpc} ${configPath}
        # Logging common
        ;;
    2)
        RpcServer ${userPath} ${UserRpc} ${configPath}
        # Logging user
        ;;
    3)
        RpcServer ${votesPath} ${VotesRpc} ${configPath}
        # Logging votes
        ;;
    4)
        RpcServer ${moviePath} ${MovieRpc} ${configPath}
        go run $GOROOT/src/go-zero-admin/rpc/oms/oms.go -f $GOROOT/src/go-zero-admin/rpc/oms/etc/oms.yaml
        ;;
    5)
        RpcServer ${bookAddPath} ${BookAdderRpc} ${configPath}
        RpcServer ${bookCheckPath} ${BookCheckerRpc} ${configPath}
        # tail -F bookstore/rpc/adder/nohup.out
        # tail -F bookstore/rpc/checker/nohup.out
        ;;
    6)
        RpcServerPlus ${getewayPath} search
        # Logging search
        ;;
    7)
        RpcServerPlus ${getewayPath} questions
        # Logging questions
        ;;
    8)
        StartServer ${getewayPath} ${geteWayApi} ${gateWayCnf}
        # tail -F nohup.out
        ;;
    9)
        /usr/local/Prometheus/prometheus --config.file=config/prometheus/config.yml
        ;;
    10)
        RpcServer ${commonPath} ${CommonRpc} ${configPath}
        RpcServer ${userPath} ${UserRpc} ${configPath}
        RpcServer ${votesPath} ${VotesRpc} ${configPath}
        RpcServer ${moviePath} ${MovieRpc} ${configPath}
        RpcServer ${bookAddPath} ${BookAdderRpc} ${configPath}
        RpcServer ${bookCheckPath} ${BookCheckerRpc} ${configPath}
        RpcServerPlus ${getewayPath} search
        RpcServerPlus ${getewayPath} questions
        StartServer ${getewayPath} ${geteWayApi} ${geteWayCnf}
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