# GoWeb_Template_CLD

一款集成常用组件和工具的goWeb模板！！！

## 项目概要

- 1、本项目是基于Gin框架的GoWeb CLD风格的项目模板。
  - CLD=Controller+Logic/Service+Dao
  - CLD的项目代码结构在Java等项目中比较常见，所以该模板比较适合习惯写Java的同学使用。 当然，项目代码的结构安排与开发语言无关，这与个人、团队的习惯有关。

- 2、本模板集成常用的工具、组件。
  - 包括：MySQL、Redis、配置文件热加载、雪花算法、日志、鉴权等。
  - 文字说明：后续会组件针对项目中的组件、工具做更详细的文字说明，更方便大家熟练使用该模板。
  - 视频说明：如果有时间，我更想通过录制视频和大家介绍该项目。

- 3、欢迎使用和提出意见，感谢start哦！

## 项目结构说明

- router `路由层`

- controller `控制层`

- logic/service `业务逻辑层`

- dao `数据库操作`
    - mysql
    - redis
    - mongo

- models `模型层`

- conf `配置文件相关`

- common `通用/公共`

- docs `文档类`


## 程序启动
- 1.克隆或拉取最新代码
  
- 2.删除go.mod、go.sum文件（如果存在的话）
  
- 3.更改配置
  - 更改配置文件(conf/config.yaml)
  - 进入 main.go 配置正确的配置文件路径
  
- 4.更新依赖： 
  - go mod init GoWeb_Template_CLD (初始化go modules)
  - go mod tidy (更新依赖)
  
- 5.启动
  - go run main (执行程序)
  
- 6.点赞
 
## 项目集成的中间件介绍
有待补充
    
## 其他说明
本项目的代码看起来还不是那么干净、甚至有点乱，后续会逐步将代码整理更整洁！

## 关联项目
后续会推出比较简洁的CS模式，CS模式开发速度会更快，敬请关注。
