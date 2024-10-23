
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o programfile .
sudo docker build -t programfile .
sudo docker tag sha256:9ee435b8c31233f66c13ccddeaf8523b5c8dff0187aa0da0bb91263d71faf01b 010309/telco-websocket:latest
sudo  docker push 010309/telco-websocket:latest

docker run -d -p 3319:8080 \
--name telco-websocket \
010309/telco-websocket:latest


docker pull 010309/telco-websocket:latest
docker stop telco-websocket
docker rm telco-websocket

docker-compose up -d 

