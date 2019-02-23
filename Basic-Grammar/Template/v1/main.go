package main

import (
	"fmt"
	"html/template"
	"os"
)

var Text = `
This is a test Template:
Name: {{.Name}}
Age:{{.Age}}
Married: {{.Married}}
School: {{.School}}
`

var TextTwo = `
This is a test Template:
Name: {{.Name |handleString | printf }}
Age: {{.Age | handleInt}}
Married: {{.Married }}
School: {{.School | handleString}}
`

var TextThree = `
{{if .Fields }}{{range .Fields }}
   Name: {{.Name}} - Age:{{.Age}} - School:{{.School}} - Married:{{.Married}}
{{ end }}{{end}}
{{if .Names}}{{ range .Names }}   {{.}} {{ end }}{{end}}
`

type Information struct {
	Name    string
	Age     int
	Married bool
	School  string
}

func (this Information) String() string {
	return fmt.Sprintf(`
This is a test String interface.
Name: %s
Age: %d
Married: %t
School: %s
`, this.Name, this.Age, this.Married, this.School)

}

func (this Information) Template(text string) {
	var (
		t *template.Template
	)
	t = template.New("New template for information")
	var (
		err error
	)

	if t, err = t.Parse(text); err != nil {
		fmt.Println("Template parse err" + err.Error())
		return
	}
	t.Execute(os.Stdout, this)
}

func (this Information) TemplateWithFuncMap(text string) {
	var (
		t   *template.Template
		err error
	)
	t = template.New("123").Funcs(InfoFuncMap())
	//t = template.New("template with funcmap")
	//t = t.Funcs(template.FuncMap{"handleString": handleString,
	//  "handleInt": handleInt})
	if t, err = t.Parse(text); err != nil {
		fmt.Println("err" + err.Error())
		return
	}
	t.Execute(os.Stdout, this)

}
func InfoFuncMap() template.FuncMap {
	return template.FuncMap{
		"handleString": handleString,
		"handleInt":    handleInt,
	}
}
func handleString(field string) string {
	if field == "" {
		return "None"
	} else {
		return " xiewei-test-for: " + field
	}
}

func handleInt(number int) int {
	return number + 20
}

type NewInformation struct {
	Fields []Information
	Names  []string
}

func RangeTemplate(args NewInformation, text string) {
	var (
		t   *template.Template
		err error
	)
	t = template.New("new template")
	if t, err = t.Parse(text); err != nil {
		fmt.Println("err" + err.Error())
		return
	}
	t.Execute(os.Stdout, args)

}

func main() {
	var (
		newInfor Information
	)
	newInfor = Information{
		Name:    "xiewei",
		Age:     18,
		School:  "shanghaiUniversity",
		Married: true,
	}
	fmt.Println("1")
	fmt.Println(newInfor)
	fmt.Println("2")
	newInfor.Template(Text)
	fmt.Println("3")
	newInfor.TemplateWithFuncMap(TextTwo)
	fmt.Println("4")
	var (
		newInformations NewInformation
	)
	newInformations = NewInformation{
		Fields: []Information{
			Information{Name: "xiexiaolu", Age: 19, Married: false, School: "pekingUniversity"},
			Information{Name: "xiexie", Age: 20, Married: true, School: "XXXXX"},
		},
		Names: []string{
			"jiewu", "baidu", "google", "aiqiyi", "wangyi",
		},
	}
	RangeTemplate(newInformations, TextThree)
}
