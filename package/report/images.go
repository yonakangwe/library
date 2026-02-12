package report

import (
	"library/package/log"
	"bytes"
	_ "embed"
	"fmt"

	"github.com/signintech/gopdf"
	"github.com/yeqown/go-qrcode"
)

//go:embed images/tz-192x192.png
var tzLogo []byte

//go:embed images/tz-192x192.png
var znzLogo []byte

func getRightLogo(pdf *gopdf.GoPdf, posX, posY float64) {
	logoImageHolder, err := gopdf.ImageHolderByBytes(znzLogo)
	if err != nil {
		log.Errorf("error crafting a holder: %v")
	}

	if err := pdf.ImageByHolder(logoImageHolder, posX, posY-10, &gopdf.Rect{W: logoSize + 20, H: logoSize + 20}); err != nil {
		log.Errorf("could not place logo: %v", err)
	}
}

func getLeftlogo(pdf *gopdf.GoPdf, posX, posY float64) {

	qrImageHolder, err := gopdf.ImageHolderByBytes(znzLogo)
	if err != nil {
		log.Errorf("error crafting a holder: %v", err)
	}

	if err := pdf.ImageByHolder(qrImageHolder, posX, posY, &gopdf.Rect{W: 90, H: 75}); err != nil {
		log.Errorf("could not create qr image: %v", err)
	}

}

func getQR(pdf *gopdf.GoPdf, posX, posY float64, qrs, doi string) {
	qrc, err := qrcode.New(qrs)

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	var qrBytes bytes.Buffer
	err = qrc.SaveTo(&qrBytes)
	if err != nil {
		fmt.Printf("error creating qr byte: %v", err)
	}

	qrImageHolder, err := gopdf.ImageHolderByBytes(qrBytes.Bytes())
	if err != nil {
		log.Errorf("error creating qr byte: %v", err)
	}

	if err := pdf.ImageByHolder(qrImageHolder, posX, posY, &gopdf.Rect{W: qrSize, H: qrSize}); err != nil {
		log.Errorf("could not create qr image: %v", err)
	}

	pdf.SetY(posY + qrSize + 10)
	pdf.SetX(posX + 1)
	setFont(pdf, 9)
	pdf.Text("DOI: " + doi)
}

// addChart
func addChart(pdf *gopdf.GoPdf, posX, posY, chartWidth, chartHeight float64, chartData []byte) {
	chartBytes, err := gopdf.ImageHolderByBytes(chartData)
	if err != nil {
		log.Errorf("error crating a holder: %v")
	}

	if err := pdf.ImageByHolder(chartBytes, posX, posY, &gopdf.Rect{W: chartWidth, H: chartHeight}); err != nil {
		log.Errorf("could not place logo: %v", err)
	}
}
