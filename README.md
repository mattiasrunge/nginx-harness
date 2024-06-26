# nginx-harness

This application wraps nginx and exposes a endpoints which can be used
to add, remove and list new proxy routes on demand.

## Updating proxy routes

The application provides several HTTP endpoints that allow you to interact with the proxy routes.

To use the HTTP endpoints, follow these steps:

1. Start the application by running the `make run` command.

2. Once the application is running, you can send HTTP requests to the following endpoints:

  - `GET /_route`: This endpoint returns a list of all the currently configured proxy routes.

  - `POST /_route`: Use this endpoint to add a new proxy route. Send a JSON payload with the necessary information, such as the target URL and the route path.

  - `DELETE /_route/{path}`: This endpoint allows you to remove a specific proxy route. Replace `{path}` with the path of the route you want to delete.

3. To interact with these endpoints, you can use tools like cURL or Postman. Here are some examples:

  - To get the list of routes:
    ```
    curl http://localhost:9898/_route
    ```

  - To add a new route:
    ```
    curl -X POST -H "Content-Type: application/json" -d '{"path": "/example", "target": "http://example.com"}' http://localhost:9898/_route
    ```

  - To delete a route:
    ```
    curl -X DELETE http://localhost:9898/_route/%2Fexample
    ```

Replace `http://localhost:9898` with the appropriate URL if you are running the application on a different host or port.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository.
2. Build the project using the `make build` command.
3. Run the application using the `make run` command.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
