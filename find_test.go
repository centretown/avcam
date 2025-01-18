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

func TestListWebcams(t *testing.T) {
	list := FindWebcams()
	err := list[0].Open(&VideoConfig{})
	t.Log(err)
}

func TestLocateCameras(t *testing.T) {
	var (
		list    []*Webcam = FindWebcams()
		listLen           = len(list)

		ctrls []v4l.ControlInfo
		cfg   v4l.DeviceConfig
		err   error
	)

	if listLen < 0 {
		t.Fatal("No cameras found")
	}

	for _, cam := range list {
		c := &VideoConfig{}
		err = cam.Open(c)
		if err != nil {
			t.Fatal("Open", err)
		}

		if cfg, err = cam.device.GetConfig(); err != nil {
			t.Fatal("GetConfig", err)
		}

		t.Log("FPS", cfg.FPS, "Width", cfg.Width, "Height", cfg.Height, "Format", FourCC(cfg.Format))

		ctrls, err = cam.device.ListControls()
		if err != nil {
			t.Fatal("ListControls", err)
		}

		for i, ctrl := range ctrls {
			v, err := cam.device.GetControl(ctrl.CID)
			if err != nil {
				t.Log(err)
			}
			t.Log(i, ctrl.Name, v)
		}
	}
}
