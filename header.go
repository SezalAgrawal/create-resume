package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func createHeader(pdf *gofpdf.Fpdf, header Header) (error){
	x, y := pdf.GetXY()

	//Add Name
	pdf.SetFontSize(nameSize)
	pdf.Cell(x, y, header.FirstName + " " + header.LastName)
	
	//Add Title
	nameSizePt := pdf.PointConvert(nameSize)
	pdf.SetXY(x,y+nameSizePt)
	pdf.SetFont("Ubuntu-Light", "", titleSize)
	pdf.Cell(x, y, header.Title)
	pdf.Ln(-1)

	//Add Summary
	pdf.SetFont("Ubuntu-Regular", "", contentSize)
	pdf.MultiCell(summaryWidth, summaryHeight, string(header.Summary), "", "TL", false)

	//Add image
	pdf.SetDrawColor(192, 192, 192)
	pdf.SetLineWidth(1)
	pdf.ClipCircle(107, 25, 16, true)
	pdf.Image(string(header.Photo), 91, 9, 32, 35, false, "PNG", 0, "")
	pdf.ClipEnd()

	if !pdf.Ok() {
		fmt.Println("Error in Header creation!", pdf.Error())
		return pdf.Error()
	} 	
	fmt.Println("Header created successfully!")
	return nil
}
