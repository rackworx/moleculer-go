package packet

type Packet struct {
	Type    string
	Target  string
	Payload interface{}
}