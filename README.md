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
curl -X GET "localhost:8090/pubg/leaderboard?accountId=account.28b08053492a44659f8bf0517d8c3580"```

Response:
```json
{"account.28b08053492a44659f8bf0517d8c3580":"{\"rank\":164,\"wins\":49,\"games\":201}"}
```
