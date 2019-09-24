package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func createWorkExperience(pdf *gofpdf.Fpdf, workEx WorkExperience) (error){	
	
	//Add Name
	pdf.SetFont("Ubuntu-Bold", "", header4)
	pdf.Write(pdf.PointConvert(header4)+4, workEx.Title)
	pdf.Ln(-1)
	
	obj := workEx.Object[0]
	pdf.SetFont("Ubuntu-Bold", "", header5)
	pdf.Write(pdf.PointConvert(header5), obj.Position)
	pdf.Ln(-1)

	pdf.SetFont("Ubuntu-Regular", "", header3)
	pdf.Write(pdf.PointConvert(header3)+2, obj.Company)
	pdf.Ln(-1)

	pdf.SetFont("Ubuntu-RegularItalic", "", header1)
	pdf.SetTextColor(textSecondaryColor.Red, textSecondaryColor.Green, textSecondaryColor.Blue)
	pdf.Write(pdf.PointConvert(header1)+3, obj.Duration.FromDate + " - " + obj.Duration.ToDate)
	pdf.SetX(marginLeft)
	pdf.CellFormat(contentLayoutWidth+4, pdf.PointConvert(header1)+3, obj.Address.City + ", " + obj.Address.Country, "", 1, "R", false, 0, "")

	pdf.Write(pdf.PointConvert(header1), "Achievements/Tasks")
	pdf.Ln(-1)

	iconX := marginLeft+2
	iconY := pdf.GetY()+2
	contentX := iconX+2
	pdf.SetFont("Ubuntu-Regular", "", header1)
	pdf.SetTextColor(textPrimaryColor.Red, textPrimaryColor.Green, textPrimaryColor.Blue)
	pdf.SetFillColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)
	pdf.Circle(iconX, iconY, 0.7, "F")
	pdf.SetX(contentX)
	pdf.MultiCell(contentLayoutWidth, 4, obj.Tasks[0], "", "TL", false)
	pdf.Ln(1)

	pdf.Circle(iconX, pdf.GetY()+1, 0.7, "F")
	pdf.SetX(contentX)
	pdf.MultiCell(contentLayoutWidth, summaryHeight, obj.Tasks[1], "", "TL", false)
	pdf.Ln(-1)

	if !pdf.Ok() {
		fmt.Println("Error in Work Experience creation!", pdf.Error())
		return pdf.Error()
	} 	
	fmt.Println("Work Experience created successfully!")
	return nil
}