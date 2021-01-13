FROM golang:1.15-alpine AS builder
WORKDIR /build
ENV GOPROXY https://goproxy.cn
ARG VERSION
ENV URL=https://github.com/hb0730/auto-sign/archive/${VERSION}.tar.gz
ADD ${URL} .
RUN tar -zxvf ${VERSION}.tar.gz && rm -f ${VERSION}.tar.gz
RUN cd auto-sign-${VERSION} && mv -f * .. && cd .. && rm -rf auto-sign-${VERSION}  && ls
RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o . && ls

FROM rodorg/rod AS final
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=builder /build/auto-sign /app/
COPY ./config /app/config
RUN apk --no-cache add tzdata zeromq \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

ENTRYPOINT ["/app/auto-sign"]