# httpecho

Minimal HTTP echo server for debugging. Returns your request back as JSON.

## Usage

```bash
go run main.go
# or
PORT=3000 go run main.go
```

## Docker

```bash
docker build -t httpecho .
docker run -p 8080:8080 httpecho
```

## Example

```bash
curl -X POST http://localhost:8080/test?foo=bar \
  -H "X-Custom: hello" \
  -d '{"msg": "hi"}'
```

Response:
```json
{
  "method": "POST",
  "path": "/test",
  "query": "foo=bar",
  "headers": {
    "X-Custom": ["hello"]
  },
  "body": "{\"msg\": \"hi\"}",
  "remote_addr": "127.0.0.1:54321",
  "timestamp": "2026-03-12T07:00:00Z"
}
```

## License

MIT
