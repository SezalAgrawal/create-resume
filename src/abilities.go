package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func createSkills(pdf *gofpdf.Fpdf, skills Skills) error {

	//Add Title
	pdf.SetFont("Ubuntu-Bold", "", header4)
	pdf.Write(pdf.PointConvert(header4)+4, skills.Title)
	pdf.Ln(-1)

	//Add Object List
	pdf.SetFont("Ubuntu-Regular", "", header1)
	x, y := pdf.GetXY()
	var j, cellHeight float64 = 0, header1
	var totalWidthUsed float64 = 0

	for _, obj := range skills.Object {
		pdf.SetXY(x+j, y)
		pdf.SetLineWidth(0.5)
		pdf.SetFillColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)
		var textWidth float64 = float64(2) + pdf.GetStringWidth(obj) + float64(2)
		if (rightContentLayoutWidth - totalWidthUsed) < (textWidth) {
			y += cellHeight
			totalWidthUsed = 0
			pdf.SetXY(x, y)
			j = 0
		}

		pdf.RoundedRect(x+j, y, textWidth, cellHeight-float64(2), 0.7, "1234", "F")
		pdf.SetXY(x+j+float64(1), y+float64(2))
		pdf.SetTextColor(textThirdColor.Red, textThirdColor.Green, textThirdColor.Blue)
		pdf.Write(pdf.PointConvert(header1), obj)
		j += (textWidth + float64(2))
		totalWidthUsed = j
	}

	pdf.SetTextColor(textPrimaryColor.Red, textPrimaryColor.Green, textPrimaryColor.Blue)
	//Add Bottom padding
	pdf.SetY(pdf.GetY() + 4 + cellHeight - float64(2))

	if !pdf.Ok() {
		fmt.Println("Error in Interest creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Interest created successfully!")
	return nil
}
