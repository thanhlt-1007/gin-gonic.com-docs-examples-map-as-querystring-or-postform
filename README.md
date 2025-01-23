# gin-gonic.com-docs-examples-map-as-querystring-or-postform

- Map as querystring or postform parameters

- Reference: https://gin-gonic.com/docs/examples/map-as-querystring-or-postform/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go mod

```sh
go mod download
```

## go run

```sh
go run .
```

## cURL

```sh
curl --location --globoff 'localhost:8080/map?ids[a]=A&ids[b]=B' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'names%5Ba%5D=A' \
--data-urlencode 'names%5Bb%5D=B'
```
