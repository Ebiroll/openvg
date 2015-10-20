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
	"time"
	"encoding/json"
)

// Data received from interface
type SLData struct {
	StatusCode  int `json:"StatusCode"`
	Message  string `json:"Message"`
	ExecutionTime int `json:"ExecutionTime"`
	ResponseData struct {
		LatestUpdate string `json:"LatestUpdate"`
		DataAge int `json:"DataAge"`
		Buses [] struct {
			JourneyDirection int `json:"JourneyDirection"`
			SiteId int `json:"SiteId"`
			TransportMode string `json:"TransportMode"` 
			StopAreaName string `json:"StopAreaName"`
			StopPointDesignation  string `json:"StopPointDesignation"`
			StopAreaNumber int `json:"StopAreaNumber"`
			StopPointNumber int `json:"StopPointNumber"`
			LineNumber string `json:"LineNumber"`
			Destination string `json:"Destination"`
 	        TimeTabledDateTime  string `json:"TimeTabledDateTime"`
			ExpectedDateTime  string `json:"ExpectedDateTime"`
			DisplayTime string `json:"DisplayTime"`
			Deviation [] struct {
				Consequence string `json:"Consequence"`
				ImportanceLevel int `json:"ImportanceLevel"`
				Text string `json:"Text"`
			} `json:"Deviations"`

 	        //GroupOfLine  string `json:"GroupOfLine"`
		} `json:"Buses"`
		Metros [] struct {
			SiteId int `json:"SiteId"`
			TransportMode string `json:"TransportMode"` 
			StopAreaName string `json:"StopAreaName"`
			StopAreaNumber int `json:"StopAreaNumber"`
			StopPointNumber int `json:"StopPointNumber"`
			LineNumber string `json:"LineNumber"`
			Destination string `json:"Destination"`
			TimeTabledDateTime  string `json:"TimeTabledDateTime"`
			ExpectedDateTime  string `json:"TimeTabledDateTime"`
			DisplayTime string `json:"DisplayTime"`
			Deviations [] struct {
				Consequence string `json:"Consequence"`
				ImportanceLevel int `json:"ImportanceLevel"`
				Text string `json:"Text"`
			} `json:"Deviations"`
		} `json:"Metros"`
		Trains [] struct {
			SiteId int `json:"SiteId"`
			TransportMode string `json:"TransportMode"` 
			StopAreaName string `json:"StopAreaName"`
			StopAreaNumber int `json:"StopAreaNumber"`
			StopPointNumber int `json:"StopPointNumber"`
			LineNumber string `json:"LineNumber"`
			Destination string `json:"Destination"`
			TimeTabledDateTime  string `json:"TimeTabledDateTime"`
			ExpectedDateTime  string `json:"TimeTabledDateTime"`
			DisplayTime string `json:"DisplayTime"`
			Deviations [] struct {
				Consequence string `json:"Consequence"`
				ImportanceLevel int `json:"ImportanceLevel"`
				Text string `json:"Text"`
			} `json:"Deviations"`
		} `json:"Trains"`
		Trams [] struct {
			SiteId int `json:"SiteId"`
			TransportMode string `json:"TransportMode"` 
			StopAreaName string `json:"StopAreaName"`
			StopAreaNumber int `json:"StopAreaNumber"`
			StopPointNumber int `json:"StopPointNumber"`
			LineNumber string `json:"LineNumber"`
			Destination string `json:"Destination"`
 	        TimeTabledDateTime  string `json:"TimeTabledDateTime"`
			ExpectedDateTime  string `json:"TimeTabledDateTime"`
			DisplayTime string `json:"DisplayTime"`
			Deviation [] struct {
				Consequence string `json:"Consequence"`
				ImportanceLevel int `json:"ImportanceLevel"`
				Text string `json:"Text"`
			}
		} `json:"Trams"`
		Ships [] struct {
			SiteId int `json:"SiteId"`
			TransportMode string `json:"TransportMode"` 
			StopAreaName string `json:"StopAreaName"`
			StopAreaNumber int `json:"StopAreaNumber"`
			StopPointNumber int `json:"StopPointNumber"`
			LineNumber string `json:"LineNumber"`
			Destination string `json:"Destination"`
 	        TimeTabledDateTime  string `json:"TimeTabledDateTime"`
			ExpectedDateTime  string `json:"TimeTabledDateTime"`
			DisplayTime string `json:"DisplayTime"`
			Deviation [] struct {
				Consequence string `json:"Consequence"`
				ImportanceLevel int `json:"ImportanceLevel"`
				Text string `json:"Text"`
			}
		} `json:"Ships"`
	}
}

