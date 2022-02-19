SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o test http1.go
go build -o test2 http2.go
go build -o p test777.go