package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type qrtype int
type cellshape int
type fileformat int
type color struct {
	R, G, B uint8
}
type qrborders struct {
	A, B, C, D int
}

// qr types
const (
	normal qrtype = iota
	withLogo
	halftone
)

// qr cell shapes
const (
	circle cellshape = iota
	rectangle
	// TODO def some other shapes. see https://github.com/yeqown/go-qrcode/blob/412857dffafc1e27900d2ac897708117c075c044/writer/standard/how-to-use-custom-shape.md
)

// qr format //TODO not implemented yet
const (
	jpeg fileformat = iota
	png
)

type QRSettings struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          int64              `bson:"user_id"`
	QRType          qrtype             `bson:"qr_type"`
	CellShape       cellshape          `bson:"cell_shape"`
	BackGroundColor color              `bson:"back_ground_color"`
	ForeGroundColor color              `bson:"fore_ground_color"`
	BorderWidth     qrborders          `bson:"border_width"`
}

// SetDefault QR settings with required User ID
func (q *QRSettings) SetDefault(userID int64) {
	q.UserID = userID
	q.QRType = normal
	q.CellShape = rectangle
	q.BackGroundColor = color{
		R: 255,
		G: 255,
		B: 255,
	}
	q.ForeGroundColor = color{
		R: 0,
		G: 0,
		B: 0,
	}
	q.BorderWidth = qrborders{
		A: 5,
		B: 5,
		C: 5,
		D: 5,
	}
}
