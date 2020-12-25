# set base image (host OS)
ARG VARIANT="20.04"
FROM ubuntu:${VARIANT}

RUN apt update
RUN apt -y upgrade

RUN apt -y install git
RUN DEBIAN_FRONTEND="noninteractive" apt -y install golang