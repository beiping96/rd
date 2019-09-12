package example

import (
	"context"
	"fmt"
	"runtime"

	"github.com/beiping96/grace"
	"github.com/beiping96/rd"
)

type (
	Client = rd.RD
	Config = rd.Config
)

func NewClient(cfg Config) (cli Client, err error) {
	for i := 0; i < 5; i++ {
		cli, err = rd.New(cfg)
		if err == nil {
			return cli, nil
		}
	}
	return nil, fmt.Errorf("Redis New Client Error %v %+v",
		err, cfg)
}

var (
	_           Client = (*ConstClient)(nil)
	constClient *ConstClient
)

type ConstClient struct{ Client }

func GetConstClient() (Client, error) {
	if constClient == nil {
		return nil, fmt.Errorf("Redis Const Client Needs Construct Firstly")
	}
	return constClient, nil
}

// Do Nothing
func (constCli *ConstClient) Close() {}

func Construct(cfg Config) {
	if constClient != nil {
		return
	}
	cli, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}
	constClient = &ConstClient{Client: cli}
	grace.Go(connCloseMonitor)
}

func connCloseMonitor(ctx context.Context) {
	<-ctx.Done()
	runtime.Gosched()
	constClient.Client.Close()
}
