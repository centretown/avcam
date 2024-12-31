package avcam

type avcamSource interface {
	Record(stop chan int)
	IsEnabled() bool
}
