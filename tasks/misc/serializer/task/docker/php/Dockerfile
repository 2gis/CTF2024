FROM php:7.4.3-fpm-alpine3.11

COPY ./docker/php/php-fpm.conf /usr/local/etc/php-fpm.d/www.conf

COPY . /app

VOLUME ["/app"]
