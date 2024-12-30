package avcam

type Control struct {
	Url        string
	Icon       string
	Multiplier int32
}

type ControlKey struct {
	Key      string
	Controls []*Control
}
