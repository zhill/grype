package template

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"text/template"

	"github.com/anchore/grype/grype/presenter/models"

	"github.com/anchore/grype/grype/match"
	"github.com/anchore/grype/grype/pkg"
	"github.com/anchore/grype/grype/vulnerability"
)

type Presenter struct {
	matches            match.Matches
	packages           []pkg.Package
	context            pkg.Context
	metadataProvider   vulnerability.MetadataProvider
	pathToTemplateFile string
}

// NewPresenter returns a new template.Presenter.
func NewPresenter(matches match.Matches, packages []pkg.Package, context pkg.Context,
	metadataProvider vulnerability.MetadataProvider, pathToTemplateFile string) *Presenter {
	return &Presenter{
		matches:            matches,
		packages:           packages,
		metadataProvider:   metadataProvider,
		context:            context,
		pathToTemplateFile: pathToTemplateFile,
	}
}

// Present creates output using a user-supplied Go template.
func (pres *Presenter) Present(output io.Writer) error {
	templateContents, err := ioutil.ReadFile(pres.pathToTemplateFile)
	if err != nil {
		return fmt.Errorf("unable to get output template: %w", err)
	}

	tmpl, err := template.New("presentation").Funcs(funcMap).Parse(string(templateContents))
	if err != nil {
		return fmt.Errorf("unable to create template: %w", err)
	}

	document, err := models.NewDocument(pres.packages, pres.context, pres.matches, pres.metadataProvider)
	if err != nil {
		return err
	}

	err = tmpl.Execute(output, document)
	if err != nil {
		return fmt.Errorf("unable to execute supplied template: %w", err)
	}

	return nil
}

// These are custom functions available to template authors.
var funcMap = template.FuncMap{
	"getLastIndex": func(collection interface{}) int {
		if v := reflect.ValueOf(collection); v.Kind() == reflect.Slice {
			return v.Len() - 1
		}

		return 0
	},
}
