FROM alpine:3.10

ARG VERSION=0.0.2

ENV AUTO_SIGN_URL=https://github.com/hb0730/auto-sign/releases/download/${VERSION}/auto-sign
ENV FILE_NAME=app

COPY ./config /opt/config

RUN curl -s ${AUTO_SIGN_URL} -o ${FILE_NAME} && mv ${FILE_NAME} /opt

ENTRYPOINT [ "/opt/${FILE_NAME}" ]