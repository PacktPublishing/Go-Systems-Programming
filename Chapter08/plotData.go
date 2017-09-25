package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

var m *image.NRGBA
var x int
var y int
var barWidth int

func plotBar(width int, height int, color color.RGBA) {
	xx := 0
	for xx < barWidth {
		yy := 0
		for yy < height {
			m.Set(xx+width, y-yy, color)
			yy = yy + 1
		}
		xx = xx + 1
	}
}

func getColor(x int) color.RGBA {
	switch {
	case x == 0:
		return color.RGBA{0, 0, 255, 255}
	case x == 1:
		return color.RGBA{255, 0, 0, 255}
	case x == 2:
		return color.RGBA{0, 255, 0, 255}
	case x == 3:
		return color.RGBA{255, 255, 0, 255}
	case x == 4:
		return color.RGBA{255, 0, 255, 255}
	case x == 5:
		return color.RGBA{0, 255, 255, 255}
	case x == 6:
		return color.RGBA{255, 100, 100, 255}
	case x == 7:
		return color.RGBA{100, 100, 255, 255}
	case x == 8:
		return color.RGBA{100, 255, 255, 255}
	case x == 9:
		return color.RGBA{255, 255, 255, 255}
	}
	return color.RGBA{0, 0, 0, 255}
}

func main() {
	var data []int
	var f *os.File
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Printf("%s X Y input\n", filepath.Base(arguments[0]))
		os.Exit(0)
	}

	if len(arguments) == 3 {
		f = os.Stdin
	} else {
		filename := arguments[3]
		fTemp, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		f = fTemp
	}
	defer f.Close()

	x, _ = strconv.Atoi(arguments[1])
	y, _ = strconv.Atoi(arguments[2])
	fmt.Println("Image size:", x, y)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			data = append(data, value)
		} else {
			fmt.Println("Error:", value)
		}
	}

	fmt.Println("Slice length:", len(data))
	if len(data)*2 > x {
		fmt.Println("Image size (x) too small!")
		os.Exit(-1)
	}

	maxValue := data[0]
	for _, temp := range data {
		if maxValue < temp {
			maxValue = temp
		}
	}

	if maxValue > y {
		fmt.Println("Image size (y) too small!")
		os.Exit(-1)
	}
	fmt.Println("maxValue:", maxValue)
	barHeighPerUnit := int(y / maxValue)
	fmt.Println("barHeighPerUnit:", barHeighPerUnit)

	PNGfile := arguments[1] + "x" + arguments[2] + ".png"
	OUTPUT, err := os.OpenFile(PNGfile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	m = image.NewNRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{x, y}})

	i := 0
	barWidth = int(x / len(data))
	fmt.Println("barWidth:", barWidth)
	for _, v := range data {
		c := getColor(v % 10)
		yy := v * barHeighPerUnit
		plotBar(barWidth*i, yy, c)
		fmt.Println("plotBar", barWidth*i, yy)
		i = i + 1
	}

	png.Encode(OUTPUT, m)
}
