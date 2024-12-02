# Go Time API with MySQL Database Logging

This Go-based API returns the current time in the Toronto timezone and logs the timestamp into a MySQL database. The API exposes an endpoint `/current-time`, which, when accessed, provides the current time in JSON format and saves the timestamp in the `time_log` table in the MySQL database.

## Features

- **Current Time API**: Fetches the current time in the Toronto timezone.
- **Database Logging**: Stores the timestamp of each request in the `time_log` table of a MySQL database.
- **Simple and Easy to Use**: Just run the Go server, and the API is ready to be used.

## Prerequisites

Before running the project, ensure you have the following installed:

1. **Go (Golang)**: Version 1.18 or higher.
2. **MySQL**: A running MySQL server (Localhost:3306), with a database named `timedb`.
3. **MySQL Driver for Go**: Install it via the command `go get github.com/go-sql-driver/mysql`.

## Setting Up MySQL

Ensure that you have a database and a table set up in MySQL to store the time logs. You can create a database `timedb` and a table `time_log` as follows:

```sql
CREATE DATABASE timedb;

USE timedb;

CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
