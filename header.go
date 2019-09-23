package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func createHeader(pdf *gofpdf.Fpdf, header Header) (error){
	pdf.MultiCell(80, 4, string(header.Summary), "", "TL", false)
	pdf.SetDrawColor(192, 192, 192)
	pdf.SetLineWidth(.9)
	pdf.ClipCircle(107, 25, 16, true)
	pdf.Image(string(header.Photo), 91, 9, 32, 35, false, "PNG", 0, "")
	pdf.ClipEnd()
	if !pdf.Ok() {
		fmt.Println("Error in Header creation!")
		return pdf.Error()
	} 	
	fmt.Println("Header created successfully!")
	return nil
}
