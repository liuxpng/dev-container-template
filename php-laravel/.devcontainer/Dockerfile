FROM php:8.3-fpm

LABEL maintainer="devcontainer-templates"
LABEL description="PHP 8.3 + Nginx + Laravel development environment"

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
    nginx \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libzip-dev \
    libicu-dev \
    libpq-dev \
    libonig-dev \
    libxml2-dev \
    libcurl4-openssl-dev \
    && rm -rf /var/lib/apt/lists/*

# Generate locale
RUN locale-gen en_US.UTF-8
ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8

# Configure and install PHP extensions
RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) \
    pdo_mysql \
    pdo_pgsql \
    mysqli \
    gd \
    zip \
    bcmath \
    opcache \
    intl \
    mbstring \
    exif \
    pcntl \
    xml \
    curl

# Install Redis extension
RUN pecl install redis && docker-php-ext-enable redis

# Install Xdebug
RUN pecl install xdebug && docker-php-ext-enable xdebug

# Configure Xdebug
RUN echo "xdebug.mode=debug,develop" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo "xdebug.start_with_request=yes" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo "xdebug.client_host=host.docker.internal" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo "xdebug.client_port=9003" >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Install Node.js (for Laravel Mix / Vite)
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

# Install global PHP tools
RUN composer global require \
    laravel/installer \
    phpunit/phpunit \
    friendsofphp/php-cs-fixer \
    phpstan/phpstan \
    && echo 'export PATH="$PATH:$HOME/.composer/vendor/bin"' >> /etc/bash.bashrc

# Create non-root user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Set working directory
WORKDIR /workspace

# Copy nginx config
COPY nginx.conf /etc/nginx/sites-available/default

EXPOSE 80 9003 5173

CMD service nginx start && php-fpm
