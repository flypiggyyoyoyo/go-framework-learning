-[x] func(http.ResponseWriter, *http.Request) 和 func(h http.ResponseWriter, r *http.Request)
  - func(http.ResponseWriter, *http.Request) 是一个函数类型的定义
  - func(h http.ResponseWriter, r *http.Request) 是一个函数声明

-[x] http.ListenAndServe(port,Hander)  
  - 第一个参数是地址。Handler是一个接口，需要实现方法 ServeHTTP ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理

- [x] FormValue(key)
  - 该方法的定义如下：
    ```go
    func (r *Request) FormValue(key string) string
    ```
    它接收一个字符串类型的 key 作为参数，表示要获取的表单数据的键。
    返回值是一个字符串，即该键对应的值。如果指定的键不存在，会返回空字符串。
- [x] c.Req.URL.Query().Get(key)
  - URL 是 http.Request 结构体的一个字段，类型为 *url.URL，它指向一个 url.URL 结构体实例。这个结构体包含了请求的 URL 信息，如协议（scheme）、主机名（host）、路径（path）、查询参数（query）等。例如，对于 URL https://example.com/search?keyword=go&page=2，URL 字段就会存储这个完整 URL 的解析结果。
  - Query() 是 url.URL 结构体的一个方法，其作用是解析 URL 中的查询参数部分，并将其转换为 url.Values 类型。url.Values 本质上是一个 map[string][]string 类型的映射，用于存储键值对形式的查询参数。例如，对于上述 URL 中的 ?keyword=go&page=2，Query() 方法会将其解析为 map[string][]string{"keyword": {"go"}, "page": {"2"}}，定义如下
    ```go
    func (u *URL) Query() Values
    ```
  - Get(key) 是 url.Values 类型的一个方法，用于从解析后的查询参数映射中获取指定键（key）对应的值。它接收一个字符串类型的 key 作为参数，并返回该键对应的第一个值（因为一个键可能对应多个值，这里只返回第一个）。如果指定的键不存在，则返回空字符串。定义如下
    ```go
      func (v Values) Get(key string) string
    ```
    
- [x] http.ResponseWriter.header()
  - header()方法用来获取响应头对象
  - 有header().set()方法来设置响应头部字段和指定相应格式
  - ``func (h Header) Set(key, value string)``
  - key：此参数类型为字符串（string），它代表 HTTP 头部字段的名称。在 HTTP 协议里，头部字段是由键值对构成的，key 就是其中的键。例如，常见的头部字段名有 "Content - Type"、"User - Agent"、"Set - Cookie" 等。
    value：同样是字符串（string）类型，它表示与 key 相对应的头部字段的值。例如，当 key 为 "Content - Type" 时，value 可能是 "application/json" 或者 "text/plain" 等
- [x] http.ResponseWriter.Write()
  - 用于向客户端发送响应体数据
- [x] encode:= json.NewEncoder(c.Writer)
  - encoding/json 包中的一个函数，用于创建一个 *json.Encoder 类型的对象。json.Encoder 是一个用于将 Go 对象编码为 JSON 格式的编码器
  - c是结构体实例，Writer是c中的一个http.ResponseWriter类型的字段
- [x] encoder.Encode(obj)
  - 调用Encode方法，把go类型的obj转换位json类型并尝试写入encoder中，返回error返回值
- [x] 动态路由
  - 路由规则（如 /p/:lang/doc）会被拆分成多个节点，每个节点对应路径中的一部分。每个节点有一个 pattern 属性，这个属性用于记录完整的路由规则。不过，只有路径中最后一个节点的 pattern 属性会被设置为完整的路由规则，其他中间节点的 pattern 属性为空。
  - 以路由规则 /p/:lang/doc 为例，它会被拆分成三个节点：p、:lang 和 doc。
    - p 节点：作为路径的第一层节点，它的 pattern 属性为空。
    - :lang 节点：作为路径的第二层节点，同样 pattern 属性为空。
    - doc 节点：作为路径的第三层节点，也是最后一层节点，它的 pattern 属性会被设置为 /p/:lang/doc。
  - 当进行路由匹配时，系统会根据节点的 pattern 属性来判断匹配是否成功。具体规则是：若匹配结束时，当前节点的 pattern 属性不为空，说明找到了对应的路由规则，匹配成功；反之，若 pattern 属性为空，则匹配失败