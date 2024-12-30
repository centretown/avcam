package avcam

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type VideoRec struct {
	Manual   any `json:"manual"`
	Circular any `json:"circular"`
	Modet    any `json:"modet"`
}

type VideoRecordStatus struct {
	Result    string   `json:"result"`
	Enabled   bool     `json:"enabled"`
	Vrec      VideoRec `json:"vrec"`
	Arec      VideoRec `json:"arec"`
	Vorec     VideoRec `json:"vorec"`
	AudioOnly bool     `json:"audio_only"`
	Mode      string   `json:"mode"`
}

type ValueMap map[string]string
type OptionMap map[string][]string

type CurVals struct {
	Orientation        string `json:"orientation"`
	Idle               string `json:"idle"`
	AudioOnly          string `json:"audio_only"`
	Overlay            string `json:"overlay"`
	Quality            string `json:"quality"`
	FocusHoming        string `json:"focus_homing"`
	IpAddress          string `json:"ip_address"`
	MainPort           string `json:"main_port"`
	Ipv6Address        string `json:"ipv6_address"`
	MotionLimit        string `json:"motion_limit"`
	AdetLimit          string `json:"adet_limit"`
	NightVision        string `json:"night_vision"`
	NightVisionAverage string `json:"night_vision_average"`
	NightVisionGain    string `json:"night_vision_gain"`
	VideoAcquisition   string `json:"video_acquisition"`
	MotionDetect       string `json:"motion_detect"`
	MotionDisplay      string `json:"motion_display"`
	VideoChunkLen      string `json:"video_chunk_len"`
	GpsActive          string `json:"gps_active"`
	VideoSize          string `json:"video_size"`
	MirrorFlip         string `json:"mirror_flip"`
	Ffc                string `json:"ffc"`
	RtspVideoFormats   string `json:"rtsp_video_formats"`
	RtspAudioFormats   string `json:"rtsp_audio_formats"`
	VideoConnections   string `json:"video_connections"`
	AudioConnections   string `json:"audio_connections"`
	IvideonStreaming   string `json:"ivideon_streaming"`
	Zoom               string `json:"zoom"`
	CropX              string `json:"crop_x"`
	CropY              string `json:"crop_y"`
	Coloreffect        string `json:"coloreffect"`
	Scenemode          string `json:"scenemode"`
	Focusmode          string `json:"focusmode"`
	Whitebalance       string `json:"whitebalance"`
	Flashmode          string `json:"flashmode"`
	Antibanding        string `json:"antibanding"`
	Torch              string `json:"torch"`
	FocusDistance      string `json:"focus_distance"`
	FocalLength        string `json:"focal_length"`
	Aperture           string `json:"aperture"`
	FilterDensity      string `json:"filter_density"`
	ExposureNs         string `json:"exposure_ns"`
	FrameDuration      string `json:"frame_duration"`
	Iso                string `json:"iso"`
	ManualSensor       string `json:"manual_sensor"`
	PhotoSize          string `json:"photo_size"`
	PhotoRotation      string `json:"photo_rotation"`
}
type Avail struct {
	Orientation        []string `json:"orientation"`
	Idle               []string `json:"idle"`
	AudioOnly          []string `json:"audio_only"`
	Overlay            []string `json:"overlay"`
	Quality            []string `json:"quality"`
	FocusHoming        []string `json:"focus_homing"`
	IpAddress          []string `json:"ip_address"`
	MainPort           []string `json:"main_port"`
	Ipv6Address        []string `json:"ipv6_address"`
	MotionLimit        []string `json:"motion_limit"`
	AdetLimit          []string `json:"adet_limit"`
	NightVision        []string `json:"night_vision"`
	NightVisionAverage []string `json:"night_vision_average"`
	NightVisionGain    []string `json:"night_vision_gain"`
	VideoAcquisition   []string `json:"video_acquisition"`
	MotionDetect       []string `json:"motion_detect"`
	MotionDisplay      []string `json:"motion_display"`
	VideoChunkLen      []string `json:"video_chunk_len"`
	GpsActive          []string `json:"gps_active"`
	VideoSize          []string `json:"video_size"`
	MirrorFlip         []string `json:"mirror_flip"`
	Ffc                []string `json:"ffc"`
	RtspVideoFormats   []string `json:"rtsp_video_formats"`
	RtspAudioFormats   []string `json:"rtsp_audio_formats"`
	VideoConnections   []string `json:"video_connections"`
	AudioConnections   []string `json:"audio_connections"`
	IvideonStreaming   []string `json:"ivideon_streaming"`
	Zoom               []string `json:"zoom"`
	CropX              []string `json:"crop_x"`
	CropY              []string `json:"crop_y"`
	Coloreffect        []string `json:"coloreffect"`
	Scenemode          []string `json:"scenemode"`
	Focusmode          []string `json:"focusmode"`
	Whitebalance       []string `json:"whitebalance"`
	Flashmode          []string `json:"flashmode"`
	Antibanding        []string `json:"antibanding"`
	Torch              []string `json:"torch"`
	FocusDistance      []string `json:"focus_distance"`
	FocalLength        []string `json:"focal_length"`
	Aperture           []string `json:"aperture"`
	FilterDensity      []string `json:"filter_density"`
	ExposureNs         []string `json:"exposure_ns"`
	FrameDuration      []string `json:"frame_duration"`
	Iso                []string `json:"iso"`
	ManualSensor       []string `json:"manual_sensor"`
	PhotoSize          []string `json:"photo_size"`
	PhotoRotation      []string `json:"photo_rotation"`
}

