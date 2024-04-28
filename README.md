# URL Shortener Service

This URL Shortener is a lightweight, efficient web service that converts long URLs into manageable links that are easier to share and manage. 
It's built using Golang and provides both a user-friendly web interface and a developer-friendly API.

## Answers on questions

### Describe your solution. What tradeoffs did you make while designing it, and why?

I developed the application to show how I use the architect

#### Design Compromises:
- Choice of Hash Generation Algorithm: I used a basic hashing algorithm for simplicity. Instead of more complex methods with collision control, I chose a simple method which could lead to collisions when scaling.
- Data Storage: For the prototype, I chose a local SQLite database, which is not optimal for high-load systems. This choice was made to simplify development and deployment.
- Error Handling: The initial version only handles basic errors. Detailed error logging and processing were simplified to save time at the early stage.
- Testing: I made no tests. Tested only manually.
- Deploy: No Docker

#### If this were a real project, how would you improve it further?
- Hash Generation Algorithm: It will be nice to guarantee that short link url will be unique. One of the decision: converting id into a short code using a higher base numbering system (letters and numbers)
- Scalability: Transition to a more scalable database system like PostgreSQL or distributed storage systems such as Redis to manage high traffic and provide fast data access. Deploy several instances with auto-scalability. 
- Security: Implement additional security measures such as input validation and sanitization to prevent XSS and SQL injection. Also, I will protect endpoints by at least rate limit  
- Error Handling: Implement a comprehensive logging and monitoring system for efficient troubleshooting and increased service reliability. Protect from infinite(long chain) redirect loops, when user use our address as a full url to shorten it.
- User Interface: Create a user-friendly web interface for users to not only shorten URLs but also manage them (view statistics, delete, etc.).
- Testing: Cover code with basic unit tests
- Deploy: Add Dockerfile to deploy it easily
- Monitoring: /health endpoint and gather metrics

#### What is the math & logic behind “short url” if there is any?
Hashing: The primary mechanism is using a hash function to convert a long URL into a unique(not really) short identifier using SHA hash function, truncated to a certain number of characters.
Length of the short url: For now it's 8 bytes. It was chosen because it's short enough and provide some options, but there is the real problem of balance. If we want really short urls, maybe we need to get user information and have autorization process. So, each user could have the own list of the shorten links. 
Uniqueness: Algorithms can be used that ensure a one-to-one correspondence between original and final data. An example of such implementation is using an incremental ID in the database and converting it into a short code using a higher base numbering system (e.g., base 62, which includes letters and numbers).

## Features

- **URL Shortening**: Instantly shorten long URLs.
- **Redirection**: Seamless redirection from shortened links to original URLs.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What you need to install the software:

- Golang 1.20+
- [NOT NOW] Docker (optional, for containerization)
- [NOT NOW] PostgreSQL

### Installing

A step by step series of examples that tell you how to get a development environment running.

#### Setup Environment

Clone the repository:

```bash
git clone https://github.com/DmSide/shortenLinker.git
cd shortenLinker
```

#### Install dependencies:
```bash
go mod tidy
```

Copy `.env.example` to `.env` and change the values

#### Setting up the Database(skip it for now - used SQLLite instead)
Create a PostgreSQL database and user:

```sql
CREATE DATABASE shortenLinker;
CREATE USER shortenLinkerAdmin WITH ENCRYPTED PASSWORD 'changeit';
GRANT ALL PRIVILEGES ON DATABASE shortenLinker TO shortenLinkerAdmin;
```

### Running the Server
To start the server, run:

```bash
go run cmd/app/main.go
```

This will start the local server on http://localhost:8081.

### Running the tests(Skip it fot now)
To run the automated tests for this system:

```bash
go test ./...
```

### Test by Curl
```bash
curl -X POST -H "Content-Type: application/json" -d '{"url":"http://example.com"}' http://localhost:8081/encode
curl -X GET http://localhost:8081/decode/:shorturl
```