package avcam

import "testing"

func TestWebcam(t *testing.T) {
	webcam := NewWebcam("/dev/video0")

	err := webcam.Open(&VideoConfig{
		Codec:  "MJPG",
		Width:  1920,
		Height: 1080,
		FPS:    30,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer webcam.Close()

	config := webcam.GetConfig()
	t.Logf("Codec:%s %dx%d @%d\n", FourCC(config.Format),
		config.Width, config.Height, config.FPS)

	err = webcam.Configure(&VideoConfig{
		Codec:  "YUYV",
		Width:  1260,
		Height: 720,
		FPS:    25,
	})
	if err != nil {
		t.Fatal(err)
	}

	config = webcam.GetConfig()
	t.Logf("Codec:%s %dx%d @%d\n", FourCC(config.Format),
		config.Width, config.Height, config.FPS)
}
