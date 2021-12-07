package etcd_defrag

import (
	"context"
	v3 "go.etcd.io/etcd/client/v3"
)

func defrag(ctx context.Context, client *v3.Client, endpoint string) error {
	status, err := client.Status(ctx, endpoint)
	if err != nil {
		return err
	}
	_, err = client.Compact(ctx, status.Header.Revision)
	if err != nil {
		return err
	}
	_, err = client.Defragment(ctx, endpoint)
	if err != nil {
		return err
	}
	return nil
}
