上下文context

> 在 Web 开发中，上下文（Context）指的是在处理 Web 请求过程中，包含与该请求相关的各种信息和状态的一个容器或环境，通常包含请求信息，响应信息和中间件处理等。

- ``context``设计
  - 初始化方法：``newCoutext``
    - http.ResponseWriter``
    - http.Request
    - path
    - Method
    - StatusCode
  - 提供表单参数查询方法：``PostForm``
  - 提供HTTP请求URL中字符串参数查询方法：``Query``
  - 提供快速构造``String``/``Data``/``JSON``/``HTML``响应的方法
- 独立路由``Router``
  - addRouter
  - handle

**思考**  
怎么理解路由匹配的全过程？开发者先设计特定的触发-反馈（**路由添加**），客户访问（提供关键词），程序查找关键词（**路由匹配**），最后反馈