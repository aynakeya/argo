package rpc

// Protocol is a set of rpc methods that aria2 daemon supports
type Protocol interface {
	AddURI(uris []string, options ...interface{}) (gid string, err error)
	AddTorrent(content string, options ...interface{}) (gid string, err error)
	AddTorrentByFilename(filename string, options ...interface{}) (gid string, err error)
	AddMetalink(filename string, options ...interface{}) (gid []string, err error)
	Remove(gid string) (g string, err error)
	ForceRemove(gid string) (g string, err error)
	Pause(gid string) (g string, err error)
	PauseAll() (ok string, err error)
	ForcePause(gid string) (g string, err error)
	ForcePauseAll() (ok string, err error)
	Unpause(gid string) (g string, err error)
	UnpauseAll() (ok string, err error)
	TellStatus(gid string, keys ...string) (info StatusInfo, err error)
	GetURIs(gid string) (infos []URIInfo, err error)
	GetFiles(gid string) (infos []FileInfo, err error)
	GetPeers(gid string) (infos []PeerInfo, err error)
	GetServers(gid string) (infos []ServerInfo, err error)
	TellActive(keys ...string) (infos []StatusInfo, err error)
	TellWaiting(offset, num int, keys ...string) (infos []StatusInfo, err error)
	TellStopped(offset, num int, keys ...string) (infos []StatusInfo, err error)
	ChangePosition(gid string, pos int, how string) (p int, err error)
	ChangeURI(gid string, fileindex int, delUris []string, addUris []string, position ...int) (p []int, err error)
	GetOption(gid string) (m Option, err error)
	ChangeOption(gid string, option Option) (ok string, err error)
	GetGlobalOption() (m Option, err error)
	ChangeGlobalOption(options Option) (ok string, err error)
	GetGlobalStat() (info GlobalStatInfo, err error)
	PurgeDownloadResult() (ok string, err error)
	RemoveDownloadResult(gid string) (ok string, err error)
	GetVersion() (info VersionInfo, err error)
	GetSessionInfo() (info SessionInfo, err error)
	Shutdown() (ok string, err error)
	ForceShutdown() (ok string, err error)
	SaveSession() (ok string, err error)
	Multicall(methods []Method) (r []interface{}, err error)
	ListMethods() (methods []string, err error)
}
