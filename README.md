## idgen

[![Build Status](https://travis-ci.org/mritd/idgen.svg?branch=master)](https://travis-ci.org/mritd/idgen)

> 一个使用 golang 编写的大陆身份证生成器，目前支持生成 姓名、身份证号、手机号、银行卡号、电子邮箱、地址信息
该工具部分代码从 [java-testdata-generator](https://github.com/binarywang/java-testdata-generator) 翻译而来，并添加了一些其他支持；
在此感谢原作者 [binarywang](https://github.com/binarywang)

## 安装

安装请直接从 release 页下载预编译的二进制文件，并放到 PATH 下即可；
docker 用户可以直接使用 `docker pull mritd/idgen` 拉取镜像；

- **自行编译**

确保已安装 go 1.14+ 和 [gox](https://github.com/mitchellh/gox)，然后执行 `make install` 即可

## 运行模式

**该工具目前支持两种运行方式:**

### 终端模式

直接命令行运行二进制文件即可生成对应信息，生成后将自动复制到系统剪切板

``` sh
➜  ~ idgen --help

This tool is used to generate Chinese name、ID number、bank card number、
mobile phone number、address and Email; automatically generate corresponding
text to the system clipboard after generation, and generate ID number by
default without sub-command

Usage:
  idgen [flags]
  idgen [command]

Available Commands:
  addr        Generate address information
  all         Generate all information
  bank        Generate bank card number
  email       Generate email address
  help        Help about any command
  idno        Generate ID number
  mobile      Generate mobile phone number
  name        Generate name
  server      Run as http server
  version     Print version

Flags:
  -h, --help      help for idgen
  -v, --version   Print version

Use "idgen [command] --help" for more information about a command.
```

### 服务器模式

使用 `server` 子命令将启动一个带有页面的 http 服务器用于浏览器访问，
同时还会启动一个 json api 接口用于其他程序调用

``` sh
➜  ~ idgen server --help

Run a simple http server to provide page access and json data return.
When the -m option is not specified, both html and json support are enabled,
and the access address is as follows:

http://BINDADDR:PORT/        return a simple html page
http://BINDADDR:PORT/api     return json format data

Usage:
  idgen server [flags]

Flags:
  -h, --help            help for server
  -l, --listen string   http listen address (default "0.0.0.0")
  -m, --mode string     server mode(html/json)
  -p, --port int        http listen port (default 8080)

Global Flags:
  -v, --version   Print version
```

docker 用户直接运行 `docker run -d -p 8080:8080 mritd/idgen` 即可
