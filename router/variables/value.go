package variables

import (
	"isso0424/racion-api/controller/action"
	"isso0424/racion-api/controller/tag"
	"isso0424/racion-api/controller/template"

	"github.com/gorilla/schema"
)

var (
	ActionController   action.ActionController
	TagController      tag.TagController
	TemplateController template.TemplateController
	Decoder            = schema.NewDecoder()
)
