# EP.GO.ESVR.LIB
## Summary | 框架概览
![avatar](https://raw.githubusercontent.com/hsu2017/EP.GO.ESVR.LIB/develop/res/esvr-structure.jpg)
* Author: Wells Hsu
* Email: wellshsu@outlook.com
* **内部框架，有限开放，仅供学习，不可商用**
* **For internal developer only**

## Feature | 功能特性
* 分布式、高承载、无状态的微服务框架
* 前后端服务分离(private/public)，有效抵御攻击
* 富前端连接，支持socket/ws/wss/http/https等
* 数据模型驱动，orm超集，支持条件表达式，快速开发
* 基于goroutine绑定的session会话结构，沙箱运行，安全稳定
* 动态负载均衡，基于UID均衡分配处理器，提高承载能力

## Usage | 安装使用
* **WIKI:https://github.com/hsu2017/EP.GO.ESVR.LIB/wiki**
* **DOCS:https://pkg.go.dev/github.com/hsu2017/EP.GO.ESVR.LIB**
* ecode-go:https://marketplace.visualstudio.com/items?itemName=wellshsu.ecode-go
* go1.18.2:https://github.com/hsu2017/ET.GO.SRC
* go get github.com/hsu2017/EP.GO.ESVR.LIB@latest
* Example: see test case.

## TODO | 待办事项
* Cache层需改造成Global缓存
* 多语言TypeScript/Lua/Java/C++/C#等横向拓展
* 重构Consul的注册模式
* xsession.GFinish性能优化

## Refer | 引用说明