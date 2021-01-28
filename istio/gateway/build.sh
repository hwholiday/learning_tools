CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .
sudo docker image build -t hwholiday/gateway:v3.01 .
sudo docker push hwholiday/gateway:v3.01