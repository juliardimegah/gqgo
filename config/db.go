package config

import "github.com/juliardimegah/gqgo/helper"

var Mongocon = helper.SetConnection(Mongostring, "geojson")