// Config file 
type Config struct {
	key  string `json:"key"`
	siteid  int `json:"siteid"`
	buses [] struct {
		lineNumber string `json:"lineNumber"`
		destination string `json:"destination"`
		timeToStop int `json:"timeToStop"`
		timeToDest int `json:"timeToDest"`		
	} `json:"buses"`		
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

func replaceAO(in string) string {
	var ret string=in
	/*
	ret=strings.Replace(in, "Å", "A", -1)
	ret=strings.Replace(ret, "Ä", "A", -1)
	ret=strings.Replace(ret, "Ö", "O", -1)
	ret=strings.Replace(ret, "å", "a", -1)
	ret=strings.Replace(ret, "ä", "a", -1)
	ret=strings.Replace(ret, "ö", "o", -1)
	*/
	return ret
}

func drawTypeOfTransport(x openvg.VGfloat,y openvg.VGfloat, w openvg.VGfloat, h int, tType string ) {

	//openvg.Background(0, 0, 255)

     openvg.FillColor("blue")    

	 switch(tType) {
		 case "METRO" :
			openvg.Text(x, y, "T" , "sans", h)		 
		 case "BUS":
			openvg.Text(x, y, "B" , "sans", h)		 		 
		 case  "TRAIN":
			openvg.Text(x, y, "J" , "sans", h)		 		 
 		 case  "TRAM":		 
			openvg.Text(x, y, "S" , "sans", h)		 		 
		 default:		 
		 	openvg.Text(x, y, "D" , "sans", h)		 		 

	 }

     openvg.FillColor("black")    

	//openvg.Background(0, 0, 0)
	
}

func main() {
	var sreenHeight  , cx, cy, cw, ch, midy int
	message := "Scrolling texts could be useful if the information does not fit on screen"

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
    var jsonData SLData

	for {	
		response, err := http.Get("http://localhost:8000")

		if err == nil {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("Error reading http data, %s", err)
			} else {
			    fmt.Printf("Got: %s\n", string(contents))	
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
	
		var TAB1 = openvg.VGfloat(4*w/20)
		var TAB2 = openvg.VGfloat(8*w/20)
	
		rx1,  rw1, rh1 := openvg.VGfloat(cx),  openvg.VGfloat(cw), openvg.VGfloat(ch)
		ty := 0
		rix := 0
		
		// Buses
		for ty = sreenHeight - (80 + int(rh1)) ; ty>0 && rix < len(jsonData.ResponseData.Buses) ; ty -= ch {
			tempy := openvg.VGfloat(ty)
			//ry := openvg.VGfloat(ty)
			if  rix%2 == 0 {
				openvg.FillRGB(0, 0, 0, .2)
				//openvg.Rect(rx1, tempy, rw1, rh1)	
				openvg.FillRGB(0, 0, 0, 1)
				tempy = tempy + 6.0
			} else {
				openvg.FillRGB(0, 0, 0, .4)
				openvg.Rect(rx1, tempy, rw1, rh1)	
				tempy = tempy + 6.0
				openvg.FillRGB(0, 0, 0, 1)			
			}
			openvg.Text(rx1, tempy, jsonData.ResponseData.Buses[rix].LineNumber , "sans", fontsize)
			drawTypeOfTransport(rx1+60,tempy+2,rw1,ch-20,jsonData.ResponseData.Buses[rix].TransportMode)
			
			openvg.Text(rx1 + 100 ,  tempy, jsonData.ResponseData.Buses[rix].StopPointDesignation, "sans", fontsize)
			openvg.Text(rx1 + TAB1, tempy, jsonData.ResponseData.Buses[rix].DisplayTime , "sans", fontsize)
			var dest = jsonData.ResponseData.Buses[rix].Destination
			dest = replaceAO(dest)
			openvg.Text(rx1 + TAB2, tempy, dest , "sans", fontsize)
	
			
			//openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
			//openvg.Background(255,255,0)
			rix = rix+1			
		}
		var trainIx=0
		for ty=ty-20; ty>0 && trainIx < len(jsonData.ResponseData.Trains) ; ty -= ch {
			tempy := openvg.VGfloat(ty)
			//ry := openvg.VGfloat(ty)
			if  rix%2 == 0 {
				openvg.FillRGB(0, 0, 0, .2)
				//openvg.Rect(rx1, tempy, rw1, rh1)	
				openvg.FillRGB(0, 0, 0, 1)
				tempy = tempy + 6.0
			} else {
				openvg.FillRGB(0, 0, 0, .4)
				openvg.Rect(rx1, tempy, rw1, rh1)	
				tempy = tempy + 6.0
				openvg.FillRGB(0, 0, 0, 1)			
			}
			openvg.Text(rx1, tempy, jsonData.ResponseData.Trains[trainIx].LineNumber , "sans", fontsize)
			drawTypeOfTransport(rx1+55,tempy+4,rw1,ch-20,jsonData.ResponseData.Trains[trainIx].TransportMode)
			openvg.Text(rx1 + TAB1, tempy, jsonData.ResponseData.Trains[trainIx].DisplayTime , "sans", fontsize)
			var dest = jsonData.ResponseData.Trains[trainIx].Destination
			dest = replaceAO(dest)
			openvg.Text(rx1 + TAB2, tempy, dest , "sans", fontsize)
	
			
			//openvg.Translate(x, ry+openvg.VGfloat(fontsize/2))
			//openvg.Background(255,255,0)
			trainIx = trainIx +1
			rix = rix+1			
		}
				
		//openvg.End()
		openvg.SaveEnd("dump.raw")
	
		time.Sleep(60*time.Second);	
	}

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
