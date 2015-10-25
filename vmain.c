

#ifdef __arm__
#include "bcm_host.h"
#endif
#include "il/ilclient.h"

int video_decode_test(char *filename,int x,int y,int w,int h);

#ifdef MAIN
int main (int argc, char **argv)
{
   if (argc < 2) {
      printf("Usage: %s <filename>\n", argv[0]);
      exit(1);
   }
   bcm_host_init();

   return video_decode_test(argv[1],100,10,400,400);
}
#endif
