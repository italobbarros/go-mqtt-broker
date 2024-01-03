#!/bin/sh

# Inicie o broker MQTT em segundo plano
/main &

# Inicie o servidor Nginx
nginx -g 'daemon off;'
