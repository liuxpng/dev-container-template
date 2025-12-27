FROM golang:1.23

LABEL maintainer="devcontainer-templates"
LABEL description="Go 1.23 development environment"

# Prevent interactive prompts
ENV DEBIAN_FRONTEND=noninteractive

# Install common development tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    curl \
    wget \
    vim \
    nano \
    htop \
    zip \
    unzip \
    jq \
    openssh-client \
    make \
    ca-certificates \
    gnupg \
    lsb-release \
    locales \
    sudo \
    zsh \
    && rm -rf /var/lib/apt/lists/*

# Generate locale
RUN locale-gen en_US.UTF-8
ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8

# Set Go environment variables
ENV GOPATH=/go
ENV PATH="${GOPATH}/bin:${PATH}"
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Install Go development tools (split for better error handling)
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/cosmtrek/air@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/cweill/gotests/gotests@latest

# Create non-root user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m -s /bin/zsh $USERNAME \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Create go directories with proper permissions
RUN mkdir -p /go/pkg /go/bin /go/src \
    && chown -R $USERNAME:$USERNAME /go

# Install Oh My Zsh for vscode user
RUN su - $USERNAME -c "curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh | sh -s -- --unattended" \
    && echo 'export PATH="$GOPATH/bin:$PATH"' >> /home/$USERNAME/.zshrc

# Set working directory
WORKDIR /workspace

EXPOSE 8080

CMD ["bash"]
