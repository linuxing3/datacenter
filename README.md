# 基于go-zero 框架写的一个数据中台中心

## 架构图
![中台系统](https://img2020.cnblogs.com/blog/203395/202012/203395-20201217094615171-335437652.jpg "中台架构")

### 已完成的功能列表
- [x] 微信公众号登陆
- [x] 七牛上传获取token
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

## 如何运行

### 先启动mysql redis etcd 服务
```shell
sh server.sh
```
### 输出如下 就显示 成功了
``` 
mysql
mysql
8d5d4b381ab7abe8947f532422255cd172f214ab4a6b0533da1619259e1cc4a5
redis
redis
Start Redis Service...
1fc187a9d82f0942dd60cac76c723a5bc531e1b67424384d04e7a69dad1362f0
etcd
etcd
98f88d81e1e218d4d53c608e5a68cd70254df221bdb34c9beab37a7473971ba0
es
es
04b37b58f10411fa8ab5894c917266cc7bc7f9a96988908fff9b3734e6259ad4
```
### 需要注意的是 elasticsearch 第一次用的时候，需要初始化密码 执行下面的操作
```bash
➜ ~ docker exec -it es /bin/bash
[root@04b37b58f104 elasticsearch]# sh /usr/share/elasticsearch/bin/elasticsearch-setup-passwords auto
...
Changed password for user elastic
PASSWORD elastic = somepassword
```
### 看到上面最后行 
```
Changed password for user elastic
PASSWORD elastic = somepassword
```
### 得到elastic 的帐号: elastic ,密码: somepassword 将这个填入search/rpc/search.yaml 文件中


    接着导入 sql.sql到 mysql数据中 ,如果有工具自行导入,下面仅参考
```
docker exec -it mysql /bin/bash
mysql -uroot -padmin
mysql > set name utf8
mysql > create databse datacenter
mysql > use datacenter
mysql > source sql.sql
```
### 然后分别把配置文件 ,文件下面分别对应了一个rpc.example.yaml的文件，复制，基本就没有问题

```
cat etc/datacenter-api.yaml #网关配置
cat user/rpc/etc/rpc.yaml #用户信息配置
cat common/rpc/etc/rpc.yaml #公共配置
cat votes/rpc/etc/rpc.yaml #投票配置
cat search/rpc/etc/search.yaml #搜索配置
cat questions/rpc/etc/questions.yaml #抽奖配置
```
### 然后启动 服务 ,应该我们要启动
```
sh restart.sh
```
### 输出如下
```
➜  datacenter.bak git:(master) ✗ sh restart.sh              
appending output to nohup.out
appending output to nohup.out
appending output to nohup.out
appending output to nohup.out 
appending output to nohup.out 
appending output to nohup.out    
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

### 在postman 导入数据 数据中台中心.postman_collection.json  就可以很愉快的玩耍了


## New RPC

greet.proto

```
Message greetReq {
	
}
service {
	rpc Login() ()
}
```

```bash
goctl rpc new greet
goctl rpc template -o greet.proto
goctl rpc proto -src greet.proto -dir .
```

## New API

movies.api

```
//请求电影信息
type MovieReq {
	Id string `json:"id"` //基本一个手机号码就完事
}

//返回电影信息
type MovieReply {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

@server(
	group: movie
)
service movie-api {
	@handler ping
	post /movies/ping ()
	
	@handler search
	post /movies/search (MovieReq) returns (MovieReply)
	
}

@server(
	group: movie
	middleware: Usercheck
)
service movie-api {
	@handler movieInfo
	get /movies/info (MovieReq) returns (MovieReply)
}
```


```bash
goctl api new movies

goctl api go -api movies.api -dir movies/api

goctl model mysql ddl -src movie.sql -dir movies/model

go run movies/api/movie.go -f movies/api/etc/movie-api.yaml
```
