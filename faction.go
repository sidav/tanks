package main

import "image/color"

func possibleFactionsNumber() int {
	return len(factionTints)
}

const zeroTiltColor = 48
const strongerTiltColor = 128

var factionTints = []color.RGBA {
	{
		R: 255,
		G: 255,
		B: zeroTiltColor,
		A: 255,
	},
	{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	},
	{
		R: zeroTiltColor,
		G: 255,
		B: zeroTiltColor,
		A: 255,
	},
	{
		R: zeroTiltColor,
		G: zeroTiltColor,
		B: 255,
		A: 255,
	},
	{
		R: 255,
		G: zeroTiltColor,
		B: zeroTiltColor,
		A: 255,
	},
	{
		R: zeroTiltColor,
		G: 255,
		B: 255,
		A: 255,
	},
	{
		R: 255,
		G: zeroTiltColor,
		B: 255,
		A: 255,
	},
	{
		R: strongerTiltColor,
		G: 255,
		B: 0,
		A: 255,
	},
}
