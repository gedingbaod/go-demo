package main

import (
	"fmt"
	"image"
	"sync"
)

var icons map[string]image.Image
var wg sync.WaitGroup
var loadIconsOnce sync.Once

func loadIcon(path string) image.Image {
	fmt.Println("load" + path)
	return nil
}
func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 是并发安全的
func Icon(name string) image.Image {
	defer wg.Done()
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Icon("left.png")
	}
	wg.Wait()
}
