FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV GOROOT /usr/local/go
RUN apk update --no-cache && apk add --no-cache ca-certificates git

RUN git clone https://github.com/feihua/zero-admin ${GOROOT}/src/go-zero-admin

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -o /app/datacenter-server ./datacenter.go

RUN cd ./common && go build -o /app/common-server ./rpc/rpc.go && cd ..
RUN cd ./user && go build -o /app/user-server ./rpc/rpc.go && cd ..
RUN cd ./votes && go build -o /app/votes-server ./rpc/rpc.go && cd ..
RUN cd ./questions && go build -o /app/questions-server ./rpc/questions.go && cd ..
RUN cd ./search && go build -o /app/search-server ./rpc/search.go && cd ..

WORKDIR ${GOROOT}/src/go-zero-admin
RUN go build -o /app/oms-server ./rpc/oms/oms.go
RUN go build -o /app/pms-server ./rpc/pms/pms.go
RUN go build -o /app/ums-server ./rpc/ums/ums.go
RUN go build -o /app/sms-server ./rpc/sms/sms.go
RUN go build -o /app/sys-server ./rpc/sys/sys.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache bash ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app /app
COPY . .
RUN chmod +x entrypoint.sh

EXPOSE 8887

ENTRYPOINT [ "bash", "/app/entrypoint.sh" ]