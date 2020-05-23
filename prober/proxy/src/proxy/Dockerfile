FROM jwilder/nginx-proxy:alpine

RUN apk add vim
COPY network_internal.conf /etc/nginx/network_internal.conf
COPY nginx.tmpl /app/nginx.tmpl
COPY nginx.conf /etc/nginx/nginx.conf
