call ent-generate.cmd
go vet -vettool=c:\users\softi\go\bin\shadow.exe
go vet 
golangci-lint run

unparam github.com/softilium/mb4
unparam github.com/softilium/mb4/api
unparam github.com/softilium/mb4/backtest
unparam github.com/softilium/mb4/config
unparam github.com/softilium/mb4/cube
unparam github.com/softilium/mb4/db
unparam github.com/softilium/mb4/domains
unparam github.com/softilium/mb4/ent
unparam github.com/softilium/mb4/pages

rem go build -ldflags "-s -w"
go build



rem go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

rem go install mvdan.cc/unparam@latest