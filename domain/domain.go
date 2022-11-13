// Package domain contain all business rules and entity
package domain

import "image/color"

type QRSettingsRepository interface {
	Store(settings QRSettings)
	FindById(id int64) QRSettingsRepository
}

type QRGenerator interface {
	GenerateQR(settings QRSettings) QR
}

type QRSettings struct {
	Id              int64
	QRType          qrType
	CellShape       cellShape
	BackGroundColor color.RGBA
	ForeGroundColor color.RGBA
	BorderWidth     qrBorders
	ImgFormat       imgFormat
}

func NewQRSettings(id int64) *QRSettings {
	return &QRSettings{
		Id: id,
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
			WidthType: Percent,
			Value:     20,
		},
		ImgFormat: Jpeg,
	}
}

type Client struct {
	Id int64
}

type QR struct {
	data []byte //TODO check how to implement image container or snt like that
}
