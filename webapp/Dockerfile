FROM swr.cn-north-1.myhuaweicloud.com/library/ubuntu:latest
MAINTAINER yangmozi
RUN echo "Asia/Shanghai" > /etc/timezone
RUN mkdir /home/ymz
COPY webapp /home/ymz
COPY start.sh /home/ymz
RUN chmod 777 /home/ymz/start.sh
WORKDIR /home/ymz
ENTRYPOINT ./start.sh
