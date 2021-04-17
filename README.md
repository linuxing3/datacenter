# 基于 go-zero 框架写的一个数据中台中心示例

这是一个微服务`grpc`的自学项目，主要是帮助熟悉`go-zero`这个新的框架

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
```

### 使用脚本启动所有的Api服务和Rpc服务

```
sh restart.sh
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

## 踩坑开始：如何加入自己的API接口？

基本上使用`goctl`就可以顺利地生成

这里生成的`logic`和`handler`可以略作修改，拷贝到`/internal/logic/handler`下的`movies`目录下

方便在统一网关中调用，因为如果跨模块调用是不允许的！！！

```bash
# 创建目录
cd movies
mkdir -p api model rpc

# 生成api模板文件
goctl api new movies

# 使用api文件生成api目录
goctl api go -api movies.api -d ./api

# 使用sql文件自动生成数据模型（包含缓存和CRUD方法）
goctl model mysql ddl -s movie.sql -d ./model -c

# 运行api服务
go run ./api/movie.go -f ./api/etc/movie-api.yaml

```

## 踩坑继续：如何创建自己的RPC微服务？

```bash
cd movies

# 使用模板生成proto定义文件
goctl rpc template -o movie.proto

# 使用proto定义文件，生成rpc目录和文件
goctl rpc proto -s movie.proto -I ./movie -d ./rpc
```

### proto的默认定义文件都在`movie.pd.go`中

>  这是一个很大的文件，做了rpc的各方面定义和配置

[movies/rpc/movie/movie.pb.go](movies/rpc/movie/movie.pb.go)

### 在配置中加入mysql和redis设置

[movies/rpc/internal/config/config.go](movies/rpc/internal/config/config.go)

### 在SVC文件中加入mysql和redis设置

[movies/rpc/internal/svc/servicecontext.go](movies/rpc/internal/svc/servicecontext.go)

### 在logic文件中加入实际业务逻辑，通常是对数据库的增删改查

[movies/rpc/internal/logic/movieslogic.go](movies/rpc/internal/logic/movieslogic.go)

主要是在这里接受前端网关传来的请求数据，调用`MoviesLogic.svcCtx.MoviesModel.FindOne`的方法查找并返回结果

```go
func (l *MoviesLogic) Movies(in *movie.MovieListReq) (*movie.MovieListResp, error) {
	movieInt := shared.StrToInt64(in.Ids[0])

	movieInfo, _ := l.svcCtx.MoviesModel.FindOne(movieInt)

	list := make([]*movie.MovieResp, 0)
	list = append(list, &movie.MovieResp{
		Title: movieInfo.Title,
		Description: movieInfo.Description,
		Url: movieInfo.Url,
	})
	return &movie.MovieListResp{
		Data: list,
	}, nil
}
```

## ~划重点~：如何把网关DataCenter-api和我们新建的RPC连起来？

> 这里的坑最多！一步都不能错！

### `internal/config/config.go`中

添加一个新配置项`MoviesRpc`，属性是`zrpc.RpcClientConf`

### `etc/datacenter-api.yaml`中

添加一个新的设置项`MoviesRpc`，用于注册到`etcd`自动服务发现中心

```yaml
MoviesRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: movie.rpc
```

### `internal/svc/servicecontext.go`上下文定义文件中

在`ServiceContext`结构体中加入一个`MoviesRpc`, 类型为`movieclient.Movies`

#### `movieclient`的定义在`movie.pb.go`文件中，实际上是一个客户端连接的接口

```go
type moviesClient struct {
	cc grpc.ClientConnInterface
}

func (c *moviesClient) Movies(ctx context.Context, in *MovieListReq, opts ...grpc.CallOption) (*MovieListResp, error) {
	out := new(MovieListResp)
	err := c.cc.Invoke(ctx, "/movies.Movies/Movies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

```

#### 生成实例对象

`ServiceContext`是`datacenter-api`的上下文对象, 可以和后端进行交互

使用`movieclient.NewMovies`就可以生成`MoviesRpc`服务器对象

```go
	mr := movieclient.NewMovies(zrpc.MustNewClient(c.MoviesRpc, zrpc.WithUnaryClientInterceptor(timeInterceptor)))
```

### `/internal/handler/routes`中添加路由和处理器

```go
// import movie "datacenter/internal/handler/movies"

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Usercheck},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/movies/info",
					Handler: movie.MovieInfoHandler(serverCtx),
				},
			}...,
		),
	)
```

### `internal/handler/movies/movieinfohandler`中添加对`logic`的调用方法

```go
		l := logic.NewMovieInfoLogic(r.Context(), ctx)
		resp, err := l.MovieInfo(req)
```

### `internal/logic/movies/movieinfologic`中添加真正处理数据的业务逻辑（这里终于可以和Rpc进行交互了！！！）

主要是通过调用`datacenter-api`上下文中的`MoviesRpc`对象的`Movies`方法向后台`Rpc`服务器发送请求

这里的请求将被`Rpc`服务器捕捉到并处理，具体过程如下：

+ 请求数据先传给Rpc的`movieserver`的`Movies`方法

+ 生成一个`MovieLogic`，调用logic对象的`Movies`方法

+ `MovieLogic`调用自身的`svcCtx`, 里面包含了`MovieModel`这个类似于`ORM`的对象，并使用`FindOne`方法和`Mysql`数据库进行查询

请检查 `movies/rpc/internal/logic/movieslogic.go`文件, 最终的处理方法也叫`Movies`~

```go
// import "datacenter/movies/rpc/movieclient"
func (l *MovieInfoLogic) MovieInfo(req types.MovieReq) (*movieclient.MovieListResp, error) {
	id:= fmt.Sprintf("%s", req.Id)
	log.Printf("从请求中获取id： %v", id);

	ids := []string{id}
	return l.svcCtx.MoviesRpc.Movies(l.ctx, &movieclient.MovieListReq{
		Ids: ids,
	})
}
```

### 教训

+ 定义文件`proto`和`api`一定要设计精确，命名规范，否则后续修改非常痛苦

建议的写法

```proto
service Movies {
  rpc Movies(MovieListReq) returns(MovieListResp);
  rpc Movie(MovieReq) returns(MovieResp);
}
```

错误的写法（把参数类型写反了！！！单个对象和对象数组前后不一致，后续要导出修改）

```proto
service Movies {
  rpc Movies(MovieReq) returns(MovieListResp);
  rpc Movie(MovieListReq) returns(MovieResp);
}
```

## 感谢

主要学习了

+ [Go-zero](https://zeromicro.github.io/go-zero/about-us.html)

+ [我是如何用go-zero 实现一个中台系统](https://www.cnblogs.com/linuxing3/p/14148518.html)

+ [开源地址](https://github.com/jackluo2012/datacenter)

### 已完成的功能列表
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