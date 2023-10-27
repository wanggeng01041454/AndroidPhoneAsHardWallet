# 概述

relayserver 是这个项目的核心。
我借助于 trezor 的开源项目 [trezord-go](https://github.com/trezor/trezord-go) 来实现主体的 转发服务器。
并使用 Fyne 库做GUI界面。


# 架构

![relayserver模块架构](../plantuml-img/docs/arch/2.relayserver.arch.svg)

# 界面设计

![界面设计](images/relayserver的UI设计.png)


# trezor server 部分

引入 trezord-go 项目，作为trezor server的实现。
```bash
go get github.com/trezor/trezord-go
```

原计划引入 trezord-go， 并 Mock 其中部分功能，以实现应用模拟硬件钱包的功能。
但是，实践中发现 trezord-go 项目中 server/core/memorywriter 三个模块直接耦合在了一起，无法Mock其中部分功能后使用，因此最终决定将 trezord-go 的相关实现代码 copy 后修改。
__copy后修改的过程中，尽量保持原命名和原结构__



# relayserver 调试记录

## 2023-10-26
通过wireshark抓包，能够观察到 metamask 与 relayserver 之间的通信（确实连接到了 21325 端口，并发送了几个http请求）
但是，如果开启 metamask 之前，没有开启 relayserver 的话，之后就观察不到发送到连接 relayser 的连接请求了。
原因未知。
目前，在 logout 之后，再次登录，就能够正常连接到 relayserver 了。

### 初始开机的请求
1. POST /
2. POST /enumerate
3. POST /enumerate
