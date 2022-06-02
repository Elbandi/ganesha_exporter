package main

import (
	"github.com/Gandi/ganesha_exporter/dbus"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clientsNfsV3RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_requested_bytes_total",
		"Number of requested bytes for NFSv3 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3TransferredDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_transferred_bytes_total",
		"Number of transferred bytes for NFSv3 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_total",
		"Number of operations for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_errors_total",
		"Number of operations in error for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV3QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v3_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv3",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_requested_bytes_total",
		"Number of requested bytes for NFSv4.0 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40TransferredDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_transferred_bytes_total",
		"Number of transferred bytes for NFSv4.0 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_total",
		"Number of operations for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_errors_total",
		"Number of operations in error for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV40QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v40_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.0",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_requested_bytes_total",
		"Number of requested bytes for NFSv4.1 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41TransferredDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_transferred_bytes_total",
		"Number of transferred bytes for NFSv4.1 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_total",
		"Number of operations for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_errors_total",
		"Number of operations in error for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV41QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v41_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsPnfsV41LayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v41_layout_operations_total",
		"Numer of layout operations for pNFSv4.1",
		[]string{"type", "clientip"}, nil,
	)
	clientsPnfsV41LayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v41_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.1",
		[]string{"type", "clientip"}, nil,
	)
	clientsPnfsV41LayoutDelayDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v41_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.1",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV42RequestedDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v42_requested_bytes_total",
		"Number of requested bytes for NFSv4.2 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV42TransferredDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v42_transferred_bytes_total",
		"Number of transferred bytes for NFSv4.2 operations",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV42OperationsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v42_operations_total",
		"Number of operations for NFSv4.2",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV42ErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v42_operations_errors_total",
		"Number of operations in error for NFSv4.2",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV42LatencyDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v42_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.2",
		[]string{"direction", "clientip"}, nil,
	)
	clientsNfsV42QueueWaitDesc = prometheus.NewDesc(
		"ganesha_clients_nfs_v42_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.2",
		[]string{"direction", "clientip"}, nil,
	)
	clientsPnfsV42LayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v42_layout_operations_total",
		"Numer of layout operations for pNFSv4.2",
		[]string{"type", "clientip"}, nil,
	)
	clientsPnfsV42LayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v42_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.2",
		[]string{"type", "clientip"}, nil,
	)
	clientsPnfsV42LayoutDelayDesc = prometheus.NewDesc(
		"ganesha_clients_pnfs_v42_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.2",
		[]string{"direction", "clientip"}, nil,
	)
)

// ClientsCollector Collector for ganesha clients
type ClientsCollector struct {
	clientMgr                                       dbus.ClientMgr
	nfsv3, nfsv40, nfsv41, pnfsv41, nfsv42, pnfsv42 *bool
}

// NewClientsCollector creates a new collector
func NewClientsCollector() ClientsCollector {
	return ClientsCollector{
		clientMgr: dbus.NewClientMgr(),
		nfsv3:     kingpin.Flag("collector.clients.nfsv3", "Activate NFSv3 stats").Default("true").Bool(),
		nfsv40:    kingpin.Flag("collector.clients.nfsv40", "Activate NFSv4.0 stats").Default("true").Bool(),
		nfsv41:    kingpin.Flag("collector.clients.nfsv41", "Activate NFSv4.1 stats").Default("true").Bool(),
		pnfsv41:   kingpin.Flag("collector.clients.pnfsv41", "Activate pNFSv4.1 stats").Default("true").Bool(),
		nfsv42:    kingpin.Flag("collector.clients.nfsv42", "Activate NFSv4.2 stats").Default("true").Bool(),
		pnfsv42:   kingpin.Flag("collector.clients.pnfsv42", "Activate pNFSv4.2 stats").Default("true").Bool(),
	}
}

// Describe prometheus description
func (ic ClientsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(ic, ch)
}

