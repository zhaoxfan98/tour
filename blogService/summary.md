# 开启博客之路

本次博客项目选用gin框架完成开发

![](./gin.png)

```
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```

首先调用`gin.Default`方法来创建默认的Engine实例，它会在初始化阶段就引入Logger和Recovery中间件，能够保障应用程序的最基本运作

- Logger:输出请求日志，并标准化日志的格式
- Recovery：异常捕获，也就是针对每次请求处理进行recovery处理，防止因为出现panic导致服务崩溃，并同时标准化异常日志的格式

```
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		AppEngine:              defaultAppEngine,
		UseRawPath:             false,
		UnescapePathValues:     true,
		MaxMultipartMemory:     defaultMultipartMemory,
		trees:                  make(methodTrees, 0, 9),
		delims:                 render.Delims{Left: "{{", Right: "}}"},
		secureJsonPrefix:       "while(1);",
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}
```

- RouterGroup:路由组，所有的路由规则都由 *RouterGroup 所属的方法进行管理，在 gin 中和 Engine 实例形成一个重要的关联组件
- RedirectTrailingSlash：是否自动重定向，如果启用了，在无法匹配当前路由的情况下，则自动重定向到带有或不带斜杠的处理程序去
- RedirectFixedPath：是否尝试修复当前请求路径，也就是在开启的情况下，gin 会尽可能的帮你找到一个相似的路由规则并在内部重定向过去，主要是对当前的请求路径进行格式清除（删除多余的斜杠）和不区分大小写的路由查找等。
- HandleMethodNotAllowed：判断当前路由是否允许调用其他方法，如果当前请求无法路由，则返回 Method Not Allowed（HTTP Code 405）的响应结果。如果无法路由，也不支持重定向其他方法，则交由 NotFound Hander 进行处理。
- ForwardedByClientIP：如果开启，则尽可能的返回真实的客户端 IP，先从 X-Forwarded-For 取值，如果没有再从 X-Real-Ip。
- UseRawPath：如果开启，则会使用 url.RawPath 来获取请求参数，不开启则还是按 url.Path 去获取。
- UnescapePathValues：是否对路径值进行转义处理。
- MaxMultipartMemory：相对应 `http.Request ParseMultipartForm` 方法，用于控制最大的文件上传大小。
- trees：多个压缩字典树（Radix Tree），每个树都对应着一种 HTTP Method。你可以理解为，每当你添加一个新路由规则时，就会往 HTTP Method 对应的那个树里新增一个 node 节点，以此形成关联关系。
- delims：用于 HTML 模板的左右定界符。

总的来讲，Engine实例就像引擎一样，与整个应用的运行、路由、对象、模板等管理和调度都有关联，另外通过上述的解析，你可以发现其实 gin 在初始化默认已经替我们做了很多事情，可以说是既定了一些默认运行基础。

```
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}
```

GET方法来将定义的路由注册进去

- 计算路由的绝对路径
- 合并现有和新注册的Handler，并创建一个函数链`HandlersChain`
- 将当前注册的路由规则（含 HTTP Method、Path、Handlers）追加到对应的树中

```
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() { debugPrintError(err) }()

	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}
```

Run会通过解析地址，再调用`http.ListenAndServe`将实例作为`Handler`注册进去，然后启动服务，开始对外提供HTTP服务

接下来在下面的小节中，我们就可以根据 RESTful API 的基本规范，针对我们的业务模块设计路由规则，从业务角度来划分多个管理接口

1. 标签管理

- 新增标签
- 删除指定标签
- 更新指定标签
- 获取标签列表

2. 文章管理

- 新增文章
- 删除指定文章
- 更新指定文章
- 获取指定文章
- 获取文章列表

本章节完成了一个项目最初始的第一步，也就是项目的标准目录创建、数据库设计、数据模型编写、接口方法的设计和接口处理方法及启动接入

## 2.3 编写公共组件

实际上在公司的项目中，度会有一类组件，常称其为基础组件，又或是公共组件，他们是不带强业务属性的，串联着整个应用程序，一般由负责基建或第一批搭建的该项目的同事进行梳理和编写，如果没有这类组件，谁都写一套，是非常糟糕的，并且这个应用程序是无法形成闭环的。

因此在这一章节我们将完成一个 Web 应用中最常用到的一些基础组件，保证应用程序的标准化，一共分为如下五个板块

- 错误码标准化
- 配置管理
- 数据库连额
- 日志写入
- 响应处理

### 配置管理

- 在启动时：可以进行一些基础应用属性、连接第三方实例（MySQL、NoSQL）等等的初始化行为
- 在运行中：可以监听文件或其他存储载体的变更来实现热更新配置的效果，例如：在发现有变更的话，就对原有配置值进行修改，以此达到相关联的一个效果。如果更深入业务使用的话，我们还可以通过配置的热更新，达到功能灰度的效果，这也是一个比较常见的场景

`Server` 配置服务，设置gin的运行模式、默认的HTTP监听端口、允许读取和写入的最大持续时间
`App` 应用配置，设置默认每页数量、所允许的最大每页数量以及默认的应用日志存储路径
`Database` 数据库配置，主要是连接实例所必需的基础参数

### 日志写入

日志组件内要使用到的第三方的开源库 lumberjack，它的核心功能是将日志写入滚动文件中，该库支持设置所允许单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数，如果出现超出设置项的情况，就会对日志文件进行滚动处理。

而我们使用这个库，主要是为了减免一些文件操作类的代码编写，把核心逻辑摆在日志标准化处理上。

### 响应处理

在应用程序中，与客户端对接的常常是服务端的接口，那客户端是怎么知道这一次的接口调用结果是怎么样的呢？一般来讲，主要是通过对返回的 HTTP 状态码和接口返回的响应结果进行判断，而判断的依据则是事先按规范定义好的响应结果。

这一节编写统一处理接口返回的响应处理方法

## 小结

本章节主要是针对项目的公共组件初始化，做了大量的规范制定、公共库编写、初始化注册等行为

## 2.4 生成接口文档

Swagger针对OpenAPI规范提供了大量与之相匹配的工具集，能够充分利用OpenAPI规范去映射生成所有与之关联的资源和操作去查看和调用RESTful接口

从功能使用上来讲，OpenAPI规范能够帮助我们描述一个API的基本信息，比如

- 有关该API的描述
- 可用路径
- 在每个路径上的可用操作
- 每个操作的输入/输出格式

Swagger 相关的工具集会根据 OpenAPI 规范去生成各式各类的与接口相关联的内容，常见的流程是编写注解 =》调用生成库-》生成标准描述文件 =》生成/导入到对应的 Swagger 工具。

## 2.6 模板开发：标签管理

