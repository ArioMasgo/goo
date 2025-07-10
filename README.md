# dbu-api

Generación de compilado en windows para linux:
```bash
set GOOS=linux
set GOARCH=amd64
go build -o miapp
```

Ejemplo de config.json:
```json
{
  "app": {
    "service_name": "DBU-API",
    "port": 4021
  },
  "router": {
    "allowed_domains": "*",
    "logger_http": true
  },
  "log": {
    "path": "./log",
    "review_interval": 60
  },
  "key": {
    "private": "./app/private_key.pem",
    "public": "./app/public_key.pem"
  },
  "db": {
    "engine": "mysql",
    "server": "127.0.0.1",
    "port": 3306,
    "name": "name",
    "user": "user",
    "password": "password",
    "instance": "",
    "is_secure": false
  }
}
```