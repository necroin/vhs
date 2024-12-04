package remote_desktop_image

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"hash/crc32"
	"image"
	"io"
)

const pngHeader = "\x89PNG\r\n\x1a\n"

type EncoderChunk struct {
	header [8]byte
	data   []byte
	crc    [4]byte
}

func NewEncoderChunk(data []byte, name string) *EncoderChunk {
	lenght := uint32(len(data))
	header := [8]byte{}
	binary.BigEndian.PutUint32(header[:4], uint32(lenght))
	header[4] = name[0]
	header[5] = name[1]
	header[6] = name[2]
	header[7] = name[3]

	crcHash := crc32.NewIEEE()
	crcHash.Write(header[4:8])
	crcHash.Write(data)

	crc := [4]byte{}
	binary.BigEndian.PutUint32(crc[:4], crcHash.Sum32())

	return &EncoderChunk{
		header: header,
		data:   data,
		crc:    crc,
	}
}

func (encoderChunk *EncoderChunk) Data() []byte {
	return append(append(encoderChunk.header[:], encoderChunk.data...), encoderChunk.crc[:]...)
}

type Encoder struct {
	out io.Writer
}

func NewEncoder(out io.Writer) (*Encoder, error) {
	return &Encoder{out: out}, nil
}

func (encoder *Encoder) Encode(img *image.RGBA) error {
	if _, err := io.WriteString(encoder.out, pngHeader); err != nil {
		return err
	}

	encoder.writeIHDR(img)

	if err := encoder.writeIDAT(img); err != nil {
		return err
	}

	encoder.writeIEND()

	return nil
}

func Encode(out io.Writer, img *image.RGBA) error {
	encoder, err := NewEncoder(out)
	if err != nil {
		return err
	}

	return encoder.Encode(img)
}

func (encoder *Encoder) Write(out io.Writer, img image.Image) error {

	return nil
}

func (encoder *Encoder) writeChunk(data []byte, name string) {
	lenght := uint32(len(data))
	header := [8]byte{}
	binary.BigEndian.PutUint32(header[:4], uint32(lenght))
	header[4] = name[0]
	header[5] = name[1]
	header[6] = name[2]
	header[7] = name[3]

	crcHash := crc32.NewIEEE()
	crcHash.Write(header[4:8])
	crcHash.Write(data)

	crc := [4]byte{}
	binary.BigEndian.PutUint32(crc[:4], crcHash.Sum32())

	encoder.out.Write(header[:])
	encoder.out.Write(data)
	encoder.out.Write(crc[:])
}

func (encoder *Encoder) writeIHDR(img *image.RGBA) {
	data := [13]byte{}
	bounds := img.Bounds()
	binary.BigEndian.PutUint32(data[0:4], uint32(bounds.Dx()))
	binary.BigEndian.PutUint32(data[4:8], uint32(bounds.Dy()))

	data[8] = 8 // bit depth
	data[9] = 2 // color type - RGB/TrueColor

	data[10] = 0 // default compression method
	data[11] = 0 // default filter method
	data[12] = 0 // non-interlaced

	encoder.writeChunk(data[:], "IHDR")
}

func (encoder *Encoder) writeIDAT(img *image.RGBA) error {
	bitsPerPixel := 24

	bounds := img.Bounds()
	size := 1 + (bitsPerPixel*bounds.Dx()+7)/8

	stride, pix := img.Stride, img.Pix

	buffer := &bytes.Buffer{}
	zlibWriter, err := zlib.NewWriterLevel(buffer, 1)
	if err != nil {
		return err
	}

	srows := make([][]byte, bounds.Max.Y-bounds.Min.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		cr := make([]uint8, size)
		cr[0] = 0

		i := 1
		if stride != 0 {
			j0 := (y - bounds.Min.Y) * stride
			j1 := j0 + bounds.Dx()*4
			for j := j0; j < j1; j += 4 {
				cr[i+0] = pix[j+0]
				cr[i+1] = pix[j+1]
				cr[i+2] = pix[j+2]
				i += 3
			}
		} else {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				cr[i+0] = uint8(r >> 8)
				cr[i+1] = uint8(g >> 8)
				cr[i+2] = uint8(b >> 8)
				i += 3
			}
		}

		srows[y-bounds.Min.Y] = cr
	}

	for _, cr := range srows {
		zlibWriter.Write(cr)
	}

	zlibWriter.Flush()
	encoder.writeChunk(buffer.Bytes(), "IDAT")
	return nil
}

func (encoder *Encoder) writeIEND() { encoder.writeChunk(nil, "IEND") }
