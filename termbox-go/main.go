package main

import (
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {

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
					// Set cell in back buffer
					// termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)

					if r.Intn(1) == 1 {
						termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorGreen)
					}
					// Flushes back buffer to screen
					termbox.Flush()
				}
			}
		}
	}
}
