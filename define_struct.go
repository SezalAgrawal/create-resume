package main

//Request object
type Request struct {
	TemplateInfo TemplateInfo `json:"templateInfo"`
	UserInfo     UserInfo     `json:"userInfo"`
}

//TemplateInfo object
type TemplateInfo struct {
	TemplateName   string         `json:"templateName"`
	TemplateDesign TemplateDesign `json:"templateDesign"`
}

//TemplateDesign object
type TemplateDesign struct {
	Name  string `json:"name"`
	Font  Font   `json:"font"`
	Theme Theme  `json:"theme"`
}

//Font object
type Font struct {
	FontType string `json:"fontType"`
	FontSize string `json:"fontSize"`
}

//Theme object
type Theme struct {
	BaseColor      string `json:"baseColor"`
	HighlightColor string `json:"highlightColor"`
}

//UserInfo object
type UserInfo struct {
	Header      Header      `json:"header"`
	Practical   Practical   `json:"practical"`
	Abilities   Abilities   `json:"abilities"`
	Theoretical Theoretical `json:"theoretical"`
	Personality Personality `json:"personality"`
}

//Header object
type Header struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Title     string  `json:"title"`
	Summary   string  `json:"summary"`
	Contact   Contact `json:"contact"`
	Photo     string  `json:"photo"`
}

//Contact object
type Contact struct {
	Main   Main   `json:"main"`
	Social Social `json:"social"`
}

//Main object
type Main struct {
	Email   string  `json:"email"`
	Contact string  `json:"contact"`
	Address Address `json:"address"`
}

//Address object
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

//Social object
type Social struct {
	SocialMedia SocialMedia `json:"socialMedia"`
	Coding      Coding      `json:"coding"`
}

//SocialMedia object
type SocialMedia struct {
	Skype    string `json:"skype"`
	Linkedin string `json:"linkedin"`
}

//Coding object
type Coding struct {
	Github string `json:"github"`
}

//Practical object
type Practical struct {
	WorkExperience WorkExperience `json:"workExperience"`
	Projects       Projects       `json:"projects"`
}

//WorkExperience object
type WorkExperience struct {
	Title  string       `json:"title"`
	Object []Experience `json:"object"`
}

//Experience object
type Experience struct {
	Position string   `json:"position"`
	Company  string   `json:"company"`
	Duration Duration `json:"duration"`
	Address  Address  `json:"address"`
	Tasks    []string `json:"tasks"`
}

//Duration object
type Duration struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

//Projects object
type Projects struct {
	Title  string    `json:"title"`
	Object []Project `json:"object"`
}

//Project object
type Project struct {
	ProjectName string   `json:"projectName"`
	Duration    Duration `json:"duration"`
	Tasks       []string `json:"tasks"`
}

//Abilities object
type Abilities struct {
	Skills Skills `json:"skills"`
}

//Skills object
type Skills struct {
	Title  string   `json:"title"`
	Object []string `json:"object"`
}

//Theoretical object
type Theoretical struct {
	Education Education `json:"education"`
}

//Education object
type Education struct {
	Title  string       `json:"title"`
	Object []University `json:"object"`
}

//University object
type University struct {
	StudyProgram string   `json:"studyProgram"`
	Institute    string   `json:"institute"`
	Duration     Duration `json:"duration"`
	CGPA         string   `json:"cgpa"`
	Courses      []string `json:"courses"`
}

//Personality object
type Personality struct {
	Interest Interest `json:"interests"`
}

//Interest object
type Interest struct {
	Title  string   `json:"title"`
	Object []string `json:"object"`
}
