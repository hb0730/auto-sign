FROM alpine:3.10

ARG VERSION=0.0.2
ARG FILE_NAME=auto-sign

ENV AUTO_SIGN_URL=https://github.com/hb0730/auto-sign/releases/download/${VERSION}/${FILE_NAME}

WORKDIR /opt

COPY ./config ./config

RUN wget ${AUTO_SIGN_URL} && ls

ENTRYPOINT [ "./auto-sign" ]