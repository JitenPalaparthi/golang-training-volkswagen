# Golang


- To run go applications

```
go run main.go
```

- if main is split into multiple files

```
go run main.go greet.go
```

```
go run .
```

- To build

```
go build main.go
```

```
go build -o helloworld main.go
```

- To get the list of distributions for cross compilation

```
go tool dist list
```

- To see Assembly language output

```
go tool compile -S main.go 
```

- To compile and link 

```
 go tool compile -o hello.o main.go
 go tool link -o hello hello.o 
```

- To see what all things happens during build

```
go build -x
```

GOOS/GOARCH -> Operating-System/Processor-Type-Architecture

- To cross compile

```
GOOS=ios GOARCH=arm64 CGO_ENABLED=1  go build -o build/hello_ios_arm64  main.go
or
GOOS=linux GOARCH=s390x go build -o build/hello_mainframe_x390x  main.go
or 
GOOS=linux GOARCH=amd64 go build -o build/hello_linux_amd64 main.go
or
GOOS=windows GOARCH=amd64 go build -o build/hello_windows_amd64 main.go
```
- To build  a release mode kind of 

```
go build -ldflags="-s -w"  -o demo  main.go
```

- To do escape analysis

```
go build -gcflags="-m" main.go
```

## go mod commands

- pulls all required dependencies from remote
```
go mod tidy
```

- go mod vendor : Some projects,it is preferred to download third party code.If not very confident on third party remote location availability

```
go mod vendor
```

## keywords 

- break,case,chan,const,continue, default,else, fallthrough, for,func,go,goto,if,import,interface,map, package,range,return,select,struct,switch,type,var (24 covered out of 25)

## builtin functions

- append,cap,clear,close,complex,copy,delete,len,print,println,make,max,min,new,real,imag (16 out of 18 covered)

```go
if condition {
    a =1
}else{
    a =2
}
println(a)
```

```
a1 =1
a2 =2 
a3= $(a,a2)
println(a3)
```


https://github.com/JitenPalaparthi/golang-training-volkswagen