package main

// PtyRequestPayload is the payload for a pty request.
type PtyRequestPayload struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	ModeList []byte
}

// ForwardTCPChannelOpenPayload is the payload for a forward-tcpip channel open request.
type ForwardTCPChannelOpenPayload struct {
	ConnectedAddress  string
	ConnectedPort     uint32
	OriginatorAddress string
	OriginatorPort    uint32
}
