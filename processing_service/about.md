

# Run the Services
Start the gRPC server:

```go
go run processing_server.go
```

Start the API gateway:

```go
go run client/api_gateway.go
```


## Test API

```PY
curl -X POST -H "Content-Type: application/json" -d "{\"request_id\":\"123\", \"data\":\"test data\"}" http://localhost:8080/process

```



* Custom Error Handling: The gRPC server returns appropriate status codes and error messages.

* Retries with Backoff: The API Gateway retries the gRPC request up to 3 times with an increasing delay between attempts.

* Detailed Error Responses: The API Gateway provides detailed error messages back to the client in case of failures.





