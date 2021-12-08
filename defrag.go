package etcd_defrag

import (
	"context"
	"crypto/tls"
	v3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Defrag struct {
	*v3.Client

	Name string
	Version string
	Endpoints []string
}

func NewDefrag(ep []string, cfg *tls.Config) (*Defrag, error) {
	c, err := v3.New(v3.Config{
		Endpoints: ep,
		TLS:cfg,
	})
	if err != nil {
		return nil, err
	}
	return &Defrag{
		Client: c,
		Endpoints: ep,
	}, nil
}

func (d *Defrag) Clean() error {
	for _, endpoint := range d.Endpoints {
		ctx, _ := context.WithTimeout(context.TODO(), time.Duration(5))
		if err := defrag(ctx, d.Client, endpoint); err != nil {
			return err
		}
	}
	return nil
}
