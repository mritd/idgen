# idgen 重构设计方案

## 概述

对 idgen 项目进行全面重构，包括：前端现代化、CLI 增强、技术栈升级。

**保留不变的核心逻辑：**
- chinaid 库的所有生成逻辑
- Cardbin（银行卡前缀）数据
- 姓名姓氏数据

---

## 1. 项目结构

```
idgen/
├── cmd/
│   ├── root.go              # 根命令 + 全局 flags (--count, --format)
│   ├── name.go
│   ├── idno.go
│   ├── mobile.go
│   ├── bank.go
│   ├── email.go
│   ├── addr.go
│   ├── all.go
│   ├── server.go            # --theme cyber|terminal (默认 cyber)
│   └── version.go
│
├── server/
│   ├── server.go            # 路由
│   ├── handler.go           # API 处理
│   └── static/
│       ├── index.html       # 单页面 + 主题切换
│       ├── style.css        # CSS 变量双主题
│       └── app.js           # 交互逻辑
│
├── utils/
│   └── formatter.go         # JSON/CSV/Table 输出格式化
│
├── main.go
├── go.mod                   # Go 1.25+
└── Taskfile.yml             # 构建任务
```

---

## 2. CLI 设计

### 全局 Flags

```bash
idgen [command] [flags]

Flags:
  -c, --count int       生成数量 (默认 1)
  -f, --format string   输出格式: table|json|csv (默认 table)
  -C, --copy            复制到剪贴板 (单条时默认开启, 批量时默认关闭)
  -h, --help            帮助信息
```

### 命令示例

```bash
# 单条生成（默认复制到剪贴板）
idgen              # 生成身份证号
idgen name         # 生成姓名
idgen all          # 生成全部信息

# 批量生成
idgen idno -c 10                  # 生成 10 个身份证号, table 格式
idgen all -c 5 -f json            # 生成 5 条完整信息, JSON 格式
idgen mobile -c 100 -f csv        # 生成 100 个手机号, CSV 格式

# 批量 + 复制
idgen name -c 10 -C               # 批量生成并复制到剪贴板
```

### 输出格式示例

**table (默认):**
```
+----+--------+-------------+--------------------+
| #  | Name   | Mobile      | IDNo               |
+----+--------+-------------+--------------------+
| 1  | 张三   | 13812345678 | 110101199001011234 |
| 2  | 李四   | 15987654321 | 320102198505052345 |
+----+--------+-------------+--------------------+
```

**json:**
```json
[{"name":"张三","mobile":"13812345678","idno":"110101199001011234"},...]
```

**csv:**
```csv
name,mobile,idno
张三,13812345678,110101199001011234
李四,15987654321,320102198505052345
```

### Server 命令

```bash
idgen server [flags]

Flags:
  -l, --listen string   监听地址 (默认 0.0.0.0)
  -p, --port int        监听端口 (默认 8080)
  -t, --theme string    默认主题: cyber|terminal (默认 cyber)
```

---

## 3. 前端页面设计

### 页面布局

```
┌─────────────────────────────────────────────────────────┐
│  [Logo/标题: IDGEN]                    [主题切换按钮]   │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  ┌─────────────────────────────────────────────────┐   │
│  │ 姓名      张三                            [复制] │   │
│  ├─────────────────────────────────────────────────┤   │
│  │ 身份证    110101199001011234              [复制] │   │
│  ├─────────────────────────────────────────────────┤   │
│  │ 手机号    13812345678                     [复制] │   │
│  ├─────────────────────────────────────────────────┤   │
│  │ 银行卡    6228480010116471234             [复制] │   │
│  ├─────────────────────────────────────────────────┤   │
│  │ 邮箱      zhangsan@example.com            [复制] │   │
│  ├─────────────────────────────────────────────────┤   │
│  │ 地址      浙江省杭州市西湖区...            [复制] │   │
│  └─────────────────────────────────────────────────┘   │
│                                                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐     │
│  │  重新生成   │  │  全部复制   │  │  批量: [10] │     │
│  └─────────────┘  └─────────────┘  └─────────────┘     │
│                                                         │
├─────────────────────────────────────────────────────────┤
│  [批量模式展开区域 - 表格展示 + CSV 导出按钮]           │
└─────────────────────────────────────────────────────────┘
```

### 赛博朋克主题 (cyber)

- 深色背景: `#0a0a0f` 渐变到 `#1a1a2e`
- 霓虹主色: 青色 `#00f5ff` + 品红 `#ff00ff` 渐变
- 发光效果: 文字和边框带 `box-shadow` / `text-shadow` 辉光
- 卡片: 半透明暗色 + 霓虹边框
- 按钮 hover: 脉冲发光动画

### 终端主题 (terminal)

