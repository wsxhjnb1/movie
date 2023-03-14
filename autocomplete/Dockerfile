FROM ubuntu:lunar-20230128

RUN apt update && apt upgrade -y
RUN apt install -y vim golang-go

VOLUME ["/autocomplete"]
VOLUME ["/elasticsearch"]
WORKDIR /autocomplete
