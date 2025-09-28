package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var funcType = flag.String("f", "f", "Funtion to plot: f, g, h, w")

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}

			z := f(xyrange*(float64(i+1)/cells-0.5), xyrange*(float64(j)/cells-0.5)) // usa la función activa si quieres
			color := colorForZ(z)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	var z float64
	switch *funcType {
	case "f":
		z = f(x, y)
	case "g":
		z = g(x, y)
	case "h":
		z = h(x, y)
	case "z":
		z = w(x, y)
	default:
		z = f(x, y)
	}

	if isNonFinite(z) {
		return math.NaN(), math.NaN()
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func colorForZ(z float64) string {
	// Normaliza z a un rango [0, 1]
	minZ, maxZ := -1.0, 1.0 // Ajusta según tu función
	norm := (z - minZ) / (maxZ - minZ)
	if norm < 0 {
		norm = 0
	}
	if norm > 1 {
		norm = 1
	}
	// Interpola entre azul y rojo
	r := int(255 * norm)
	b := int(255 * (1 - norm))
	return fmt.Sprintf("#%02x00%02x", r, b)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func g(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 4
}

func h(x, y float64) float64 {
	return math.Sin(x) + math.Sin(y)
}

func w(x, y float64) float64 {
	return (x*x - y*y) / 40
}

func isNonFinite(f float64) bool {
	return math.IsInf(f, 0)
}
