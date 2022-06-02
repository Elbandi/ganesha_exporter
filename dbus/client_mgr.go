package dbus

import (
	"github.com/godbus/dbus"
	"golang.org/x/sys/unix"
	"log"
)

// Client Structure of the output of ShowClients dbus call
type Client struct {
	Client   string
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

// ClientMgr is a handle to dbus object ClientMgr
type ClientMgr struct {
	dbusObject dbus.BusObject
}

// NewClientMgr Get a new ClientMgr
func NewClientMgr() ClientMgr {
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Panic(err)
	}
	return ClientMgr{
		dbusObject: conn.Object(
			"org.ganesha.nfsd",
			"/org/ganesha/nfsd/ClientMgr",
		),
	}
}

func (mgr ClientMgr) ShowClients() (unix.Timespec, []Client) {
	var clients []Client
	utime := unix.Timespec{}
	err := mgr.dbusObject.
		Call("org.ganesha.nfsd.clientmgr.ShowClients", 0).
		Store(&utime, &clients)
	if err != nil {
		log.Panic(err)
	}
	return utime, clients
}

func (mgr ClientMgr) getNFSIO(ipaddr string, callName string, gandi bool) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats."+callName, 0, ipaddr)
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

func (mgr ClientMgr) getNFSLayouts(ipaddr string, callName string) PNFSOperations {
	out := PNFSOperations{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats."+callName, 0, ipaddr)
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

func (mgr ClientMgr) GetNFSv3IO(ipaddr string) BasicStats {
	return mgr.getNFSIO(ipaddr, "GetNFSv3IO", false)
}

func (mgr ClientMgr) GetNFSv40IO(ipaddr string) BasicStats {
	return mgr.getNFSIO(ipaddr, "GetNFSv40IO", false)
}

func (mgr ClientMgr) GetNFSv41IO(ipaddr string) BasicStats {
	return mgr.getNFSIO(ipaddr, "GetNFSv41IO", Gandi)
}

func (mgr ClientMgr) GetNFSv41Layouts(ipaddr string) PNFSOperations {
	return mgr.getNFSLayouts(ipaddr, "GetNFSv41Layouts")
}

func (mgr ClientMgr) GetNFSv42IO(ipaddr string) BasicStats {
	return mgr.getNFSIO(ipaddr, "GetNFSv42IO", Gandi)
}

func (mgr ClientMgr) GetNFSv42Layouts(ipaddr string) PNFSOperations {
	return mgr.getNFSLayouts(ipaddr, "GetNFSv42Layouts")
}
