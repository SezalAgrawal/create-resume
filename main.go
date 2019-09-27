package main

import (
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"os"
	"strconv"
)

//create global variables for fontSize
var header1, header2, header3, header4, header5, marginLeft, marginTop, marginRight, summaryHeight, pageWidth, pageHeight, layoutWidth, layoutHeight, headerLayoutWidth, contentLayoutWidth, rightContentLayoutWidth, headerPercent float64
var textPrimaryColor, textSecondaryColor, textThirdColor, primaryColor, secondaryColor Color

func main() {

	//..........................
	// Step 1: Read resume JSON object stored in json file
	//..........................

	// Open resume json file
	jsonFile, err := os.Open("resume.json")
	if err != nil {
		fmt.Println("Error", err)
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
	}

	//..........................
	// Step 2: Create PDF
	//..........................

	pdf := gofpdf.New("P", "mm", "A4", "font")

	//Set font type and size

	//Make this generic, can use bytes in one file.go
	pdf.AddFont("Ubuntu-Light", "", "Ubuntu-Light.json")
	pdf.AddFont("Ubuntu-LightItalic", "", "Ubuntu-LightItalic.json")
	pdf.AddFont("Ubuntu-Regular", "", "Ubuntu-Regular.json")
	pdf.AddFont("Ubuntu-RegularItalic", "", "Ubuntu-RegularItalic.json")
	pdf.AddFont("Ubuntu-Medium", "", "Ubuntu-Medium.json")
	pdf.AddFont("Ubuntu-MediumItalic", "", "Ubuntu-MediumItalic.json")
	pdf.AddFont("Ubuntu-Bold", "", "Ubuntu-Bold.json")
	pdf.AddFont("Ubuntu-BoldItalic", "", "Ubuntu-BoldItalic.json")

	pdf.AddPage()
	header1, _ = strconv.ParseFloat(req.TemplateInfo.TemplateDesign.Font.FontSize, 32)
	pdf.SetFont("Ubuntu-Regular", "", header1)

	//Set global variables
	setGlobalVariables()

	//Set Margin
	pdf.SetMargins(marginLeft, marginTop, marginRight)
	pdf.SetXY(marginLeft, marginTop)

	//Set basic width/height
	pageWidth, pageHeight = pdf.GetPageSize()
	layoutWidth = pageWidth - (marginLeft + marginRight)
	layoutHeight = pageHeight - (2 * marginTop)
	headerLayoutWidth = layoutWidth/2 - 23
	contentLayoutWidth = layoutWidth/2 - 5
	rightContentLayoutWidth = contentLayoutWidth - 2

	//Create header
	err = createHeader(pdf, req.UserInfo.Header)
	if err != nil {
		os.Exit(1)
	}

	//Create Line
	err = createLine(pdf)
	if err != nil {
		os.Exit(1)
	}

	//Set XY for left side
	pdf.SetXY(marginLeft, layoutHeight*headerPercent+4)

	//Create Work Experience
	err = createWorkExperience(pdf, req.UserInfo.Practical.WorkExperience)
	if err != nil {
		os.Exit(1)
	}

	//Create Education
	err = createEducation(pdf, req.UserInfo.Theoretical.Education)
	if err != nil {
		os.Exit(1)
	}

	//Set XY for right side
	pdf.SetLeftMargin(marginLeft + layoutWidth/2 + 5)
	pdf.SetXY(marginLeft+layoutWidth/2+5, layoutHeight*headerPercent+4)

	//Create Skills
	err = createSkills(pdf, req.UserInfo.Abilities.Skills)
	if err != nil {
		os.Exit(1)
	}

	//Create Projects
	err = createProjects(pdf, req.UserInfo.Practical.Projects)
	if err != nil {
		os.Exit(1)
	}

	//Create Interests
	err = createInterests(pdf, req.UserInfo.Personality.Interest)
	if err != nil {
		os.Exit(1)
	}

	//Create PDF
	err = pdf.OutputFileAndClose(req.TemplateInfo.TemplateDesign.Name + ".pdf")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully created resume!")
	}

}
