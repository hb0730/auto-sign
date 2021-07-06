FROM rodorg/rod
WORKDIR /app
ENV FLAG ""
COPY ./config .
COPY auto-sign .
ENTRYPOINT /app/auto-sign ${FLAG}