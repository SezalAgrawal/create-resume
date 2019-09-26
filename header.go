package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"strings"
)

func createHeader(pdf *gofpdf.Fpdf, header Header) error {
	_, y := pdf.GetXY()

	//Add Name
	pdf.SetFontSize(header5)
	nameSizePt := pdf.PointConvert(header5)
	pdf.Write(nameSizePt+1, header.FirstName+" "+header.LastName)
	pdf.Ln(-1)

	//Add Title
	pdf.SetFont("Ubuntu-Regular", "", header3)
	pdf.SetTextColor(textSecondaryColor.Red, textSecondaryColor.Green, textSecondaryColor.Blue)
	titleSizePt := pdf.PointConvert(header3)
	pdf.Write(titleSizePt, header.Title)

	//Add Summary
	pdf.SetY(y + nameSizePt + titleSizePt + 3)
	pdf.SetFont("Ubuntu-Regular", "", header1)
	pdf.SetTextColor(textPrimaryColor.Red, textPrimaryColor.Green, textPrimaryColor.Blue)
	pdf.MultiCell(headerLayoutWidth, summaryHeight, header.Summary, "", "TL", false)

	//Add image
	pdf.SetDrawColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)
	pdf.SetLineWidth(1)
	pdf.ClipCircle(107, 25, 16, true)
	pdf.Image(header.Photo, 91, 9, 32, 35, false, "PNG", 0, "")
	pdf.ClipEnd()

	//Add contact info
	pdf.SetY(y)
	contentSizePt := pdf.PointConvert(header1)
	//setting 4 as default width of icons
	iconX := pageWidth - (4 + marginRight)
	infoX := pageWidth - (4 + marginRight + headerLayoutWidth + 2)

	pdf.Image("image/email.png", iconX, y, 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(headerLayoutWidth, contentSizePt, header.Contact.Main.Email, "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/phone.png", iconX, y+3+contentSizePt, 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(headerLayoutWidth, contentSizePt, header.Contact.Main.Contact, "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/location.png", iconX, y+2*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(headerLayoutWidth, contentSizePt, header.Contact.Main.Address.City+", "+header.Contact.Main.Address.Country, "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/linkedin.png", iconX, y+3*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(headerLayoutWidth, contentSizePt, strings.Replace(header.Contact.Social.SocialMedia.Linkedin, "https://www.", "", -1), "", 1, "R", false, 0, string(header.Contact.Social.SocialMedia.Linkedin))
	pdf.Ln(-1)

	pdf.Image("image/skype.png", iconX, y+4*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(headerLayoutWidth, contentSizePt, header.Contact.Social.SocialMedia.Skype, "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/github.png", iconX, y+5*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(headerLayoutWidth, contentSizePt, strings.Replace(header.Contact.Social.Coding.Github, "https://", "", -1), "", 1, "R", false, 0, string(header.Contact.Social.Coding.Github))
	pdf.Ln(-1)

	if !pdf.Ok() {
		fmt.Println("Error in Header creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Header created successfully!")
	return nil
}

func createLine(pdf *gofpdf.Fpdf) error {
	pdf.SetDrawColor(secondaryColor.Red, secondaryColor.Green, secondaryColor.Blue)
	pdf.SetLineWidth(0.4)
	pdf.Line(0, layoutHeight*headerPercent, pageWidth, layoutHeight*headerPercent)
	if !pdf.Ok() {
		fmt.Println("Error in Line creation!", pdf.Error())
		return pdf.Error()
	}
	fmt.Println("Line created successfully!")
	return nil
}
