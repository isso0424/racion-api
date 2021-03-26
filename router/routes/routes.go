package routes

import (
	"isso0424/racion-api/router"
	"isso0424/racion-api/router/tag"
	"isso0424/racion-api/router/template"
)

var Routes = []router.Route{
	tag.TagCreating{},
	tag.Get{},
	tag.Put{},
	tag.Delete{},
	template.Create{},
}
