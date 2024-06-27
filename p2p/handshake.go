package p2p

type HandshakeFunc func (any) error
fufnc NOPHandshakeFunc(any) error { return nil }