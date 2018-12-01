# Web Assembly built by Golang
#### build server:
> go build -o server server.go
#### build lib.wasm:
> GOARCH=wasm GOOS=js go build -o lib.wasm main.go
### run server:
> ./server
