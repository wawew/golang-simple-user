# Simple Profile Commenting System

## Table of Contents
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Base URL](#base-url)
  - [Endpoints](#endpoints)
    - [Users](#users)

## Setup

### Prerequisites
- Docker Compose
- Go (1.22.2 or later)
- PostgreSQL (13 or later)

### Installation
1. Clone the repository:
   ```
   git clone https://github.com/wawew/golang-simple-user.git
   ```
2. Navigate to the project directory:
   ```
   cd golang-simple-user
   ```
3. Install the dependencies:
   ```
   go mod download
   ```
4. Create a `.env` file in the root directory and add your environment variables. You can use the `.env.example` file as a template.

## Running the Application
To run the application with docker compose, use the following command:
```
docker compose up
```
The application will start, and you can access it at `http://localhost:3000`.

## API Endpoints

### Base URL
`http://localhost:3000`

### Endpoints

#### Users
- **POST /DisplayUser**: Get a user by ID.
  - Request Body:
    ```json
    {
        "Userid": "string"
    }
    ```
  - Response:
    ```json
    {
        "Userid": "string",
        "Name": "string"
    }
    ```
  - Screenshot:
    ![DisplayUser](https://github.com/wawew/golang-simple-user/assets/69183538/e43a5a13-5acd-40df-8a93-89020c3e54d2)
- **POST /DisplayAllUsers**: Get all users.
  - Response:
    ```json
    [
        {
            "Userid": "string",
            "Name": "string"
        }
    ]
    ```
  - Screenshot:
    ![DisplayAllUsers](https://github.com/wawew/golang-simple-user/assets/69183538/0461e383-dd25-4d43-a06e-001ae1bb31e9)
