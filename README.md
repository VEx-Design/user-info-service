# User Info Service

## Introduction
The User Info Service is designed to manage and retrieve user information efficiently. It provides a structured approach to handle user data, ensuring secure storage and easy access to user-related information.

## Features
- Secure storage of user information
- Efficient retrieval of user data
- Easy integration with other services

## Installation Dependencies
Ensure you have Go installed. After cloning the project, install the necessary packages and libraries with the following command:

```bash
go mod tidy
```

## Environment Variables
Create a `.env` file in the root directory and configure it based on `.env.example`. This file should contain all necessary environment-specific configurations.

## Running the Project
To run the project, use the following command:

```bash
make restart
```

This command will build and start the service.

## Docker
### Building the Docker Image and Running the Container
To build the Docker image and run the container, execute:

```bash
make up
```


