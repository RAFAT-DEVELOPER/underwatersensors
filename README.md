# Under water Sensor Backend API Implementation

## Overview

This project aims to implement a backend API for a set of underwater sensors. The sensors, though not physically installed yet, will be simulated to test the functionality of the API.

## Project Components

### Tools Used

- **Go**: The API will be implemented using the Go programming language.
- **Docker/Docker Compose**: Docker containers will be utilized for local development and testing.
- **PostgreSQL**: Data storage will be handled through PostgreSQL.
- **Redis**: Caching will be employed to enhance data retrieval performance.
- **Gin**: The Gin framework will be used for routing and handling HTTP requests.
- **Gorm**: The Gorm library will serve as the Object-Relational Mapping (ORM) tool for database interactions.
- **Viper**: Configuration settings will be managed using Viper.
- **Swagger**: API documentation will be generated using Swagger to ensure clear and accessible documentation.
- **Air**: Live reloading will be facilitated by the Air tool to streamline development and testing.

### Simulation

Since the physical sensors are not yet installed, a simulation environment will be created to mimic sensor data and interactions. This will allow for the development and testing of the API's functionality in a controlled environment.

## API Documentation

Detailed API documentation can be found in the Swagger documentation, which will be generated as part of the project. This documentation will provide insights into the available endpoints, request parameters, and responses.



## Table of Contents
- [Prerequisites](#prerequisites)
- [Requirements](#requirements)
- [Setup](#setup)
- [PostgreSQL Database](#postgresql-database)
- [API Endpoints](#api-endpoints)
- [Caching](#caching)
- [Swagger Documentation](#swagger-documentation)
- [Postman Collection](#postman-collection)
- [Tools Used](#tools-used)

## Requirements

The requirements for the underwatersensors project are as follows:
- Generate fake data for the sensors, including temperature, transparency, and fish species counts.
- Implement an API to access the sensor data.
- Use Docker Compose to provide a convenient way to set up the project locally.
- Store the data in PostgresSQL.
- Implement a caching for some API endpoint using Redis.
- Include a Swagger specification for the API.
- Optionally, add end-to-end tests.

## Prerequisites

Before you begin working with the underwatersensors project, please ensure that your system meets the following requirements:

- **Go**: You'll need to have Go installed on your machine. You can download and install it from [the official Go website](https://golang.org/dl/).

- **Docker**: To run the underwatersensors project locally, you'll need Docker. You can obtain Docker from [the official Docker website](https://docs.docker.com/get-docker/).

- **PostgreSQL Client**: To interact with the database, you should have a PostgreSQL client installed. You can choose from tools like pgAdmin or DataGrip.

- **Redis Client**: For Redis-related tasks, it's essential to have a Redis client installed. You can use Redis Desktop Manager or a similar tool.

- **Gin Package**: To work with the Gin web framework for Go, install it using the following command:
  ```
  go get github.com/gin-gonic/gin
  ```

- **Gorm Package**: The Gorm Go ORM library is used in this project. Install it with the command:
  ```
  go get gorm.io/gorm
  ```

- **Air Package**: Air is used for enabling hot-reloading in the Gin server. Install it with the following command:
  ```
  go install github.com/cosmtrek/air@latest
  ```

- **Viper Package**: Viper, a Go configuration library, is required. Install it using the following command:
  ```
  go get github.com/spf13/viper
  ```

These requirements are essential for setting up and running the underwatersensors project on your local system. If you encounter any issues or have questions, please don't hesitate to reach out for assistance.
   
## Setup

To set up the "underwatersensors" project locally, follow these structured steps:

### 1. Clone the Repository

To start, clone the project repository from its GitHub source:

```bash
git clone https://github.com/RAFAT-DEVELOPER/underwatersensors.git
```

### 2. Navigate to the Project Directory

Change your current working directory to the project's root folder:

```bash
cd underwatersensors
```

### 3. Start the Project with Docker Compose

Initiate the project by launching the required containers using Docker Compose:

```bash
docker-compose up -d
```

### 4. Run the Project with Air

Execute the project using "Air" to enable automatic server reloading upon file changes:

```bash
air
```

The API will be accessible at http://localhost:8000.

### 5. Migrate the Database

Run the database migration process to set up the database schema:

```bash
go run migrate/migrate.go
```

### 6. Generate Sensor Data

Generate sensor data by running the provided script:

```bash
go run ./generatedata/sensors/main.go
```


### 7. Insert Seed Data into the Database
Connect to the database and add generated or provided data for sensors, sensor groups, and fish species. You can use tools like pgAdmin or DataGrip with the credentials available in the `app.env` file. The SQL file with sensor groups and fish species can be found in the assets folder with the following name:

1. assets/DataCreation.sql
2. assets/initialSQL.sql

- Run these query in the database
- 24 sensor groups will create, each with a unique Greek letter name (e.g., alpha, beta, gamma, etc.).
- Sensor groups cover various ocean zones: Epipelagic, Mesopelagic, Bathypelagic, Abyssopelagic, and Trenches.
- Fish species data is obtained from [this source](https://oceana.org/ocean-fishes).
- Sensor coordinates are generated based on the specific zone.

### 9. Start Continuous Data Generation

- Temperature and transparency data are continuously generated based on the zone, depth, and time of the day.
- Fish data is generated based on their habitat zone.
Commence the continuous data generation process for sensor observations using the provided script:

```bash
go run ./generatedata/continious/main.go
```
Data will written in batches when the buffer is full or every 10 minutes.

## PostgreSQL Database

- Database migrations are managed using the Gorm library.
- The database schema is illustrated in the provided image: [Database Schema](assets/Postgres_DB_Schema.PNG).


## API Endpoints

The underwatersensors API provides a range of endpoints for accessing sensor data and gathering statistics, including:

- `GET /group/<groupName>/transparency/average`: Retrieve the current average transparency inside the group.

- `GET /group/<groupName>/temperature/average`: Retrieve the current average temperature inside the group.

- `GET /group/<groupName>/species`: Retrieve the complete list of species (with counts) currently detected inside the group.

- `GET /group/<groupName>/species/top/<N>`: Retrieve the top N species (with counts) currently detected inside the group.

- `GET /region/temperature/min`: Retrieve the current minimum temperature inside the specified region.

- `GET /region/temperature/max`: Retrieve the current maximum temperature inside the specified region.

- `GET /sensor/<codeName>/temperature/average`: Retrieve the average temperature detected by a particular sensor between specified date/time pairs (UNIX timestamps).

For more details on request/response formats and endpoints, please refer to the Swagger specification, which is accessible at http://localhost:8000/api/docs/index.html.

Additionally, specific CRUD operations are available; consult the provided Postman collection in the assets folder for further details.

## Caching

The API employs caching for specific endpoints in Redis with a Time-To-Live (TTL) of 10 seconds. Cached endpoints include:

- `GET /group/<groupName>/transparency/average`

- `GET /group/<groupName>/temperature/average`

## Swagger Documentation

The Swagger documentation for the API is provided at http://localhost:8000/api/docs/index.html. It comprehensively covers statistics endpoints.

## Postman Collection

You can find the Postman collection for the API in the assets folder.

Please follow these steps to set up and explore the "underwatersensors" project locally. Should you encounter any issues or have inquiries, please do not hesitate to contact us for assistance.