# User Service

## Description

The User Service is a robust and versatile Golang-based application meticulously crafted using the Gin framework, designed to empower developers with a comprehensive set of user management and authentication features. This service offers a seamless user experience by providing the following key functionalities:

- User Registration: Users can effortlessly create their accounts by providing essential information, including email and a secure password.

- Email and Phone Number Verification: To ensure security and trust, the User Service offers both email and phone number verification mechanisms. Users receive verification codes via email and SMS, respectively, to confirm their identity.

- User Login and Logout: With a user-friendly login process, users can access their accounts securely using their registered email and password. Additionally, the service facilitates convenient logout to ensure data privacy.

- Password Reset: Users can regain access to their accounts if they forget their passwords by initiating a password reset process. A password reset email is sent to the user's registered email address.

- User Profile Management: The User Service enables users to maintain and personalize their profiles, including essential details such as first name, last name, profile picture, and more.

- User Search and Listing: Users can search for other users based on various criteria and view user listings with optional pagination and filtering options.

- Sending Notifications: Authenticated users can send notifications to others through multiple communication channels, such as email, SMS, and push notifications, making it easier to stay connected.

- Bearer Token-based Authentication: The service utilizes Bearer tokens to authenticate users, ensuring secure and authorized access to protected resources. Users receive a Bearer token upon successful login, which they include in subsequent API requests for authorization.

- Phone Number Validation (Firebase-based): Leveraging Firebase, the User Service offers phone number validation through SMS. Users can request and verify their phone numbers by entering received verification codes.

