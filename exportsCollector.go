package main

import (
	"github.com/Gandi/ganesha_exporter/dbus"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/alecthomas/kingpin.v2"
	"strconv"
)

var (
	nfsV3RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_requested_bytes_total",
		"Number of requested bytes for NFSv3 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3TransferredDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_transferred_bytes_total",
		"Number of transferred bytes for NFSv3 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_total",
		"Number of operations for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_errors_total",
		"Number of operations in error for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_requested_bytes_total",
		"Number of requested bytes for NFSv4.0 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40TransferredDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_transferred_bytes_total",
		"Number of transferred bytes for NFSv4.0 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_total",
		"Number of operations for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_errors_total",
		"Number of operations in error for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_requested_bytes_total",
		"Number of requested bytes for NFSv4.1 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41TransferredDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_transferred_bytes_total",
		"Number of transferred bytes for NFSv4.1 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_total",
		"Number of operations for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_errors_total",
		"Number of operations in error for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	pnfsV41LayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v41_layout_operations_total",
		"Numer of layout operations for pNFSv4.1",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsV41LayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v41_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.1",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsV41LayoutDelayDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v41_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_requested_bytes_total",
		"Number of requested bytes for NFSv4.2 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42TransferredDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_transferred_bytes_total",
		"Number of transferred bytes for NFSv4.2 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_total",
		"Number of operations for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_errors_total",
		"Number of operations in error for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	pnfsV42LayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v42_layout_operations_total",
		"Numer of layout operations for pNFSv4.2",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsV42LayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v42_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.2",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsV42LayoutDelayDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v42_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
)

// ExportsCollector Collector for ganesha exports
type ExportsCollector struct {
	exportMgr                                       dbus.ExportMgr
	nfsv3, nfsv40, nfsv41, pnfsv41, nfsv42, pnfsv42 *bool
}

// NewExportsCollector creates a new collector
func NewExportsCollector() ExportsCollector {
	return ExportsCollector{
		nfsv3:     kingpin.Flag("collector.exports.nfsv3", "Activate NFSv3 stats").Default("true").Bool(),
		nfsv40:    kingpin.Flag("collector.exports.nfsv40", "Activate NFSv4.0 stats").Default("true").Bool(),
		nfsv41:    kingpin.Flag("collector.exports.nfsv41", "Activate NFSv4.1 stats").Default("true").Bool(),
		pnfsv41:   kingpin.Flag("collector.exports.pnfsv41", "Activate pNFSv4.1 stats").Default("true").Bool(),
		nfsv42:    kingpin.Flag("collector.exports.nfsv42", "Activate NFSv4.2 stats").Default("true").Bool(),
		pnfsv42:   kingpin.Flag("collector.exports.pnfsv42", "Activate pNFSv4.2 stats").Default("true").Bool(),
	}
}

func (ic *ExportsCollector) InitDBus() {
    ic.exportMgr = dbus.NewExportMgr()
}

// Describe prometheus description
func (ic ExportsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(ic, ch)
}

func (ic ExportsCollector) exportNFSStatsIO(ch chan<- prometheus.Metric, export dbus.Export, stats dbus.BasicStats,
	nfsRequestedDesc *prometheus.Desc,
	nfsTransferredDesc *prometheus.Desc,
	nfsOperationsDesc *prometheus.Desc,
	nfsErrorsDesc *prometheus.Desc,
	nfsLatencyDesc *prometheus.Desc,
	nfsQueueWaitDesc *prometheus.Desc,
) {
	exportid := strconv.FormatUint(uint64(export.ExportID), 10)
	path := export.Path
	ch <- prometheus.MustNewConstMetric(
		nfsRequestedDesc,
		prometheus.CounterValue,
		float64(stats.Read.Requested),
		"read", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsTransferredDesc,
		prometheus.CounterValue,
		float64(stats.Read.Transferred),
		"read", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsOperationsDesc,
		prometheus.CounterValue,
		float64(stats.Read.Total),
		"read", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsErrorsDesc,
		prometheus.CounterValue,
		float64(stats.Read.Errors),
		"read", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsLatencyDesc,
		prometheus.CounterValue,
		float64(stats.Read.Latency)/1e9,
		"read", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsQueueWaitDesc,
		prometheus.CounterValue,
		float64(stats.Read.QueueWait)/1e9,
		"read", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsRequestedDesc,
		prometheus.CounterValue,
		float64(stats.Write.Requested),
		"write", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsTransferredDesc,
		prometheus.CounterValue,
		float64(stats.Write.Transferred),
		"write", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsOperationsDesc,
		prometheus.CounterValue,
		float64(stats.Write.Total),
		"write", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsErrorsDesc,
		prometheus.CounterValue,
		float64(stats.Write.Errors),
		"write", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsLatencyDesc,
		prometheus.CounterValue,
		float64(stats.Write.Latency)/1e9,
		"write", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		nfsQueueWaitDesc,
		prometheus.CounterValue,
		float64(stats.Write.QueueWait)/1e9,
		"write", exportid, path)
}

