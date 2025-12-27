# DevContainer Templates

English | [中文](README.md)

An open-source collection of Dev Container templates providing ready-to-use development environments for PHP, Python, and Go.

## Features

- **Ready to Use** - Clone, open in VS Code, start developing immediately
- **Fully Equipped** - Pre-installed development tools, no manual configuration needed
- **Independent Environments** - Each language environment is isolated, use as needed
- **Customizable** - Easy to modify versions and configurations

## Quick Start

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

### Usage

```bash
# 1. Clone the repository
git clone https://github.com/your-username/devcontainer-templates.git

# 2. Navigate to desired language directory
cd devcontainer-templates/python  # or php / go

# 3. Open with VS Code
code .

# 4. VS Code will prompt "Reopen in Container", click to proceed
```

## Available Environments

### PHP 8.3 + Nginx

| Tool | Version | Description |
|------|---------|-------------|
| PHP | 8.3 | with FPM |
| Nginx | latest | Web server |
| Composer | latest | Package manager |
| Xdebug | latest | Debugger |
| PHPUnit | latest | Unit testing |
| PHP CS Fixer | latest | Code formatter |
| PHPStan | latest | Static analysis |

**Ports**: 80 (HTTP), 9003 (Xdebug)

**Example**: Visit http://localhost after starting

### Python 3.12

| Tool | Version | Description |
|------|---------|-------------|
| Python | 3.12 | Latest stable |
| pip | latest | Package manager |
| Poetry | 1.8.x | Modern dependency management |
| uv | latest | Fast package installer |
| pytest | latest | Unit testing |
| ruff | latest | Fast linter |
| black | latest | Code formatter |
| mypy | latest | Type checking |
| ipython | latest | Enhanced shell |

**Port**: 8000

**Example**:
```bash
cd src && python main.py
# Visit http://localhost:8000
```

### Go 1.23

| Tool | Version | Description |
|------|---------|-------------|
| Go | 1.23 | Latest stable |
| golangci-lint | latest | Code linting |
| delve | latest | Debugger |
| gopls | latest | Language server |
| air | latest | Hot reload |
| mockgen | latest | Mock generation |
| swag | latest | Swagger docs |

**Port**: 8080

**Example**:
```bash
cd src && go run main.go
# Visit http://localhost:8080
```

## Common Tools

All environments include these tools:

| Tool | Description |
|------|-------------|
| git | Version control |
| curl / wget | HTTP requests |
| vim / nano | Text editing |
| htop | Process monitoring |
| zip / unzip | Compression |
| jq | JSON processing |
| ssh | Remote connection |
| make | Build tool |

## Customization

### Change Language Version

Edit `.devcontainer/Dockerfile` in the corresponding directory:

```dockerfile
# PHP
FROM php:8.2-fpm  # Change to 8.2

# Python
FROM python:3.11-slim  # Change to 3.11

# Go
FROM golang:1.22  # Change to 1.22
```

After modification, rebuild the container: `Dev Containers: Rebuild Container`

### Add VS Code Extensions

Edit `.devcontainer/devcontainer.json`:

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

### Add Port Forwarding

```json
{
  "forwardPorts": [80, 3000, 5432]
}
```

## Project Structure

```
devcontainer-templates/
├── README.md
├── README_EN.md
├── LICENSE
├── .gitignore
├── scripts/
│   └── install-common-tools.sh
├── php/
│   ├── .devcontainer/
│   │   ├── devcontainer.json
│   │   ├── Dockerfile
│   │   └── nginx.conf
│   └── src/
│       └── index.php
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

## Troubleshooting

### Container Fails to Start

1. Ensure Docker is running
2. Try `Dev Containers: Rebuild Container Without Cache`
3. Check Docker logs

### Port Conflicts

Modify `forwardPorts` in `devcontainer.json`, or stop the service using that port.

### File Permission Issues

The container uses the `vscode` user. If you encounter permission issues:
```bash
sudo chown -R vscode:vscode /workspace
```

## Contributing

Issues and Pull Requests are welcome!

## License

[MIT License](LICENSE)
