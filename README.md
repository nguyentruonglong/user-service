# User Service

## Description

The User Service is a robust and versatile Golang-based application meticulously crafted using the Gin framework, designed to empower developers with a comprehensive set of user management and authentication features. This service offers a seamless user experience by providing the following key functionalities:

- User Registration: Users can effortlessly create their accounts by providing essential information, including email and a secure password.

- Email and Phone Number Verification: To ensure security and trust, the User Service offers both email and phone number verification mechanisms. Users receive verification codes via email and SMS, respectively, to confirm their identity.

- Phone Number Validation (Twilio-based): Leveraging Twilio, the User Service now offers phone number validation through SMS. Users can request and verify their phone numbers by entering received verification codes.

- User Login and Logout: With a user-friendly login process, users can access their accounts securely using their registered email and password. Additionally, the service facilitates convenient logout to ensure data privacy.

- Password Reset: Users can regain access to their accounts if they forget their passwords by initiating a password reset process. A password reset email is sent to the user's registered email address.

- User Profile Management: The User Service enables users to maintain and personalize their profiles, including essential details such as first name, last name, profile picture, and more.

- User Search and Listing: Users can search for other users based on various criteria and view user listings with optional pagination and filtering options.

- Sending Notifications: Authenticated users can send notifications to others through multiple communication channels, such as email, SMS, and push notifications, making it easier to stay connected.

This comprehensive suite of features caters to a wide range of user management and authentication needs, making the User Service a reliable and indispensable component for building secure and user-centric applications.

### Project Directory Structure

The User Service project is organized with a clear directory structure to promote code modularity, maintainability, and separation of concerns. Here's an overview of the project directory structure:

```
    user-service/
    |-- api/
    |   |-- v1/
    |   |   |-- controllers/
    |   |   |   |-- search_controller.go              # Search Controller
    |   |   |   |-- user_login_controller.go          # User Login Controller
    |   |   |   |-- user_logout_controller.go         # User Logout Controller
    |   |   |   |-- user_register_controller.go       # User Register Controller
    |   |   |   |-- user_verification_controller.go   # User Verification Controller
    |   |   |   |-- notification_controller.go        # Notification Controller
    |   |   |
    |   |   |-- routes/
    |   |   |   |-- user_routes.go                    # User Routes
    |   |   |   |-- search_routes.go                  # Search Routes
    |   |   |   |-- verification_routes.go            # Verification Routes
    |   |   |   |-- notification_routes.go            # Notification Routes
    |   |   |
    |   |   |-- validators/
    |   |   |   |-- user_validator.go                 # User Validator
    |   |   |   |-- search_validator.go               # Search Validator
    |   |   |   |-- verification_validator.go         # Verification Validator
    |   |   |   |-- notification_validator.go         # Notification Validator
    |   |
    |   |-- middlewares/
    |   |   |-- auth_middleware.go                    # Authentication Middleware
    |   |
    |   |-- models/
    |   |   |-- user.go                               # User Model
    |   |
    |   |-- services/
    |   |   |-- user_service.go                       # User Service
    |   |   |-- search_service.go                     # Search Service
    |   |   |-- verification_service.go               # Verification Service
    |   |   |-- notification_service.go               # Notification Service
    |
    |-- config/
    |   |-- dev_config.yaml                           # Development Configuration
    |   |-- prod_config.yaml                          # Production Configuration
    |   |-- config.go                                 # Configuration Module
    |
    |---database/
    |   |-- database.go                               # Database Connection
    |   |   
    |   \-- migrations/
    |
    |---docs
    |   |-- docs.go                                   # Documentation Module
    |   |-- swagger.json                              # Swagger JSON Configuration
    |   |-- swagger.yaml                              # Swagger YAML Configuration
    |   |   
    |   \-- images/                                   # Image Resources for Documentation
    |   \-- swagger/                                  # Swagger UI Assets
    |
    |-- firebase/
    |   |-- firebase.go                               # Firebase Configuration
    |   |-- auth_service.go                           # Firebase Authentication Service
    |   |-- database_service.go                       # Firebase Realtime Database Service
    |
    |-- email/
    |   |-- email_service.go                          # Email Service
    |   |-- verification_email.go                     # Email Verification Service
    |   |-- password_reset_email.go                   # Password Reset Email Service
    |
    |-- sms/
    |   |-- sms_service.go                            # SMS Service
    |   |-- verification_sms.go                       # SMS Verification Service
    |
    |-- utils/
    |   |-- helpers.go                                # Utility Functions
    |
    |-- tests/
    |   |-- user_login_controller_test.go             # Unit Tests for Login Controller
    |   |-- user_logout_controller_test.go            # Unit Tests for Logout Controller
    |   |-- user_register_controller_test.go          # Unit Tests for Register Controller  
    |   |-- search_controller_test.go                 # Unit Tests for Search Controller
    |   |-- verification_controller_test.go           # Unit Tests for Verification Controller
    |   |-- notification_controller_test.go           # Unit Tests for Notification Controller
    |   |-- user_service_test.go                      # Unit Tests for User Service
    |   |-- search_service_test.go                    # Unit Tests for Search Service
    |   |-- verification_service_test.go              # Unit Tests for Verification Service
    |   |-- notification_service_test.go              # Unit Tests for Notification Service
    |
    |-- main.go                                       # Main Application Entry Point
    |-- go.mod                                        # Go Module File
    |-- go.sum                                        # Go Module Dependencies Sum File
```

- `api/`: Contains the API-related components.
  - `v1/`: Represents API version 1.
    - `controllers/`: Houses the controller files responsible for handling HTTP requests and responses for different API endpoints.
    - `routes/`: Defines API routes and their associated handlers.
    - `validators/`: Contains validation logic for request data to ensure data integrity and security.
  - `middleware/`: Contains middleware functions, such as authentication middleware.
  - `models/`: Defines data models used throughout the application.
  - `services/`: Houses the business logic and service implementations for various functionalities, such as user management and notifications.

