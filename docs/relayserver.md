# 概述

relayserver 是这个项目的核心。
我借助于 trezor 的开源项目 [trezord-go](https://github.com/trezor/trezord-go) 来实现主体的 转发服务器。
并使用 Fyne 库做GUI界面。

## 安装编译问题
由于 GUI 界面使用 Fyne，在windows平台下编译时，需要先安装 [mysys2](https://www.msys2.org/)，并使用 pacman 安装 gcc 和 make 工具。
同时，必须保证gcc在环境变量中。
否则无法编译通过
