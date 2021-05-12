// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collect

import (
	"context"
	"crypto/tls"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// GetConnection returns a gRPC client connection to the onos service
func GetConnection(address, certPath, keyPath string, noTls bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	if noTls {
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
		}
	} else {
		if certPath != "" && keyPath != "" {
			cert, err := tls.LoadX509KeyPair(certPath, keyPath)
			if err != nil {
				return nil, err
			}
			opts = []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				})),
			}
		} else {
			// Load default Certificates
			cert, err := tls.X509KeyPair([]byte(certs.DefaultClientCrt), []byte(certs.DefaultClientKey))
			if err != nil {
				return nil, err
			}
			opts = []grpc.DialOption{
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				})),
			}
		}
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewContextWithAuthHeaderFromFlag - use from the CLI with --auth-header flag
func NewContextWithAuthHeaderFromFlag(ctx context.Context, authHeaderFlag *pflag.Flag) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if authHeaderFlag != nil && authHeaderFlag.Value != nil && authHeaderFlag.Value.String() != "" {
		md := make(metadata.MD)
		md.Set("authorization", authHeaderFlag.Value.String())
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}
