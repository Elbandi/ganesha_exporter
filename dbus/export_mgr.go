package dbus

import (
	"github.com/godbus/dbus"
	"golang.org/x/sys/unix"
	"log"
)

// Export Structure of the output of ShowExports dbus call
type Export struct {
	ExportID uint32
	Path     string
	NFSv3    bool
	MNTv3    bool
	NLMv4    bool
	RQUOTA   bool
	NFSv40   bool
	NFSv41   bool
	NFSv42   bool
	Plan9    bool
	LastTime unix.Timespec
}

// ExportMgr is a handle to dbus object ExportMgr
type ExportMgr struct {
	dbusObject dbus.BusObject
}

// NewExportMgr Get a new ExportMgr
func NewExportMgr() ExportMgr {
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Panic(err)
	}
	return ExportMgr{
		dbusObject: conn.Object(
			"org.ganesha.nfsd",
			"/org/ganesha/nfsd/ExportMgr",
		),
	}
}

func (mgr ExportMgr) ShowExports() (unix.Timespec, []Export) {
	var exports []Export
	utime := unix.Timespec{}
	err := mgr.dbusObject.
		Call("org.ganesha.nfsd.exportmgr.ShowExports", 0).
		Store(&utime, &exports)
	if err != nil {
		log.Panic(err)
	}
	return utime, exports
}

func (mgr ExportMgr) getNFSIO(exportID uint32, callName string, gandi bool) BasicStats {
	out := BasicStats{}
	var call *dbus.Call
	if gandi {
		call = mgr.dbusObject.Call("org.ganesha.nfsd.exportstats."+callName, 0, exportID)
	} else {
		call = mgr.dbusObject.Call("org.ganesha.nfsd.exportstats."+callName, 0, uint16(exportID))
	}
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if gandi {
		if err := call.Store(
			&out.Status, &out.Error, &out.Time,
			&out.Read, &out.Write,
			&out.Open, &out.Close, &out.Getattr, &out.Lock,
		); err != nil {
			log.Panic(err)
		}
	} else {
		if err := call.Store(
			&out.Status, &out.Error, &out.Time,
			&out.Read, &out.Write,
		); err != nil {
			log.Panic(err)
		}
	}
	return out
}

func (mgr ExportMgr) getNFSLayouts(exportID uint32, callName string, gandi bool) PNFSOperations {
	out := PNFSOperations{}
	var call *dbus.Call
	if gandi {
		call = mgr.dbusObject.Call("org.ganesha.nfsd.exportstats."+callName, 0, exportID)
	} else {
		call = mgr.dbusObject.Call("org.ganesha.nfsd.exportstats."+callName, 0, uint16(exportID))
	}
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if err := call.Store(
		&out.Status, &out.Error, &out.Time,
		&out.Getdevinfo, &out.LayoutGet, &out.LayoutCommit, &out.LayoutReturn, &out.LayoutRecall,
	); err != nil {
		log.Panic(err)
	}
	return out
}
func (mgr ExportMgr) GetNFSv3IO(exportID uint32) BasicStats {
	return mgr.getNFSIO(exportID, "GetNFSv3IO", Gandi)
}

func (mgr ExportMgr) GetNFSv40IO(exportID uint32) BasicStats {
	return mgr.getNFSIO(exportID, "GetNFSv40IO", Gandi)
}

func (mgr ExportMgr) GetNFSv41IO(exportID uint32) BasicStats {
	return mgr.getNFSIO(exportID, "GetNFSv41IO", Gandi)
}

func (mgr ExportMgr) GetNFSv41Layouts(exportID uint32) PNFSOperations {
	return mgr.getNFSLayouts(exportID, "GetNFSv41Layouts", Gandi)
}

func (mgr ExportMgr) GetNFSv42IO(exportID uint32) BasicStats {
	return mgr.getNFSIO(exportID, "GetNFSv42IO", Gandi)
}

func (mgr ExportMgr) GetNFSv42Layouts(exportID uint32) PNFSOperations {
	return mgr.getNFSLayouts(exportID, "GetNFSv42Layouts", Gandi)
}
