#!/bin/bash
set -e

echo "Installing common development tools..."

# Update package list and install common tools
apt-get update && apt-get install -y --no-install-recommends \
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
    sudo

# Clean up
rm -rf /var/lib/apt/lists/*

# Generate locale
locale-gen en_US.UTF-8

# Create non-root user if not exists
USERNAME=${1:-vscode}
if ! id "$USERNAME" &>/dev/null; then
    echo "Creating user: $USERNAME"
    useradd -m -s /bin/bash "$USERNAME"
    echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME
    chmod 0440 /etc/sudoers.d/$USERNAME
fi

echo "Common tools installation completed!"
