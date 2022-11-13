package main

import (
	"fmt"
	"github.com/1azar/QRChan/domain"
	"github.com/1azar/go-qrcode/writer/standard"
	"github.com/yeqown/go-qrcode/v2"
	"image"
	"image/color"
	"os"
)

type ct func(interface{}) standard.ImageOption

type bufWriteCloser struct {
}

func main() {

	// Read image from file that already exists
	existingImageFile, err := os.Open("C:/Users/User/GolandProjects/github.com/1azar/QRChan/test/k40.jpg")
	existingImageFileHT, err2 := os.Open("test/ypic.jpeg")
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	defer existingImageFile.Close()
	defer existingImageFileHT.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)
	imageDataHT, imageTypeHT, err2 := image.Decode(existingImageFileHT)
	if err != nil {
		panic(err)
	}
	//fmt.Println(imageData)
	fmt.Println(imageType)
	fmt.Println(imageTypeHT)

	b := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 0xff,
	}
	f := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0xff,
	}

	qrc, err := qrcode.New("https://github.com/yeqown/go-qrcode")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}
	optFuns := []standard.ImageOption{}
	optFuns2 := []standard.ImageOption{}
	optFuns = append(optFuns, standard.WithBgColor(b))
	optFuns = append(optFuns, standard.WithFgColor(f))
	optFuns = append(optFuns, standard.WithLogoImage(imageData))
	optFuns = append(optFuns, standard.WithHalftoneImage(imageDataHT))
	optFuns2 = append(optFuns2, standard.WithHalftoneFile("test/sbLogo.png"))
	optFuns = append(optFuns, standard.WithQRWidth(50))
	//optFuns = append(optFuns, standard.WithBorderWidth(20))

	w, err := standard.New("testQR50.jpeg", optFuns...)

	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	// save file

	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}

type QRGeneratorByYeqown struct {
}

func (qrg QRGeneratorByYeqown) Generate(settings domain.QRSettings) domain.QR {
	return domain.QR{}
}
