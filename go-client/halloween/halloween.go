//
// clip: test rectangular clipping
//
package main

import (
	"bufio"
	"github.com/Ebiroll/openvg"
	"os"
)

func main() {
	var cx, cy, cw, ch, midy int
	var message []string
	message = []string {  " Black" , "cats" , "skitter" ,  "and" , "ghouls" , "patter" , "Happy" , "Halloween" }

	w, h := openvg.Init()
	//var speed openvg.VGfloat = 0.5
	var x openvg.VGfloat = 0
	var midxx openvg.VGfloat = openvg.VGfloat(w/2)
	midy = (h / 2)
	fontsize := w / 8
	cx = 0
	ch = fontsize * 2
	cw = w
	cy = midy - (ch / 2)
	var redness uint8
	var index int

	rx, ry, rw, rh := openvg.VGfloat(cx), openvg.VGfloat(cy), openvg.VGfloat(cw), openvg.VGfloat(ch)
	// scroll the text, only in the clipping rectangle
	index = 0
	for {
		for redness = 255; redness > 32; redness-=16 {
			openvg.Start(w, h)
			openvg.Background(redness, 0, 0)
			openvg.FillRGB(0, 0, 0, .2)
			openvg.Rect(rx, ry, rw, rh)
			//openvg.ClipRect(cx, cy, cw, ch)
			openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
			openvg.FillRGB(0, 0, 0, 1)
			openvg.TextMid(midxx, 0, message[index], "shf", fontsize)
			openvg.ClipEnd()
			openvg.End()
		}
		index=(index+1)%8
		if (index==0) {
				openvg.Video(100.0,100.0,800.0,600.0,"test.h264");
		}
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	openvg.Finish()
	os.Exit(0)
}
