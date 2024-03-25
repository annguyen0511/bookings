package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/annguyen0511/bookings/pkg/config"
	"github.com/annguyen0511/bookings/pkg/models"
)

// option 1: more simple but not easy to fix for a long time
// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//create inMap variable to check to see if we already have the template in our cache
// 	_, inMap := templateCache[t]

// 	if !inMap {
// 		//need to create the template and add to cache if inMap is not exist
// 		log.Println("creating the template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//have the template in the cache if inMap is exist
// 		log.Println("using cached template")
// 	}

// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./template/%s", t),
// 		"./template/base.layout.tmpl",
// 	}
// 	// parse the template
// 	parseTemplate, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	templateCache[t] = parseTemplate
// 	return nil
// }

/*
	work flow:

create a map as a cache -> use function filepath.Glob(./template/*.page.tmpl) to find results matching with pattern and return a slice of the files name  *.page.tmpl from ./templates
-> use for loop to get a name of last element in path -> create a new template with a name and parse the old path
-> use function filepath.Glob(./template/*.layout.tmpl) to find results matching with pattern and return a slice of the files name *.layout.tmpl from ./template
-> if at least 1 results is found by filepath.Glob(./template/*.layout.tmpl), parse the path of result into a new template
-> assign a map with name is key and a new template is value -> return a map
*/
var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultTemplateData(d *models.TemplateData) *models.TemplateData {
	return d
}

// renderTemplate renders templates using html/template is parsing template html and return it as a request
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// // create a template cache
	// templateCache, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultTemplateData(td)

	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// excute direct data from templateCache[tmpl] to ResponseWriter
	// err = t.Execute(w, nil)
	// if err != nil {
	// 	log.Println(err)
	// }

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	//create myCache as a map for a storage of template( the same declare: myCache := make(map[string]*template.Template) )
	myCache := map[string]*template.Template{}

	//get all of the files name  *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
