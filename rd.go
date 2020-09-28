package rd

import (
	"math"
	"time"

	"github.com/pkg/errors"
)

// RD declare redis client method
type RD interface {
	// ErrNil declare NIL-error
	ErrNil() error

	// Do execute redis command
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
}

// Config desc redis client method
type Config struct {
	// if is redis-cluster
	IsCluster bool

	// required redis addr
	Address string
	// if has password
	Password string

	// default is 10 seconds
	DialTimeout time.Duration
	// default is 10 seconds
	ReadTimeout time.Duration
	// default is 10 seconds
	WriteTimeout time.Duration

	// default is 20
	OpenConns int
	// default is keep alive
	Lifetime time.Duration

	// Prefix will wrap redis-key at head
	// default is nil
	Prefix func(key string) (keyWithPrefix string)

	// DB select redis database
	// default is nil
	// useless in cluster
	DB int
}

// New redis client with config
func New(cfg Config) (cli RD, err error) {
	if len(cfg.Address) == 0 {
		return nil, errors.Errorf("RD Config NIL Address")
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

	if cfg.OpenConns <= 0 {
		cfg.OpenConns = 20
	}
	if cfg.Lifetime <= 0 {
		cfg.Lifetime = time.Duration(math.MaxInt16) * time.Hour
	}

	if cfg.Prefix == nil {
		cfg.Prefix = func(key string) (keyWithPrefix string) { return key }
	}

	if cfg.DB < 0 {
		cfg.DB = 0
	}

	if cfg.IsCluster {
		clusterCli := new(cluster)
		if err := clusterCli.construct(cfg); err != nil {
			return nil, errors.Wrapf(err, "RD Construct Cluster %+v",
				cfg)
		}
		cli = clusterCli
	} else {
		redisCli := new(redis)
		if err := redisCli.construct(cfg); err != nil {
			return nil, errors.Wrapf(err, "RD Construct Redis %+v",
				cfg)
		}
		cli = redisCli
	}

	return cli, nil
}
