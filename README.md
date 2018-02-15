# Go Protobuf OneOf JSON
This project illustrates the issue that occurs when marshalling a protobuf
object to JSON when that object has a field that's an array and the elements
of the array utilize `OneOf`.

## Getting Started
To illustrate the issue, run the following command:

```shell
$ go get github.com/akutz/go-proto-oneof-json && \
  go test -v github.com/akutz/go-proto-oneof-json
```
