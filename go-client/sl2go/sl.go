//
// sl - Displays departures from a specific location based on configuration
//   Uses openVG to display the data in form of a list
//
package main

import (
	"bufio"
	"github.com/Ebiroll/openvg"
	"log"
	"os"
	"os/exec"
)

func Show(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var cx, cy, cw, ch, midy int
	message := "Now is the time for all good men to come to the aid of the party"

	w, h := openvg.Init()
	var speed openvg.VGfloat = 0.5
	var x openvg.VGfloat = 0
	midy = (h / 2)
	fontsize := w / 50
	cx = 0
	ch = fontsize * 2
	cw = w
	cy = midy - (ch / 2)
	openvg.Start(w, h)
	imgw,imgh := 0 , 0
	openvg.Image(0, 450, imgw, imgh, "../img/SL.jpg")

	openvg.Background(255, 255, 255)


    rx1,  rw1, rh1 := openvg.VGfloat(cx),  openvg.VGfloat(cw), openvg.VGfloat(ch)
	ty := 0
	tix := 0
	for ty = 400 ; ty>0; ty -= 20 {
	    tempy := openvg.VGfloat(ty)
		//ry := openvg.VGfloat(ty)
		tix++
		if  tix%1 == 0 {
		  openvg.FillRGB(0, 0, 0, .2)
		} else {
		  openvg.FillRGB(0, 0, 0, .4)			
		}
		//openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
		//openvg.Background(255,255,0)
		openvg.Text(rx1, tempy, "591    12:22", "sans", fontsize)

		openvg.Rect(rx1, tempy, rw1, rh1)	
		
	}
	openvg.End()

	bufio.NewReader(os.Stdin).ReadBytes('\n')


	rx, ry, rw, rh := openvg.VGfloat(cx), openvg.VGfloat(cy), openvg.VGfloat(cw), openvg.VGfloat(ch)
	// scroll the text, only in the clipping rectangle
	for x = 0; x < rw+speed; x += speed {
		openvg.Start(w, h)
		openvg.Background(255, 255, 255)
		openvg.FillRGB(0, 0, 0, .2)
		openvg.Rect(rx, ry, rw, rh)
		openvg.ClipRect(cx, cy, cw, ch)
		openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
		openvg.FillRGB(0, 0, 0, 1)
		openvg.Text(0, 0, message, "sans", fontsize)
		openvg.ClipEnd()
		openvg.End()
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	openvg.Finish()
	os.Exit(0)
}
