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
---
- 2023.04.06 
  - [x] 重构数据验证，采用了validator
  - [x] 重构查看列表，添加显示数据总数，帮助前端处理
  - [x] 完成了跨域处理
---
- 2023.04.06
  - [x] 搭建好前端环境
---
- 2023.04.07
  - [x] 完善好前端路由配置与element ui导入
  - [x] 编写好了login页面样式
---
- 2023.04.10
- [x] 完成前端数据验证部分
- [x] 完善前端与后端交互
- [x] 解决了跨域问题
---
- 2023.04.11
  - [x] 完成login业务逻辑
  - [x] 实现了sessionstorage的存储与进行页面认证判断
  - [x] 分配了admin布局与组件
---
- 2023.04.12
  - [x] 规划admin页面内容
  - [x] 完成退出登录功能
---
- 2023.04.13
  - [x] 基本完成用户页面布局
  - [x] 实现了前端显示用户列表
  - [x] 初步完成分页功能
---
- 2023.04.14
  - [x] 修复了一些issue提到的问题
  - [x] 完成了根据用户输入用户名模糊查询列表
  - [x] 完成了删除用户功能
  - [x] 利用对话框实现了删除确认功能
---
- 2023.04.15
  - [x] 修复了一些issue提到的问题
  - [x] 完成了搜索功能表单设计
---
- 2023.04.17
  - [x] 修复了一些issue问题（role处理问题） 
  - [x] 完成了增加用户功能
  - [x] 完成了编辑用户功能
  - [x] 重构了编辑用户的后端代码（添加了判断用户名是否修改，如未修改不影响编辑；添加了关于是否修改用户密码的操作）
---
- 2023.04.18
  - [x] 完善了页面中中权限列的展示，采用了el-tag进行美化与名称格式化，去掉了原先的格式化方法
  - [x] 通过代码复用，完成了分类管理页面
  - [x] 通过代码复用，基本完成了文章管理页面
---
- 2023.04.19
  - [x] 完成了文章管理页面中按分类筛选功能
  - [x] 完成了编写文章页面的基本内容
  - [x] 内置了MdEditor
  - [x] 重构了一下关于新建文章的api，现在可以直接通过登录的token获取用户uid，
  - [x] 新增通过username查询对应用户id的service，方便通过解析token获得username后直接获取用户id
---
- 2023.04.20
  - [x] 完成了编写文章页面基本功能的编写
  - [x] 实现了在编写文章页面通过一个发布按钮判断是否存在文章，如果存在则是编辑，否则为新增（作者为最终编辑的用户）
  - [x] 修改文章列表按更新事件降序排列，保证最新的文章始终在最前面
  - [x] 修复了一些issue中的问题
---
- 2023.04.21
  - [x] 决定前台页面依旧采用ElementUI，完成项目搭建
  - [x] 基本确定前台样式，完成了基本布局
  - [x] 实现了文章卡片展示
---
- 2023.04.22
  - [x] 实现了前台调用MdEditor显示文章
  - [x] 实现了文章详情页的编写
  - [x] 完成了资料展示卡片的布局
  - [x] 修复了一些bug
  - [x] 通过代码复用，实现了关于评论的增删改查api
  - [x] 重新规划了一些数据库字段设计，进行了细节完善
---
- 2023.04.24
  - [x] 通过代码复用，实现了关于友链的增删改查api
  - [x] 重新规划了数据库设计，并完成了友链和标签的数据库设计
  - [x] 修复了一些之前代码bug
- 2023.04.25
  - [x] 实现了前端后台友链的增删改查
  - [x] 修复了一些bug
  - [x] 完善了前端前台主页展示，基本要素全部显示成功
  - [x] 完善了前端后台评论页面，实现了功能完备
- 2023.04.26
  - [x] 修复了一些现有bug
  - [x] 通过代码复用，完成了标签的增删改查api
  - [x] 实现了前端后台标签的增删改查
---
- 2023.05.09
  - [x] 完成了博客数据页面的编写
  - [x] 修复了一些bug
---
- 2023.05.10
  - [x] 完成了个人资料页面的展示代码编写
  - [x] 重构了用户模型，完善了一些api和service代码
