FROM mysql:latest

RUN mkdir -p /var/log/mysql/ \
&& chown -R mysql:root /var/log/mysql/ /var/lib/mysql/

COPY my.cnf /etc/mysql/conf.d/my.cnf

ENV LANG C.UTF-8