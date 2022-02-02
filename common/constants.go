package common

const (
	EndpointCtrlDisconnect = iota
	EndpointCtrlAddRTunnel
	EndpointCtrlAddTunnel
	EndpointCtrlSocksProxy
	EndpointCtrlSocksProxyAck
	EndpointCtrlSocksKill
	EndpointCtrlDeleteTunnel
	EndpointCtrlCMD
	EndpointCtrlCMDResult
)

const (
	TunnelCtrlConnect = iota
	TunnelCtrlAck
	TunnelCtrlDisconnect
)

const (
	ConnectionStatusConnected = iota
	ConnectionStatusClosed
)

const (
	TaskCMD = iota
	TaskRunSC
	TaskUploadFile
	TaskDownloadFile
	TaskTCPScan
)

const (
	TaskStatusRunning = iota
	TaskStatusFileDone
	TaskStatusDone
	TaskStatusError
)
