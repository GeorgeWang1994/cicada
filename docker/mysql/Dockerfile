FROM mysql:5.7

RUN mkdir -p /var/log/mysql/ \
&& chown -R mysql:root /var/log/mysql/ /var/lib/mysql/
COPY mysql.conf /etc/mysql/conf.d/my.cnf

ENV LANG C.UTF-8