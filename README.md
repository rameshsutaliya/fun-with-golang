# fun-with-golang
Just go lang stuff

To Execute this repo:
`go run main.go`

To run all the test cases:
`go test ./...`

To run a specific test case:
`go test -run <function_name>`

To see all the test cases verbose:
```sh
go test ./... -v
```

change the build.sh access mode:
``` sh
chmod +x build.sh
 ```

Cross Compilation:
Web Assembly: 
``` sh
gogio -target js fun-with-golang
```
To Run this at server: 
```sh
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("fun-with-golang")))'
```