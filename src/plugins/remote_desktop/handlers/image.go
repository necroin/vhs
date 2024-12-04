package remote_desktop_handlers

import (
	"io"
	plugins_core "vhs/src/plugins/core"
	remote_desktop_image "vhs/src/plugins/remote_desktop/image"
)

func ImageHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	desktopImage, err := remote_desktop_image.CaptureDesktopImage()
	if desktopImage == nil || err != nil {
		return nil
	}

	remote_desktop_image.Encode(out, desktopImage)

	return nil
}
