# User Management CLI


### Deploy and run locally
1. go mod tidy
2. go build

### Environtment Variable
API URL can be override
```
API_URL=https://api.test.com/1,https://api.test.com/2
```

### Save User data to csv command
`go run cmd/main.go save`

### Search User by tags
`go run main.go search --tag=[tags]`

### Run test
`go test github.com/fahrurben/users-management-cli/internal`