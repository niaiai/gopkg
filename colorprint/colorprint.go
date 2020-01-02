package colorprint

import (
	"fmt"
	"math/rand"
	"time"
)

type ColorItem struct {
	Control    int
	Background int
	FontColor  int
}

type Color struct {
	Colors []ColorItem
	Index  int
}

func (color *Color) Print(str string) string {
	c := color.Colors[color.Index]
	color.Index += 1
	if color.Index >= len(color.Colors) {
		color.Index = 0
	}
	if c.Background == 0 {
		return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", c.Control,
			c.FontColor, str)
	} else {
		return fmt.Sprintf("\x1b[%d;%d;%dm%s\x1b[0m", c.Control, c.Background,
			c.FontColor, str)
	}
}

// Fisher–Yates shuffle 洗牌算法
func FisherYatesShuffle(arr []ColorItem) {
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i >= 0; i-- {
		index := rand.Intn(i + 1)
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func GetColor() Color {
	c := []int{0, 1, 4, 7, 8} // 5 为闪烁
	b := []int{40, 41, 42, 43, 44, 45, 46, 47}
	f := []int{30, 31, 32, 33, 34, 35, 36, 37}
	items := make([]ColorItem, 0, 300)
	for _, i := range c {
		for _, k := range f {
			items = append(items, ColorItem{i, 0, k})
			for _, j := range b {
				if j+k == 80 || j == k+10 {
					continue
				}
				items = append(items, ColorItem{i, j, k})
			}
		}
	}
	FisherYatesShuffle(items)
	return Color{items, 0}
}
