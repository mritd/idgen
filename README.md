## idgen

[![Build Status](https://travis-ci.org/mritd/idgen.svg?branch=master)](https://travis-ci.org/mritd/idgen)

> 一个使用 golang 编写的大陆身份证生成器，目前支持生成 姓名、身份证号、手机号、银行卡号、电子邮箱、地址信息

## 运行

**该工具目前支持两种运行方式:**

### 直接运行

直接命令行运行二进制文件即可生成对应信息，生成后将自动复制到系统剪切板

``` sh
➜  ~ idgen -h

该工具用于生成中国大陆 姓名 身份证号 银行卡号 手机号 地址 Email
生成后自动复制相应文本到系统剪切板，不使用子命令则默认生成身份证号

Usage:
  idgen [flags]
  idgen [command]

Available Commands:
  addr        生成地址
  all         生成所有信息
  bank        生成银行卡号
  email       生成 Email
  help        Help about any command
  idno        生成身份证号
  mobile      生成手机号
  name        生成姓名
  server      启动 http server

Flags:
      --config string   config file (default is $HOME/.idgen.yaml)
  -h, --help            help for idgen

Use "idgen [command] --help" for more information about a command.
```

### 服务器模式

使用 `server` 子命令将启动一个带有页面的 http 服务器用于浏览器访问，
同时还会启动一个 json api 接口用于其他程序调用

``` sh
➜  ~ idgen server -h

启动一个简单的 http server 用于提供页面访问以及 json 数据返回，
当不指定 -m 选项则同时开启 html 和 json 支持，访问地址如下:

http://BINDADDR:PORT/        返回一个简单的 html 页面
http://BINDADDR:PORT/api     返回 json 格式数据

Usage:
  idgen server [flags]

Flags:
  -h, --help            help for server
  -l, --listen string   http 监听地址 (default "0.0.0.0")
  -m, --mode string     server 运行模式(http/json)
  -p, --port int        http 监听端口 (default 8080)

Global Flags:
      --config string   config file (default is $HOME/.idgen.yaml)
```
