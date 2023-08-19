# Go Quotes Scraper API 🚀

This API scrapes quotes from `quotes.toscrape.com` and returns them in JSON format.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

## Getting Started

### Running locally

1. Clone this repository:
```bash
git clone https://github.com/marcusziade/goscrapequotes.git
cd goscrapequotes
```

2. Download Go dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run .
```

Visit `http://localhost:8080/quotes` in your browser or using tools like [Postman](https://www.postman.com/) to see the scraped quotes.

### Running with Docker

1. Build the Docker image:
```bash
docker build -t goscrapequotes .
```

2. Run the Docker container:
```bash
docker run -p 8080:8080 goscrapequotes
```

Visit `http://localhost:8080/quotes` to access the API.

## Testing

Run the tests using the command:
```bash
go test
```

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Write your code and tests.
4. Open a pull request against the `master` branch.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
