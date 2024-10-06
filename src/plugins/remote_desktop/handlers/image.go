package remote_desktop_handlers

import (
	"image/jpeg"
	"io"
	plugins_core "vhs/src/plugins/core"
	remote_desktop_image "vhs/src/plugins/remote_desktop/image"
)

func ImageHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	desktopImage, err := remote_desktop_image.CaptureDesktopImage()
	if desktopImage == nil || err != nil {
		return nil
	}
	jpeg.Encode(out, desktopImage, &jpeg.Options{Quality: 75})

	// encoder := png.Encoder{
	// 	CompressionLevel: png.NoCompression,
	// }
	// encoder.Encode(out, desktopImage)

	return nil
}
