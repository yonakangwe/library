package report

import (
	"github.com/signintech/gopdf"
)

func AddText(pdf *gopdf.GoPdf, x, y float64, text string) {
	pdf.SetXY(x, y)
	pdf.Text(text)
}
