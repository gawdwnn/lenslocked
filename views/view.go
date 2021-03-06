package views

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	fmt.Println(files)
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil) 
}

// Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	switch data.(type) {
	case Data:
		// do nothing
	default:
		data = Data{
			Yield: data,
		}
	}

	var buff bytes.Buffer

	if err := v.Template.ExecuteTemplate(&buff, v.Layout, data); err != nil {
		http.Error(w, "Something went wrong, if problem persist contact us", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buff)
}

// layoutFiles returns a slice of strings representing
// the layout files used in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// addTemplatePath Takes in a slice of string
// representing file paths for templates, and it prepends
// the TemplateDir directory to each string in the slice
//
// eg the input ["home"] will result to ["views/home"]
// if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt Takes in a slice of string
// representing file paths for templates, and it appends
// the TemplateExt extension to each string in the slice
//
// eg the input ["home"] will result to ["home.gohtml"]
// if TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
