# onos-exporter
The exporter for ONOS SD-RAN (µONOS Architecture) to scrape, format, and export KPIs to TSDB databases (e.g., Prometheus).

## Overview
The onos-exporter realizes the collection of KPIs from multiple ONOS SD-RAN components via gRPC interfaces, properly label them according to their namespace and subsystem, and turn them available to be pulled (or pushed to) TSDBs. Currently the implementation supports Prometheus.