package model

//template data hold the data sent  from the handlers to the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRToken  string
	Flash     string
	Warning   string
	Error     string
}
