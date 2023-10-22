# AndroidPhoneAsHardWallet
基于此项目将废弃的安卓手机变为硬件钱包

# 系统简介

系统结构如图所示：
![系统结构](./plantuml-img/docs/arch/系统总体架构.svg)

relayserver 是我们的开发部分；


# 使用方式

1. 编译 relayserver
2. 在桌面上启动 relayserver


# 安装编译注意事项

## relayserver 编译注意事项
由于 GUI 界面使用 Fyne，在windows平台下编译时，需要先安装 [mysys2](https://www.msys2.org/)，并使用 pacman 安装 gcc 和 make 工具。
同时，必须保证gcc在环境变量中。
否则无法编译通过

## 关于中文字体
因为界面中使用了中文字体， relayserver 启动前必须设置 `FYNE_FONT` 环境变量指向中文字体，否则无法显示中文。
例如：
```bash
# linux
export FYNE_FONT="/usr/share/fonts/msyh.ttc"

# windows
set FYNE_FONT="C:\Windows\Fonts\msyh.ttc"
```