func (ic ClientsCollector) exportNFSStatsIO(ch chan<- prometheus.Metric, client dbus.Client, stats dbus.BasicStats,
	clientsNfsRequestedDesc *prometheus.Desc,
	clientsNfsTransferredDesc *prometheus.Desc,
	clientsNfsOperationsDesc *prometheus.Desc,
	clientsNfsErrorsDesc *prometheus.Desc,
	clientsNfsLatencyDesc *prometheus.Desc,
	clientsNfsQueueWaitDesc *prometheus.Desc,
) {
	clientip := client.Client
	ch <- prometheus.MustNewConstMetric(
		clientsNfsRequestedDesc,
		prometheus.CounterValue,
		float64(stats.Read.Requested),
		"read", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsTransferredDesc,
		prometheus.CounterValue,
		float64(stats.Read.Transferred),
		"read", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsOperationsDesc,
		prometheus.CounterValue,
		float64(stats.Read.Total),
		"read", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsErrorsDesc,
		prometheus.CounterValue,
		float64(stats.Read.Errors),
		"read", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsLatencyDesc,
		prometheus.CounterValue,
		float64(stats.Read.Latency)/1e9,
		"read", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsQueueWaitDesc,
		prometheus.CounterValue,
		float64(stats.Read.QueueWait)/1e9,
		"read", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsRequestedDesc,
		prometheus.CounterValue,
		float64(stats.Write.Requested),
		"write", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsTransferredDesc,
		prometheus.CounterValue,
		float64(stats.Write.Transferred),
		"write", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsOperationsDesc,
		prometheus.CounterValue,
		float64(stats.Write.Total),
		"write", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsErrorsDesc,
		prometheus.CounterValue,
		float64(stats.Write.Errors),
		"write", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsLatencyDesc,
		prometheus.CounterValue,
		float64(stats.Write.Latency)/1e9,
		"write", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsNfsQueueWaitDesc,
		prometheus.CounterValue,
		float64(stats.Write.QueueWait)/1e9,
		"write", clientip)
}

func (ic ClientsCollector) exportNFSStatsLayouts(ch chan<- prometheus.Metric, client dbus.Client, stats dbus.PNFSOperations,
	clientsPnfsLayoutOperationsDesc *prometheus.Desc,
	clientsPnfsLayoutErrorsDesc *prometheus.Desc,
	clientsPnfsLayoutDelayDesc *prometheus.Desc,
) {
	clientip := client.Client
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.Getdevinfo.Total),
		"getdevinfo", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.Getdevinfo.Errors),
		"getdevinfo", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.Getdevinfo.Delays)/1e9,
		"getdevinfo", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutGet.Total),
		"get", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutGet.Errors),
		"get", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutGet.Delays)/1e9,
		"get", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutCommit.Total),
		"commit", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutCommit.Errors),
		"commit", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutCommit.Delays)/1e9,
		"commit", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutReturn.Total),
		"return", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutReturn.Errors),
		"return", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutReturn.Delays)/1e9,
		"return", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutRecall.Total),
		"recall", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutRecall.Errors),
		"recall", clientip)
	ch <- prometheus.MustNewConstMetric(
		clientsPnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutRecall.Delays)/1e9,
		"recall", clientip)
}

func (ic ClientsCollector) collectNFSv3IO(ch chan<- prometheus.Metric, client dbus.Client) bool {
	var stats dbus.BasicStats
	if client.NFSv3 {
		stats = ic.clientMgr.GetNFSv3IO(client.Client)
	}
	ic.exportNFSStatsIO(ch, client, stats,
		clientsNfsV3RequestedDesc,
		clientsNfsV3TransferredDesc,
		clientsNfsV3OperationsDesc,
		clientsNfsV3ErrorsDesc,
		clientsNfsV3LatencyDesc,
		clientsNfsV3QueueWaitDesc,
	)
	return true
}

