// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collect

import (
	"context"
	"io"

	adminapi "github.com/onosproject/onos-api/go/onos/e2t/admin"
	"github.com/onosproject/onos-exporter/pkg/kpis"
	"google.golang.org/grpc"
)

func onose2tListConnections(conn *grpc.ClientConn) (*kpis.OnosE2tConnections, error) {
	e2tconnections := kpis.NewOnosE2tConnections()

	request := adminapi.ListE2NodeConnectionsRequest{}
	client := adminapi.NewE2TAdminServiceClient(conn)
	stream, err := client.ListE2NodeConnections(context.Background(), &request)

	if err != nil {
		return e2tconnections, err
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return e2tconnections, err
		}

		e2tconnections.NumberConnections += 1
	}

	return e2tconnections, nil
}
