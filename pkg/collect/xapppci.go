// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package collect

import (
	"bytes"
	"context"
	"fmt"

	pciapi "github.com/onosproject/onos-api/go/onos/pci"
	"github.com/onosproject/onos-exporter/pkg/kpis"
	"google.golang.org/grpc"
)

// xappPciCollector is the onos xapp pci collector.
// It extracts all the pci related kpis using the Collect method.
type xappPciCollector struct {
	collector
}

// Collect implements the Collector interface behavior for
// XappPciCollector, returning a list of kpis.KPI.
func (col *xappPciCollector) Collect() ([]kpis.KPI, error) {
	kpis := []kpis.KPI{}

	if len(col.config.getAddress()) == 0 {
		return kpis, fmt.Errorf("XappPciCollector Collect missing service address")
	}

	conn, err := GetConnection(
		col.config.getAddress(),
		col.config.getCertPath(),
		col.config.getKeyPath(),
		col.config.noTLS(),
	)
	if err != nil {
		return kpis, err
	}
	defer conn.Close()

	cellInfoKPI, err := listCellInfo(conn)
	if err != nil {
		return kpis, err
	}

	conflictsKPI, err := listResolvedConflictsAll(conn)
	if err != nil {
		return kpis, err
	}

	kpis = append(kpis, cellInfoKPI)
	kpis = append(kpis, conflictsKPI)

	return kpis, err
}

// listCellInfo receives a connection to a pci xapp service
// to retrieve the pci conflicts and store them according to the
// data structure of the kpis.XappPciNumConflicts KPI.
func listCellInfo(conn *grpc.ClientConn) (kpis.KPI, error) {
	numConflictsKPI := kpis.XappPciNumConflicts()
	numConflictsKPI.Cells = make(map[string]kpis.CellInfo)

	request := pciapi.GetConflictsRequest{}
	client := pciapi.NewPciClient(conn)
	response, err := client.GetConflicts(context.Background(), &request)
	if err != nil {
		return numConflictsKPI, err
	}

	for _, cell := range response.GetCells() {

		cellID := fmt.Sprintf("%x", cell.Id)
		nodeID := cell.NodeId
		cellType := cell.CellType.String()
		cellPci := fmt.Sprintf("%d", cell.Pci)

		cellDlearfcn := float64(cell.Dlearfcn)

		neighbors := neighborsAsCSV(cell)

		cInfo := kpis.CellInfo{
			CellID:        cellID,
			NodeID:        nodeID,
			CellType:      cellType,
			CellPci:       cellPci,
			CellDlearfcn:  cellDlearfcn,
			CellNeighbors: neighbors,
		}
		numConflictsKPI.Cells[cellID] = cInfo
	}

	return numConflictsKPI, nil
}

func neighborsAsCSV(cell *pciapi.PciCell) string {
	var buffer bytes.Buffer
	first := true
	for _, neighbor := range cell.NeighborIds {
		if !first {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("%x", neighbor))
		first = false
	}
	return buffer.String()
}

// listNumConflictsAll receives a connection to a pci xapp service
// to retrieve the pci conflicts and store them according to the
// data structure of the kpis.XappPciNumConflicts KPI.
func listResolvedConflictsAll(conn *grpc.ClientConn) (kpis.KPI, error) {
	resolvedConflictsKPI := kpis.XappPciResolvedConflicts()
	resolvedConflictsKPI.Cells = make(map[string]kpis.CellConflict)

	request := pciapi.GetResolvedConflictsRequest{}
	client := pciapi.NewPciClient(conn)
	response, err := client.GetResolvedConflicts(context.Background(), &request)
	if err != nil {
		return resolvedConflictsKPI, err
	}

	for _, cell := range response.GetCells() {
		cellID := fmt.Sprintf("%x", cell.Id)

		cInfo := kpis.CellConflict{
			CellID:            cellID,
			OriginalPci:       fmt.Sprintf("%d", cell.OriginalPci),
			ResolvedPci:       fmt.Sprintf("%d", cell.ResolvedPci),
			ResolvedConflicts: float64(cell.ResolvedConflicts),
		}
		resolvedConflictsKPI.Cells[cellID] = cInfo
	}

	return resolvedConflictsKPI, nil
}
