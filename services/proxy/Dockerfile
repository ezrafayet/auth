FROM openresty/openresty:alpine AS proxy

COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf
COPY locations /usr/local/openresty/nginx/conf/locations
COPY lua /usr/local/openresty/site/lualib
