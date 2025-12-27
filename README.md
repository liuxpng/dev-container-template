# DevContainer Templates

[English](README_EN.md) | 中文

一个开源的 Dev Container 模板集合，提供 PHP、Python、Go 三种语言的即用型开发环境。

## 特性

- **即开即用** - 克隆仓库，用 VS Code 打开，一键启动开发环境
- **工具齐全** - 预装常用开发工具，无需手动配置
- **独立环境** - 各语言环境相互独立，按需使用
- **可定制** - 轻松修改版本和配置

## 快速开始

### 前置要求

- [Docker](https://www.docker.com/get-started)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers 扩展](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

### 使用方法

```bash
# 1. 克隆仓库
git clone https://github.com/your-username/devcontainer-templates.git

# 2. 进入想要的语言目录
cd devcontainer-templates/python  # 或 php / go

# 3. 用 VS Code 打开
code .

# 4. VS Code 会提示 "Reopen in Container"，点击即可
```

## 可用环境

### PHP 8.3 + Nginx

| 工具 | 版本 | 说明 |
|------|------|------|
| PHP | 8.3 | 带 FPM |
| Nginx | latest | Web 服务器 |
| Composer | latest | 包管理器 |
| Xdebug | latest | 调试器 |
| PHPUnit | latest | 单元测试 |
| PHP CS Fixer | latest | 代码格式化 |
| PHPStan | latest | 静态分析 |

**端口**: 80 (HTTP), 9003 (Xdebug)

**示例**: 启动后访问 http://localhost

### Laravel (PHP 8.3 + Nginx)

专为 Laravel 框架优化的开发环境。

| 工具 | 版本 | 说明 |
|------|------|------|
| PHP | 8.3 | 带 FPM |
| Nginx | latest | 配置指向 public/ |
| Node.js | 20 | Vite/Mix 前端构建 |
| Composer | latest | 包管理器 |
| Laravel Installer | latest | Laravel 安装器 |
| Xdebug | latest | 调试器 |

**端口**: 80 (HTTP), 9003 (Xdebug), 5173 (Vite)

**特性**: 空目录自动初始化 Laravel 项目

**使用方法**:
```bash
# 创建空目录
mkdir my-laravel-app && cd my-laravel-app

# 复制配置
cp -r /path/to/devcontainer-templates/php-laravel/.devcontainer .

# 用 VS Code 打开，选择 "Reopen in Container"
code .

# 容器启动后自动创建 Laravel 项目
# 访问 http://localhost
```

### Python 3.12

| 工具 | 版本 | 说明 |
|------|------|------|
| Python | 3.12 | 最新稳定版 |
| pip | latest | 包管理器 |
| Poetry | 1.8.x | 现代依赖管理 |
| uv | latest | 高速包管理器 |
| pytest | latest | 单元测试 |
| ruff | latest | 快速 linter |
| black | latest | 代码格式化 |
| mypy | latest | 类型检查 |
| ipython | latest | 增强 shell |

**端口**: 8000

**示例**:
```bash
cd src && python main.py
# 访问 http://localhost:8000
```

### Go 1.23

| 工具 | 版本 | 说明 |
|------|------|------|
| Go | 1.23 | 最新稳定版 |
| golangci-lint | latest | 代码检查 |
| delve | latest | 调试器 |
| gopls | latest | 语言服务器 |
| air | latest | 热重载 |
| mockgen | latest | Mock 生成 |
| swag | latest | Swagger 文档 |

**端口**: 8080

**示例**:
```bash
cd src && go run main.go
# 访问 http://localhost:8080
```

## 通用工具

所有环境都预装以下工具：

| 工具 | 说明 |
|------|------|
| git | 版本控制 |
| curl / wget | HTTP 请求 |
| vim / nano | 文本编辑 |
| htop | 进程监控 |
| zip / unzip | 压缩解压 |
| jq | JSON 处理 |
| ssh | 远程连接 |
| make | 构建工具 |

## 自定义配置

### 修改语言版本

编辑对应目录下的 `.devcontainer/Dockerfile`：

```dockerfile
# PHP
FROM php:8.2-fpm  # 改为 8.2

# Python
FROM python:3.11-slim  # 改为 3.11

# Go
FROM golang:1.22  # 改为 1.22
```

修改后重新构建容器：`Dev Containers: Rebuild Container`

### 添加 VS Code 扩展

编辑 `.devcontainer/devcontainer.json`：

```json
{
  "customizations": {
    "vscode": {
      "extensions": [
        "existing.extension",
        "your.new-extension"
      ]
    }
  }
}
```

### 添加端口转发

```json
{
  "forwardPorts": [80, 3000, 5432]
}
```

## 项目结构

```
devcontainer-templates/
├── README.md
├── README_EN.md
├── LICENSE
├── .gitignore
├── scripts/
│   └── install-common-tools.sh
├── php/                         # 通用 PHP 环境
│   ├── .devcontainer/
│   │   ├── devcontainer.json
│   │   ├── Dockerfile
│   │   └── nginx.conf
│   └── src/
│       └── index.php
├── php-laravel/                 # Laravel 专用环境
│   └── .devcontainer/
│       ├── devcontainer.json
│       ├── Dockerfile
│       ├── nginx.conf
│       └── init-laravel.sh
├── python/
│   ├── .devcontainer/
│   │   ├── devcontainer.json
│   │   └── Dockerfile
│   └── src/
│       └── main.py
└── go/
    ├── .devcontainer/
    │   ├── devcontainer.json
    │   └── Dockerfile
    └── src/
        └── main.go
```

## 常见问题

### 容器启动失败

1. 确保 Docker 正在运行
2. 尝试 `Dev Containers: Rebuild Container Without Cache`
3. 检查 Docker 日志

### 端口冲突

修改 `devcontainer.json` 中的 `forwardPorts`，或停止占用该端口的服务。

### 文件权限问题

容器内使用 `vscode` 用户，如遇权限问题：
```bash
sudo chown -R vscode:vscode /workspace
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

[MIT License](LICENSE)
