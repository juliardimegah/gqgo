package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	geoquery "github.com/juliardimegah/gqgo"
)

func init() {
	functions.HTTP("Init", geoquery.PostGeoIntersects)
}
