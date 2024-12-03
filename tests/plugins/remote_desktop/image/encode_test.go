package tests

import (
	"os"
	"testing"
	remote_desktop_image "vhs/src/plugins/remote_desktop/image"
)

const (
	outputPath = "image.png"
)

func TestEncode(t *testing.T) {
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
	// 	CompressionLevel: png.BestCompression,
	// }
	// encoder.Encode(out, img)
	remote_desktop_image.Encode(out, img)
}
