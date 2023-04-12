# StarBlog
一款基于gin+vue的前后端分离的博客系统

# 进度日志

- [x] 建立文件夹
- server：后端代码
  - config：存放配置文件，管理网站配置参数
  - api：控制器，存放api接口，可细分api版本
  - middleware：考虑到前后端分离，后续需要中间件来解决跨域问题,同时存jwt鉴权
  - model：数据库的存储和访问，包含一些数据库关键组建
    - 数据库实体类：这些类通常用于描述数据库表的结构，包含表中每个字段的名称、数据类型、约束等信息。在使用ORM框架时，这些实体类通常会被自动生成。 
    - 数据库连接池：这是一个用于管理数据库连接的组件，它可以帮助应用程序更加高效地使用数据库连接资源，避免频繁地创建和销毁连接，从而提高应用程序的性能和可伸缩性。
  - routes：路由接口
  - service: 业务逻辑代码
  - utils：公共工具包，需要全局使用的模块放这里，作为功能模块
  - upload：上传文件夹
  - asset：资源文件夹
- web：前端代码
- [x] 初始化项目和配置参数
- server
  - config：新建config.ini文件，使用ini封装语法进行io操作
  - routers：新建router.go文件，进行初始化路由，先随便测试了一下接口
  - utils：新建setting.go文件，进行数据处理和配置文件初始化
- [x] 配置数据库 
---
- 2023.02.22 
  - [x] 完善api架构
  - [x] 错误模块构建
  - [x] 基本路由接口构建
---
- 2023.02.24 开始后端核心业务编程：用户模块接口
  - [x] 实现新增用户接口
  - [x] 实现后端验证注册用户时账号密码是否为空
  - [x] 实现采用分页的方式查看用户列表
  - [x] 实现了用户密码加盐存入数据库
  - [x] 修改后端验证注册用户时判断账号密码是否为空的逻辑
  - [x] 实现删除用户接口
  - [x] 将业务逻辑抽离到Service层，model层现在只负责与数据库的交互，存放实体类
  - [x] 实现编辑用户接口，暂时只有username，username要求不能为空，且不可和已有的重复
---
- 2023.03.28 开始编写分类模块接口
  - [x] 通过代码复用，完成了新增分类、查看分类列表、编辑分类、删除分类接口
  - [x] 修改了分页逻辑，目前两个参数可以均不填写，设定pagesize后，如果不设定pagenum则默认为第一页
---
- 2023.03.29 开始编写文章模块接口
  - [x] 通过代码复用，完成了添加文章、编辑文章、删除文章的接口
  - [x] 完成了查看单个文章信息、查看分类下所有文章、查看文章列表等功能
--- 
- 2023.03.31 开始编写登录模块
  - [x] 完成了jwt生成、解析、中间件的编写
  - [x] 完成了jwt鉴权，重新路由分类，编写完管理员登录接口，实现了登录验证密码是否正确
---
- 2023.04.03 开始编写上传模块
  - [x] 完成了上传到本地的代码
  - [x] 重构了router分类，更规整
---
- 2023.04.04 开始编写日志处理系统
  - [x] 新建log文件夹，用来存放日志
  - [x] 实现了自定义log文件
  - [x] 实现了log的分割，方便查询
- 2023.04.06 
  - [x] 重构数据验证，采用了validator
  - [x] 重构查看列表，添加显示数据总数，帮助前端处理
  - [x] 完成了跨域处理
- 2023.04.06
  - [x] 搭建好前端环境
- 2023.04.07
  - [x] 完善好前端路由配置与element ui导入
  - [x] 编写好了login页面样式
- 2023.04.10
- [x] 完成前端数据验证部分
- [x] 完善前端与后端交互
- [x] 解决了跨域问题
- 2023.04.11
  - [x] 完成login业务逻辑
  - [x] 实现了sessionstorage的存储与进行页面认证判断
  - [x] 分配了admin布局与组件
- 2023.04.12
  - [x] 规划admin页面内容
  - [x] 完成退出登录功能

# Todo
- [x] ~~采用第三方库validator来进行数据校验~~
- [ ] 更换为使用钩子函数来完成密码加盐的操作
- [x] ~~实现查看单个分类下所有文章~~
- [ ] 实现添加文章时可以使用多个标签和分类
- [ ] 通过反射分离结构体属性
- [ ] 修改分类与文章的关联模式为many2many，用户与文章的关联模式为belong。
- [ ] 完善token错误处理
- [ ] 增加文件上传到oss
- [ ] 完善权限控制机制
- [ ] 设计好各页面样式
- [ ] 完善跨域配置
- [ ] 实现token判断
- [ ] 实现记住密码

- 

# Issue

- [ ] 解决result.Error != gorm.ErrRecordNotFound问题
- [x] ~~分页功能存在问题，需要解决关于无法定位当前页的问题~~/原因：代码顺序写错了，把find写在了offset之前，Limit方法限制每页返回的记录数，Offset方法确定从哪个记录开始返回结果。
- [ ] 查看分类下所有文章时，没有判断如果分类不存在时如何，如果分类不存在仍然会返回空切片。
- [x] ~~处理密码验证~~/采用了随机salt，并将salt存入数据库，之后通过生成对比。
- [x] ~~StandardClaims结构体中的NotBefore和ExpiresAt无法直接使用时间戳~~/通过查看jwt-go/v4库官方文档，发现新版采用了jwt.NewTime函数，并且接受一个float64时间戳。
- [ ] 关于role处理存在一些问题，应当解决如何创建管理员用户，目前是强制将role修改为2；
- [x] ~~跨域中间件存在一些问题~~ /初步解决方案是允许所有跨域，之后再进行完善



 


