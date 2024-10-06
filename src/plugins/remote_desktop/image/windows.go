//go:build windows

package remote_desktop_image

import (
	"fmt"
	"image"
	"unsafe"
	"vhs/src/utils/winapi"

	"golang.org/x/sys/windows"
)

func CaptureImage(desktopHWND windows.HWND, x int, y int, width int, height int) (image.Image, error) {
	desktopHDC, err := winapi.GetWindowDC(desktopHWND)
	if err != nil {
		return nil, fmt.Errorf("failed get desktop device context: %s", err)
	}
	defer winapi.ReleaseDC(desktopHWND, desktopHDC)

	desktopCompatibleHDC, err := winapi.CreateCompatibleDC(desktopHDC)
	if err != nil {
		return nil, fmt.Errorf("failed create compatible device context: %s", err)
	}
	defer winapi.DeleteDC(desktopCompatibleHDC)

	bitmap, err := winapi.CreateCompatibleBitmap(desktopHDC, int32(width), int32(height))
	if err != nil {
		return nil, fmt.Errorf("failed create compatible bitmap: %s", err)
	}
	defer winapi.DeleteObject(winapi.HGDIOBJ(bitmap))

	if _, err := winapi.SelectObject(desktopCompatibleHDC, winapi.HGDIOBJ(bitmap)); err != nil {
		return nil, fmt.Errorf("failed select bitmap: %s", err)
	}

	bitmapHeader := winapi.BITMAPINFOHEADER{}
	bitmapHeader.BiSize = uint32(unsafe.Sizeof(bitmapHeader))
	bitmapHeader.BiPlanes = 1
	bitmapHeader.BiBitCount = 32
	bitmapHeader.BiWidth = int32(width)
	bitmapHeader.BiHeight = -int32(height)
	bitmapHeader.BiCompression = winapi.BI_RGB
	bitmapHeader.BiSizeImage = 0

	bitmapDataSize := uint32(((int64(width)*int64(bitmapHeader.BiBitCount) + 31) / 32) * 4 * int64(height))
	memptr, err := windows.LocalAlloc(windows.LMEM_FIXED, bitmapDataSize)
	if err != nil {
		return nil, fmt.Errorf("failed LocalAlloc: %s", err)
	}
	defer windows.LocalFree(windows.Handle(memptr))

	if err := winapi.BitBlt(desktopCompatibleHDC, 0, 0, int32(width), int32(height), desktopHDC, int32(x), int32(y), winapi.SRCCOPY|winapi.CAPTUREBLT); err != nil {
		return nil, fmt.Errorf("failed bit blt: %s", err)
	}

	if err := winapi.GetDIBits(desktopCompatibleHDC, bitmap, 0, uint32(height), (*uint8)(unsafe.Pointer(memptr)), (*winapi.BITMAPINFO)(unsafe.Pointer(&bitmapHeader)), winapi.DIB_RGB_COLORS); err != nil {
		return nil, fmt.Errorf("failed GetDIBits: %s", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	i := 0
	src := uintptr(unsafe.Pointer(memptr))
	for y := int32(0); y < int32(height); y++ {
		for x := int32(0); x < int32(width); x++ {
			B := *(*uint8)(unsafe.Pointer(src))
			G := *(*uint8)(unsafe.Pointer(src + 1))
			R := *(*uint8)(unsafe.Pointer(src + 2))

			img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = R, G, B, 255

			i += 4
			src += 4
		}
	}

	return img, nil
}

func RectWidth(value windows.Rect) int32 {
	return value.Right - value.Left
}

func RectHeight(value windows.Rect) int32 {
	return value.Bottom - value.Top
}

func CaptureDesktopImage() (image.Image, error) {
	desktopHWND := winapi.GetDesktopWindow()

	desktopRect, err := winapi.GetWindowRect(desktopHWND)
	if err != nil {
		return nil, fmt.Errorf("failed get desktop rect: %s", err)
	}

	return CaptureImage(
		desktopHWND,
		int(desktopRect.Left),
		int(desktopRect.Top),
		int(RectWidth(desktopRect)),
		int(RectHeight(desktopRect)),
	)
}
