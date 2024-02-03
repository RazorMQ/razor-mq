package message

type Message struct {
	Data       []byte
	OriginHost string
	Topic      string
}

type EnqueuedMessage struct {
	ProducerHost string
	Topic        string
	Timestamp    string
	Index        int64
	Data         []byte
}
