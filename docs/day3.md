前缀树实现动态路由

所谓前缀树，就是每一个节点的所有的子节点都拥有相同的前缀，请看[极客兔兔老师的图解](https://geektutu.com/post/gee-day3.html)，非常清晰，一目了然

- Trie树实现
  - node节点设计
    - pattern 完整的待匹配路由路径
    - part 将完整路由路径按照 ``/`` 分割后的一个片段
    - []*node 子节点
    - isWild 是否精确匹配
   - 路由注册（Trie树插入节点）
     - insert  递归插入，遍历查询能匹配当前``part``的节点，若没有则新建一个，插入
     - matchChild  遍历查询实现
   - 路由匹配
     - search 遍历当前``part``的全部子节点，递归匹配
     - matchChildren  查询全部子节点实现
- Router优化
  - 初始化
    - roots 存每种方法的根节点
    - handlers 存每种请求方式的``HandlerFunc``处理函数
    ```go
    // roots key eg, roots['GET'] roots['POST']
    // handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
    ```
  - parsePattern 路由路径按照``/``分开存储到切片中
  - addRouter  添加路由（调用Trie树的路由注册函数）
  - getRouter  根据HTTP请求的方法和路径，在路由树中查找匹配的路由，若找到则返回该路由最后的节点和该条路由上的动态参数map（动态参数：实际参数）
  - handle方法添加给Params赋值语句
- context优化
  - 添加``Params  map[string]string``存储参数
  - Params方法，提供对参数的访问
