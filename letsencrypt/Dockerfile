FROM python:2.7
#FROM quay.io/letsencrypt/letsencrypt:latest

LABEL org.3r1c.maintainer="Eric Muellenbach"

ADD https://dl.eff.org/certbot-auto /usr/local/bin/certbot

RUN chmod +x /usr/local/bin/certbot && \
    certbot renew --non-interactive

ADD letsencrypt-helper /letsencrypt-helper

CMD ["/letsencrypt-helper"]