type DeviceInfo struct {
	FreeSpaceBytes      string  `json:"freeSpaceBytes"`
	UserSuppliedId      string  `json:"userSuppliedId"`
	BatteryPercent      int     `json:"batteryPercent"`
	BatteryVoltage      float64 `json:"batteryVoltage"`
	BatteryTemperatureC float64 `json:"batteryTemperatureC"`
	BatteryCharging     string  `json:"batteryCharging"`
}

type IPWebcamVariant struct {
	VideoConnections int                `json:"video_connections"`
	AudioConnections int                `json:"audio_connections"`
	VideoStatus      *VideoRecordStatus `json:"video_status"`
	CurVals          *CurVals           `json:"curvals"`
	Avail            *Avail             `json:"avail"`
	DeviceInfo       *DeviceInfo        `json:"deviceInfo"`
}

type IPWebcamStatus struct {
	VideoConnections int                `json:"video_connections"`
	AudioConnections int                `json:"audio_connections"`
	VideoStatus      *VideoRecordStatus `json:"video_status"`
	Options          ValueMap           `json:"curvals"`
	OptionMap        OptionMap          `json:"avail"`
	DeviceInfo       *DeviceInfo        `json:"deviceInfo"`
}

type IPCWConfig struct {
	Command   string
	InputType string
	Min       int
	Max       int
	Step      int
	Translate func(string) int32
}

type IpWebcamOptions struct {
	Key       string
	Value     string
	Options   []string
	Command   string
	InputType string
}

func LoadIpWebCamStatus(url string) (ipcwStat *IPWebcamStatus, err error) {
	ipcwStat = &IPWebcamStatus{}

	var (
		client = &http.Client{}
		req    *http.Request
		resp   *http.Response
		buf    []byte
	)

	req, err = http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Println("NewRequest", url, err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Println("Do Request", url, err)
		return
	}

	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll", url, err)
		return
	}

	err = json.Unmarshal(buf, ipcwStat)
	if err != nil {
		log.Println("Unmarshal", url, err)
		log.Println(string(buf))
		return
	}

	return

}

type IpWebcam struct {
	VideoConnections int
	AudioConnections int
	DeviceInfo       *DeviceInfo
	Properties       map[string]*IpWebcamOptions
}

func NewIpWebCam() *IpWebcam {
	return &IpWebcam{
		Properties: make(map[string]*IpWebcamOptions),
	}
}

func (ipwc *IpWebcam) Load(path string, configs map[string]*IPCWConfig) (err error) {
	var (
		ipcwStat *IPWebcamStatus
		url      = path + "/status.json"
		all      = len(ipwc.Properties) == 0
	)

	if all {
		url += "?show_avail=1"
	}

	ipcwStat, err = LoadIpWebCamStatus(url)
	if err != nil {
		return err
	}

	if !all {
		for k, v := range ipcwStat.Options {
			ipwc.Properties[k].Value = v
		}
		return
	}

	ipwc.VideoConnections = ipcwStat.VideoConnections
	ipwc.AudioConnections = ipcwStat.AudioConnections
	ipwc.DeviceInfo = ipcwStat.DeviceInfo
	ipwc.Properties = make(map[string]*IpWebcamOptions)

	for k, v := range ipcwStat.Options {

		opts := &IpWebcamOptions{
			Key:     k,
			Value:   v,
			Options: ipcwStat.OptionMap[k],
		}
		c, ok := configs[k]
		if ok {
			opts.Command = c.Command
			opts.InputType = c.InputType
		}

		ipwc.Properties[k] = opts
	}

	return
}
