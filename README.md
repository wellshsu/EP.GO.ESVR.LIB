# EP.GO.ESVR.LIB
## SUMMARY | 框架概览
![avatar](https://raw.githubusercontent.com/hsu2017/EP.GO.ESVR.LIB/develop/res/esvr-structure.jpg)
* Author: Wells Hsu
* Email: wellshsu@outlook.com
* **内部框架，仅供学习(For internal developer only)**

## FEATURE | 功能特性
* 分布式、高承载、无状态的微服务框架
* 前后端服务分离(private/public)，有效抵御攻击
* 富前端连接，支持socket/ws/wss/http/https等
* 数据模型驱动，orm超集，支持条件表达式，快速开发
* 基于goroutine绑定的session会话结构，内存沙箱，安全稳定
* 动态负载均衡，基于UID均衡分配处理器，提高承载能力

## USAGE | 安装使用
* **[wiki](https://github.com/hsu2017/EP.GO.ESVR.LIB/wiki): https://github.com/hsu2017/EP.GO.ESVR.LIB/wiki**
* **[godoc](https://pkg.go.dev/github.com/hsu2017/EP.GO.ESVR.LIB): https://pkg.go.dev/github.com/hsu2017/EP.GO.ESVR.LIB**
* [ecode-go](https://marketplace.visualstudio.com/items?itemName=wellshsu.ecode-go): https://marketplace.visualstudio.com/items?itemName=wellshsu.ecode-go
* [go1.18.2](https://github.com/hsu2017/ET.GO.SRC): https://github.com/hsu2017/ET.GO.SRC
* installation: go get github.com/hsu2017/EP.GO.ESVR.LIB@latest
* example: see test case.

## TODO | 待办事项
* Cache层需改造成Global缓存
* 多语言(TypeScript/Lua/Java/C++/C#)横向拓展
* 重构Consul的注册模式，先注册，再启动
* xsession.GFinish性能优化
* xcollect、xorm等泛型重构
* docker部署流程待完善
* log和pipe线程在极端情况下会满负荷运行，待优化

## REFER | 引用说明