func (ic ExportsCollector) exportNFSStatsLayouts(ch chan<- prometheus.Metric, export dbus.Export, stats dbus.PNFSOperations,
	pnfsLayoutOperationsDesc *prometheus.Desc,
	pnfsLayoutErrorsDesc *prometheus.Desc,
	pnfsLayoutDelayDesc *prometheus.Desc,
) {
	exportid := strconv.FormatUint(uint64(export.ExportID), 10)
	path := export.Path
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.Getdevinfo.Total),
		"getdevinfo", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.Getdevinfo.Errors),
		"getdevinfo", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.Getdevinfo.Delays)/1e9,
		"getdevinfo", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutGet.Total),
		"get", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutGet.Errors),
		"get", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutGet.Delays)/1e9,
		"get", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutCommit.Total),
		"commit", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutCommit.Errors),
		"commit", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutCommit.Delays)/1e9,
		"commit", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutReturn.Total),
		"return", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutReturn.Errors),
		"return", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutReturn.Delays)/1e9,
		"return", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutOperationsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutRecall.Total),
		"recall", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutErrorsDesc,
		prometheus.CounterValue,
		float64(stats.LayoutRecall.Errors),
		"recall", exportid, path)
	ch <- prometheus.MustNewConstMetric(
		pnfsLayoutDelayDesc,
		prometheus.CounterValue,
		float64(stats.LayoutRecall.Delays)/1e9,
		"recall", exportid, path)
}

func (ic ExportsCollector) collectNFSv3IO(ch chan<- prometheus.Metric, export dbus.Export) bool {
	var stats dbus.BasicStats
	if export.NFSv3 {
		stats = ic.exportMgr.GetNFSv3IO(export.ExportID)
	}
	ic.exportNFSStatsIO(ch, export, stats,
		nfsV3RequestedDesc,
		nfsV3TransferredDesc,
		nfsV3OperationsDesc,
		nfsV3ErrorsDesc,
		nfsV3LatencyDesc,
		nfsV3QueueWaitDesc,
	)
	return true
}

func (ic ExportsCollector) collectNFSv40IO(ch chan<- prometheus.Metric, export dbus.Export) bool {
	stats := dbus.BasicStats{}
	if export.NFSv40 {
		stats = ic.exportMgr.GetNFSv40IO(export.ExportID)
	}
	ic.exportNFSStatsIO(ch, export, stats,
		nfsV40RequestedDesc,
		nfsV40TransferredDesc,
		nfsV40OperationsDesc,
		nfsV40ErrorsDesc,
		nfsV40LatencyDesc,
		nfsV40QueueWaitDesc,
	)
	return true
}

func (ic ExportsCollector) collectNFSv41IO(ch chan<- prometheus.Metric, export dbus.Export) bool {
	stats := dbus.BasicStats{}
	if export.NFSv41 {
		stats = ic.exportMgr.GetNFSv41IO(export.ExportID)
	}
	ic.exportNFSStatsIO(ch, export, stats,
		nfsV41RequestedDesc,
		nfsV41TransferredDesc,
		nfsV41OperationsDesc,
		nfsV41ErrorsDesc,
		nfsV41LatencyDesc,
		nfsV41QueueWaitDesc,
	)
	return true
}

func (ic ExportsCollector) collectNFSv41Layouts(ch chan<- prometheus.Metric, export dbus.Export) bool {
	stats := dbus.PNFSOperations{}
	if export.NFSv41 {
		stats = ic.exportMgr.GetNFSv41Layouts(export.ExportID)
	}
	ic.exportNFSStatsLayouts(ch, export, stats,
		pnfsV41LayoutOperationsDesc,
		pnfsV41LayoutErrorsDesc,
		pnfsV41LayoutDelayDesc,
	)
	return true
}

func (ic ExportsCollector) collectNFSv42IO(ch chan<- prometheus.Metric, export dbus.Export) bool {
	stats := dbus.BasicStats{}
	if export.NFSv42 {
		stats = ic.exportMgr.GetNFSv42IO(export.ExportID)
	}
	ic.exportNFSStatsIO(ch, export, stats,
		nfsV42RequestedDesc,
		nfsV42TransferredDesc,
		nfsV42OperationsDesc,
		nfsV42ErrorsDesc,
		nfsV42LatencyDesc,
		nfsV42QueueWaitDesc,
	)
	return true
}

func (ic ExportsCollector) collectNFSv42Layouts(ch chan<- prometheus.Metric, export dbus.Export) bool {
	stats := dbus.PNFSOperations{}
	if export.NFSv42 {
		stats = ic.exportMgr.GetNFSv42Layouts(export.ExportID)
	}
	ic.exportNFSStatsLayouts(ch, export, stats,
		pnfsV42LayoutOperationsDesc,
		pnfsV42LayoutErrorsDesc,
		pnfsV42LayoutDelayDesc,
	)
	return true
}

// Collect do the actual job
func (ic ExportsCollector) Collect(ch chan<- prometheus.Metric) {
	ok := true
	_, exports := ic.exportMgr.ShowExports()
	for _, export := range exports {
		if *ic.nfsv3 {
			ok = ic.collectNFSv3IO(ch, export) && ok
		}
		if *ic.nfsv40 {
			ok = ic.collectNFSv40IO(ch, export) && ok
		}
		if *ic.nfsv41 {
			ok = ic.collectNFSv41IO(ch, export) && ok
		}
		if *ic.pnfsv41 {
			ok = ic.collectNFSv41Layouts(ch, export) && ok
		}
		if *ic.nfsv42 {
			ok = ic.collectNFSv42IO(ch, export) && ok
		}
		if *ic.pnfsv42 {
			ok = ic.collectNFSv42Layouts(ch, export) && ok
		}
	}
}
