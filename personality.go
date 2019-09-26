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
	for _, obj := range interest.Object {
		pdf.SetX(pdf.GetX() + 2)
		pdf.MultiCell(contentLayoutWidth+2, summaryHeight, obj, "1", "TM", false)
		pdf.Ln(-1)
	}

	if !pdf.Ok() {
		fmt.Println("Error in Interest creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Interest created successfully!")
	return nil
}
