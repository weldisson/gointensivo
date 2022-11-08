# Go Intensivo


## Steps

### Init repository
``` 
go mod init github.com/weldisson/gointensivo 
```

### internal
- internal resources
- screaming archtecture
```
 . internal/
 . . order/
 . . . entity/ # entity layer
 . . . . order.go 
 . . . infra/ 
 . . . . database/ #database layer
 . . . . . order_repository.go 
 . . . usecase/ #usecase layer
 . . . . calculate_price.go
 . . . . get_total.go
```

### downloading external packages
```
# add you external package in your code and run
go mod tidy
```

###  run all tests
```
go test ./...
```