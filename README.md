# User Service

# Installations

```
go get github.com/activechoon/gin-gorm-filter
```

# Starting the Application Server

```
go run *.go
```

```
curl --location --request POST 'http://localhost:3005/api/token/' \
--header 'Authorization: Bearer dgshhj' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "longnguyen@gmail.com",
    "password": "123456"
}'
```

```
curl --location --request GET 'http://localhost:3005/api/v1/user/' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxvbmduZ3V5ZW5AZ21haWwuY29tIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2NjQxMDkxMzl9.aO7kl5TKWdwwm-Q8ujsaJc7EjuNChdWGTLnjWTppLEk'
```