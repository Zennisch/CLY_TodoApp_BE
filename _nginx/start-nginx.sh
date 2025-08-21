#!/bin/bash

echo "Starting Nginx reverse proxy..."

nginx -g 'daemon off;' &
NGINX_PID=$!

sleep 5

# Kiểm tra xem có cần lấy chứng chỉ SSL không
if [ ! -f "/etc/letsencrypt/live/cty-todo-app-be.duckdns.org/fullchain.pem" ] || [ ! -s "/etc/letsencrypt/live/cty-todo-app-be.duckdns.org/fullchain.pem" ]; then
    echo "Getting SSL certificate..."
    
    certbot certonly \
        --webroot \
        --webroot-path=/var/www/certbot \
        --email ${LETSENCRYPT_EMAIL:-admin@example.com} \
        --agree-tos \
        --no-eff-email \
        --non-interactive \
        -d cty-todo-app-be.duckdns.org
    
    if [ $? -eq 0 ]; then
        echo "SSL certificate obtained successfully. Reloading nginx..."
        nginx -s reload
    else
        echo "Failed to obtain SSL certificate. Continuing with self-signed cert..."
    fi
else
    echo "SSL certificate already exists"
fi

# Thiết lập cron job để gia hạn chứng chỉ
echo "0 12 * * * certbot renew --quiet && nginx -s reload" | crontab -

wait $NGINX_PID
