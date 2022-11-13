package domain

import "image"

// A borderWidthType represent border size measure unit
type borderWidthType int

const (
	Pixel borderWidthType = iota
	Percent
)

// A qrBorders represent border thickness of qr
type qrBorders struct {
	WidthType borderWidthType
	Value     int
}

// A qrType represent the type of qr
type qrTypeName int

const (
	Normal qrTypeName = iota
	WithLogo
	Halftone
)

type qrType struct {
	Name qrTypeName
	Img  image.Image
}

// A cellShape represent cell shapes of qr
type cellShape int

const (
	Circle cellShape = iota
	Rectangle
)

// A imgFormat represent resulting qr image format
type imgFormat int

const (
	Jpeg imgFormat = iota
	Png
)

// A color is color type with R G B fields
//type color struct {
//	R, G, B uint8
//}
