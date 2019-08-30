package rd

import (
	"fmt"
	"math"
	"time"
)

type RD interface {
	ErrNil() error

	Do(cmd string, args ...interface{}) (interface{}, error)

	DoBool(cmd string, args ...interface{}) (bool, error)
	DoBytes(cmd string, args ...interface{}) ([]byte, error)
	DoFloat64(cmd string, args ...interface{}) (float64, error)
	DoInt(cmd string, args ...interface{}) (int, error)
	DoInt64(cmd string, args ...interface{}) (int64, error)
	DoInts(cmd string, args ...interface{}) ([]int, error)
	DoString(cmd string, args ...interface{}) (string, error)
	DoStringMap(cmd string, args ...interface{}) (map[string]string, error)
	DoStrings(cmd string, args ...interface{}) ([]string, error)

	Values(cmd string, args ...interface{}) ([]interface{}, error)
	Scan(src []interface{}, dst ...interface{}) ([]interface{}, error)

	Close()

	construct(Config) error
}

type Config struct {
	IsCluster bool

	Address  string
	Password string

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	MaxOpenConns int
	MaxLifetime  time.Duration

	Prefix func(key string) (keyWithPrefix string)
}

func New(cfg Config) (cli RD, err error) {
	if len(cfg.Address) == 0 {
		return nil, fmt.Errorf("RD Config NIL Address")
	}

	if cfg.DialTimeout <= 0 {
		cfg.DialTimeout = time.Duration(10) * time.Second
	}
	if cfg.ReadTimeout <= 0 {
		cfg.ReadTimeout = time.Duration(10) * time.Second
	}
	if cfg.WriteTimeout <= 0 {
		cfg.WriteTimeout = time.Duration(10) * time.Second
	}

	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = 20
	}
	if cfg.MaxLifetime <= 0 {
		cfg.MaxLifetime = time.Duration(math.MaxInt16) * time.Hour
	}

	if cfg.Prefix == nil {
		cfg.Prefix = func(key string) (keyWithPrefix string) { return key }
	}

	if cfg.IsCluster {
		cli = new(cluster)
	} else {
		cli = new(redis)
	}
	if err := cli.construct(cfg); err != nil {
		return nil, fmt.Errorf("RD Construct Error %v %+v",
			err, cfg)
	}
	return cli, nil
}