This comprehensive suite of features caters to a wide range of user management and authentication needs, making the User Service a reliable and indispensable component for building secure and user-centric applications.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Credits](#credits)
- [License](#license)
- [Documentation](#documentation)

## Installation

Before launching the User Service, make sure you have Go (Golang) installed on your system.

1. Clone this repository:

   ```sh
   git clone https://github.com/nguyentruonglong/user-service.git
    ```

2. Change to the project directory:

   ```sh
    cd user-service
   ```

3. Install the project dependencies using Go modules:

   ```sh
    go mod tidy
   ```

## Usage

### Launching

Windows

To run the User Service on Windows, follow these steps:

1. Open a command prompt.

2. Navigate to the project directory:

    ```
    cd user-service
    ```

3. Build and run the application:

    ```
    go run main.go
    ```

Ubuntu

To run the User Service on Ubuntu or a Linux-based system, follow these steps:

1. Open a terminal.

2. Navigate to the project directory:

    ```
    $ cd user-service
    ```

3. Build the application:

    ```
    $ go build -o user-service main.go
    ```

4. Run the application:

    ```
    $ ./user-service
    ```

### Functions

The User Service provides the following functions:

- User registration and management.
- Email and phone number verification.
- User login and logout.
- Password reset.
- User profile management.
- User search and listing.
- Sending notifications to users.
- Bearer token-based authentication.
- Phone number validation using Firebase.

### APIs

User Management APIs

#### 1. User Registration

- Endpoint: /api/v1/register (POST)

- Description: Allows users to register by providing their information.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/register -d '{"email": "user@example.com", "password": "secure_password"}'
```

#### 2. Email Verification

- Endpoint: /api/v1/verify-email (POST)

- Description: Allows users to verify their email addresses using a verification code sent via email.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/verify-email -d '{"email": "user@example.com", "verification_code": "123456"}'
```

#### 3. User Login

- Endpoint: /api/v1/login (POST)

- Description: Allows users to log in and obtain a Bearer token by providing their email and password.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/login -d '{"email": "user@example.com", "password": "secure_password"}'
```

#### 4. User Logout

- Endpoint: /api/v1/logout (POST)

- Description: Allows users to log out, effectively invalidating their Bearer token.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/logout -H "Authorization: Bearer <your-access-token>"
```

#### 5. Password Reset

- Endpoint: /api/v1/reset-password (POST)

- Description: Allows users to reset their password by sending a password reset email.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/reset-password -d '{"email": "user@example.com"}'
```

#### 6. User Profile Management

- Endpoint: /api/v1/profile (GET, PUT)

- Description: Allows users to retrieve their own profile information and update it, including their profile picture.

- Sample cURL Requests:

##### Retrieve Profile:

```curl
curl -X GET http://localhost:8080/api/v1/profile -H "Authorization: Bearer <your-access-token>"
```

##### Update Profile:

```curl
curl -X PUT http://localhost:8080/api/v1/profile -d '{"first_name": "John", "last_name": "Doe"}' -H "Authorization: Bearer <your-access-token>"
```

### User Search and Listing APIs

#### 7. User Search

- Endpoint: /api/v1/user-search (GET)

- Description: Allows users to search for other users based on specified criteria.

- Sample cURL Request:

```curl
curl -X GET 'http://localhost:8080/api/v1/user-search?query=John&country=US'
```

#### 8. User Listing

- Endpoint: /api/v1/user-list (GET)

- Description: Lists users with optional pagination and filtering options.

- Sample cURL Request:

```curl
curl -X GET 'http://localhost:8080/api/v1/user-list?page=1&per_page=10'
```

### Account Verification API

#### 9. Account Verification

- Endpoint: /api/v1/verify-account/:verification_code (POST)

- Description: Allows users to verify their accounts by providing a verification code received via email.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/verify-account/123456
```

### User Notification APIs

#### 10. Send Notification

- Endpoint: /api/v1/send-notification (POST)

- Description: Allows authorized users to send notifications to others via various channels (email, SMS, push, etc.).

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/send-notification -d '{"recipient": "user@example.com", "message": "Hello, user!"}' -H "Authorization: Bearer <your-access-token>"
```

### Bearer Authentication APIs

#### 11. User Authentication

- Endpoint: /api/v1/authenticate (POST)

- Description: Allows users to authenticate by providing their credentials (email and password) and receive a Bearer token.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/authenticate -d '{"email": "user@example.com", "password": "secure_password"}'
```

#### 12. Access Token Validation

- Endpoint: /api/v1/validate-token (GET)

- Description: Allows other services to validate a Bearer token.

- Sample cURL Request:

```curl
curl -X GET http://localhost:8080/api/v1/validate-token -H "Authorization: Bearer <your-access-token>"
```

#### 13. Token Refresh

- Endpoint: /api/v1/refresh-token (POST)

- Description: Allows users to refresh their Bearer token using a refresh token.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/refresh-token -d '{"refresh_token": "<your-refresh-token>"}'
```

#### 14. Token Revocation

- Endpoint: /api/v1/revoke-token (POST)

- Description: Allows users or services to revoke a Bearer token.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/revoke-token -d '{"token": "<token-to-revoke>"}' -H "Authorization: Bearer <your-access-token>"
```

### Phone Number Validation APIs (Firebase-based)

#### 15. Send Verification Code via SMS
- Endpoint: /api/v1/send-verification-sms (POST)

- Description: Allows users to request a verification code to be sent to their phone number via SMS.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/send-verification-sms -d '{"phone_number": "+1234567890"}' -H "Authorization: Bearer <your-access-token>"
```

#### 16. Verify Phone Number with Code

- Endpoint: /api/v1/verify-phone (POST)

- Description: Allows users to verify their phone number by providing the received verification code.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/verify-phone -d '{"phone_number": "+1234567890", "verification_code": "123456"}' -H "Authorization: Bearer <your-access-token>"
```

## Contributing

If you wish to contribute to this project, please follow the guidelines outlined in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## Credits

- Author: Nguyen Truong Long

## License

This project is licensed under a closed-source license agreement. See the [LICENSE](LICENSE) file for more details. For inquiries, contact Nguyen Truong Long.

## Documentation

- [Golang Documentation](https://golang.org/doc/)
- [Gin Framework Documentation](https://pkg.go.dev/github.com/gin-gonic/gin)
- [GORM Documentation](https://gorm.io/docs/)
- [Firebase Documentation](https://firebase.google.com/docs)
