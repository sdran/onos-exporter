// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collect

import (
	"fmt"
	"strings"

	"github.com/onosproject/onos-exporter/pkg/kpis"
)

// Collector defines an interface for Collectors to retrieve
// a list of kpis.KPI via the Collect method.
type Collector interface {
	Collect() ([]kpis.KPI, error)
}

// KPIs retrieves the list of kpis.KPI from each Collector.
func KPIs(collectors []Collector) ([]kpis.KPI, error) {
	var errstrings []string
	kpis := []kpis.KPI{}

	for _, col := range collectors {
		colKPIs, err := col.Collect()

		if err != nil {
			errstrings = append(errstrings, err.Error())

		} else {
			kpis = append(kpis, colKPIs...)
		}

	}

	if len(errstrings) > 0 {
		err := fmt.Errorf(strings.Join(errstrings, "\n"))
		return kpis, err
	}

	return kpis, nil
}
