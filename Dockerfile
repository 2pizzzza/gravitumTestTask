FROM ubuntu:latest
LABEL authors="sosal"

ENTRYPOINT ["top", "-b"]