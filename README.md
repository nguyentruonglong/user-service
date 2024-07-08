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
- Sending Notifications: Authenticated users can send notifications to others through multiple channels.
- Role Management: Manage user roles within the system to control access permissions.
- Group Management: Organize users into groups for collective management and permissions.
- User Activity Logs: Track and record user activities for audit and monitoring purposes.

This comprehensive suite of features caters to a wide range of user management and authentication needs, making the User Service a reliable and indispensable component for building secure and user-centric applications.

### Project Directory Structure

The User Service project is organized with a clear directory structure to promote code modularity, maintainability, and separation of concerns. Here's an overview of the project directory structure:

```
user-service/
    |-- api/
    |   |-- v1/
    |   |   |-- controllers/
    |   |   |   |-- notification_controller.go        # Notification Controller
    |   |   |   |-- search_controller.go              # Search Controller
    |   |   |   |-- user_login_controller.go          # User Login Controller
    |   |   |   |-- user_logout_controller.go         # User Logout Controller
    |   |   |   |-- user_register_controller.go       # User Register Controller
    |   |   |   |-- phone_verification_controller.go  # Phone Verification Controller
    |   |   |   |-- email_verification_controller.go  # Email Verification Controller
    |   |
    |   |   |-- routes/
    |   |   |   |-- notification_routes.go            # Notification Routes
    |   |   |   |-- search_routes.go                  # Search Routes
    |   |   |   |-- user_routes.go                    # User Routes
    |   |   |   |-- verification_routes.go            # Verification Routes
    |   |
    |   |   |-- validators/
    |   |   |   |-- notification_validator.go         # Notification Validator
    |   |   |   |-- search_validator.go               # Search Validator
    |   |   |   |-- user_validator.go                 # User Validator
    |   |   |   |-- verification_validator.go         # Verification Validator
    |   |
    |-- middlewares/
    |   |-- auth_middleware.go                        # Authentication Middleware
    |
    |-- models/                          
    |   |-- email_template.go                         # Email Template Model
    |   |-- group.go                                  # Group Model
    |   |-- permission.go                             # Permission Model
    |   |-- role.go                                   # Role Model
    |   |-- token.go                                  # Token Model
    |   |-- user.go                                   # User Model
    |
    |-- services/
    |   |-- notification_service.go                   # Notification Service
    |   |-- search_service.go                         # Search Service
    |   |-- user_service.go                           # User Service
    |   |-- verification_service.go                   # Verification Service
    |
    |-- config/
    |   |-- config.go                                 # Configuration Module
    |   |-- dev_config.yaml                           # Development Configuration
    |   |-- prod_config.yaml                          # Production Configuration
    |
    |-- database/
    |   |-- database.go                               # Database Connection
    |   |-- seed.go                                   # Database Seed Data
    |
    |-- development.sqlite3                           # SQLite Database File
    |-- docker-compose.yml                            # Docker Compose Configuration
    |
    |-- docs/
    |   |-- docs.go                                   # Documentation Module
    |   |-- swagger.json                              # Swagger JSON Configuration
    |   |-- swagger.yaml                              # Swagger YAML Configuration
    |
    |-- email_services/
    |   |-- email_service.go                          # Email Service
    |   |-- password_reset_email.go                   # Password Reset Email Service
    |   |-- verification_email.go                     # Email Verification Service
    |
    |-- firebase_services/
    |   |-- auth_service.go                           # Firebase Authentication Service
    |   |-- database_service.go                       # Firebase Realtime Database Service
    |   |-- firebase.go                               # Firebase Configuration
    |   |-- realtime_db.go                            # Firebase Realtime Database Interactions
    |
    |-- go.mod                                        # Go Module File
    |-- go.sum                                        # Go Module Dependencies Sum File
    |-- main.go                                       # Main Application Entry Point
    |
    |-- sms_services/
    |   |-- sms_service.go                            # SMS Service
    |   |-- verification_sms.go                       # SMS Verification Service
    |
    |-- tasks/
    |   |-- email_task.go                             # Email Task
    |   |-- worker.go                                 # Task Worker
    |
    |-- tests/
    |   |-- notification_controller_test.go           # Unit Tests for Notification Controller
    |   |-- notification_service_test.go              # Unit Tests for Notification Service
    |   |-- search_controller_test.go                 # Unit Tests for Search Controller
    |   |-- search_service_test.go                    # Unit Tests for Search Service
    |   |-- user_login_controller_test.go             # Unit Tests for User Login Controller
    |   |-- user_logout_controller_test.go            # Unit Tests for User Logout Controller
    |   |-- user_register_controller_test.go          # Unit Tests for User Register Controller
    |   |-- user_service_test.go                      # Unit Tests for User Service
    |   |-- verification_controller_test.go           # Unit Tests for Verification Controller
    |   |-- verification_service_test.go              # Unit Tests for Verification Service
    |
    |-- user-service                                  # Compiled Binary
    |
    |-- utils/
    |   |-- helpers.go                                # Utility Functions
    |
    |-- Dockerfile                                    # Docker Configuration
    |-- LICENSE                                       # License File
    |-- README.md                                     # Project Documentation
```


