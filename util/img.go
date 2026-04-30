package util

import (
	"image"
)

func FromRGBA(id uint32, positionX, positionY uint16, img *image.RGBA) []byte {
	return FromBytesSigned(id, positionX, positionY, uint16(img.Bounds().Dx()), uint16(img.Bounds().Dy()), uint32(len(img.Pix)), img.Pix)
}

func FromBytesSigned(id uint32, positionX, positionY, width, height uint16, length uint32, data []byte) []byte {
	buffer := make([]byte, 0)
	buffer = append(buffer, Uint32ToBytesBE(id)...)
	buffer = append(buffer, Uint16ToBytesBE(positionX)...)
	buffer = append(buffer, Uint16ToBytesBE(positionY)...)

	buffer = append(buffer, Uint16ToBytesBE(width)...)
	buffer = append(buffer, Uint16ToBytesBE(height)...)
	buffer = append(buffer, Uint32ToBytesBE(length)...)
	buffer = append(buffer, data...)
	return buffer
}

func FromBytes(id int, positionX, positionY, width, height int, length int, data []byte) []byte {
	return FromBytesSigned(uint32(id), uint16(positionX), uint16(positionY), uint16(width), uint16(height), uint32(length), data)
}
