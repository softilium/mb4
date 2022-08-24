call ent-generate.cmd
go vet -vettool=c:\users\softi\go\bin\shadow.exe
go vet 
golangci-lint run
rem go build -ldflags "-s -w"
go build



rem go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
