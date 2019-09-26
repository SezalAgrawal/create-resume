package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func createEducation(pdf *gofpdf.Fpdf, ed Education) error {

	//Add Title
	pdf.SetFont("Ubuntu-Bold", "", header4)
	pdf.Write(pdf.PointConvert(header4)+4, ed.Title)
	pdf.Ln(-1)

	var iconX, iconY, contentX float64
	iconRadius := 0.7

	//Add Object List
	for _, obj := range ed.Object {

		//Add Position
		pdf.SetFont("Ubuntu-Bold", "", header3)
		pdf.Write(pdf.PointConvert(header3), obj.StudyProgram)
		pdf.Ln(-1)

		//Add Company
		pdf.SetFont("Ubuntu-Regular", "", header3)
		pdf.Write(pdf.PointConvert(header3)+2, obj.Institute)
		pdf.Ln(-1)

		//Add Duration and Location
		pdf.SetFont("Ubuntu-RegularItalic", "", header1)
		pdf.SetTextColor(textSecondaryColor.Red, textSecondaryColor.Green, textSecondaryColor.Blue)
		pdf.Write(pdf.PointConvert(header1)+3, obj.Duration.FromDate+" - "+obj.Duration.ToDate)
		pdf.SetX(marginLeft)
		pdf.CellFormat(contentLayoutWidth+4, pdf.PointConvert(header1)+3, obj.CGPA, "", 1, "R", false, 0, "")

		pdf.Write(pdf.PointConvert(header1)+1, "Courses")
		pdf.Ln(-1)

		//Add Tasks
		pdf.SetFont("Ubuntu-Regular", "", header1)
		pdf.SetTextColor(textPrimaryColor.Red, textPrimaryColor.Green, textPrimaryColor.Blue)
		pdf.SetFillColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)

		iconX = pdf.GetX() + 2
		iconY = pdf.GetY() + 2
		contentX = iconX + 2
		factor := 1.0

		for _, course := range obj.Courses {
			pdf.Circle(iconX, iconY, iconRadius, "F")
			pdf.SetX(contentX)
			pdf.Write(pdf.PointConvert(header1)+1, course)

			iconX = iconX + factor*(contentLayoutWidth/2+6)
			contentX = iconX + 2
			if factor == -1.0 {
				pdf.Ln(-1)
				iconY = pdf.GetY() + 2
			}
			factor *= -1.0

		}

		//Add Bottom padding
		pdf.SetY(pdf.GetY() + 4)
	}

	if !pdf.Ok() {
		fmt.Println("Error in Education creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Education created successfully!")
	return nil
}
