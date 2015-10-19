//
// sl - Displays departures from a specific location based on configuration
//   Uses openVG to display the data in form of a list
//
package main

import (
	"bufio"
	"github.com/Ebiroll/openvg"
	"log"
	"fmt"
    "net/http"
    "io/ioutil"
	"os"
	"os/exec"
	"encoding/json"
)

type SLData struct {
	StatusCode  int `json:"StatusCode"`
	Message  string `json:"Message"`
	ExecutionTime int `json:"ExecutionTime"`
	ResponseData struct {
		LatestUpdate string `json:"LatestUpdate"`
		DataAge int `json:"DataAge"`
		Buses [] struct {
			JourneyDirection int `json:"JourneyDirection"`
 	        GroupOfLine  string `json:"GroupOfLine"`
			StopAreaName string `json:"StopAreaName"`
			StopAreaNumber int `json:"StopAreaNumber"`
			StopPointNumber int `json:"StopPointNumber"`
 	        StopPointDesignation  string `json:"StopPointDesignation"`
 	        TimeTabledDateTime  string `json:"TimeTabledDateTime"`
			ExpectedDateTime  string `json:"TimeTabledDateTime"`
			DisplayTime string `json:"DisplayTime"`
			Deviations string `json:"Deviations"`			
		}
	}
}





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
	var sreenHeight  , cx, cy, cw, ch, midy int
	message := "Now is the time for all good men to come to the aid of the party"

	w, h := openvg.Init()
	sreenHeight= h
	var speed openvg.VGfloat = 0.5
	var x openvg.VGfloat = 0
	midy = (h / 2)
	fontsize := w / 50
	cx = 0
	ch = fontsize * 2
	cw = w
	cy = midy - (ch / 2)
	
	response, err := http.Get("http://localhost:8000")
	
	if err == nil {
	    defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("Error reading http data, %s", err)
        } else {
           fmt.Printf("Got: %s\n", string(contents))	
		   
		    var jsonData SLData

		   
		   if err := json.Unmarshal(contents, &jsonData); err != nil {
              panic(err)
           }
           fmt.Println(jsonData)
		   
		   
	    }
	}
	
	openvg.Start(w, h)
	imgw,imgh := 0 , 0
	openvg.Background(255, 255, 255)
	
    //SLHeight = 60
	var imgPosY = openvg.VGfloat(sreenHeight - 70 )
	openvg.Image(4, imgPosY , imgw, imgh, "SL.jpg")

    rx1,  rw1, rh1 := openvg.VGfloat(cx),  openvg.VGfloat(cw), openvg.VGfloat(ch)
	ty := 0
	rix := 0
	for ty = sreenHeight - (80 + int(rh1)) ; ty>0; ty -= ch {
	    tempy := openvg.VGfloat(ty)
		//ry := openvg.VGfloat(ty)
		rix = rix+1
		if  rix%2 == 0 {
		  openvg.FillRGB(0, 0, 0, .2)
		  //openvg.Rect(rx1, tempy, rw1, rh1)	
		  openvg.FillRGB(0, 0, 0, 1)
		  openvg.Text(rx1, tempy, "591    \t12:29", "sans", fontsize)
		} else {
		  openvg.FillRGB(0, 0, 0, .4)
		  openvg.Rect(rx1, tempy, rw1, rh1)	
  		  openvg.FillRGB(0, 0, 0, 1)			
		  openvg.Text(rx1, tempy, "593    \t12:22", "sans", fontsize)
		}
		//openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
		//openvg.Background(255,255,0)

		
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
