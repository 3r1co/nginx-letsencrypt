FROM python:2.7

LABEL org.3r1c.maintainer="Eric Muellenbach"

ADD https://dl.eff.org/certbot-auto /usr/local/bin/certbot

RUN chmod +x /usr/local/bin/certbot 

RUN certbot renew --non-interactive

ADD letsencrypt-helper /letsencrypt-helper
ADD le_wrapper.sh /le_wrapper.sh

RUN chmod +x /le_wrapper.sh

CMD ["/letsencrypt-helper"]
