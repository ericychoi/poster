package main

type PipeMessage struct {
	payload       []byte
	ackHandler    func() error
	rejectHandler func(error)
}

func (p *PipeMessage) Ack() error {
	return p.ackHandler()
}

func (p *PipeMessage) Reject(e error) {
	p.rejectHandler(e)
	return
}

func (p *PipeMessage) ID() string {
	return "id"
}

func (p *PipeMessage) Payload() []byte {
	return p.payload
}

func (p *PipeMessage) SetAckHandler(f func() error) {
	p.ackHandler = f
}

func (p *PipeMessage) SetRejectHandler(f func(error)) {
	p.rejectHandler = f
}