- `config/`: Stores configuration files for different environments (e.g., dev_config.yaml and prod_config.yaml) and a central configuration module (config.go) to manage environment-specific settings.

- `database/`: Manages database-related components, including migrations for database schema updates (migrations/) and the database connection setup (database.go).

- `firebase/`: Contains Firebase-related code, including Firebase Authentication Service (auth_service.go) and Firebase Realtime Database Service (database_service.go).

- `email/`: Houses email-related services, such as the Email Service (email_service.go), Email Verification Service (verification_email.go), and Password Reset Email Service (password_reset_email.go).

- `sms/`: Stores SMS-related services, including the SMS Service (sms_service.go) and SMS Verification Service (verification_sms.go).

- `utils/`: Provides utility functions and helper methods that can be used across the application (helpers.go).

- `main.go`: The main entry point of the application.

- `go.mod` and `go.sum`: Go module files for managing dependencies.


In this approach, the Firebase, email, and SMS service files are organized in their respective directories. This organization can be beneficial if:

- Each service is relatively complex and involves multiple files.
- I want to keep service-specific code well-isolated and organized.
- I prefer a clear separation between different types of services.

The current structure follows this approach, ensuring that service-specific code is organized within its respective directory for clarity and maintainability.

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

3. Initialize a Go module for your project:

    ```sh
    $ go mod init user-service
    ```

    This command creates a go.mod file in your project directory and sets the module path to "user-service." You can replace "user-service" with your desired module name.

4. Install the project dependencies using Go modules:

   ```sh
    go mod tidy
   ```

5. Run swag init to generate the Swagger documentation. Navigate to your project's root directory and run:

   ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
   ```

   ```sh
    swag init
   ```

## Usage

### Running the Server

To run the User Service on different environments, follow these instructions:

#### Development Environment (Dev):

1. Open a terminal.

2. Navigate to the project directory:

    ```
    $ cd user-service
    ```

3. Build the application:

    - On Windows:

    Open a command prompt, navigate to the project directory, and run the following command to build the application:

    ```cmd
    go build -o user-service.exe main.go
    ```

    - On Ubuntu (or Linux):

    Open a terminal, navigate to the project directory, and run the following command to build the application:

    ```sh
    $ go build -o user-service main.go
    ```

4. Run the server in development mode using the dev configuration:

    - On Windows:

    After building the application, you can run the server in development mode with the dev configuration using the following command:

    ```cmd
    user-service.exe --config=config\dev_config.yaml
    ```

    - On Ubuntu (or Linux):

    After building the application, you can run the server in development mode with the dev configuration using the following command:

    ```sh
    $ ./user-service --config=config/dev_config.yaml
    ```

####  Production Environment (Prod):

1. Open a terminal.

2. Navigate to the project directory:

    ```
    $ cd user-service
    ```

3. Build the application:

    - On Windows:

    Open a command prompt, navigate to the project directory, and run the following command to build the application:

    ```cmd
    go build -o user-service.exe main.go
    ```

    - On Ubuntu (or Linux):

    Open a terminal, navigate to the project directory, and run the following command to build the application:

    ```sh
    $ go build -o user-service main.go
    ```

4. Run the server in production mode using the prod configuration:

    - On Windows:

    After building the application, you can run the server in production mode with the prod configuration using the following command:

    ```cmd
    docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
    user-service.exe --config=config\prod_config.yaml
    ```

    - On Ubuntu (or Linux):
    
    After building the application, you can run the server in production mode with the prod configuration using the following command:

    ```sh
    $ docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
    $ ./user-service --config=config/prod_config.yaml
    ```

These commands will launch the User Service with the specified configuration, whether in development or production mode. Be sure to customize the configurations in `dev_config.yaml` and `prod_config.yaml` to suit your environment settings.

### Functions

The User Service provides the following functions:

- User registration and management.
- Email and phone number verification.
- Phone number validation using Twilio.
- User login and logout.
- Password reset.
- User profile management.
- User search and listing.
- Sending notifications to users.
- Bearer token-based authentication.

### APIs

User Management APIs

#### 1. User Registration

- Endpoint: /api/v1/register (POST)

- Description: Allows users to register by providing their information.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/register -d '{
  "email": "user@example.com",
  "first_name": "John",
  "middle_name": "Doe",
  "last_name": "Smith",
  "password": "secure_password",
  "date_of_birth": "1990-01-01T00:00:00Z",
  "phone_number": "1234567890",
  "address": "1234 Elm St, Some City, Some Country",
  "country": "Some Country",
  "province": "Some Province",
  "avatar_url": "http://example.com/avatar.jpg"
}'
```

#### 2. Email Verification

- Endpoint: /api/v1/verify-email (POST)

- Description: Allows users to verify their email addresses using a verification code sent via email.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/send-verification-email -d '{"email": "example@example.com", "first_name": "John", "verification_link": "http://example.com/verify?token=123456"}'
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
curl -X POST http://localhost:8080/api/v1/logout -H "Authorization: Bearer <your-access-token>" -d '{"refresh_token": "<your-refresh-token>"}'
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

### Phone Number Validation APIs (Twilio-based)

#### 11. Send Verification Code via SMS
- Endpoint: /api/v1/send-verification-sms (POST)

- Description: Allows users to request a verification code to be sent to their phone number via SMS.

- Sample cURL Request:

```curl
curl -X POST http://localhost:8080/api/v1/send-verification-sms -H "Authorization: Bearer <your-access-token>"
```

#### 12. Verify Phone Number with Code

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
