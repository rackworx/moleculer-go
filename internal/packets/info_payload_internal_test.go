package packets

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoPayload(t *testing.T) {
	payload := &InfoPayload{
		Ver: "1",
		IpList: []string{
			"127.0.0.1",
			"192.168.255.255",
		},
		Hostname:   "hostname",
		InstanceID: "instanceId",
		Seq:        1,
		Client: &ClientPayload{
			Type:        "golang",
			Version:     "1",
			LangVersion: "1.15",
		},
	}

	content, err := json.Marshal(payload)

	assert.Nil(t, err)

	np := &InfoPayload{}

	err = json.Unmarshal(content, np)

	assert.Nil(t, err)

	assert.Equal(t, "hostname", np.Hostname)
}