- `api/`: Contains the API-related components.
  - `v1/`: Represents API version 1.
    - `controllers/`: Houses the controller files responsible for handling HTTP requests and responses for different API endpoints.
    - `routes/`: Defines API routes and their associated handlers.
    - `validators/`: Contains validation logic for request data to ensure data integrity and security.
  - `middlewares/`: Contains middleware functions, such as authentication middleware.
  - `models/`: Defines data models used throughout the application.
  - `services/`: Houses the business logic and service implementations for various functionalities, such as user management and notifications.

- `config/`: Stores configuration files for different environments (e.g., dev_config.yaml and prod_config.yaml) and a central configuration module (config.go) to manage environment-specific settings.

- `database/`: Manages database-related components, including migrations for database schema updates (migrations/) and the database connection setup (database.go).

- `docs/`: Contains documentation-related files, including Swagger configuration files (swagger.json, swagger.yaml) and the main documentation module (docs.go).

- `email_services/`: Houses email-related services, such as the Email Service (email_service.go), Email Verification Service (verification_email.go), and Password Reset Email Service (password_reset_email.go).

- `firebase_services/`: Contains Firebase-related code, including Firebase Authentication Service (auth_service.go) and Firebase Realtime Database Service (database_service.go).

- `sms_services/`: Stores SMS-related services, including the SMS Service (sms_service.go) and SMS Verification Service (verification_sms.go).

- `utils/`: Provides utility functions and helper methods that can be used across the application (helpers.go).

- `tasks/`: Manages background tasks and workers, including email tasks (email_task.go) and the worker implementation (worker.go).

- `tests/`: Contains unit tests for various components of the application, ensuring code quality and reliability.

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

### Swagger generation:

Ensure you have the required dependencies installed. You'll need `gin-swagger` and `swag`:

```bash
$ go install github.com/swaggo/swag/cmd/swag@latest
$ export PATH=$PATH:$(go env GOPATH)/bin
$ source ~/.zshrc
```

```bash
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```

### Functions

The User Service provides the following functions:

- **User Registration and Management:** Allows users to register and manage their accounts with email verification and secure password management.
- **Email and Phone Number Verification:** Ensures users can verify their contact details through secure codes sent via email and SMS.
- **User Login and Logout:** Supports secure login mechanisms and clean logout processes, helping maintain user session integrity.
- **Password Reset:** Facilitates users in resetting their forgotten passwords through a secure email-based process.
- **User Profile Management:** Enables users to update and maintain their profiles with personal information and profile pictures.
- **User Search and Listing:** Provides functionality to search for other users by various criteria and view listings, supporting pagination and filtering.
- **Sending Notifications:** Allows authenticated users to send notifications through multiple channels.
- **Role Management:** Admins can create, modify, and assign roles to users, managing permissions systematically.
- **Group Management:** Supports the organization of users into groups for more streamlined management.
- **User Activity Logs:** Tracks user activities within the system for audit and monitoring purposes.

### APIs

#### User Management APIs

1. **User Registration**
   - **Endpoint:** `/api/v1/register` (POST)
   - **Description:** Allows users to register by providing their information.
   - **Sample cURL Request:**
     ```bash
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

2. **Send Verification Code via Email**
   - **Endpoint:** `/api/v1/send-verification-email` (POST)
   - **Description:** Allows users to request a verification code to be sent to their email.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/send-verification-email -H "Authorization: Bearer <access-token>"
     ```

3. **Verify Email with Code**
   - **Endpoint:** `/api/v1/verify-email/` (POST)
   - **Description:** Allows users to verify their email addresses by providing the received verification code.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/verify-email -d '{"verification_code": "123456"}' -H "Authorization: Bearer <access-token>"
     ```

4. **Send Verification Code via SMS**
   - **Endpoint:** `/api/v1/send-verification-sms` (POST)
   - **Description:** Allows users to request a verification code to be sent to their phone number via SMS.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/send-verification-sms -H "Authorization: Bearer <access-token>"
     ```

5. **Verify Phone Number with Code**
   - **Endpoint:** `/api/v1/verify-phone-number` (POST)
   - **Description:** Allows users to verify their phone number by providing the received verification code.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/verify-phone-number -d '{"verification_code": "123456"}' -H "Authorization: Bearer <access-token>"
     ```

6. **User Login**
   - **Endpoint:** `/api/v1/login` (POST)
   - **Description:** Allows users to log in and obtain a Bearer token by providing their email and password.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/login -d '{"email": "user@example.com", "password": "secure_password"}'
     ```

