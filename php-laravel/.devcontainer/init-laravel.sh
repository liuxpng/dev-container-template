#!/bin/bash
set -e

# Increase composer process timeout (default is 300s)
export COMPOSER_PROCESS_TIMEOUT=900

echo "Initializing Laravel project..."

# Check if Laravel is already installed
if [ -f "artisan" ]; then
    echo "Laravel already installed, running composer install..."
    composer install --no-scripts
    php artisan package:discover --ansi || true

    # Generate key if .env exists but APP_KEY is empty
    if [ -f ".env" ] && grep -q "APP_KEY=$" .env; then
        php artisan key:generate
    fi
else
    echo "Creating new Laravel project..."

    # Create Laravel project in current directory
    composer create-project laravel/laravel tmp --prefer-dist --no-scripts

    # Move files from tmp to workspace root
    shopt -s dotglob
    mv tmp/* .
    rmdir tmp

    echo "Laravel project created successfully!"
    echo ""
    echo "Next steps:"
    echo "  1. Visit http://localhost to see your Laravel app"
    echo "  2. Run 'npm install && npm run dev' for Vite"
    echo "  3. Configure your database in .env"
fi

echo "Done!"
