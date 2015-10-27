INCLUDEFLAGS=-I/opt/vc/include -I/opt/vc/include/interface/vmcs_host/linux -I/opt/vc/include/interface/vcos/pthreads

#FONTLIB=/usr/share/fonts/truetype/ttf-dejavu DejaVuSans.ttf
FONTLIB=/usr/share/fonts/TTF/
FONTFILES=DejaVuSans.inc  DejaVuSansMono.inc DejaVuSerif.inc

CFLAGS+=-DSTANDALONE -D__STDC_CONSTANT_MACROS -D__STDC_LIMIT_MACROS -DTARGET_POSIX -D_LINUX -fPIC -DPIC -D_REENTRANT -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -U_FORTIFY_SOURCE -Wall -g -DHAVE_LIBOPENMAX=2 -DOMX -DOMX_SKIP64BIT -ftree-vectorize -pipe -DUSE_EXTERNAL_OMX -DHAVE_LIBBCM_HOST -DUSE_EXTERNAL_LIBBCM_HOST -DUSE_VCHIQ_ARM -Wno-psabi


LIBFLAGS+=-L/opt/vc/lib/ -lGLESv2 -lEGL -lopenmaxil -lbcm_host -lvcos -lvchiq_arm -lpthread -lrt -lm 


INCLUDES+=-I/opt/vc/include/ -I/opt/vc/include/interface/vcos/pthreads -I/opt/vc/include/interface/vmcs_host/linux -I./ 


CGOFILES = openvg.go

all:	font2openvg fonts library hello video listcomponents

listcomponents: omx/test/listcomponents,c
	gcc -g -DRASPBERRY_PI -I /opt/vc/include/IL -I /opt/vc/include  -I /opt/vc/include/interface/vcos/pthreads   -o listcomponents omx/test/listcomponents.c  -L /opt/vc/lib -l openmaxil -l bcm_host

libshapes.o:	libshapes.c shapes.h fontinfo.h fonts
	gcc $(CFLAGS) -O2 -fPIC -Wall $(INCLUDEFLAGS) -D BCMHOST -c libshapes.c

gopenvg:	openvg.go
	go install .

oglinit.o:	oglinit.c
	gcc -O2 -fPIC -Wall $(INCLUDEFLAGS) -D BCMHOST -c oglinit.c

hello:	hello.o video.o oglinit.o libshapes.o
	gcc  -g  $(LIBFLAGS) -lbcm_host -lvchiq_arm -ljpeg -lpthread -lrt -lm  hello.o libshapes.o oglinit.o video.o   -o hello

hello.o:	client/hellovg.c
	gcc $(CFLAGS) -g -fPIC -Wall $(INCLUDEFLAGS) -I . -o hello.o -c client/hellovg.c


video:	video.o vmain.c
	gcc  $(CFLAGS) -DMAIN  $(INCLUDEFLAGS) $(LIBFLAGS) video.o -Wl,--no-whole-archive -rdynamic -lbcm_host -lpthread  -o video vmain.c 

video.o:	video.c
	gcc $(CFLAGS) -fPIC -Wall $(INCLUDEFLAGS) -I . -o video.o -c video.c



font2openvg:	fontutil/font2openvg.cpp
	g++ -I/usr/include/freetype2 fontutil/font2openvg.cpp -o font2openvg -lfreetype

fonts:	$(FONTFILES)

DejaVuSans.inc: font2openvg $(FONTLIB)/DejaVuSans.ttf
	./font2openvg $(FONTLIB)/DejaVuSans.ttf DejaVuSans.inc DejaVuSans

DejaVuSerif.inc: font2openvg $(FONTLIB)/DejaVuSerif.ttf
	./font2openvg $(FONTLIB)/DejaVuSerif.ttf DejaVuSerif.inc DejaVuSerif

DejaVuSansMono.inc: font2openvg $(FONTLIB)/DejaVuSansMono.ttf
	./font2openvg $(FONTLIB)/DejaVuSansMono.ttf DejaVuSansMono.inc DejaVuSansMono

clean:
	rm -f *.o *.inc *.so font2openvg *.c~ *.h~ video hello
	#indent -linux -c 60 -brf -l 132  libshapes.c oglinit.c shapes.h fontinfo.h

library: oglinit.o libshapes.o video.o
	gcc $(LIBFLAGS) -fPIC -shared -o libshapes.so oglinit.o libshapes.o video.o

install:
	install -m 755 -p font2openvg /usr/bin/
	install -m 755 -p libshapes.so /usr/lib/libshapes.so.1.0.0
	strip --strip-unneeded /usr/lib/libshapes.so.1.0.0
	ln -f -s /usr/lib/libshapes.so.1.0.0 /usr/lib/libshapes.so
	ln -f -s /usr/lib/libshapes.so.1.0.0 /usr/lib/libshapes.so.1
	ln -f -s /usr/lib/libshapes.so.1.0.0 /usr/lib/libshapes.so.1.0
	install -m 644 -p shapes.h /usr/include/
	install -m 644 -p fontinfo.h /usr/include/

uninstall:
	rm -f /usr/bin/font2openvg
	rm -f /usr/lib/libshapes.so.1.0.0 /usr/lib/libshapes.so.1.0 /usr/lib/libshapes.so.1 /usr/lib/libshapes.so
	rm -f /usr/include/shapes.h /usr/include/fontinfo.h