- 纯黑背景: `#000000`
- 主色调: 经典绿 `#00ff00` 或琥珀 `#ffb000`
- 字体: 等宽字体 (JetBrains Mono / Fira Code / monospace)
- 效果: 扫描线叠加、光标闪烁、打字机动画
- 边框: ASCII 风格 `┌──┐`

### 交互功能

- 单行复制: 点击 [复制] 按钮, 显示 "已复制" 提示
- 全部复制: 复制所有字段, 格式化为多行文本
- 主题切换: 右上角按钮, 切换后存 `localStorage`
- 批量生成: 选择数量后展开表格区域
- CSV 导出: 批量模式下显示导出按钮

---

## 4. HTTP API 设计

### API 端点

```
+----------------------+--------+--------------------------------+
|  Path                |  Method|  Description                   |
+----------------------+--------+--------------------------------+
|  /                   |  GET   |  HTML 页面                     |
|  /api/v1/generate    |  GET   |  生成单条数据                  |
|  /api/v1/batch       |  GET   |  批量生成                      |
|  /api/v1/export      |  GET   |  CSV 导出                      |
+----------------------+--------+--------------------------------+
```

### 单条生成

```
GET /api/v1/generate
```

响应:
```json
{
  "name": "张三",
  "idno": "110101199001011234",
  "mobile": "13812345678",
  "bank": "6228480010116471234",
  "email": "zhangsan@example.com",
  "address": "浙江省杭州市西湖区文三路 123 号"
}
```

### 批量生成

```
GET /api/v1/batch?count=10
```

响应:
```json
{
  "count": 10,
  "data": [
    {"name": "张三", "idno": "...", ...},
    {"name": "李四", "idno": "...", ...}
  ]
}
```

### CSV 导出

```
GET /api/v1/export?count=10
Content-Type: text/csv
Content-Disposition: attachment; filename="idgen_export.csv"
```

响应:
```csv
name,idno,mobile,bank,email,address
张三,110101199001011234,13812345678,...
李四,320102198505052345,15987654321,...
```

---

## 5. 依赖和构建

### go.mod

```go
module github.com/mritd/idgen

go 1.25

require (
    github.com/mritd/chinaid    latest    // 核心生成逻辑
    github.com/mritd/logrus     latest    // 日志 (import _ 初始化)
    github.com/spf13/cobra      latest    // CLI 框架
    github.com/atotto/clipboard latest    // 剪贴板
)
```

所有依赖升级到最新版本。

### Taskfile.yml

```yaml
version: '3'

vars:
  VERSION:
    sh: git describe --tags --always 2>/dev/null || echo "dev"
  BUILD_COMMIT:
    sh: git rev-parse --short HEAD 2>/dev/null || echo "unknown"
  BUILD_DATE:
    sh: date "+%F %T"

tasks:
  default:
    cmds: [task: build]

  clean:
    desc: Clean build artifacts
    cmds: [rm -rf dist]

  build:
    desc: Build for current platform
    cmds:
      - |
        go build -trimpath -o dist/idgen -ldflags \
        "-w -s -X 'github.com/mritd/idgen/cmd.version={{.VERSION}}' \
        -X 'github.com/mritd/idgen/cmd.commit={{.BUILD_COMMIT}}' \
        -X 'github.com/mritd/idgen/cmd.buildDate={{.BUILD_DATE}}'"

  test:
    desc: Run tests
    cmds: [go test -v -race ./...]

  lint:
    desc: Run linter
    cmds: [golangci-lint run ./...]

  fmt:
    desc: Format code
    cmds:
      - gofmt -w -s .
      - goimports -w .

  release:
    desc: Cross-compile all platforms
    cmds:
      - task: clean
      - task: linux-386
      - task: linux-amd64
      - task: linux-amd64-v3
      - task: linux-amd64-v4
      - task: linux-armv5
      - task: linux-armv6
      - task: linux-armv7
      - task: linux-arm64
      - task: darwin-amd64
      - task: darwin-arm64
      - task: checksums

  checksums:
    desc: Generate checksums
    cmds:
      - cd dist && shasum -a 256 idgen-* > checksums.txt
```

### 版本信息

```go
// cmd/version.go
var (
    version   = "dev"
    commit    = "unknown"
    buildDate = "unknown"
)
```

---

## 6. 实施计划

1. 清理旧代码 (Makefile, .cross_compile.sh, server/html.go)
2. 创建新目录结构
3. 实现 utils/formatter.go
4. 重构 cmd/*.go 添加全局 flags
5. 重构 server/ 实现新 API
6. 实现前端页面 (HTML + CSS + JS)
7. 创建 Taskfile.yml
8. 升级 go.mod 依赖
9. 测试和调试
