FROM ubuntu:20.04

RUN apt update && apt upgrade

CMD tail -f > /dev/null

