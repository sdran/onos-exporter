// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collect

import (
	"github.com/onosproject/onos-exporter/pkg/kpis"
)

// Var defitions of Configuration needed for each collector.
// Multiple collectors can be defined in this file, with their
// respective configs specified here.
var (
	onose2tConfig = InitConfig("onos-e2t")
)

// Onose2tCollector implements the collector of the onos e2t service kpis.
// It uses the function(s) defined in onose2t.go to extract the kpis and return
// a list of them.
// This function can create go routines if needed in order to extract multiple
// onos e2t kpis using the same connection and multiple calls to functions
// defined in the file onose2t.go.
func Onose2tCollector(e2tServiceAddress string) ([]kpis.KPI, error) {
	kpis := []kpis.KPI{}

	err := onose2tConfig.set(map[string]string{addressKey: e2tServiceAddress})
	if err != nil {
		return kpis, err
	}

	conn, err := GetConnection(
		onose2tConfig.getAddress(),
		onose2tConfig.getCertPath(),
		onose2tConfig.getKeyPath(),
		onose2tConfig.noTLS(),
	)
	if err != nil {
		return kpis, err
	}
	defer conn.Close()

	e2tconnectionsKPI, err := onose2tListConnections(conn)
	if err != nil {
		return kpis, err
	}

	kpis = append(kpis, e2tconnectionsKPI)

	return kpis, nil
}
