FROM mysql:8

RUN echo "##### command install ##########"
RUN echo RUN apt update -qq && apt install -y vim
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]
COPY my.cnf /etc/mysql/my.cnf
RUN chmod 644 /etc/mysql/my.cnf