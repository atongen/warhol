package main

import colorful "github.com/lucasb-eyer/go-colorful"

func getLabs(hue, numHues, numShadesRoot int) []*LAB {
	// rotate starting hue for each image,
	// get complimentary hue opposite starting hue
	hue = hue % 360
	portion := 360 / numHues
	hues := make([]int, numHues)
	for i := 0; i < numHues; i++ {
		hues[i] = (hue + (portion * i)) % 360
	}

	numShades := numShadesRoot * numShadesRoot

	// generate num shades of each hue
	result := make([]*LAB, numHues*numShades)
	for i, h := range hues {
		for j := 0; j < numShadesRoot; j++ {
			for k := 0; k < numShadesRoot; k++ {
				x, y := floatGrid(j, k, numShadesRoot)
				result[i*numShades+j*numShadesRoot+k] = getLabFromHcl(float64(h), x, y)
			}
		}
	}
	return result
}

func floatGrid(j, k, n int) (x, y float64) {
	pad := float64(0.25)
	size := (float64(1.0) - pad*float64(2.0)) / float64(n)
	x = pad + float64(j)*size + (size / float64(2.0))
	y = pad + float64(k)*size + (size / float64(2.0))
	return
}

func getLabFromHcl(h, c, l float64) *LAB {
	col := colorful.Hcl(h, c, l)
	l, a, b := col.Lab()
	return newLab(l, a, b)
}
