# Light Gate

The program starts a simple HTTP server to serve files and handle requests. It allows setting the server port, enabling verbose logging, and choosing a directory to serve files from. The server stops execution when the help flag is provided, showing usage information instead.

## Install

### Source code

```
go install github.com/nrhox/lightgate
```

## Usage

```
Usage: lightgate [options]

Options:
  -d    Directory file
  -i    Index file endpoint (optional)
  -n    404 file (optional)
  -p    Port server
  -r    Path to _redirects file (optional)
  -v    Show current version
  -ver  Display detailed logs for each request
```

## Example

```
lightgate -p 8080 -d ./web
```

## License

MIT
