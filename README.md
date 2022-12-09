# User Service

# Installations

```
go get github.com/activechoon/gin-gorm-filter
```

# Starting the Application Server

### Ubuntu
go run *.go

### Windows
choco install mingw
go run ./

### User Registration API:

```
curl --location --request POST 'http://localhost:3005/api/auth/register' \
--header 'Authorization;' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "longnguyen@gmail.com",
    "password": "123456",
    "name": "Long Nguyen",
    "phone": "+84583594331",
    "address": "168 Truong Van Bang, Thanh My Loi ward, Thu Duc City",
    "country": "Vietnam",
    "zipcode": "71350"
}'
```

<figure>
    <center>
        <img src="docs/images/Screenshot from 2022-12-09 15-38-48.png">
        <figcaption><i>Demo User Registration API</i></figcaption>
    </center>
</figure>

### User Login API:

```
curl --location --request POST 'http://localhost:3005/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "longnguyen@gmail.com",
    "password": "123456"
}'
```

<figure>
    <center>
        <img src="docs/images/Screenshot from 2022-12-09 15-37-27.png">
        <figcaption><i>Demo User Login API</i></figcaption>
    </center>
</figure>

### Users Search API:

```
curl --location --request GET 'http://localhost:3005/api/v1/users/search' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxvbmduZ3V5ZW5AZ21haWwuY29tIiwicGFzc3dvcmQiOiIkMmEkMTQkQ1Z0N2lPamVOZmhOZHFvNlZzdDFHTzZVWE5UWXBuRGRZdkNuazk3R1IyNUxmNkpmL1kzTjYiLCJleHAiOjE2NzA1Nzc1NDJ9.zXioVjnMZPXUitSqWrOZvYBVmMyPgn7IS2GcycnLjhA'
```

<figure>
    <center>
        <img src="docs/images/Screenshot from 2022-12-09 15-37-10.png">
        <figcaption><i>Demo Users Search API</i></figcaption>
    </center>
</figure>

### User Detail API:

```
curl --location --request GET 'http://localhost:3005/api/v1/users/detail/9' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxvbmduZ3V5ZW5AZ21haWwuY29tIiwicGFzc3dvcmQiOiIkMmEkMTQkQ1Z0N2lPamVOZmhOZHFvNlZzdDFHTzZVWE5UWXBuRGRZdkNuazk3R1IyNUxmNkpmL1kzTjYiLCJleHAiOjE2NzA1Nzc1NDJ9.zXioVjnMZPXUitSqWrOZvYBVmMyPgn7IS2GcycnLjhA'
```

<figure>
    <center>
        <img src="docs/images/Screenshot from 2022-12-09 15-37-19.png">
        <figcaption><i>Demo User Detail API</i></figcaption>
    </center>
</figure>