## idgen

[![Build Status](https://travis-ci.org/mritd/idgen.svg?branch=master)](https://travis-ci.org/mritd/idgen)

> 一个使用 golang 编写的大陆身份证生成器，目前支持生成 姓名、身份证号、手机号、银行卡号、电子邮箱、地址信息
该工具部分代码从 [java-testdata-generator](https://github.com/binarywang/java-testdata-generator) 翻译而来，在此感谢原作者

## 安装

安装请直接从 release 页下载预编译的二进制文件，并放到 PATH 下即可；
docker 用户可以直接使用 `docker pull mritd/idgen` 拉取镜像；

- **自行编译**

注意: 由于本项目采用了 [SQLite3](https://github.com/mattn/go-sqlite3)，所以编译需要开启
cgo 支持(`CGO_ENABLE=1`),否则将会导致在使用 `idgen name` 命令同时存在 `~/.idgen/data.db`
文件时出现 SQLite3 初始化错误

## 运行模式

**该工具目前支持两种运行方式:**

### 终端模式

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
  init        初始化配置
  mobile      生成手机号
  name        生成姓名
  server      启动 http server
  version     显示当前版本

Flags:
      --config string   config file (default is $HOME/.idgen/idgen.yaml)
  -h, --help            help for idgen
  -v, --version         显示当前版本

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
  -m, --mode string     server 运行模式(html/json)
  -p, --port int        http 监听端口 (default 8080)

Global Flags:
      --config string   config file (default is $HOME/.idgen/idgen.yaml)
  -v, --version         显示当前版本
```

docker 用户直接运行 `docker run -d -p 8080:8080 mritd/idgen` 即可

### 关于姓名生成

v0.0.3 版本增加 sqlite 数据库支持，目前该数据库中存放了大量抓取的名字信息，工具在生成姓名
时将**自动检测 `~/.idgen/data.db` 是否存在，如果存在将会自动从中进行随机姓名生成**；此种
方式生成的姓名更加真实，但是由于数据量比较大(db文件 11MB)，**所以默认该数据库不包含在 release**
文件中，需要使用 `idgen init` 命令初始化(自动从 github 下载)