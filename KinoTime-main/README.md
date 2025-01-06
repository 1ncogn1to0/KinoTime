# KinoTime

## Project Description

**KinoTime** is a web application built using Go, designed to manage and display movie information. It features functionalities such as creating, reading, updating, and deleting movie data, interacting with a PostgreSQL database, and serving static files like HTML, CSS, and JavaScript for the frontend.

### Purpose

The purpose of this project is to practice building a RESTful API using Go and managing movie data. It targets developers who are interested in learning how to build web servers with Go, CRUD operations, and working with databases in a practical context.

### Features

- **CRUD Operations** for managing movie data
- **POST and GET** requests for movie information
- **Static file serving** for front-end files (HTML, CSS, JS)
- **Error handling** and JSON response formatting
- **Database integration** with PostgreSQL

## Team Members

- Aliya Akmagambetova
- Samat Rashatov

## Screenshot

![Main Page Screenshot](https://github.com/1ncogn1to0/KinoTime/blob/main/page.jpg)  

## Project Setup and Installation

Follow these steps to set up and run the project locally.

### Prerequisites

- Go 1.21 or later
- PostgreSQL (Ensure that the database is running)
- A `.env` file or adjust the database connection string accordingly in the code.

### Tools and Resources Used

- **Go**: For building the server and handling requests.
- **PostgreSQL**: For database management.
- **Gorilla Mux**: For routing in Go.
- **GORM**: For interacting with the PostgreSQL database.
- **JSON**: For data exchange format.
- **HTML/CSS/JS**: For frontend and static file serving.

## How to Start the Project

### Step 1: Clone the repository

Clone the project repository to your local machine.

bash
git clone https://github.com/your-username/KinoTime.git
cd KinoTime

## How to Start the Project

### Step 2: Install dependencies

Install the required Go dependencies:

bash
go mod tidy

This command will download and install all necessary libraries for the project.

### Step 3: Set up the database
Make sure PostgreSQL is installed and running.
Ensure the connection string in the code (db.NewDb()) is correct. You may want to set up a .env file with the necessary database credentials.
Example .env file:

env

DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=KinoTime
DB_HOST=localhost
DB_PORT=5432

Run the migration to create necessary tables in the database:
bash

go run main.go
This will automatically run the migrations for the database (using gorm), setting up the required tables such as movies.

### Step 4: Run the server
Start the server by running:

bash
go run main.go
This will start the server, and you will see an output like:

bash

Server started: http://localhost:8080
### Step 5: Visit the web page
Open your browser and go to http://localhost:8080. The server will handle requests to the website and perform CRUD operations.

GET request: Retrieve movie data.
POST request: Add new movies.
PUT request: Update movie data.
DELETE request: Delete a movie entry.

### License
This project is licensed under the MIT License - see the LICENSE file for details.
