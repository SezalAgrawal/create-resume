package main

import (
	"encoding/json"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"os"
	"strconv"
	//"github.com/create-resume/font"
)

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

	pdf := gofpdf.New("P", "mm", "A4", "")
	
	//Make this generic, can use bytes in one file.go
	pdf.AddFont("Ubuntu", "", "font/Ubuntu-Regular.json")

	pdf.AddPage()
	fontSize, _ := strconv.ParseFloat(req.TemplateInfo.TemplateDesign.Font.FontSize, 32)
	pdf.SetFont("Ubuntu", "", fontSize)
	pdf.Cell(40, 10, "Hello, world")
	err = pdf.OutputFileAndClose(req.TemplateInfo.TemplateDesign.Name + ".pdf")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully created resume!")
	}

}
