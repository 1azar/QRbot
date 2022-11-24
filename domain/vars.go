package domain

import "image"

// A borderWidthType represent border size measure unit
type borderWidthType int

const (
	Pixel borderWidthType = iota
	Percent
)

func (b borderWidthType) String() string {
	switch b {
	case Pixel:
		return "pix"
	case Percent:
		return "%"
	default:
		return "Unknown"
	}
}

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

func (qt qrTypeName) String() string {
	switch qt {
	case Normal:
		return "Normal"
	case WithLogo:
		return "WithLogo"
	case Halftone:
		return "Halftone"
	default:
		return "Unknown"
	}
}

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

func (c cellShape) String() string {
	switch c {
	case Circle:
		return "Circle"
	case Rectangle:
		return "Rectangle"
	default:
		return "Unknown"
	}
}

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
