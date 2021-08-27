### usage

```
curl localhost:8080/user/create -X POST -H "Content-Type: application/json" -d '{"name": "test"}'
curl localhost:8080/user/get -X GET -H "Content-Type: application/json" -H "x-token:a"
curl localhost:8080/user/update -X PATCH -H "Content-Type: application/json" -H "x-token:a" -d '{"name": "test2"}'
```


```
dco up
go run cmd/migrate/main.go -e up
go run cmd/server/main.go
```