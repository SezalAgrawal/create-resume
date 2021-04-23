package main

import (
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"os"
)

//create global variables for fontSize
var header1, header2, header3, header4, header5, marginLeft, marginTop, marginRight, summaryHeight, pageWidth, pageHeight, layoutWidth, layoutHeight, headerLayoutWidth, contentLayoutWidth, rightContentLayoutWidth, headerPercent, boxWidth, boxHeight float64
var textPrimaryColor, textSecondaryColor, textThirdColor, primaryColor, secondaryColor Color

func main() {

	//..........................
	// Step 1: Read resume JSON object stored in json file
	//..........................

	// Open resume json file.
	jsonFile, err := os.Open("input/resume.json")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		fmt.Println("Successfully opened resume.json")
	}

	// defer the closing of our jsonFile so that it can be parsed later on
	defer jsonFile.Close()

	// read contents
	var req Request

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//..........................
	// Step 2: Create PDF
	//..........................

	pdf := gofpdf.New("P", "mm", "A4", "font")

	//Set font type and size
	pdf.AddFont("Ubuntu-Light", "", "Ubuntu-Light.json")
	pdf.AddFont("Ubuntu-LightItalic", "", "Ubuntu-LightItalic.json")
	pdf.AddFont("Ubuntu-Regular", "", "Ubuntu-Regular.json")
	pdf.AddFont("Ubuntu-RegularItalic", "", "Ubuntu-RegularItalic.json")
	pdf.AddFont("Ubuntu-Medium", "", "Ubuntu-Medium.json")
	pdf.AddFont("Ubuntu-MediumItalic", "", "Ubuntu-MediumItalic.json")
	pdf.AddFont("Ubuntu-Bold", "", "Ubuntu-Bold.json")
	pdf.AddFont("Ubuntu-BoldItalic", "", "Ubuntu-BoldItalic.json")

	
	template, err := createProfessionalTemplate(pdf, req)
	if err != nil {
		fmt.Println("Unable to cerate Professional Template!")
		os.Exit(1)
	}

	pdf.AddPage()
	pdf.UseTemplate(template)

	//Create PDF
	err = pdf.OutputFileAndClose("output/" + req.TemplateInfo.TemplateDesign.Name + ".pdf")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully created resume!")
	}

}
