//go:build !windows

package remote_desktop_image

import (
	"fmt"
	"image"
)

func CaptureDesktopImage() (image.Image, error) {
	return nil, fmt.Errorf("unsupported")
}
