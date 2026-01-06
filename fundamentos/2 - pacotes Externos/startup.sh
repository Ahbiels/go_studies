go mod init external/packages
go get github.com/badoux/checkmail

go run main.go
go build

./packages