7. **User Logout**
   - **Endpoint:** `/api/v1/logout` (POST)
   - **Description:** Allows users to log out, effectively invalidating their Bearer token.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/logout -H "Authorization: Bearer <access-token>" -d '{"refresh_token": "<refresh-token>"}'
     ```

8. **Password Reset**
   - **Endpoint:** `/api/v1/reset-password` (POST)
   - **Description:** Allows users to reset their password by sending a password reset email.
   - **Sample cURL Request:**
     ```bash
     curl -X POST http://localhost:8080/api/v1/reset-password -d '{"email": "user@example.com"}'
     ```

9. **User Profile Management**
   - **Endpoints:** `/api/v1/profile` (GET, PUT)
   - **Description:** Allows users to retrieve and update their profile information.
   - **Sample cURL Requests:**
     ```bash
     # Retrieve Profile
     curl -X GET http://localhost:8080/api/v1/profile -H "Authorization: Bearer <access-token>"
     # Update Profile
     curl -X PUT http://localhost:8080/api/v1/profile -d '{"first_name": "John", "last_name": "Doe"}' -H "Authorization: Bearer <access-token>"
     ```

10. **User Search**
    - **Endpoint:** `/api/v1/user-search` (GET)
    - **Description:** Allows users to search for other users based on specified criteria.
    - **Sample cURL Request:**
      ```bash
      curl -X GET 'http://localhost:8080/api/v1/user-search?query=John&country=US'
      ```

11. **User Listing**
    - **Endpoint:** `/api/v1/user-list` (GET)
    - **Description:** Lists users with optional pagination and filtering options.
    - **Sample cURL Request:**
      ```bash
      curl -X GET 'http://localhost:8080/api/v1/user-list?page=1&per_page=10'
      ```

#### Role Management APIs

12. **Create Role**
    - **Endpoint:** `/api/v1/roles` (POST)
    - **Description:** Allows creation of a new role with specific permissions.
    - **Sample cURL Request:**
      ```bash
      curl -X POST http://localhost:8080/api/v1/roles -d '{
        "name": "Administrator",
        "permissions": ["create_user", "delete_user"]
      }'
      ```

13. **Update Role**
    - **Endpoint:** `/api/v1/roles/{role_id}` (PUT)
    - **Description:** Updates an existing role's permissions.
    - **Sample cURL Request:**
      ```bash
      curl -X PUT http://localhost:8080/api/v1/roles/{role_id} -d '{
        "permissions": ["update_user"]
      }'
      ```

14. **Delete Role**
    - **Endpoint:** `/api/v1/roles/{role_id}` (DELETE)
    - **Description:** Deletes a role.
    - **Sample cURL Request:**
      ```bash
      curl -X DELETE http://localhost:8080/api/v1/roles/{role_id}
      ```

#### Group Management APIs

15. **Create Group**
    - **Endpoint:** `/api/v1/groups` (POST)
    - **Description:** Allows creation of a new user group.
    - **Sample cURL Request:**
      ```bash
      curl -X POST http://localhost:8080/api/v1/groups -d '{
        "name": "Support Team",
        "user_ids": [1, 2, 3]
      }'
      ```

16. **Update Group**
    - **Endpoint:** `/api/v1/groups/{group_id}` (PUT)
    - **Description:** Updates an existing group's details.
    - **Sample cURL Request:**
      ```bash
      curl -X PUT http://localhost:8080/api/v1/groups/{group_id} -d '{
        "name": "Technical Support",
        "user_ids": [1, 2, 3, 4]
      }'
      ```

17. **Delete Group**
    - **Endpoint:** `/api/v1/groups/{group_id}` (DELETE)
    - **Description:** Deletes a user group.
    - **Sample cURL Request:**
      ```bash
      curl -X DELETE http://localhost:8080/api/v1/groups/{group_id}
      ```

#### User Activity Logs APIs

18. **View User Activity**
    - **Endpoint:** `/api/v1/users/{user_id}/activities` (GET)
    - **Description:** Retrieves the activity logs of a specific user.
    - **Sample cURL Request:**
      ```bash
      curl -X GET http://localhost:8080/api/v1/users/{user_id}/activities
      ```

#### Notification APIs

19. **Send Notification**
    - **Endpoint:** `/api/v1/notifications/send` (POST)
    - **Description:** Allows users or admins to send notifications to other users.
    - **Sample cURL Request:**
      ```bash
      curl -X POST http://localhost:8080/api/v1/notifications/send -d '{
        "user_id": "1",
        "message": "Hello, your account has been updated."
      }' -H "Authorization: Bearer <access-token>"
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
