package etcd_defrag

import (
	"crypto/tls"
	"go.etcd.io/etcd/client/pkg/v3/transport"
)

const (
	CertFile = iota
	KeyFile
	TrustedCAFile
	Padding
)

type CA struct {
	File [Padding]string
	TLSInfo *transport.TLSInfo
}

func NewCA(cert, key, trusted string) *CA {
	file := [Padding]string{
		cert, key, trusted,
	}
	TLSInfo := transport.TLSInfo{
		CertFile: file[CertFile],
		KeyFile: file[KeyFile],
		TrustedCAFile: file[TrustedCAFile],
	}
	return &CA{
		file,
		&TLSInfo,
	}
}

func (r *CA) GetConfig () (*tls.Config, error) {
	return r.TLSInfo.ClientConfig()
}
