/* Plate tectonics
 * subtlepseudonym (subtlepseudonym@gmail.com)
 */

package atlas

import (
	"log"

 	"github.com/pzsz/voronoi"
)

const (
	width int = 512
	height int = 512
	cellCount int = 2048
	relaxPassCount int = 8
)

type PlateBoundary struct {
	Edges []*voronoi.Edge
	Plates []*Plate // the plates separated by this boundary
	Force [2]int // force exerted by boundary
}

type Plate struct {
	// Cells on tectonic plate
	Cells []*voronoi.Cell
	// Plate boundaries with other plates
	Boundaries []*PlateBoundary
}

// This function mostly copied from github.com/peterhellberg/karta
func NewDiagram(w, h, c, r int) *voronoi.Diagram {
	bbox := voronoi.NewBBox(0, w, 0, h)
	sites := utils.RandomSites(bbox, c)
	diag := voronoi.ComputeDiagram(sites, bbox, true)

	// Max number of iterations is 16
	if r > 16 {
		r = 16
	}

	// Relax using Lloyd's algorithm
	for i := 0; i < r; i++ {
		sites = utils.LloydRelaxation(diag.Cells)
		diag = voronoi.ComputeDiagram(sites, bbox, true)
	}

	return diag
}

func BuildPlate(diag *voronoi.Diagram) *Plate {
	// TODO: decide plate data structure first
}

func init() {
	log.Println("tectonics")
}