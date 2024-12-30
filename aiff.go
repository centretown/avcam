package avcam

import (
	"encoding/binary"
	"log"
	"os"
)

const (
	OffsetTotalBytes   int64 = 4
	OffsetTotalSamples int64 = 22
	OffsetSize         int64 = 42
	BitsPerSample      int16 = 32
)

func InitAIFF(file *os.File, sampleRate float64, channelCount int16) (err error) {
	// form chunk
	_, err = file.WriteString("FORM")
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(0)) //total bytes
	if err != nil {
		return
	}

	_, err = file.WriteString("AIFF")
	if err != nil {
		return
	}
	// common chunk
	_, err = file.WriteString("COMM")
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(18)) // file size
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, channelCount) //channels
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(0)) //number of samples
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, BitsPerSample) //bits per sample
	if err != nil {
		return
	}

	sr := uint16(sampleRate)
	b := []byte{0x40, 0x0e, byte(sr >> 8), byte(sr & 0xff), 0, 0, 0, 0, 0, 0}
	_, err = file.Write(b) //80-bit sample rate
	if err != nil {
		return
	}

	// sound chunk
	_, err = file.WriteString("SSND")
	if err != nil {
		return
	}

	err = binary.Write(file, binary.BigEndian, int32(0)) //size
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(0)) //offset
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(0)) //block
	if err != nil {
		return
	}

	return
}

func finalizeAIFF(file *os.File, nSamples int) (err error) {
	log.Println("fill in missing sizes")
	totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
	_, err = file.Seek(OffsetTotalBytes, 0)
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(totalBytes))
	if err != nil {
		return
	}

	_, err = file.Seek(OffsetTotalSamples, 0)
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(nSamples))
	if err != nil {
		return
	}
	_, err = file.Seek(OffsetSize, 0)
	if err != nil {
		return
	}
	err = binary.Write(file, binary.BigEndian, int32(4*nSamples+8))
	if err != nil {
		return
	}
	err = file.Close()
	return
}
