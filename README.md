### usage

```
curl localhost:8080/user/create -X POST -H "Content-Type: application/json" -d '{"name": "test"}'
curl localhost:8080/user/get -X GET -H "Content-Type: application/json" -H "x-token:a"
curl localhost:8080/user/update -X PATCH -H "Content-Type: application/json" -H "x-token:a" -d '{"name": "test2"}'
```


```
dco up
dco` exec api bash
go run cmd/migrate/main.go -e up
go run cmd/server/main.go
```

```
https://qiita.com/locomotive/items/b2bcca90f738b2bca378
https://qiita.com/daiki-murakami/items/c8f9df8defc937e185ee
https://github.com/auth0/go-jwt-middleware
https://qiita.com/gold-kou/items/45a95d61d253184b0f33
https://github.com/golang-jwt/jwt
https://otiai10.hatenablog.com/entry/2016/10/03/044300
```