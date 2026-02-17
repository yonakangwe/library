package report

import (
	"strings"

	"github.com/signintech/gopdf"
)

func mainHeader(pdf *gopdf.GoPdf, mainTitle, title, qrs, doi string) {
	//get positions
	xp := pdf.GetX() / 2
	yp := pdf.GetY() / 2

	//get logo
	getRightLogo(pdf, xp, yp)

	//wiret header
	xp = leftMargin + logoSize
	yp = topMargin

	setFont(pdf, 14)

	titleWidth := availablePageWidth - logoSize - qrSize // - 50

	x, y := addMultiLineBlock(pdf, xp, yp, titleWidth, 30, strings.ToUpper(mainTitle), true)
	pdf.SetTextColor(0, 122, 204)
	setFont(pdf, 12)
	subtitle := ""
	splTitle := strings.Split(title, "\n")
	if len(splTitle) > 1 {
		title = splTitle[0]
		subtitle = splTitle[1]
	}
	if subtitle != "" {
		xs, ys := addMultiLineBlock(pdf, x, y, titleWidth, 20.0, title, true)
		pdf.SetTextColor(0, 0, 0)
		addMultiLineBlock(pdf, xs, ys, titleWidth, 15.0, subtitle, true)
	} else {
		addMultiLineBlock(pdf, x, y, titleWidth, 15.0, title, true)
	}
	xp = pageWidth - rightMargin - qrSize
	yp = topMargin

	getQR(pdf, xp, yp, qrs, doi)
	//getLeftlogo(pdf, xp, yp)

	addHr(pdf, topMargin+logoSize+15)
}

func reportHeader(pdf *gopdf.GoPdf, mainTitle, title, qrs, doi string) {
	//get positions
	xp := pdf.GetX()
	yp := pdf.GetY()

	//get tz logo
	getRightLogo(pdf, xp, yp)

	//wiret header
	xp = leftMargin + logoSize
	yp = topMargin - 10
	setFontBold(pdf, 12)

	titleWidth := availablePageWidth - logoSize - qrSize

	x, y := addMultiLineBlock(pdf, xp, yp, titleWidth, 30, "THE UNITED REPUBLIC OF TANZANIA", true)

	pdf.SetTextColor(0, 122, 204)
	x, y = addMultiLineBlock(pdf, x, y, titleWidth, 30, "MINISTRY OF EDUCATION, SCIENCE AND TECHNOLOGY", true)

	pdf.SetTextColor(0, 0, 0)
	x, y = addMultiLineBlock(pdf, x+10, y, titleWidth, 15.0, strings.ToUpper(mainTitle), true)

	addHr(pdf, topMargin+logoSize+15)
	addHr(pdf, topMargin+logoSize+16)

	addMultiLineBlock(pdf, x, y+32, titleWidth, 15.0, title, true)

	//qr code
	xp = pageWidth - rightMargin - qrSize
	yp = topMargin
	getQR(pdf, xp, yp, qrs, doi)

	addHr(pdf, topMargin+logoSize+15)
}
