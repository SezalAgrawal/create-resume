package main

import (
	"github.com/jung-kurt/gofpdf"
	"strconv"
)

func createProfessionalTemplate(pdf *gofpdf.Fpdf, req Request) (gofpdf.Template, error) {
	var err error
	var template gofpdf.Template
	template = pdf.CreateTemplate(func(pdf *gofpdf.Tpl) {

		header1, _ = strconv.ParseFloat(req.TemplateInfo.TemplateDesign.Font.FontSize, 32)
		pdf.SetFont("Ubuntu-Regular", "", header1)

		// //Set global variables
		setGlobalVariables()

		// //Set Margin
		pdf.SetMargins(marginLeft, marginTop, marginRight)
		pdf.SetXY(marginLeft, marginTop)

		// //Set basic width/height
		pageWidth, pageHeight = pdf.GetPageSize()
		layoutWidth = pageWidth - (marginLeft + marginRight)
		layoutHeight = pageHeight - (2 * marginTop)
		headerLayoutWidth = layoutWidth/2 - 23
		contentLayoutWidth = layoutWidth/2 - 5
		rightContentLayoutWidth = contentLayoutWidth - 2

		// //Create header
		err = createHeader(pdf, req.UserInfo.Header)
		if err != nil {
			return
		}

		//Create Line
		err = createLine(pdf)
		if err != nil {
			return
		}

		//Set XY for left side
		pdf.SetXY(marginLeft, layoutHeight*headerPercent+4)

		//Create Work Experience
		err = createWorkExperience(pdf, req.UserInfo.Practical.WorkExperience)
		if err != nil {
			return
		}

		//Create Education
		err = createEducation(pdf, req.UserInfo.Theoretical.Education)
		if err != nil {
			return
		}

		//Set XY for right side
		pdf.SetLeftMargin(marginLeft + layoutWidth/2 + 5)
		pdf.SetXY(marginLeft+layoutWidth/2+5, layoutHeight*headerPercent+4)

		//Create Skills
		err = createSkills(pdf, req.UserInfo.Abilities.Skills)
		if err != nil {
			return
		}

		//Create Projects
		err = createProjects(pdf, req.UserInfo.Practical.Projects)
		if err != nil {
			return
		}

		//Create Interests
		err = createInterests(pdf, req.UserInfo.Personality.Interest)
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	return template, nil
}
