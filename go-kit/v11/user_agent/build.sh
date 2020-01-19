CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o user_agent .
