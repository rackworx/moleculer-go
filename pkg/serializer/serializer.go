package serializer

type Serializer interface {
	Marshal(payload interface{}) ([]byte, error)
	Unmarshal(data []byte, payload interface{}) error
}
