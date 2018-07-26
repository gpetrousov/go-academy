package main

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {

	digital_raindrops := []string{
		"ｱ", "ｲ", "ｳ", "ｴ", "ｵ", "ｶ", "ｷ",
		"ｸ", "ｹ", "ｺ", "ｻ", "ｼ", "ｽ", "ｾ", "ｿ", "ﾀ", "ﾁ", "ﾂ",
		"ﾃ", "ﾄ", "ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ", "ﾊ", "ﾋ", "ﾌ", "ﾍ",
		"ﾎ", "ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ", "ﾔ", "ﾕ", "ﾖ", "ﾗ", "ﾘ",
		"ﾙ", "ﾚ", "ﾛ", "ﾜ", "ﾝ"}

	fmt.Println(digital_raindrops)

	// Initialize terminal
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Clear internal back buffer
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// Get window size
	w, h := termbox.Size()
	// Setup go-routine to poll for events (key press)
	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	// Random seeding
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

mainloop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break mainloop
			}
		default:
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					if r.Intn(2) == 1 {
						dropToDisplay := digital_raindrops[r.Intn(len(digital_raindrops))]
						r := []rune(dropToDisplay)
						termbox.SetCell(x, y, r[0], termbox.ColorGreen, termbox.ColorBlack)
					}
					termbox.Flush()
				}
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
