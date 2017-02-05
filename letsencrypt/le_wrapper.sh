#!/bin/sh
certbot certonly --verbose --debug --webroot -w "$LE_WWW" --agree-tos --non-interactive --expand -m "$1" -d "$2"