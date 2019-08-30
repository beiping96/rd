package rd

import (
	"context"
	"fmt"

	redisDriver "github.com/gomodule/redigo/redis"
)

var _ RD = (*redis)(nil)

type redis struct {
	*redisDriver.Pool
	cfg Config
}

func (i *redis) ErrNil() error {
	return redisDriver.ErrNil
}

func (i *redis) Do(cmd string, args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		args[0] = i.cfg.Prefix(fmt.Sprintf("%v", args[0]))
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		i.cfg.DialTimeout)
	defer cancel()
	conn, err := i.Pool.GetContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("RD %v",
			err)
	}
	defer conn.Close()
	return conn.Do(cmd, args...)
}

func (i *redis) DoBool(cmd string, args ...interface{}) (bool, error) {
	return redisDriver.Bool(i.Do(cmd, args...))
}
func (i *redis) DoBytes(cmd string, args ...interface{}) ([]byte, error) {
	return redisDriver.Bytes(i.Do(cmd, args...))
}
func (i *redis) DoFloat64(cmd string, args ...interface{}) (float64, error) {
	return redisDriver.Float64(i.Do(cmd, args...))
}
func (i *redis) DoInt(cmd string, args ...interface{}) (int, error) {
	return redisDriver.Int(i.Do(cmd, args...))
}
func (i *redis) DoInt64(cmd string, args ...interface{}) (int64, error) {
	return redisDriver.Int64(i.Do(cmd, args...))
}
func (i *redis) DoInts(cmd string, args ...interface{}) ([]int, error) {
	return redisDriver.Ints(i.Do(cmd, args...))
}
func (i *redis) DoString(cmd string, args ...interface{}) (string, error) {
	return redisDriver.String(i.Do(cmd, args...))
}
func (i *redis) DoStringMap(cmd string, args ...interface{}) (map[string]string, error) {
	return redisDriver.StringMap(i.Do(cmd, args...))
}
func (i *redis) DoStrings(cmd string, args ...interface{}) ([]string, error) {
	return redisDriver.Strings(i.Do(cmd, args...))
}

func (i *redis) Values(cmd string, args ...interface{}) ([]interface{}, error) {
	return redisDriver.Values(i.Do(cmd, args...))
}
func (i *redis) Scan(src []interface{}, dst ...interface{}) ([]interface{}, error) {
	return redisDriver.Scan(src, dst...)
}

func (i *redis) Close() { i.Pool.Close() }

func (i *redis) construct(cfg Config) error {
	options := []redisDriver.DialOption{
		redisDriver.DialKeepAlive(cfg.Lifetime),
	}
	if len(cfg.Password) > 0 {
		options = append(options, redisDriver.DialPassword(cfg.Password))
	}
	if cfg.DialTimeout > 0 {
		options = append(options, redisDriver.DialConnectTimeout(cfg.DialTimeout))
	}
	if cfg.ReadTimeout > 0 {
		options = append(options, redisDriver.DialReadTimeout(cfg.ReadTimeout))
	}
	if cfg.WriteTimeout > 0 {
		options = append(options, redisDriver.DialWriteTimeout(cfg.WriteTimeout))
	}
	dial := func() (redisDriver.Conn, error) {
		return redisDriver.Dial("tcp", cfg.Address, options...)
	}
	*i = redis{
		Pool: redisDriver.NewPool(dial, cfg.OpenConns),
		cfg:  cfg,
	}
	return nil
}
