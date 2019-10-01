package main

import (
	"github.com/jung-kurt/gofpdf"
)

func setGlobalVariables() {
	//set header
	header2 = header1 + 1
	header3 = header1 + 3
	header4 = header1 + 5
	header5 = header1*2 + 2

	//set margin
	marginLeft = 12
	marginTop = 8
	marginRight = 12

	summaryHeight = 4
	headerPercent = 0.17
	boxWidth = marginLeft - 3
	boxHeight = marginLeft

	//set Color
	textPrimaryColor = Color{0, 0, 0}
	textSecondaryColor = Color{105, 105, 105}
	textThirdColor = Color{255, 255, 255}
	primaryColor = Color{255, 69, 0}
	secondaryColor = Color{169, 169, 169}
}

func createBox(pdf *gofpdf.Tpl, color Color, x, y, w, h float64) {
	pdf.SetFillColor(color.Red, color.Green, color.Blue)
	pdf.Rect(x, y, w, h, "F")
}
