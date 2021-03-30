package presenter

import (
	"strings"
)

const (
	JSONPresenter      Format = "json"
	TablePresenter     Format = "table"
	CycloneDxPresenter Format = "cyclonedx"
)

// Format is a dedicated type to represent a specific kind of presenter output format.
type Format string

func (f Format) String() string {
	return string(f)
}

// Parse returns the presenter.Format specified by the given user input.
func Parse(userInput string) Format {
	switch strings.ToLower(userInput) {
	case strings.ToLower(JSONPresenter.String()):
		return JSONPresenter
	case strings.ToLower(TablePresenter.String()):
		return TablePresenter
	case strings.ToLower(CycloneDxPresenter.String()):
		return CycloneDxPresenter
	default:
		pathToTemplateFile := userInput
		return Format(pathToTemplateFile)
	}
}

// FormatOptions is a list of presenter format options available to users.
var FormatOptions = []string{
	JSONPresenter.String(),
	TablePresenter.String(),
	CycloneDxPresenter.String(),
	"./path/to/a/custom-template.tmpl",
}
