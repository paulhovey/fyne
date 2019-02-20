package widget

import (
	"testing"
	"time"

	"fyne.io/fyne"
	"github.com/stretchr/testify/assert"
)

func TestInfiniteProgressBar_Creation(t *testing.T) {
	bar := NewInfiniteProgressBar()
	// ticker should be nil when created
	assert.Nil(t, bar.ticker)
}

func TestInfiniteProgressBar_Ticker(t *testing.T) {
	bar := NewInfiniteProgressBar()

	bar.Show()
	// Show() starts a goroutine, so pause for it to initialize
	time.Sleep(10 * time.Millisecond)
	assert.NotNil(t, bar.ticker)
	bar.Hide()
	assert.Nil(t, bar.ticker)

	// make sure it restarts when re-shown
	bar.Show()
	// Show() starts a goroutine, so pause for it to initialize
	time.Sleep(10 * time.Millisecond)
	assert.NotNil(t, bar.ticker)
	bar.Hide()
	assert.Nil(t, bar.ticker)
}

func TestInfiniteProgressRenderer_Layout(t *testing.T) {
	bar := NewInfiniteProgressBar()
	width := 100
	bar.Resize(fyne.NewSize(width, 10))

	render := Renderer(bar).(*infProgressRenderer)
	// width of bar is 1/5 of width
	assert.Equal(t, width/5, render.bar.Size().Width)

}