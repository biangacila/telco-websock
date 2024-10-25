
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o programfile .
sudo docker build -t programfile .
sudo docker tag sha256:e2c912676caaf228903bd9d0e5db5ca61b49c9b13069064eb5fde4ac38626201 010309/telco-websocket:latest
sudo  docker push 010309/telco-websocket:latest

docker run -d -p 3319:8080 \
--name telco-websocket \
010309/telco-websocket:latest


docker pull 010309/telco-websocket:latest
docker stop telco-websocket
docker rm telco-websocket

docker-compose up -d 

