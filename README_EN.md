# DevContainer Templates

English | [中文](README.md)

An open-source collection of Dev Container templates providing ready-to-use development environments for PHP, Python, Go, and more.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

---

## Quick Start

This project uses **pre-built images** by default, requiring only **30-60 seconds** for first-time startup (pulling image), no local build needed.

| Startup Method | First-time Duration | Use Case |
|---------------|---------------------|----------|
| Pre-built image (default) | 30-60 seconds | Recommended, ready to use |
| Local build | 5-10 minutes | Need to customize Dockerfile |

### Switch to Local Build

If you need to customize the Dockerfile, edit `.devcontainer/devcontainer.json`:

```jsonc
{
  // Comment out this line
  // "image": "ghcr.io/liuxpng/devcontainer-php:latest",

  // Uncomment the following
  "build": {
    "dockerfile": "Dockerfile",
    "context": "."
  }
}
```

---

## Usage Scenarios

### Scenario 1: Starting a New Project from Scratch

**For**: You just created an empty directory and want to start a new project.

```bash
# 1. Create project directory
mkdir my-project && cd my-project

# 2. Copy the corresponding devcontainer configuration
#    For Laravel:
cp -r /path/to/devcontainer-templates/php-laravel/.devcontainer .

#    For regular PHP:
cp -r /path/to/devcontainer-templates/php/.devcontainer .

#    For Python:
cp -r /path/to/devcontainer-templates/python/.devcontainer .

#    For Go:
cp -r /path/to/devcontainer-templates/go/.devcontainer .

# 3. Open with VS Code
code .

# 4. VS Code will prompt "Reopen in Container", click to confirm
#    Once the container starts, your dev environment is ready!
```

**Laravel Note**: When using the `php-laravel` template, `composer create-project laravel/laravel` runs automatically after container startup.

---

### Scenario 2: Adding Dev Environment to an Existing Project

**For**: You have an existing project and want team members to set up the environment with one click.

```bash
# 1. Navigate to your project directory
cd /path/to/your-existing-project

# 2. Copy devcontainer configuration
cp -r /path/to/devcontainer-templates/php/.devcontainer .

# 3. Modify configuration as needed (see below)

# 4. Commit to version control
git add .devcontainer
git commit -m "Add devcontainer configuration"

# 5. Open with VS Code and start container
code .
```

**Common Modifications**:

| Project Type | Configuration to Modify |
|-------------|------------------------|
| Laravel | Change `root` in `nginx.conf` to `/workspace/public` |
| Symfony | Change `root` in `nginx.conf` to `/workspace/public` |
| Regular PHP | Change `root` in `nginx.conf` to your entry directory |
| Python | Change `postCreateCommand` in `devcontainer.json` to `pip install -r requirements.txt` |
| Go | Change `postCreateCommand` in `devcontainer.json` to `go mod download` |

---

### Scenario 3: Cloning a Project with DevContainer

**For**: You cloned a project that already has devcontainer configured.

```bash
# 1. Clone the project
git clone https://github.com/someone/some-project.git
cd some-project

# 2. Open with VS Code
code .

# 3. VS Code detects .devcontainer and prompts "Reopen in Container"
#    Click to confirm and wait for container build to complete
```

---

## Available Templates

| Template | Directory | Use Case |
|----------|-----------|----------|
| **PHP** | `php/` | General PHP projects |
| **Laravel** | `php-laravel/` | Laravel framework (auto-initializes) |
| **Python** | `python/` | Python projects |
| **Go** | `go/` | Go projects |

---

## Template Details

### PHP 8.3 + Nginx

```bash
cp -r devcontainer-templates/php/.devcontainer /your-project/
```

| Pre-installed Tools | Description |
|--------------------|-------------|
| PHP 8.3 + FPM | Runtime |
| Nginx | Web server |
| Composer | Package manager |
| Xdebug | Debugger |
| PHPUnit | Unit testing |
| PHP CS Fixer | Code formatter |
| PHPStan | Static analysis |

**Ports**: 80 (HTTP), 9003 (Xdebug)

### Laravel (PHP 8.3 + Nginx + Node.js)

```bash
cp -r devcontainer-templates/php-laravel/.devcontainer /your-project/
```

| Pre-installed Tools | Description |
|--------------------|-------------|
| PHP 8.3 + FPM | Runtime |
| Nginx | Configured for public/ |
| Node.js 20 | Vite/Mix frontend builds |
| Composer | Package manager |
| Laravel Installer | Quick project creation |
| Xdebug | Debugger |

**Ports**: 80 (HTTP), 9003 (Xdebug), 5173 (Vite)

**Auto-initialization**: Runs `composer create-project laravel/laravel` in empty directories

### Python 3.12

```bash
cp -r devcontainer-templates/python/.devcontainer /your-project/
```

| Pre-installed Tools | Description |
|--------------------|-------------|
| Python 3.12 | Runtime |
| pip | Package manager |
| Poetry | Modern dependency management |
| uv | Fast package installer |
| pytest | Unit testing |
| ruff | Fast linter |
| black | Code formatter |
| mypy | Type checking |
| ipython | Enhanced shell |

**Port**: 8000

### Go 1.23

```bash
cp -r devcontainer-templates/go/.devcontainer /your-project/
```

| Pre-installed Tools | Description |
|--------------------|-------------|
| Go 1.23 | Runtime |
| golangci-lint | Code linting |
| delve | Debugger |
| gopls | Language server |
| air | Hot reload |
| mockgen | Mock generation |
| swag | Swagger docs |

**Port**: 8080

---

## Customization

### Changing Language Version

Edit `.devcontainer/Dockerfile`:

```dockerfile
# PHP: Change to 8.2
FROM php:8.2-fpm

# Python: Change to 3.11
FROM python:3.11-slim

# Go: Change to 1.22
FROM golang:1.22
```

After modification: `Dev Containers: Rebuild Container`

### Changing Nginx Root Directory

Edit `.devcontainer/nginx.conf`:

```nginx
server {
    # Laravel / Symfony
    root /workspace/public;

    # ThinkPHP
    root /workspace/public;

    # Regular PHP
    root /workspace;
}
```

### Adding Port Forwarding

Edit `.devcontainer/devcontainer.json`:

```json
{
  "forwardPorts": [80, 3306, 6379]
}
```

### Adding Startup Commands

```json
{
  "postCreateCommand": "composer install && php artisan migrate",
  "postStartCommand": "sudo service nginx start"
}
```

---

## Troubleshooting

### Container Fails to Start

1. Ensure Docker is running
2. Try `Dev Containers: Rebuild Container Without Cache`
3. Check Docker logs for issues

### Port Already in Use

Modify `forwardPorts` in `devcontainer.json`, or stop the local service using that port.

### File Permission Issues

```bash
sudo chown -R vscode:vscode /workspace
```

### How to Access Database?

Containers don't include pre-installed databases. Recommended approaches:
1. Install database locally, connect from container via `host.docker.internal`
2. Use Docker Compose to add database services

---

## Project Structure

```
devcontainer-templates/
├── .github/
│   └── workflows/          # CI/CD auto-build images
├── php/                    # General PHP
├── php-laravel/            # Laravel specific
├── python/                 # Python
├── go/                     # Go
├── README.md
└── LICENSE
```

## Contributing

Issues and Pull Requests are welcome!

## License

[MIT License](LICENSE)
