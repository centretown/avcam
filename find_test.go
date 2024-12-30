package avcam

import (
	"testing"

	"github.com/korandiz/v4l"
)

func TestFind(t *testing.T) {
	list := v4l.FindDevices()
	for i, info := range list {
		t.Log(i, info.Camera, info.Path, info.DriverName, info.DeviceName)
	}
}
