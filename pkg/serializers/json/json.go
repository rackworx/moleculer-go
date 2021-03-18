package json

import (
	"encoding/json"

	"github.com/rackworx/moleculer-go/pkg/serializer"
)

type jsonSerializer struct{}

func New() serializer.Serializer {
	return &jsonSerializer{}
}

func (s *jsonSerializer) Marshal(payload interface{}) ([]byte, error) {
	return json.Marshal(&payload)
}

func (s *jsonSerializer) Unmarshal(data []byte, payload interface{}) error {
	return json.Unmarshal(data, &payload)
}
