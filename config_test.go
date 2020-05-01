package config4go

import (
	"testing"
	"time"
)

func Test_LoadConfig(t *testing.T) {
	type mysqlOptions struct {
		Hostname           string
		Port               string
		User               string
		Password           string
		DBName             string
		TablePrefix        string
		MaxOpenConnections int
		MaxIdleConnections int
		ConnMaxLifetime    time.Duration
		Debug              bool
	}

	type redisOptions struct {
		Host        string
		Port        string
		Password    string
		IdleTimeout time.Duration
		MaxIdle     int
		MaxActive   int
	}

	type testConf struct {
		Name  string
		Mysql mysqlOptions
		Redis redisOptions
	}

	var conf testConf
	testFile := ""
	err := LoadConfig(testFile, &conf)
	if err != nil {
		t.Log(err)
	}

	testFile = "./testdata/app_test.toml"
	err = LoadConfig(testFile, &conf)
	if err != nil {
		t.Error(err)
	}
	t.Logf("conf content: %#v", conf)

	testFile = "./testdata/app_test_no_exist.toml"
	err = LoadConfig(testFile, &conf)
	if err != nil {
		t.Log(err)
	}

	testFile = "./testdata/app_test_bad.toml"
	var conf1 testConf
	err = LoadConfig(testFile, &conf1)
	if err != nil {
		t.Log(err)
	}

	t.Logf("conf content: %#v", conf)

}
