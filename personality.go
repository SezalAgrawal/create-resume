package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func createInterests(pdf *gofpdf.Fpdf, interest Interest) error {

	//Add Title
	pdf.SetFont("Ubuntu-Bold", "", header4)
	pdf.Write(pdf.PointConvert(header4)+4, interest.Title)
	pdf.Ln(-1)

	//Add Object List
	pdf.SetFont("Ubuntu-Regular", "", header1)
	x, y := pdf.GetXY()
	x = x + 2
	for _, obj := range interest.Object {
		pdf.SetLineWidth(0.4)
		pdf.SetDrawColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)
		pdf.RoundedRect(x, y, rightContentLayoutWidth, summaryHeight*float64(4.5), 0.7, "1234", "D")
		pdf.SetXY(x+float64(1), y+float64(1))
		pdf.MultiCell(rightContentLayoutWidth, summaryHeight, obj, "", "TL", false)
		pdf.Ln(-1)

		//Add Bottom padding
		y = pdf.GetY()
	}

	if !pdf.Ok() {
		fmt.Println("Error in Interest creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Interest created successfully!")
	return nil
}
