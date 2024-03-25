package models

// TemplateData holds data sents from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	// CSRF = Cross-site Request Forgery
	CSRFToken string
	Success   string
	Warning   string
	Error     string
}
