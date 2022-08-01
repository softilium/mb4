call ent-generate.cmd
go vet -vettool=c:\users\softi\go\bin\shadow.exe
go vet 
rem go build -ldflags "-s -w"
go build
