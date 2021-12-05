// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpis

import "github.com/prometheus/client_golang/prometheus"

// KPI interface defines the methods that format the behavior
// of a kpi. It includes that a kpi must provide those methods
// in order to support its content to be exported to a particular
// TSDB.
type KPI interface {
	PrometheusFormat() ([]prometheus.Metric, error)
}

// Const definitions of kpis name and description.
// Name and description are used to define a particular KPI.
const (
	onosE2tConnectionsKPIName        = "connections"
	onosE2tConnectionsKPIDescription = "The number of e2t connections"

	xappPciNumConflictsKPIName     = "info"
	xappPciNumConflictsDescription = "The xapp pci cell info"

	xappPciResolvedConflictsKPIName     = "conflicts"
	xappPciResolvedConflictsDescription = "The xapp pci resolved cell conflicts"

	xappkpimonKPIName     = "kpm"
	xappkpimonDescription = "The KPM related metrics"

	topoEntitiesKPIName        = "entities"
	topoEntitiesKPIDescription = "The onos topo entities"

	topoRelationsKPIName        = "relations"
	topoRelationsKPIDescription = "The onos topo relations"

	OnosUenibUEsKPIName        = "aspects"
	OnosUenibUEsKPIDescription = "The uenib aspects "
)

// OnosE2tConnections defines the factory implementation of a kpi
// onosE2tConnections having a well defined name and description.
func OnosE2tConnections() *onosE2tConnections {
	return &onosE2tConnections{
		name:        onosE2tConnectionsKPIName,
		description: onosE2tConnectionsKPIDescription,
	}
}

// XappKpiMon defines the factory implementation of a kpi
// onosE2subs having a well defined name and description.
func XappKpiMon() *xappkpimon {
	return &xappkpimon{
		name:        xappkpimonKPIName,
		description: xappkpimonDescription,
	}
}

// XappPciNumConflicts defines the factory implementation of a kpi
// xappPciNumConflicts having a well defined name and description.
func XappPciNumConflicts() *xappPciNumConflicts {
	return &xappPciNumConflicts{
		name:        xappPciNumConflictsKPIName,
		description: xappPciNumConflictsDescription,
	}
}

// XappPciResolvedConflicts defines the factory implementation of a kpi
// xappPciResolvedConflicts having a well defined name and description.
func XappPciResolvedConflicts() *xappPciResolvedConflicts {
	return &xappPciResolvedConflicts{
		name:        xappPciResolvedConflictsKPIName,
		description: xappPciResolvedConflictsDescription,
	}
}

// OnosTopoEntities defines the factory implementation of a kpi
// topoEntities having a well defined name and description.
func OnosTopoEntities() *topoEntities {
	return &topoEntities{
		name:        topoEntitiesKPIName,
		description: topoEntitiesKPIDescription,
	}
}

// OnosTopoRelations defines the factory implementation of a kpi
// topoRelations having a well defined name and description.
func OnosTopoRelations() *topoRelations {
	return &topoRelations{
		name:        topoRelationsKPIName,
		description: topoRelationsKPIDescription,
	}
}

// OnosUenibUEs defines the factory implementation of a kpi
// onosUenibUEs having a well defined name and description.
func OnosUenibUEs() *onosUenibUEs {
	return &onosUenibUEs{
		name:        OnosUenibUEsKPIName,
		description: OnosUenibUEsKPIDescription,
	}
}