func (ic ClientsCollector) collectNFSv40IO(ch chan<- prometheus.Metric, client dbus.Client) bool {
	stats := dbus.BasicStats{}
	if client.NFSv40 {
		stats = ic.clientMgr.GetNFSv40IO(client.Client)
	}
	ic.exportNFSStatsIO(ch, client, stats,
		clientsNfsV40RequestedDesc,
		clientsNfsV40TransferredDesc,
		clientsNfsV40OperationsDesc,
		clientsNfsV40ErrorsDesc,
		clientsNfsV40LatencyDesc,
		clientsNfsV40QueueWaitDesc,
	)
	return true
}

func (ic ClientsCollector) collectNFSv41IO(ch chan<- prometheus.Metric, client dbus.Client) bool {
	stats := dbus.BasicStats{}
	if client.NFSv41 {
		stats = ic.clientMgr.GetNFSv41IO(client.Client)
	}
	ic.exportNFSStatsIO(ch, client, stats,
		clientsNfsV41RequestedDesc,
		clientsNfsV41TransferredDesc,
		clientsNfsV41OperationsDesc,
		clientsNfsV41ErrorsDesc,
		clientsNfsV41LatencyDesc,
		clientsNfsV41QueueWaitDesc,
	)
	return true
}

func (ic ClientsCollector) collectNFSv41Layouts(ch chan<- prometheus.Metric, client dbus.Client) bool {
	stats := dbus.PNFSOperations{}
	if client.NFSv41 {
		stats = ic.clientMgr.GetNFSv41Layouts(client.Client)
	}
	ic.exportNFSStatsLayouts(ch, client, stats,
		clientsPnfsV41LayoutOperationsDesc,
		clientsPnfsV41LayoutErrorsDesc,
		clientsPnfsV41LayoutDelayDesc,
	)
	return true
}

func (ic ClientsCollector) collectNFSv42IO(ch chan<- prometheus.Metric, client dbus.Client) bool {
	stats := dbus.BasicStats{}
	if client.NFSv42 {
		stats = ic.clientMgr.GetNFSv42IO(client.Client)
	}
	ic.exportNFSStatsIO(ch, client, stats,
		clientsNfsV42RequestedDesc,
		clientsNfsV42TransferredDesc,
		clientsNfsV42OperationsDesc,
		clientsNfsV42ErrorsDesc,
		clientsNfsV42LatencyDesc,
		clientsNfsV42QueueWaitDesc,
	)
	return true
}

func (ic ClientsCollector) collectNFSv42Layouts(ch chan<- prometheus.Metric, client dbus.Client) bool {
	stats := dbus.PNFSOperations{}
	if client.NFSv42 {
		stats = ic.clientMgr.GetNFSv42Layouts(client.Client)
	}
	ic.exportNFSStatsLayouts(ch, client, stats,
		clientsPnfsV42LayoutOperationsDesc,
		clientsPnfsV42LayoutErrorsDesc,
		clientsPnfsV42LayoutDelayDesc,
	)
	return true
}

// Collect do the actual job
func (ic ClientsCollector) Collect(ch chan<- prometheus.Metric) {
	ok := true
	_, clients := ic.clientMgr.ShowClients()
	for _, client := range clients {
		if *ic.nfsv3 {
			ok = ic.collectNFSv3IO(ch, client) && ok
		}
		if *ic.nfsv40 {
			ok = ic.collectNFSv40IO(ch, client) && ok
		}
		if *ic.nfsv41 {
			ok = ic.collectNFSv41IO(ch, client) && ok
		}
		if *ic.pnfsv41 {
			ok = ic.collectNFSv41Layouts(ch, client) && ok
		}
		if *ic.nfsv42 {
			ok = ic.collectNFSv42IO(ch, client) && ok
		}
		if *ic.pnfsv42 {
			ok = ic.collectNFSv42Layouts(ch, client) && ok
		}
	}
}
