package avcam

import (
	"os"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/gordonklaus/portaudio"
)

var tmpl = template.Must(template.New("").Parse(
	`{{. | len}} host APIs: {{range .}}
	Name:                   {{.Name}}
	{{if .DefaultInputDevice}}Default input device:   {{.DefaultInputDevice.Name}}{{end}}
	{{if .DefaultOutputDevice}}Default output device:  {{.DefaultOutputDevice.Name}}{{end}}
	Devices: {{range .Devices}}
		Name:                      {{.Name}}
		MaxInputChannels:          {{.MaxInputChannels}}
		MaxOutputChannels:         {{.MaxOutputChannels}}
		DefaultLowInputLatency:    {{.DefaultLowInputLatency}}
		DefaultLowOutputLatency:   {{.DefaultLowOutputLatency}}
		DefaultHighInputLatency:   {{.DefaultHighInputLatency}}
		DefaultHighOutputLatency:  {{.DefaultHighOutputLatency}}
		DefaultSampleRate:         {{.DefaultSampleRate}}
	{{end}}
{{end}}`,
))

func TestEnumerate(t *testing.T) {
	portaudio.Initialize()
	defer portaudio.Terminate()
	hostApis, err := portaudio.HostApis()
	if err != nil {
		t.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, hostApis)
	if err != nil {
		t.Fatal(err)
	}

	for i, api := range hostApis {
		t.Log(i, api.Name)
		for j, v := range api.Devices {
			if strings.HasPrefix(v.Name, "NexiGo N660") {
				t.Logf("%d '%s'", j, v.Name)
			}
		}

		in := api.DefaultInputDevice
		if in != nil {
			t.Log(in.Name)
		}
		out := api.DefaultOutputDevice
		if out != nil {
			t.Log(out.Name)
		}
	}

}

func TestRecord(t *testing.T) {
	au := Newavcam()
	stop := make(chan int)
	device, err := au.CurrentDevice()
	if err != nil {
		t.Fatal(err)
	}

	go au.RecordX(device, "TestRecord_13.aiff", stop)
	time.Sleep(10 * time.Second)
	stop <- 1
	// give go routine time to finalize
	time.Sleep(time.Second)
}

func TestSearch(t *testing.T) {
	au := Newavcam()
	lst := au.FindDevices("usb")
	t.Log("USB DEVICES", len(lst))
	for i, v := range lst {
		t.Log(i, v.Name)
	}

	lst = au.ListAllDevices()
	t.Log("ALL DEVICES", len(lst))
	for i, d := range lst {
		t.Log(i, d.Name)
	}
}

func TestavcamFind(t *testing.T) {
	au := Newavcam()
	device, err := au.FindDevice("NexiGo N660")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(device.Name)
	stop := make(chan int)
	go au.RecordX(device, "TestFind_04.aiff", stop)
	time.Sleep(20 * time.Second)
	stop <- 1
	// give go routine time to finalize
	time.Sleep(time.Second)
}

func TestDecodeSampleRate(t *testing.T) {
	bit80 := []byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}
	t.Log(bit80)
	sa := 0x400eac44
	t.Log(sa, 0x400e, 0xac44)
	x := 44100
	a, b := x>>8, x&0xff
	t.Logf("a:%x    b:%x\n", a, b)
	x = 16000
	a, b = x>>8, x&0xff
	t.Logf("a:%x    b:%x\n", a, b)
}
