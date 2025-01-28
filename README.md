# Project from Let's Go Further by Alex Edwards.

All commits will follow each chapter name.

## 1 - Getting Started

- Create a skeleton directory structure for the project and explain at a high-level how our Go code and other assets will be organized.

- Establish a HTTP server to listen for incoming HTTP requests.

- Introduce a sensible pattern for managing configuration settings (via command-line flags) and using dependency injection to make dependencies available to our handlers.

- Use the `httprouter` package to help implement a standard RESTful structure for the API endpoints.

## 2 - Sending JSON Responses

- How to send JSON responses from your REST API (including error responses).

- How to encode native Go objects into JSON using the encoding/json package.

- Different techniques for customizing how Go objects are encoded to JSON — first by
using struct tags, and then by leveraging the json.Marshaler interface.

- How to create a reusable helper for sending JSON responses, which will ensure that all
your API responses have a sensible and consistent structure.

## 3 - Parsing JSON Requests

- How to read a request body and decode it to a native Go object using the encoding/json package.

- How to deal with bad requests from clients and invalid JSON, and return clear, actionable, error messages.

- How to create a reusable helper package for validating data to ensure it meets your business rules.

- Different techniques for controlling and customizing how JSON is decoded.

## 4 - Database Setup and Configuration

- How to install and set up PostgreSQL on you local machine

- How to use the psql interactil tool to create databases, PostgreSQL extensions and user accounts.

- How to initialize a database connection pool in Go and configure its settings to improve performance and stability.

## 5 - SQL Migrations

- The high-level principles behind SQL migrations and why they are useful.

- How to use the command-line migrate tool to programmatically manage changes to your database schema.

## 6 - CRUD Operations

- How to create a database model which isolates all the logic for executing SQL queries against your database.

- How to implement the four basic CRUD (create, read, update and delete) operations on a specific resource in the context of an API.

## 7 - Advanced CRUD Operations

- How to support partial updates to a resource (so that the client only needs to send the data that they want to change).

- How to use optimistic concurrency control to avoid race conditions when two clients try to update the same resource at the same time.

- How to use context timeouts to terminate long-running database queries and prevent unnecessary resource use.