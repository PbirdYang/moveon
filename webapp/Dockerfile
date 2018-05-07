FROM 100.125.0.198:20202/library/ubuntu:latest
MAINTAINER yangmozi
RUN echo "Asia/Shanghai" > /etc/timezone
RUN mkdir /home/ymz
WORKDIR /home/ymz
COPY play.sh .
ADD abc.tar.gz .
COPY app1 .
ENTRYPOINT ./app1
