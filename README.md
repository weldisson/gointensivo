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

### build app
```
go build cmd/consumer/main.go
# or for windows os system
GOOS=windows go build cmd/consumer/main.go 
```
### run docker
```
docker-compose up -d
```

### configure grafana
- Login with admin/admin in http://localhost:3000
- Click in skip 
- Click in Add your data source
- Click in Prometheus
- Configure prometheus with URL:http://prometheus:9090
- Click in save and test
- Create a dashboard
- Go to grafana dashboards website: https://grafana.com/grafana/dashboards/10991-rabbitmq-overview/
- Click in "Copy ID to clipboard" (this ID is 10991)
- Go to menu dashboad in grafana, click in "Import" or click here (http://localhost:3000/dashboard/import) 
- Paste the dashboard ID (10991) in "import via grafana.com" and click in the first "load" button.
- In the next step click in "select a Prometheus data source" and choose Prometheus
- Click in "Import"
- go to home and save you Dashboard!


### creating docker image
```
docker build -t weldissonaraujo/gointensivo:latest -f Dockerfile.prod .
docker push weldissonaraujo/gointensivo:latest
```
### run kubernetes
- create cluster:
```
kind create cluster --name=gointensivo
kubectl cluster-info --context kind-gointensivo
```
- run :
```
kubectl apply -f k8s
kubectl get deployment
```
--- 