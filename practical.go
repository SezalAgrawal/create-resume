package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func createWorkExperience(pdf *gofpdf.Fpdf, workEx WorkExperience) error {

	//Add Title
	pdf.SetFont("Ubuntu-Bold", "", header4)
	pdf.Write(pdf.PointConvert(header4)+4, workEx.Title)
	pdf.Ln(-1)

	//Set Parameters to be used in Object Creation
	iconX := pdf.GetX() + 2
	iconY := 0.0
	contentX := iconX + 2
	iconRadius := 0.7

	//Add Object List
	for _, obj := range workEx.Object {

		//Add Position
		pdf.SetFont("Ubuntu-Bold", "", header3)
		pdf.Write(pdf.PointConvert(header3), obj.Position)
		pdf.Ln(-1)

		//Add Company
		pdf.SetFont("Ubuntu-Regular", "", header3)
		pdf.Write(pdf.PointConvert(header3)+2, obj.Company)
		pdf.Ln(-1)

		//Add Duration and Location
		pdf.SetFont("Ubuntu-RegularItalic", "", header1)
		pdf.SetTextColor(textSecondaryColor.Red, textSecondaryColor.Green, textSecondaryColor.Blue)
		pdf.Write(pdf.PointConvert(header1)+3, obj.Duration.FromDate+" - "+obj.Duration.ToDate)
		pdf.SetX(marginLeft)
		pdf.CellFormat(contentLayoutWidth+4, pdf.PointConvert(header1)+3, obj.Address.City+", "+obj.Address.Country, "", 1, "R", false, 0, "")

		pdf.Write(pdf.PointConvert(header1)+1, "Achievements/Tasks")
		pdf.Ln(-1)

		//Add Tasks
		pdf.SetFont("Ubuntu-Regular", "", header1)
		pdf.SetTextColor(textPrimaryColor.Red, textPrimaryColor.Green, textPrimaryColor.Blue)
		pdf.SetFillColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)

		for _, task := range obj.Tasks {
			iconY = pdf.GetY() + 2
			pdf.Circle(iconX, iconY, iconRadius, "F")
			pdf.SetX(contentX)
			pdf.MultiCell(contentLayoutWidth, summaryHeight, task, "", "TL", false)
			pdf.Ln(1)
		}

		//Add Bottom padding
		pdf.SetY(pdf.GetY() + 4)
	}

	if !pdf.Ok() {
		fmt.Println("Error in Work Experience creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Work Experience created successfully!")
	return nil
}

func createProjects(pdf *gofpdf.Fpdf, projects Projects) error {

	//Add Title
	pdf.SetFont("Ubuntu-Bold", "", header4)
	pdf.Write(pdf.PointConvert(header4)+4, projects.Title)
	pdf.Ln(-1)

	//Set Parameters to be used in Object Creation
	iconX := pdf.GetX() + 2
	iconY := 0.0
	contentX := iconX + 2
	iconRadius := 0.7
	duration := ""
	//Add Object List
	for _, obj := range projects.Object {

		//Add Project Name and Duration
		pdf.SetFont("Ubuntu-Regular", "", header2)
		if len(obj.Duration.FromDate) == 0 {
			duration = ""
		} else {
			duration = " (" + obj.Duration.ToDate + " - " + obj.Duration.ToDate + ")"
		}
		pdf.Write(pdf.PointConvert(header2)+1, obj.ProjectName+duration)
		pdf.Ln(-1)

		//Add Tasks
		pdf.SetFont("Ubuntu-Regular", "", header1)
		pdf.SetTextColor(textPrimaryColor.Red, textPrimaryColor.Green, textPrimaryColor.Blue)
		pdf.SetFillColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)

		for _, task := range obj.Tasks {
			iconY = pdf.GetY() + 2
			pdf.Circle(iconX, iconY, iconRadius, "F")
			pdf.SetX(contentX)
			pdf.MultiCell(contentLayoutWidth, summaryHeight, task, "", "TL", false)
			pdf.Ln(1)
		}

		//Add Bottom padding
		pdf.SetY(pdf.GetY() + 3)
	}

	if !pdf.Ok() {
		fmt.Println("Error in Projects creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Projects created successfully!")
	return nil
}
