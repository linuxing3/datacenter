
# 整合现有的rpc代码

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

添加一个新配置项`MoviesRpc`，属性是`zrpc.RpcClientConf`，这是我们用来调用后台rpc服务器的客户端

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

#### `movieclient`结构体定义在`movie.pb.go`文件中，包含了一个客户端连接的~接口~

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

### `/internal/handler/routes`中添加路由route和处理器handler

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

### `internal/handler/movies/movieinfohandler`中添加对业务逻辑代码`logic`的调用方法

```go
		l := logic.NewMovieInfoLogic(r.Context(), ctx)
		resp, err := l.MovieInfo(req)
```

### `internal/logic/movies/movieinfologic`中才是真正处理数据的业务逻辑（这里终于可以和Rpc进行交互了！！！）

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

## 整合`go-zero-admin`

比如想使用一个不错的rpc代码库`go-zero-admin`，但这个库是本地模块的，而没有`github.com`，使用`go mod tidy`是无法从github抓取的


###  直接放入GOROOT

一个变通的办法就是把这个库下载到`GOROOT`，然后就可以直接运行

```bash
git clone github.com/feihua/zero-admin $GOROOT/go-zero-admin
go run $GOROOT/src/go-zero-admin/rpc/oms/oms.go -f $GOROOT/src/go-zero-admin/rpc/oms/etc/oms.yaml
```

然后就可以直接导入了！
```go
import "go-admin-zero/rpc/oms/omsclient"
```

### 如何使用前端api调用handler和logic代码？

拷贝相关代码到实际的`internal/handler`和`internal/logic`目录下

> 注意在每个文件中修改导入文件的路径

```go
	// logic "<原来的api目录>/internal/logic/member/address"
	logic "datacenter/internal/logic/member/address"
	// "<原来的api目录>/internal/svc"
	"datacenter/internal/svc"
	// "<原来的api目录>/internal/types"
	"datacenter/internal/types"
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