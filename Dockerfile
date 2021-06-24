FROM  rodorg/rod
WORKDIR /app
COPY ./config .
COPY  auto-sign .
ENTRYPOINT ["/app/auto-sign"]