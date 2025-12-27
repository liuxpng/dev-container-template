# DevContainer Templates

[English](README_EN.md) | 中文

一个开源的 Dev Container 模板集合，提供 PHP、Python、Go 等语言的即用型开发环境。

## 前置要求

- [Docker](https://www.docker.com/get-started)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers 扩展](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

---

## 快速开始

本项目默认使用**预构建镜像**，首次启动仅需 **30-60 秒**（拉取镜像），无需本地构建。

| 启动方式 | 首次耗时 | 适用场景 |
|---------|---------|---------|
| 预构建镜像（默认） | 30-60 秒 | 推荐，开箱即用 |
| 本地构建 | 5-10 分钟 | 需要定制 Dockerfile |

### 切换到本地构建

如需自定义 Dockerfile，编辑 `.devcontainer/devcontainer.json`：

```jsonc
{
  // 注释掉这行
  // "image": "ghcr.io/liuxpng/devcontainer-php:latest",

  // 取消注释以下内容
  "build": {
    "dockerfile": "Dockerfile",
    "context": "."
  }
}
```

---

## 使用场景

### 场景一：从零开始新项目

**适用于**：你刚创建了一个空目录，准备开始一个新项目。

```bash
# 1. 创建项目目录
mkdir my-project && cd my-project

# 2. 复制对应的 devcontainer 配置
#    Laravel 项目：
cp -r /path/to/devcontainer-templates/php-laravel/.devcontainer .

#    普通 PHP 项目：
cp -r /path/to/devcontainer-templates/php/.devcontainer .

#    Python 项目：
cp -r /path/to/devcontainer-templates/python/.devcontainer .

#    Go 项目：
cp -r /path/to/devcontainer-templates/go/.devcontainer .

# 3. 用 VS Code 打开
code .

# 4. VS Code 提示 "Reopen in Container"，点击确认
#    容器启动后，开发环境就绑了！
```

**Laravel 特别说明**：使用 `php-laravel` 模板时，容器启动后会自动执行 `composer create-project laravel/laravel`，无需手动安装。

---

### 场景二：为已有项目添加开发环境

**适用于**：你已经有一个项目，想让团队成员能一键配置开发环境。

```bash
# 1. 进入你的项目目录
cd /path/to/your-existing-project

# 2. 复制 devcontainer 配置到项目中
cp -r /path/to/devcontainer-templates/php/.devcontainer .

# 3. 根据项目需要修改配置（见下方说明）

# 4. 提交到版本控制
git add .devcontainer
git commit -m "Add devcontainer configuration"

# 5. 用 VS Code 打开并启动容器
code .
```

**常见修改**：

| 项目类型 | 需要修改的配置 |
|---------|---------------|
| Laravel | `nginx.conf` 中 `root` 改为 `/workspace/public` |
| Symfony | `nginx.conf` 中 `root` 改为 `/workspace/public` |
| 普通 PHP | `nginx.conf` 中 `root` 改为你的入口目录 |
| Python | `devcontainer.json` 中 `postCreateCommand` 改为 `pip install -r requirements.txt` |
| Go | `devcontainer.json` 中 `postCreateCommand` 改为 `go mod download` |

---

### 场景三：克隆包含 devcontainer 的项目

**适用于**：你克隆了一个已配置 devcontainer 的项目。

```bash
# 1. 克隆项目
git clone https://github.com/someone/some-project.git
cd some-project

# 2. 用 VS Code 打开
code .

# 3. VS Code 检测到 .devcontainer，提示 "Reopen in Container"
#    点击确认，等待容器构建完成即可
```

---

## 可用模板

| 模板 | 目录 | 适用场景 |
|------|------|---------|
| **PHP** | `php/` | 通用 PHP 项目 |
| **Laravel** | `php-laravel/` | Laravel 框架项目（自动初始化） |
| **Python** | `python/` | Python 项目 |
| **Go** | `go/` | Go 项目 |

---

## 模板详情

### PHP 8.3 + Nginx

```bash
cp -r devcontainer-templates/php/.devcontainer /your-project/
```

| 预装工具 | 说明 |
|---------|------|
| PHP 8.3 + FPM | 运行环境 |
| Nginx | Web 服务器 |
| Composer | 包管理器 |
| Xdebug | 调试器 |
| PHPUnit | 单元测试 |
| PHP CS Fixer | 代码格式化 |
| PHPStan | 静态分析 |

**端口**: 80 (HTTP), 9003 (Xdebug)

### Laravel (PHP 8.3 + Nginx + Node.js)

```bash
cp -r devcontainer-templates/php-laravel/.devcontainer /your-project/
```

| 预装工具 | 说明 |
|---------|------|
| PHP 8.3 + FPM | 运行环境 |
| Nginx | 已配置指向 public/ |
| Node.js 20 | Vite/Mix 前端构建 |
| Composer | 包管理器 |
| Laravel Installer | 快速创建项目 |
| Xdebug | 调试器 |

**端口**: 80 (HTTP), 9003 (Xdebug), 5173 (Vite)

**自动初始化**: 空目录启动时自动运行 `composer create-project laravel/laravel`

### Python 3.12

```bash
cp -r devcontainer-templates/python/.devcontainer /your-project/
```

| 预装工具 | 说明 |
|---------|------|
| Python 3.12 | 运行环境 |
| pip | 包管理器 |
| Poetry | 现代依赖管理 |
| uv | 高速包管理器 |
| pytest | 单元测试 |
| ruff | 快速 linter |
| black | 代码格式化 |
| mypy | 类型检查 |
| ipython | 增强 shell |

**端口**: 8000

### Go 1.23

```bash
cp -r devcontainer-templates/go/.devcontainer /your-project/
```

| 预装工具 | 说明 |
|---------|------|
| Go 1.23 | 运行环境 |
| golangci-lint | 代码检查 |
| delve | 调试器 |
| gopls | 语言服务器 |
| air | 热重载 |
| mockgen | Mock 生成 |
| swag | Swagger 文档 |

**端口**: 8080

---

## 自定义配置

### 修改语言版本

编辑 `.devcontainer/Dockerfile`：

```dockerfile
# PHP: 改为 8.2
FROM php:8.2-fpm

# Python: 改为 3.11
FROM python:3.11-slim

# Go: 改为 1.22
FROM golang:1.22
```

修改后执行：`Dev Containers: Rebuild Container`

### 修改 Nginx 根目录

编辑 `.devcontainer/nginx.conf`：

```nginx
server {
    # Laravel / Symfony
    root /workspace/public;

    # ThinkPHP
    root /workspace/public;

    # 普通 PHP
    root /workspace;
}
```

### 添加端口转发

编辑 `.devcontainer/devcontainer.json`：

```json
{
  "forwardPorts": [80, 3306, 6379]
}
```

### 添加启动命令

```json
{
  "postCreateCommand": "composer install && php artisan migrate",
  "postStartCommand": "sudo service nginx start"
}
```

---

## 常见问题

### 容器启动失败

1. 确保 Docker 正在运行
2. 尝试 `Dev Containers: Rebuild Container Without Cache`
3. 查看 Docker 日志排查问题

### 端口被占用

修改 `devcontainer.json` 中的 `forwardPorts`，或停止占用该端口的本地服务。

### 文件权限问题

```bash
sudo chown -R vscode:vscode /workspace
```

### 如何访问数据库？

容器内没有预装数据库。推荐方案：
1. 本地安装数据库，容器内通过 `host.docker.internal` 连接
2. 使用 Docker Compose 添加数据库服务

---

## 项目结构

```
devcontainer-templates/
├── .github/
│   └── workflows/          # CI/CD 自动构建镜像
├── php/                    # 通用 PHP
├── php-laravel/            # Laravel 专用
├── python/                 # Python
├── go/                     # Go
├── README.md
└── LICENSE
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

[MIT License](LICENSE)
