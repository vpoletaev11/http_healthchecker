# http_healthchecker

## Step 1: Clone project

```shell
$ git clone https://github.com/vpoletaev11/http_healthchecker.git
```

## Step 2: Test project & Install dependencies
```shell
$ cd http_healthchecker
$ go test ./...
```

## Step 3: Run project

```shell
$ go run cmd/healthchecker/main.go -config_path testdata/config.json -req_timeout_sec 20
```
