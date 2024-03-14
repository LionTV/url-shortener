# Go URL Shortener

This project is a simple URL shortener written in Go, allowing users to shorten long URLs and share them via the generated short link.

## What does it?

The program creates a SQLite database in the current user's directory to store the shortened URLs.
After creating the SQLite database, a web server is set up using the `net/http` package. This server hosts the various endpoints required for the functionality of the URL shortener.

## Features

- Shortening long URLs into custom short links.
- Tracking clicks for each short link.
- Customizing the port of the web server via a command-line parameter.


## API Routes

- `/api/create?url=yourwebsite`: Creates a short URL for the specified long URL.
- `/api/delete?short=yourshort`: Deletes the short URL corresponding to the given short code.
- `/api/all`: Returns all short URLs in JSON format.
- `/api/clicks?short=yourshort`: Returns the click count of the specified short URL.
- `/short`: Redirects to the long url.

## Installation and Usage

### 1. Download the Source Code

You can clone the source code using Git:

```bash
[git clone https://github.com/username/url-shortener.git](https://github.com/LionTV/url-shortener.git)
```

### 2. Install Dependencies

Make sure you have the latest go version. Then, navigate to the project directory and run the following command to install dependencies:

```bash
cd url-shortener
go mod tidy
```

### 3. Compile the program

```bash
go build
```

### 4. Run the program

Run the compiled program and specify the desired port of the web server as a command-line parameter. By default, port 80 will be used.

```bash
./url-shortener.exe -port=80
```

### 5. Using the URL Shortener

Open your web browser and navigate to **`http://localhost:<port>`**, where **`<port>`** is the port you specified (default is **8080**). You will then see the user interface of the URL shortener, where you can enter long URLs to shorten them.

## Contribution and Collaboration

Contributions are always welcome! If you find a bug or want to add a feature, feel free to open a pull request.
