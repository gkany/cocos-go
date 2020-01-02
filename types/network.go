package types

//go:generate ffjson $GOFILE

type PeerInfo struct {
	Addr                       string `json:"addr"`
	AddrLocal                  string `json:"addrlocal"`
	RPCPort                    int    `json:"rpc-port"`
	Services                   int    `json:"services"`
	LastSend                   uint64 `json:"lastsend"`
	LastRecv                   uint64 `json:"lastrecv"`
	BytesSent                  uint64 `json:"bytessent"`
	BytesRecv                  uint64 `json:"bytesrecv"`
	Conntime                   Time   `json:"conntime"`
	Version                    string `json:"version"`
	Subver                     string `json:"subver"`
	Inbound                    bool   `json:"inbound"`
	FirewallStatus             string `json:"firewall_status"`
	FCGitRevisionSha           string `json:"fc_git_revision_sha"`
	FCGitRevisionUnixTimestamp Time   `json:"fc_git_revision_unix_timestamp"`
	FCGitRevisionAge           string `json:"fc_git_revision_age"`
	Platform                   string `json:"platform"`
	CurrentHeadBlock           string `json:"current_head_block"`
	CurrentHeadBlockNumber     uint64 `json:"current_head_block_number"`
	CurrentHeadBlockTime       Time   `json:"current_head_block_time"`
}

type NetWorkPeer struct {
	Version int      `json:"version"`
	Host    string   `json:"host"`
	Info    PeerInfo `json:"info"`
}

type NetWorkPeers []NetWorkPeer
