// first OpenVG program
// Anthony Starks (ajstarks@gmail.com)
#include <stdio.h>
#include <stdlib.h>
#ifndef _WIN32
#include <unistd.h>
#endif
#include "VG/openvg.h"
#include "VG/vgu.h"
#include "fontinfo.h"
#include "shapes.h"

int main() {
	int width, height;
	char s[3];

	init(&width, &height);				   // Graphics initialization
    float color[]={0.0,1.0,0.0,1.0};

        char testSWE[14] = {'H','e','j',',',' ','v', 0xc3, 0xa4,'r' , 'l','d' ,'e','n',0};

	Start(width, height);				   // Start the picture
	Background(0, 0, 0);				   // Black background
	Fill(44, 77, 232, 1);				   // Big blue marble
	Circle(width / 2, 0, width);			   // The "world"
	Fill(255, 255, 255, 1);				   // White text
	TextMid(width / 2, height / 2, "hello, world", SerifTypeface, width / 10);	// Greetings 
	TextMid(width / 2, height / 2 - width/5,testSWE , SerifTypeface, width / 10);	// Greetings in swedish
	End();						   // End the picture
	
	Video(0.0f,0.0f,100.0,100.0,"test.h264");



	fgets(s, 2, stdin);				   // look at the pic, end with [RETURN]
	finish();					   // Graphics cleanup
	exit(0);
}
