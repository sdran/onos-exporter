// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collect

import (
	"context"
	"strconv"

	pciapi "github.com/onosproject/onos-api/go/onos/pci"
	"github.com/onosproject/onos-exporter/pkg/kpis"
	"google.golang.org/grpc"
)

var (
	xappPciConfig = InitConfig("onos-xappPci")
)

// XappPciCollector is the onos xapp pci collector.
// It extracts all the pci related kpis using the Collect method.
type XappPciCollector struct {
	XappPciServiceAddress string
}

// Collect implements the Collector interface behavior for
// XappPciCollector, returning a list of kpis.KPI.
func (col XappPciCollector) Collect() ([]kpis.KPI, error) {
	kpis := []kpis.KPI{}

	err := xappPciConfig.set(map[string]string{addressKey: col.XappPciServiceAddress})
	if err != nil {
		return kpis, err
	}

	conn, err := GetConnection(
		xappPciConfig.getAddress(),
		xappPciConfig.getCertPath(),
		xappPciConfig.getKeyPath(),
		xappPciConfig.noTLS(),
	)
	if err != nil {
		return kpis, err
	}
	defer conn.Close()

	numConflictsKPI, err := listNumConflictsAll(conn)
	if err != nil {
		return kpis, err
	}

	kpis = append(kpis, numConflictsKPI)

	return kpis, err
}

// listNumConflictsAll receives a connection to a pci xapp service
// to retrieve the pci conflicts and store them according to the
// data structure of the kpis.XappPciNumConflicts KPI.
func listNumConflictsAll(conn *grpc.ClientConn) (kpis.KPI, error) {
	numConflictsKPI := kpis.XappPciNumConflicts()
	numConflictsKPI.NumberConflicts = make(map[string]float64)

	request := pciapi.GetRequest{
		Id: "pci",
	}

	client := pciapi.NewPciClient(conn)

	response, err := client.GetNumConflictsAll(context.Background(), &request)

	if err != nil {
		return numConflictsKPI, err
	}

	for cellID, numConflictsStr := range response.GetObject().GetAttributes() {
		numConflicts, err := strconv.ParseInt(numConflictsStr, 0, 64)
		if err != nil {
			return numConflictsKPI, err
		}
		numConflictsKPI.NumberConflicts[cellID] = float64(numConflicts)
	}

	return numConflictsKPI, nil
}
