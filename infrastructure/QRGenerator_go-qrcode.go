package infrastructure

import (
	"github.com/1azar/QRChan/domain"
	"github.com/1azar/go-qrcode/writer/standard"
)

type QRGenerator struct {
}

func (qg QRGenerator) GenerateQR(settings domain.QRSettings) domain.QR {
	opt := []standard.ImageOption{}
	switch settings.QRType.Name {
	case domain.Normal:
		// no actions required due to the builder of standard.ImageOption
	case domain.WithLogo:
		opt = append(opt, standard.WithLogoImage(settings.QRType.Img))
	case domain.Halftone:
		opt = append(opt, standard.WithHalftoneImage(settings.QRType.Img))
	}
	switch settings.CellShape {
	case domain.Circle:
		opt = append(opt, standard.WithCircleShape())
	case domain.Rectangle:
		// no action required due to the builder of standard.ImageOption
	}
	opt = append(opt, standard.WithBgColor(settings.BackGroundColor))
	opt = append(opt, standard.WithFgColor(settings.ForeGroundColor))

	return domain.QR{}
}
