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


func TestLocateCameras(t *testing.T) {
	var (
		list    []*Webcam = FindWebcams()
		listLen           = len(list)

		ctrls []v4l.ControlInfo
		cfg   v4l.DeviceConfig
		err   error
	)

	if listLen < 0 {
		t.Log("No cameras found")
	}

	for _, cam := range list {
		c := &VideoConfig{}
		err = cam.Open(c)
		if err != nil {
			t.Log("Open", err)
			continue
		}

		if cfg, err = cam.device.GetConfig(); err != nil {
			t.Log("GetConfig", err)
			cam.Close()
			continue
		}

		t.Log("FPS", cfg.FPS, "Width", cfg.Width, "Height", cfg.Height, "Format", FourCC(cfg.Format))

		ctrls, err = cam.device.ListControls()
		if err != nil {
			t.Log("ListControls", err)
			cam.Close()
			continue
		}

		for i, ctrl := range ctrls {
			v, err := cam.device.GetControl(ctrl.CID)
			if err != nil {
				t.Log(err)
			}
			t.Log(i, ctrl.Name, v)
		}

		cam.Close()
	}
}
