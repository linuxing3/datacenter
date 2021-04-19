# 基于 go-zero 框架写的一个数据中台中心示例

这是一个微服务`grpc`的自学项目，主要是熟悉练习`go-zero`这个新的框架

+ [开源地址](https://github.com/linuxing3/datacenter)

## 关于`go-zero`

go-zero 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

go-zero 包含极简的 API 定义和生成工具 goctl，可以根据定义的 api 文件一键生成 Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript 代码，并可直接运行。

## 中台架构图

![中台系统](https://img2020.cnblogs.com/blog/203395/202012/203395-20201217094615171-335437652.jpg "中台架构")


## 如何运行？

### 先启动mysql、redis、es、 etcd 等服务，这里都使用docker

```shell
sh server.sh
```

输出如下 就显示 成功了

``` 
mysql
Start Redis Service...
etcd
es
```

### elasticsearch 第一次需要初始化密码 

执行下面的操作

```bash
# 进入docker
docker exec -it es /bin/bash
# 执行密码设置命令
[root@04b37b58f104 elasticsearch]# sh /usr/share/elasticsearch/bin/elasticsearch-setup-passwords auto
...
# 可以查看密码，得到elastic 的帐号: elastic ,密码: somepassword
Changed password for user elastic
PASSWORD elastic = somepassword
```

填入`search/rpc/search.yaml` 配置文件中

#### 踩坑

注意添加`--privileged=true`，如果发现无法正常启动无权限则，运行 `chmod -R 777 [mydata]`

### 导入 sql.sql到 mysql数据中 

+ 直接登录到Docker中的数据库

```
docker exec -it mysql /usr/bin/mysql -uroot -padmin
```

+ 导入数据库设置，生成数据表

```sql
mysql > set name utf8
mysql > create databse datacenter
mysql > use datacenter
mysql > source sql.sql
```

### 根据需要更新每个Rpc模块下的配置文件

```
cat etc/datacenter-api.yaml #网关配置，这里是api总的入口

cat user/rpc/etc/rpc.yaml #用户信息配置
cat common/rpc/etc/rpc.yaml #公共配置
cat votes/rpc/etc/rpc.yaml #投票配置
cat search/rpc/etc/search.yaml #搜索配置
cat questions/rpc/etc/questions.yaml #抽奖配置

cat oms/rpc/etc/oms.yaml #gozero配置
cat pms/rpc/etc/pms.yaml #gozero配置
cat ums/rpc/etc/ums.yaml #gozero配置
cat sys/rpc/etc/sys.yaml #gozero配置
```

### 使用脚本启动所有的Api服务和Rpc服务

```
sh restart-microservice.sh
```

### 可以分别查看是否启动成功
```
tail -F nohup.out  #网关的服务
tail -F user/rpc/nohup.out #用户的rpc服务
tail -F common/rpc/nohup.out #公共的
tail -F votes/rpc/nohup.out #投票的
tail -F search/rpc/nohup.out #搜索的
tail -F questions/rpc/nohup.out #问答抽奖
```

### 在postman 导入数据

`postman_collection.json`定义了基本的`postman`项目


## 扩展

### [整合现有的rpc代码](/docs/integrate.md)

### [jwt验证功能](/docs/jwt.md)


## 感谢

参考项目：

+ [Go-zero](https://zeromicro.github.io/go-zero/about-us.html)

+ [我是如何用go-zero 实现一个中台系统](https://www.cnblogs.com/linuxing3/p/14148518.html)

+ [开源地址](https://github.com/jackluo2012/datacenter)


## 功能列表

### 已完成的功能列表
- [x] go-zero-admin
    - [x] oms
    - [x] ums
    - [x] pms
    - [x] sys
- [x] 微信公众号登陆
- [x] 七牛上传获取token
- [x] 电影
    - [x] 查询信息
- [x] 投票
    - [x] 报名
    - [x] 报名列表
    - [x] 投票
- [x] 抽奖问答
    - [x] 活动信息
    - [x] 问答列表
    - [x] 提交答案 
    - [x] 获取得分 
    - [x] 抽奖
    - [x] 填写中奖人信息    
- [x] 搜索
    - [x] 基于elasticsearch

### 未完成的        
- [ ] 微信支付宝登陆
- [ ] 微信支付宝支付