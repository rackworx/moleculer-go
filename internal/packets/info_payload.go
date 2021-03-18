package packets

type InfoPayload struct {
	Config     interface{}    `json:"Config"`
	Metadata   interface{}    `json:"Metadata"`
	Ver        string         `json:"ver"`
	Sender     string         `json"sender"`
	IpList     []string       `json:"ipList"`
	Hostname   string         `json:"hostname"`
	InstanceID string         `json:"instanceID"`
	Seq        int32          `json:"seq"`
	Client     *ClientPayload `json:"client"`
}

type ClientPayload struct {
	Type        string `json:"type"`
	Version     string `json:"version"`
	LangVersion string `json:"langVersion"`
}