# Todo
- [x] ~~采用第三方库validator来进行数据校验~~
- [ ] 更换为使用钩子函数来完成密码加盐的操作
- [x] ~~实现查看单个分类下所有文章~~
- [ ] 实现添加文章时可以使用多个标签
- [ ] 通过反射分离结构体属性
- [ ] 修改分类与文章的关联模式为many2many，用户与文章的关联模式为belong。
- [ ] 完善token错误处理
- [ ] 增加文件上传到oss
- [x] ~~完善权限控制机制,添加新增用户时权限的判断~~
- [x] ~~设计好各页面样式~~
- [ ] 完善跨域配置
- [x] ~~实现token判断~~
- [ ] 实现记住密码
- [ ] 实现响应式布局，适应移动端
- [x] ~~实现更好的分页方式~~
- [ ] 优化模糊查询代码
- [x] ~~当输入搜索的用户名后，将当前页重新设置为1~~
- [ ] 将ElementUI导入方式更改为按需导入 
- [ ] 完善编写文章页面的业务逻辑
- [ ] 实现显示文章目录
- [ ] 完成文章与标签的互动
- [ ] 完善鉴权，实现token过期之后的处理(回登录页面)

- 

# Issue

- [x] ~~解决result.Error != gorm.ErrRecordNotFound问题~~ :/`使用了多个err，并且将err全局声明`
- [x] ~~分页功能存在问题，需要解决关于无法定位当前页的问题~~ :/`原因：代码顺序写错了，把find写在了offset之前，Limit方法限制每页返回的记录数，Offset方法确定从哪个记录开始返回结果。`
- [ ] 查看分类下所有文章时，没有判断如果分类不存在时如何，如果分类不存在仍然会返回空切片。
- [x] ~~处理密码验证~~ :/`采用了随机salt，并将salt存入数据库，之后通过生成对比。`
- [x] ~~StandardClaims结构体中的NotBefore和ExpiresAt无法直接使用时间戳~~ :/`通过查看jwt-go/v4库官方文档，发现新版采用了jwt.NewTime函数，并且接受一个float64时间戳。`
- [x] ~~关于role处理存在一些问题，应当解决如何创建管理员用户，目前是强制将role修改为2;~~ :/`目前解决方法是多创建了一个需要token认证的接口，通过判断是否有正确的token来确定能否修改role，如果有正确的token，则可以随意设定role为管理员或普通用户，如果没有role只能为2（即普通用户，同不需要token认证的接口）`
- [x] ~~跨域中间件存在一些问题~~ :/`初步解决方案是允许所有跨域，之后再进行完善`
- [x] ~~路由存在一些问题，点击菜单项之后，它会在根目录下补全，而不是在admin下补全~~ :/~~初步解决方案是将admin目录设置为根目录，之后再进一步完善~~/`将跳转的时候设定为绝对路径跳转（包括在nav中的跳转）`
- [x] ~~点击菜单项目之后，刷新后会出现菜单active位置错误~~ :/~~解决方案为获取地址栏中`/`后的字符串，之后设定active~~/`更新方案为初始页面默认active数据博客页面，之后点击菜单项目后通过绝对路径进行active`
- [x] ~~分页功能存在问题，当点击其他页面时，会立刻又将currentPage变为第一页，从而无法换页~~ :/~~后端代码存在问题，.count位置不当，导致第二页时total出现问题~~/`最新解决方案：增加了model参数，先进行查询获得总数，之后再分页`
- [x] ~~进行搜索查询时，总数始终是全部用户总数~~ :/`和上面的一样解决方案，最开始采用的是直接model(&users),始终查询整个表，现在采用了上面的最新方案，先查询获得总数之后再分页，这样就变正常了`
- [ ] 当在非第一页刷新时，会返回第一页(因为初始页为1页，考虑是否进行修改)
- [ ] category页面的新增分类，在输入框中回车会导致页面跳转（且不会携带数据跳转到到想要的页面）
- [x] ~~新建文章或编辑文章时，会提示外键有问题，article的模型中是引用了uid和cid作为外键，但是会提示无法更新或修改（实际是成功了，但是response中仍然为400错误，且新增或编辑的response中category和user的数据为0）~~ :/`后端article的新增文章api返回的json写错了，导致固定返回400错误`
- [x] ~~在/write-article/:id页面直接点击nav中的编写文章，只去掉/:id，但是不刷新页面，也不清空数据~~ :/`是由于两个表单使用同一个ref，导致vue缓存，通过给显示内容的el-main部分的vue-view添加一个key（通过$route.path），来保证上一个路由的独一无二`
- [ ] 无法复用其他文件提供的内容，导致代码存在臃肿现象
- [ ] 当token过期后，虽然有的页面会提示token过期，但是不会跳转回登录页面
- [ ] 返回的data段包含的信息太多，甚至包含隐私信息
 


 


