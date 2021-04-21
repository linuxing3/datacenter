#!/bin/bash
## author by linuxing3
## Email linuxing3@qq.com
## wechat linuxing3

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
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup /app/${myserver} -f /app/${mydir}${mycnf} &
}

# 附加Rpc服务启动函数
RpcServerPlus(){
    mydir=$1/$2/rpc
    myserver=${2}-server
    mycnf=/etc/${2}.yaml
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup /app/${myserver} -f /app/${mydir}${mycnf} &
}

# 附加Rpc服务启动函数
RpcServerExtra(){
    # go run $GOROOT/src/go-zero-admin/rpc/oms/oms.go -f $GOROOT/src/go-zero-admin/rpc/oms/etc/oms.yaml 
    myserver=${1}-server
    mycnf=${1}/rpc/etc/${1}.yaml
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup /app/${myserver} -f /app/${mycnf} &
}

# 启动Api服务
StartServer(){
    mydir=$1
    myserver=$2
    mycnf=$3
    kill -9 $(ps -ef|grep "${myserver}"|awk '{print $2}') >/dev/null 2>&1
    nohup /app/${myserver} -f /app/${mydir}${mycnf} &
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

StartAllServer
