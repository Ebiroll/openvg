//
// clip: test rectangular clipping
//
package main

import (
	"bufio"
	"github.com/Ebiroll/openvg"
	"os"
    "math/rand"	
)

func main() {
	var  cy, ch, midy int
	var message []string
	message = []string {  " Black" , "cats" , "skitter" ,  "and" , "ghouls" , "patter" , "Happy" , "Halloween" }

	w, h := openvg.Init()
	//var speed openvg.VGfloat = 0.5
	var x openvg.VGfloat	= 0
	var vw,vh openvg.VGfloat 
	var midxx openvg.VGfloat = openvg.VGfloat(w/2)
	midy = (h / 2)
	fontsize := w / 8
	//cx = 0
	ch = fontsize * 2
	//cw = w
	cy = midy - (ch / 2)
	var redness uint8
	var index int

	//rx , rw, rh := openvg.VGfloat(cx) , openvg.VGfloat(cw), openvg.VGfloat(ch)
	ry := openvg.VGfloat(cy)
	// scroll the text, only in the clipping rectangle
	index = 0
	for {
		for redness = 255; redness > 16; redness-=2 {
			
			ch = fontsize * 2
	        //cw = w
	        cy = midy - (ch / 2)
			openvg.Start(w, h)
			openvg.Background(redness, 0, 0)
			openvg.FillRGB(0, 0, 0, .2)
			//openvg.Rect(rx, ry, rw, rh)
			//openvg.ClipRect(cx, cy, cw, ch)
			openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
			openvg.FillRGB(0, 0, 0, 1)
			var fsiz int 
			fsiz =  int(255-redness) 
			openvg.TextMid(midxx, 0, message[index], "shf",fsiz )
			openvg.ClipEnd()
			openvg.End()
		}
		index=(index+1)%8
		if (index==0) {
		        vw=openvg.VGfloat(w)-200
				vh=openvg.VGfloat(h)-200
				openvg.Video(100.0,100.0, vw,vh,"test.h264");
		}
		if (index==0) {		
			var test , oix int
			var xpos [40]float32
			var ypos [40]float32
			var rotate [40]float32
			var spdx [40]float32
			var spdy [40]float32

			
			// Init positions
			for test=0 ;test < 40 ;test++ {
				var rot  = rand.Float32()
				var rax = rand.Float32()
				rax = float32(w) * rax 
				var ray = rand.Float32()
				ray = float32(h) * ray

  			    spdx[test]=float32(test)*rand.Float32()
			    spdy[test]=float32(test)*rand.Float32()
				
				ypos[test] = ray
				xpos[test] = rax
				rot=0
				rotate[test] = rot			
				
			}

			
			
			// Move around
			for oix=0 ;oix < 100 ;oix++ {
				openvg.Start(w, h)
				openvg.Background(0, 0, 0)
			    openvg.FillColor("red")  
				
				for test=0 ;test < 40 ;test++ {
					var rot  = rand.Float32()
					var rax = rand.Float32() *  float32(4.0) - float32(2.0) 
					var ray = float32(4.0) * rand.Float32() - float32(2.0)  
					
					spdy[test]+=ray
					spdx[test]+=rax

				    xpos[test] = xpos[test]  + spdx[test]														
				    ypos[test] =  ypos[test] + spdy[test]
					rotate[test]=rotate[test]+float32(rot*4-2)
		
		            openvg.Rotate(openvg.VGfloat(rotate[test]))
					openvg.Translate(openvg.VGfloat(xpos[test]), openvg.VGfloat(ypos[test]))

					openvg.TextMid(0, 0, "Happy Halloween", "shf", 30 )

					openvg.Translate(-openvg.VGfloat(xpos[test]), -openvg.VGfloat(ypos[test]))
				    openvg.Rotate(-openvg.VGfloat(rotate[test]))



				}
			    openvg.End()
			}
						
		}
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	openvg.Finish()
	os.Exit(0)
}
