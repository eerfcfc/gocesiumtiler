package io

import (
	"github.com/mfbonfigli/gocesiumtiler/structs/octree"
	"github.com/mfbonfigli/gocesiumtiler/structs/tiler"
)

// Contains the minimal data needed to produce a single 3d tile, i.e. a binary content.pnts file and a tileset.json file
type WorkUnit struct {
	OctNode  *octree.OctNode
	Opts     *tiler.TilerOptions
	BasePath string
}
