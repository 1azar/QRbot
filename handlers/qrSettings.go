package handlers

type qrtype int
type cellshape int
type fileformat int
type color struct {
	R, G, B uint8
}
type qrborders struct {
	a, b, c, d int
}

const (
	normal qrtype = iota
	withLogo
	halftone
)
const (
	circle cellshape = iota
	rectangle
	// TODO def some other shapes. see https://github.com/yeqown/go-qrcode/blob/412857dffafc1e27900d2ac897708117c075c044/writer/standard/how-to-use-custom-shape.md
)
const (
	jpeg fileformat = iota
	png
)

type QRSettings struct {
	QRType          qrtype    `json:"qr_type"`
	CellShape       cellshape `json:"cell_shape"`
	BackGroundColor color     `json:"back_ground_color"`
	ForeGroundColor color     `json:"fore_ground_color"`
	BorderWidth     qrborders `json:"border_width"`
}
