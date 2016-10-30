/* Initial map generation
 * subtlepseudonym (subtlepseudonym@gmail.com)
 */

package atlas

import (
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/leesper/go_rng"
 	"github.com/peterhellberg/karta"
 	"github.com/peterhellberg/karta/diagram"
 	"github.com/peterhellberg/karta/noise"
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

const (
	plateSeed int64 = 6
	width int = 512
	height int = 512
	numPlates int = 2048 // Earth has 7 major plates and a bunch of smaller ones
)

type AtlasKarta struct {
	*karta.Karta
	CellGroups []*karta.Cells // TODO
}

// New instantiates a new Karta
func newAtlasKarta(w, h, c, r int) *AtlasKarta {
	return &AtlasKarta{
		&karta.Karta {
			Width:   w,
			Height:  h,
			Unit:    float64(math.Min(float64(w), float64(h)) / 20),
			Cells:   karta.Cells{},
			Diagram: newKartaDiagram(float64(w), float64(h), c, r),
			Noise:   noise.New(rand.Int63n(int64(w * h))),
		},
	}
}

// New generates a Voronoi diagram, relaxed by Lloyd's algorithm
func newKartaDiagram(w, h float64, c, r int) *diagram.Diagram {
	bbox := voronoi.NewBBox(0, w, 0, h)
	sites := randomSites(bbox, c)

	// Compute voronoi diagram.
	d := voronoi.ComputeDiagram(sites, bbox, true)

	// Max number of iterations is 16
	if r > 16 {
		r = 16
	}

	// Relax using Lloyd's algorithm
	for i := 0; i < r; i++ {
		sites = utils.LloydRelaxation(d.Cells)
		d = voronoi.ComputeDiagram(sites, bbox, true)
	}

	center := voronoi.Vertex{float64(w / 2), float64(h / 2)}

	return &diagram.Diagram{d, center}
}

// Generate random sites in given bounding box for plate topography
func randomSites(bbox voronoi.BBox, count int) []voronoi.Vertex {
	sites := make([]voronoi.Vertex, count)
	w := bbox.Xr - bbox.Xl
	h := bbox.Yb - bbox.Yt
	gauss := rng.NewGaussianGenerator(1)
	for j := 0; j < count; j++ {
		sites[j].X = rand.Float64() * w + bbox.Xl
		// Using ExpFloat64 rather than Float64 to simulate tectonic plate distribution
		sites[j].Y = gauss.Gaussian(0.5, 0.1) * h + bbox.Yt
	}
	return sites
}

func AtlasTest() {
	k := newAtlasKarta(width, height, numPlates, 3)

	if k.Generate() == nil {
		file, err := os.Create("exptest.png")
		if err != nil { log.Fatal(err) }
		defer file.Close()

		if err := png.Encode(file, k.Image); err != nil {
			log.Fatal(err)
		}
	}
}
