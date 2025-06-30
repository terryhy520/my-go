# my-go

本项目为 Go 语言开发的应用，包含基础的服务结构和 GitHub Actions 自动化配置。

## 目录结构

```
go.mod
go.sum
makefile
myapp
myapp.exe
cmd/
    main.go
internal/
    config/
        config.go
    service/
        greeting.go
        greeting_test.go
.github/
    workflows/
        mcp-server.yml
```

## 快速开始

### 1. 克隆仓库

```sh
git clone https://github.com/terryhy520/my-go.git
cd my-go
```

### 2. 构建与运行

```sh
go build -o myapp ./cmd
./myapp
```

### 3. 运行测试

```sh
go test ./...
```

## MCP Server 自动化

本项目已配置 GitHub Actions workflow（`.github/workflows/mcp-server.yml`），每次 push 或 PR 到 main 分支时会自动运行 MCP Server：

- 安装 Python 环境
- 安装 `mcp_server_time` 依赖
- 运行 `python -m mcp_server_time --local-timezone=America/Los_Angeles`

## 贡献

欢迎提交 issue 和 PR。

## License

MIT
