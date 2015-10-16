cmake_minimum_required(VERSION 2.4)
PROJECT(go_shiva_openvg)
#SET (CMAKE_BUILD_TYPE Release)
#SET (CMAKE_BUILD_TYPE RelWithDebInfo)
SET (CMAKE_BUILD_TYPE Debug)

find_path(GLEW_INCLUDE_DIR GL/glew.h)
find_library(GLEW_LIBRARY NAMES GLEW glew32 glew glew32s PATH_SUFFIXES lib64)

include(CheckIncludeFile)
CHECK_INCLUDE_FILE(/opt/vc/include/bcm_host.h BCMHOST)

add_definitions(-DCMAKE_BUILD)

INCLUDE_DIRECTORIES(../include/vg ../include)

INCLUDE_DIRECTORIES(shivavg .)


SET(shiva_SRC
shivavg/shgl.h
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

IF(WIN32)

find_package(OpenGL REQUIRED)
include_directories(${OPENGL_INCLUDE_DIR})

IF(MSVC)

# Ignore precompile headers
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /Y-")

ENDIF(MSVC)


ADD_EXECUTABLE(shivavg
libshapes.c
shivavg/glw.c
client/hellovg.c
shivavg/VG/openvg.h
shivavg/VG/vgu.h
${shiva_SRC}
)


TARGET_LINK_LIBRARIES(shivavg ${OPENGL_LIBRARIES})


ELSE(WIN32)

#Linux
ADD_EXECUTABLE(shivavg
shivavg/glx.c
libshapes.c
#client/hellovg.c
client/shapedemo.c
${shiva_SRC}
)

TARGET_LINK_LIBRARIES(shivavg X11 GLU glut dl GL m jpeg)

ENDIF(WIN32)



