package routes

import (
	"isso0424/racion-api/router"
	"isso0424/racion-api/router/tag"
)

var Routes = []router.Route{
	tag.TagCreating{},
}
