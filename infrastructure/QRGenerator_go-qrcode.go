package infrastructure

import (
	"bytes"
	"github.com/1azar/QRChan/domain"
	"github.com/1azar/QRChan/usecases"
	"github.com/1azar/go-qrcode/writer/standard"
	"github.com/yeqown/go-qrcode/v2"
	"io"
)

type nopWriteCloser struct {
	io.Writer
}

func (_ nopWriteCloser) Close() error {
	return nil
}

func NopWriteCloser(writer io.Writer) io.WriteCloser {
	return nopWriteCloser{writer}
}

// GenerateQR function generates QR struct.
func GenerateQR(settings domain.QRSettings, logger usecases.Logger) domain.QR {

	qrc, err := qrcode.New(settings.Text)
	if err != nil {
		logger.Error("could not generate QRCode: ", err)
		return domain.QR{}
	}

	opt := parsDomainQRSettings(settings)

	var buf bytes.Buffer
	var QRBuf io.WriteCloser = NopWriteCloser(&buf)

	w := standard.NewWithWriter(QRBuf, opt...)

	if err = qrc.Save(w); err != nil {
		logger.Error("could not write image: ", err)
		return domain.QR{}
	}

	return domain.QR{
		Data: buf.Bytes(),
	}
}

func parsDomainQRSettings(settings domain.QRSettings) []standard.ImageOption {
	var opt []standard.ImageOption
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

	// TODO add handler for BorderWidth type (chose between percent and pixels etc.)
	// Temporary solution:
	opt = append(opt, standard.WithBorderWidth(0))

	// TODO jpeg and png formats encoder read from settings.:
	//opt = append(opt, ...

	return opt
}
