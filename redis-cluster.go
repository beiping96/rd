package rd

import (
	"fmt"

	clusterDriver "github.com/chasex/redis-go-cluster"
)

var _ RD = (*cluster)(nil)

type cluster struct {
	*clusterDriver.Cluster
	cfg Config
}

func (i *cluster) ErrNil() error {
	return clusterDriver.ErrNil
}

func (i *cluster) Do(cmd string, args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		args[0] = i.cfg.Prefix(fmt.Sprintf("%v", args[0]))
	}
	return i.Cluster.Do(cmd, args...)
}

func (i *cluster) DoBool(cmd string, args ...interface{}) (bool, error) {
	return clusterDriver.Bool(i.Do(cmd, args...))
}
func (i *cluster) DoBytes(cmd string, args ...interface{}) ([]byte, error) {
	return clusterDriver.Bytes(i.Do(cmd, args...))
}
func (i *cluster) DoFloat64(cmd string, args ...interface{}) (float64, error) {
	return clusterDriver.Float64(i.Do(cmd, args...))
}
func (i *cluster) DoInt(cmd string, args ...interface{}) (int, error) {
	return clusterDriver.Int(i.Do(cmd, args...))
}
func (i *cluster) DoInt64(cmd string, args ...interface{}) (int64, error) {
	return clusterDriver.Int64(i.Do(cmd, args...))
}
func (i *cluster) DoInts(cmd string, args ...interface{}) ([]int, error) {
	return clusterDriver.Ints(i.Do(cmd, args...))
}
func (i *cluster) DoString(cmd string, args ...interface{}) (string, error) {
	return clusterDriver.String(i.Do(cmd, args...))
}
func (i *cluster) DoStringMap(cmd string, args ...interface{}) (map[string]string, error) {
	return clusterDriver.StringMap(i.Do(cmd, args...))
}
func (i *cluster) DoStrings(cmd string, args ...interface{}) ([]string, error) {
	return clusterDriver.Strings(i.Do(cmd, args...))
}

func (i *cluster) Values(cmd string, args ...interface{}) ([]interface{}, error) {
	return clusterDriver.Values(i.Do(cmd, args...))
}
func (i *cluster) Scan(src []interface{}, dst ...interface{}) ([]interface{}, error) {
	return clusterDriver.Scan(src, dst...)
}

func (i *cluster) construct(cfg Config) error {
	options := new(clusterDriver.Options)
	options.StartNodes = []string{cfg.Address}
	if cfg.DialTimeout > 0 {
		options.ConnTimeout = cfg.DialTimeout
	}
	if cfg.ReadTimeout > 0 {
		options.ReadTimeout = cfg.ReadTimeout
	}
	if cfg.WriteTimeout > 0 {
		options.WriteTimeout = cfg.WriteTimeout
	}
	options.KeepAlive = cfg.OpenConns
	options.AliveTime = cfg.Lifetime
	c, err := clusterDriver.NewCluster(options)
	if err != nil {
		return fmt.Errorf("github.com/chasex/redis-go-cluster NewCluster error %v",
			err)
	}
	*i = cluster{
		Cluster: c,
		cfg:     cfg,
	}
	return nil
}
