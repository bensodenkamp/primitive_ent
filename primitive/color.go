package primitive

import (
	"fmt"
	"image/color"
	"strings"
)

// Color Models an RGBA color
type Color struct {
	R, G, B, A int
}

// MakeColor creates non-premultiplied from a premultiplied color
func MakeColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	return Color{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// MakeHexColor makes a color from an RGBA style hex string
func MakeHexColor(x string) Color {
	x = strings.Trim(x, "#")
	var r, g, b, a int
	// Default mask is 255.
	a = 255
	var err error
	switch len(x) {
	case 3:
		_, err = fmt.Sscanf(x, "%1x%1x%1x", &r, &g, &b)
		if err != nil {
			return Color{0, 0, 0, 0}
		}
		r = (r << 4) | r
		g = (g << 4) | g
		b = (b << 4) | b
	case 4:
		_, err = fmt.Sscanf(x, "%1x%1x%1x%1x", &r, &g, &b, &a)
		if err != nil {
			return Color{0, 0, 0, 0}
		}
		r = (r << 4) | r
		g = (g << 4) | g
		b = (b << 4) | b
		a = (a << 4) | a
	case 6:
		_, err = fmt.Sscanf(x, "%02x%02x%02x", &r, &g, &b)
		if err != nil {
			return Color{0, 0, 0, 0}
		}
	case 8:
		_, err = fmt.Sscanf(x, "%02x%02x%02x%02x", &r, &g, &b, &a)
		if err != nil {
			return Color{0, 0, 0, 0}
		}
	}
	return Color{r, g, b, a}
}

// NRGBA returns the values of a non premultiplied color as uint8 values
func (c *Color) NRGBA() color.NRGBA {
	return color.NRGBA{uint8(c.R), uint8(c.G), uint8(c.B), uint8(c.A)}
}
