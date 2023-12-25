# Database Monitoring Tool

## Overview

This simple Go application provides a basic database monitoring tool for Microsoft SQL Server. It connects to a specified SQL Server instance, queries a specified table's row count at regular intervals, and outputs the count along with a timestamp. This tool can be useful for tracking changes in the number of records in a table over time.

## Prerequisites

Before using this tool, ensure you have the following:

- Go programming language installed
- Access to a Microsoft SQL Server database
- Necessary connection details (server address, port, username, password, and database name)
- The go-mssqldb package, which can be installed using:
  ```bash
  go get github.com/denisenkom/go-mssqldb
# Usage

1. **Clone the repository:**

    ```bash
    git clone https://github.com/karelpelcak/GoLang_UserCounter
    ```

2. **Navigate to the project directory:**

    ```bash
    cd GoLang_UserCounter
    ```

3. **Modify the `main.go` file with your SQL Server connection details and target table name.**

4. **Run the application:**

    ```bash
    go run main.go
    ```

# Configuration

Edit the `main.go` file to set the following variables:

```go
server := "server_address"      // Server address
port := 1433                    // Default port
user := "username"              // Your DB username
password := "password"          // Your DB password
database := "database_name"      // Name of your Database
tableName := "table_name"        // Name of the target table
