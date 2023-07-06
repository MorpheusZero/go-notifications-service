# go-notifications-service
A simple demo of a Notifications Microservice written in Golang.

## Running Locally
```shell
DISCORD_WEBHOOK_URL="1234" API_KEY="aRandomString123" go run main.go
```

This example uses only a Discord Provider, but you can add more providers.

## API Endpoints
### GET /health
```text
ok
```

### POST /v1/notify

BODY:
```json
{
  "notification_type": "discord",
  "data": "{\"content\":\"Hello World!\"}"
}
```

Response:
```text
ok
```