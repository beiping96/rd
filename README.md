# RD

Redis (Cluster) Driver

[![Build Status](https://travis-ci.com/beiping96/rd.svg?branch=master)](https://travis-ci.com/beiping96/rd)
[![GoDoc](https://godoc.org/github.com/beiping96/rd?status.svg)](https://godoc.org/github.com/beiping96/rd)
[![Go Report Card](https://goreportcard.com/badge/github.com/beiping96/rd)](https://goreportcard.com/report/github.com/beiping96/rd)

## Simple Usage

``` go
package main

import (
    "github.com/beiping96/rd"
)

func main() {
    cfg := rd.Config{
        IsCluster: true,
        Address: "127.0.0.1:6379",
    }
    cli, err := rd.NewClient(cfg)
    if err != nil {
        // handle err
    }

    defer cli.Close()

    _, err := cli.DoString("GET", "not_exist_key")
    if err == cli.ErrNil() {
        // key not exist
    } else if err != nil {
        // handle err
    }
}
```

## With Instance

[Instance Example](https://github.com/beiping96/rd/blob/master/example/redis.go)
