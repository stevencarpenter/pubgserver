# pubgserver

A small webserver exposing one endpoint `/pubg/leaderboard` that takes one query parameter `accountId` and returns that player's stats from redis.

## Build
To build the container directly into minikube

```shell
eval $(minikube docker-env)
docker build -t pubgserver . 
```

## Usage
```shell
curl -X GET "localhost:8090/pubg/leaderboard?accountId=player1"
```

Response:
```json
{"player1":"{\"rank\":6,\"wins\":7,\"games\":10}"}
```
