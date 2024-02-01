### INSTALL

```
brew doctor

brew install protobuf

```

1. Install the protocol compiler plugins for Go using the following commands
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
2. Update your PATH so that the protoc compiler can find the plugins
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Install grpc to project
```
go get -u google.golang.org/grpc
```

4. Start evans to check service of grpc
```
make evans
```
