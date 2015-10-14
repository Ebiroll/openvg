cmake_minimum_required(VERSION 2.4)
PROJECT(go_shiva_openvg)
#SET (CMAKE_BUILD_TYPE Release)
#SET (CMAKE_BUILD_TYPE RelWithDebInfo)
SET (CMAKE_BUILD_TYPE Debug)

find_path(GLEW_INCLUDE_DIR GL/glew.h)
find_library(GLEW_LIBRARY NAMES GLEW glew32 glew glew32s PATH_SUFFIXES lib64)

include(CheckIncludeFile)
CHECK_INCLUDE_FILE(/opt/vc/include/bcm_host.h BCMHOST)

INCLUDE_DIRECTORIES(../include/vg ../include)

INCLUDE_DIRECTORIES(shivavg .)


SET(shiva_SRC
shivavg/shDefs.h
shivavg/shExtensions.h
shivavg/shArrayBase.h
shivavg/shArrays.h
shivavg/shVectors.h
shivavg/shPath.h
shivavg/shImage.h
shivavg/shPaint.h
shivavg/shGeometry.h
shivavg/shContext.h
shivavg/shExtensions.c
shivavg/shArrays.c
shivavg/shVectors.c
shivavg/shPath.c
shivavg/shImage.c
shivavg/shPaint.c
shivavg/shGeometry.c
shivavg/shPipeline.c
shivavg/shParams.c
shivavg/shContext.c
shivavg/shVgu.c)

ADD_EXECUTABLE(shivavg
shivavg/glx.cpp
#../examples/test_tiger.c
#../examples/test_tiger_paths.c
#../examples/test.c
libshapes.c
client/hellovg.c
${shiva_SRC}
)

TARGET_LINK_LIBRARIES(shivavg X11 GLU glut dl GL m jpeg)
