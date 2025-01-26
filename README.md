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


