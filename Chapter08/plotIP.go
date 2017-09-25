package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var m *image.NRGBA
var x int
var y int
var barWidth int

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

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
	arguments := os.Args
	if len(arguments) < 4 {
		fmt.Printf("%s X Y IP input\n", filepath.Base(arguments[0]))
		os.Exit(0)
	}

	x, _ = strconv.Atoi(arguments[1])
	y, _ = strconv.Atoi(arguments[2])
	WANTED := arguments[3]
	fmt.Println("Image size:", x, y)

	for _, filename := range arguments[4:] {
		count := 0
		fmt.Println(filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "Error in file: %s\n", err)
				continue
			}
			ip := findIP(line)
			if ip == WANTED {
				count = count + 1
			}
		}
		data = append(data, count)
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
	PNGfile := WANTED + ".png"
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
