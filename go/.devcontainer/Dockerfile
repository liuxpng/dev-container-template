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

# Install Go development tools
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/air-verse/air@latest \
    && go install go.uber.org/mock/mockgen@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/cweill/gotests/gotests@latest

# Create non-root user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Create go directories with proper permissions
RUN mkdir -p /go/pkg /go/bin /go/src \
    && chown -R $USERNAME:$USERNAME /go

# Configure PATH for user
RUN echo 'export PATH="$GOPATH/bin:$PATH"' >> /home/$USERNAME/.bashrc

# Set working directory
WORKDIR /workspace

EXPOSE 8080

CMD ["bash"]
