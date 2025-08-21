FROM nginx:alpine

RUN apk add --no-cache certbot certbot-nginx openssl bash

RUN mkdir -p /var/www/certbot
RUN mkdir -p /etc/letsencrypt

COPY ssl-setup.sh /usr/local/bin/ssl-setup.sh
COPY start-nginx.sh /usr/local/bin/start-nginx.sh
RUN chmod +x /usr/local/bin/ssl-setup.sh
RUN chmod +x /usr/local/bin/start-nginx.sh

COPY nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /etc/letsencrypt/live/cty-todo-app-be.duckdns.org && \
    openssl req -x509 -nodes -days 1 -newkey rsa:2048 \
    -keyout /etc/letsencrypt/live/cty-todo-app-be.duckdns.org/privkey.pem \
    -out /etc/letsencrypt/live/cty-todo-app-be.duckdns.org/fullchain.pem \
    -subj "/C=US/ST=State/L=City/O=Organization/CN=cty-todo-app-be.duckdns.org"

EXPOSE 80 443

CMD ["/usr/local/bin/start-nginx.sh"]
