# MyGram - A Free Photo Sharing App Written in Go

MyGram is a photo sharing app built with Go. It allows users to share, view, and comment on photos posted by other users. MyGram is a final project submission for the "Scalable Web Services with Go" course offered by Digitalent in collaboration with Hacktiv8.

## Features

- Users can create an account by registering an email address and selecting a username.
- Users can share photos with other users.
- Users can view and comment on photos shared by other users.

## Getting Started

To get started with MyGram, follow the steps below:

1. Clone the repository:

    ```bash
    git clone https://github.com/szczynk/MyGram.git
    ```

1. Install all required dependencies:

    ```bash
    cd MyGram && go mod tidy
    ```

1. Copy the example config file and adjust the config file:

    ```bash
    cp config.example.yaml config.yaml
    ```

1. Start the server:

    ```bash
    go run main.go
    ```

1. Check the API documentation:

    ```bash
    http://localhost:8080/swagger/index.html
    ```

## Contributing

Contributions are welcome! If you'd like to contribute to MyGram, please fork the repository and create a pull request.

## License

MyGram is open source and available under the MIT License.
