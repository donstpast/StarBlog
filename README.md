# StarBlog
一款基于gin+vue的前后端分离的博客系统
# 进度日志

- [x] 建立文件夹
- server：后端代码
  - config：存放配置文件，管理网站配置参数
  - api：控制器，存放api接口，可细分api版本
  - middleware：考虑到前后端分离，后续需要中间件来解决跨域问题
  - model：管理数据库存储、读写
  - routes：路由接口
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

- 2023.02.22 
  - [x] 完善api架构
  - [x] 错误模块构建
  - [x] 基本路由接口构建
- 2023.02.24 开始核心业务编程
  - [x] 实现新增用户接口
  - [x] 实现后端验证注册用户时账号密码是否为空
  - [x] 实现采用分页的方式查看用户列表
  - [x] 实现了用户密码加盐存入数据库
  - [x] 修改后端验证注册用户时判断账号密码是否为空的逻辑
  - [x] 实现删除用户接口
  - [x] 将业务逻辑抽离到Service层，model层现在只负责与数据库的交互，存放实体类

# Todo
- [ ] 解决result.Error != gorm.ErrRecordNotFound问题
- [ ] 采用第三方库validator来进行数据校验
- [ ] 更换为使用钩子函数来完成密码加盐的操作

 


