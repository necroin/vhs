package tests

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"
	remote_desktop_image "vhs/src/plugins/remote_desktop/image"
)

const (
	outputPath = "image.png"
)

func TestEncode(t *testing.T) {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	img, err := remote_desktop_image.CaptureDesktopImage()
	if err != nil {
		t.Fatal(err)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	// encoder := png.Encoder{
	// 	CompressionLevel: 1,
	// }
	// encoder.Encode(out, img)
	remote_desktop_image.Encode(out, img)
}
