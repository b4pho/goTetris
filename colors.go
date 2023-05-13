package main

import "image/color"

type colorName int

// Note: exotic color names are taken from color-name.com but this color scheme is based on dracula (draculatheme.com)
const (
	outerSpace colorName = iota
	keyLime
	veryLightTangelo
	lightMalachiteGreen
	persianPink
	paleViolet
	gunMetal
	darkGunMetal
	transparent
)

func (c colorName) color() color.RGBA {
	switch c {
	case outerSpace:
		return color.RGBA{68, 71, 90, 255}
	case keyLime:
		return color.RGBA{241, 250, 140, 255}
	case veryLightTangelo:
		return color.RGBA{255, 184, 108, 255}
	case lightMalachiteGreen:
		return color.RGBA{80, 250, 123, 255}
	case persianPink:
		return color.RGBA{255, 121, 198, 255}
	case paleViolet:
		return color.RGBA{189, 147, 249, 255}
	case gunMetal:
		return color.RGBA{40, 42, 54, 255}
	case darkGunMetal:
		return color.RGBA{29, 31, 40, 255}
	case transparent:
		return color.RGBA{0, 0, 0, 0}
	default:
		return color.RGBA{0, 0, 0, 0}
	}
}

func (c colorName) equals(colour color.RGBA) bool {
	colour2 := c.color()
	return colour.R == colour2.R &&
		colour.G == colour2.G &&
		colour.B == colour2.B &&
		colour.A == colour2.A
}

func toGrayscale(c color.RGBA) color.RGBA {
	r, g, b := float64(c.R), float64(c.G), float64(c.B)
	gray := uint8(0.2126*r + 0.7152*g + 0.0722*b)
	return color.RGBA{gray, gray, gray, c.A}
}
