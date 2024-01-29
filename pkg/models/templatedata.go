package models

// TemplateData holds data sent from handlers to
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CDRFToken string
	Flash     string
	Warning   string
	Error     string
}