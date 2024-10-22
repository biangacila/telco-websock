

sudo docker build -t programfile .
sudo docker tag sha256:fbff41617ef0a5b315f8bb5238bade3e54416acba507c9a13a1daaa74ef39eb5 010309/telco-websocket:latest
sudo  docker push 010309/telco-websocket:latest

docker run -d -p 3319:8080 \
--name telco-websocket \
010309/telco-websocket:latest