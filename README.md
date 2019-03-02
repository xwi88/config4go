# config

If you use govendor or blocked by the GFW, you can ref `https://github.com/xwi88/config`.

The rep support module and without vendor dir.

## support

support format: toml, ymal, json, ini, properties and so on.

## usage

`go get -u github.com/xwi88/config4go`

## tips

If go get timeout, you can try set follows:

```bash
export GO111MODULE=on #go1.12 default set on
export GOPROXY=https://goproxy.io
```

## example

```go
	type redisOptions struct {
		Host        string
		Port        string
		Password    string
		IdleTimeout int
		MaxIdle     int
		MaxActive   int
	}

	type testConf struct {
		Name  string
		Redis redisOptions
	}

	var conf map[string]interface{}

	testfile := "./testdata/app_test.toml"
	err := LoadConfig(testfile, &conf)
	if err != nil {
		// some deal
    }
    
    fmt.Println(conf.Host)
```
