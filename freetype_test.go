package gocv

import (
	"image"
	"image/color"
	"testing"
)

func TestFreeTypeLoadFontData(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("/System/Library/Fonts/ArialHB.ttc", 0)
}

func TestFreeTypeGetTextSize(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("/System/Library/Fonts/ArialHB.ttc", 0)

	baseLine := 0
	size := ft.GetTextSize("test", 60, 2, &baseLine)
	t.Log(size)
	t.Log(baseLine)

	if size.X != 169 {
		t.Error("Invalid text size width")
	}

	if size.Y != 40 {
		t.Error("Invalid text size height")
	}

	if baseLine != 0 {
		t.Errorf("invalid base. expected %d, actual %d", 0, baseLine)
	}
}

func TestFreeTypePutText(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("/System/Library/Fonts/Times.ttc", 0)

	img := NewMatWithSize(150, 500, MatTypeCV8UC3)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
	defer img.Close()

	pt := image.Pt(80, 80)
	ft.PutText(&img, "Testing", pt, 60, color.RGBA{255, 255, 255, 0}, 2, 8, true)

	if img.Empty() {
		t.Error("Error in PutText test")
	}
}

func TestFreeTypeSetSplitNumber(t *testing.T) {
	ft := NewFreeType2()
	defer ft.Close()

	ft.LoadFontData("/System/Library/Fonts/Times.ttc", 0)
	ft.SetSplitNumber(10)
}
