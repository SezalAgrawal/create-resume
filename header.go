package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"strings"
)

func createHeader(pdf *gofpdf.Fpdf, header Header) (error){
	_, y := pdf.GetXY()

	//Add Name
	pdf.SetFontSize(nameSize)
	nameSizePt := pdf.PointConvert(nameSize)
	pdf.Write(nameSizePt, header.FirstName + " " + header.LastName)
	pdf.Ln(-1)

	//Add Title
	pdf.SetFont("Ubuntu-Light", "", titleSize)
	titleSizePt := pdf.PointConvert(titleSize)
	pdf.Write(titleSizePt, header.Title)

	//Add Summary
	pdf.SetY(y+nameSizePt+titleSizePt+3)
	pdf.SetFont("Ubuntu-Regular", "", contentSize)
	pdf.MultiCell(summaryWidth, summaryHeight, string(header.Summary), "", "TL", false)

	//Add image
	pdf.SetDrawColor(192, 192, 192)
	pdf.SetLineWidth(1)
	pdf.ClipCircle(107, 25, 16, true)
	pdf.Image(string(header.Photo), 91, 9, 32, 35, false, "PNG", 0, "")
	pdf.ClipEnd()

	//Add contact info
	pdf.SetY(y)
	contentSizePt := pdf.PointConvert(contentSize)
	pageWidth, _ := pdf.GetPageSize()
	//setting 4 as default width of icons	
	iconX := pageWidth - (4 + marginRight)
	infoX := pageWidth - (4 + marginRight + summaryWidth + 2)

	pdf.Image("image/email.png", iconX, y, 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(summaryWidth, contentSizePt, string(header.Contact.Main.Email), "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/phone.png", iconX, y+3+contentSizePt, 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(summaryWidth, contentSizePt, string(header.Contact.Main.Contact), "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/location.png", iconX, y+2*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(summaryWidth, contentSizePt, string(header.Contact.Main.Address.City + ", " + header.Contact.Main.Address.Country), "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/linkedin.png", iconX, y+3*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(summaryWidth, contentSizePt, strings.Replace(header.Contact.Social.SocialMedia.Linkedin, "https://www.", "", -1), "", 1, "R", false, 0, string(header.Contact.Social.SocialMedia.Linkedin))
	pdf.Ln(-1)

	pdf.Image("image/skype.png", iconX, y+4*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(summaryWidth, contentSizePt, string(header.Contact.Social.SocialMedia.Skype), "", 1, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.Image("image/github.png", iconX, y+5*(3+contentSizePt), 3, 3, false, "PNG", 0, "")
	pdf.SetX(infoX)
	pdf.CellFormat(summaryWidth, contentSizePt, strings.Replace(header.Contact.Social.Coding.Github, "https://", "", -1), "", 1, "R", false, 0, string(header.Contact.Social.Coding.Github))
	pdf.Ln(-1)

	if !pdf.Ok() {
		fmt.Println("Error in Header creation!", pdf.Error())
		return pdf.Error()
	} 	
	fmt.Println("Header created successfully!")
	return nil
}
//rgba(235, 71, 6, 1)