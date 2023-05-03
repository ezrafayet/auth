FROM openresty/openresty:alpine AS proxy

COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf
COPY locations /usr/local/openresty/nginx/conf/locations
COPY headers /usr/local/openresty/nginx/conf/headers
COPY lualib /usr/local/openresty/site/lualib
