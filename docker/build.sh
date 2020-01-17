CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o agent .

sudo docker-compose -f docker-compose.yml up

sudo docker-compose -f docker-compose.yml up agent

sudo docker-compose -f docker-compose.yml down