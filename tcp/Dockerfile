##构建程序CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o howie_tcp .
FROM scratch
ADD howie_tcp /var/howie_tcp
WORKDIR /var
ENTRYPOINT ["/var/howie_tcp"]
##docker build -t tcp:v1 .
