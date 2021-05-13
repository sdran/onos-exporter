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

// onose2tListConnections implements the extraction of the kpi OnosE2tConnections
// from the component onose2t. It connects to onos e2t service list the e2NodeConnections
// and fill the proper fields of the OnosE2tConnectionsKPI.
// Other functions must be implemented similar to this one in order to extract other
// kpis from onos e2t service.
func onose2tListConnections(conn *grpc.ClientConn) (kpis.KPI, error) {
	OnosE2tConnectionsKPI := kpis.OnosE2tConnections()

	request := adminapi.ListE2NodeConnectionsRequest{}
	client := adminapi.NewE2TAdminServiceClient(conn)
	stream, err := client.ListE2NodeConnections(context.Background(), &request)

	if err != nil {
		return OnosE2tConnectionsKPI, err
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return OnosE2tConnectionsKPI, err
		}

		OnosE2tConnectionsKPI.NumberConnections += 1
	}

	return OnosE2tConnectionsKPI, nil
}
