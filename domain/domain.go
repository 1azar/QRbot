// Package domain contain all business rules and entity
package domain

import "image/color"

type QRSettingsRepository interface {
	Store(settings QRSettings) error
	FindById(id int64) (QRSettings, error)
}

type QRGenerator interface {
	GenerateQR(settings QRSettings) QR
}

type QRSettings struct {
	ID              int64 //ID must be the same for User id
	Text            string
	QRType          qrType
	CellShape       cellShape
	BackGroundColor color.RGBA
	ForeGroundColor color.RGBA
	BorderWidth     qrBorders
	ImgFormat       imgFormat
}

func NewQRSettings(id int64) QRSettings {
	return QRSettings{
		ID: id,
		QRType: qrType{
			Name: Normal,
			Img:  nil,
		},
		CellShape: Rectangle,
		BackGroundColor: color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0xff,
		},
		ForeGroundColor: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 0xff,
		},
		BorderWidth: qrBorders{
			WidthType: Pixel,
			Value:     20,
		},
		ImgFormat: Jpeg,
	}
}

type QR struct {
	Data []byte
}
