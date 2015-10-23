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

INCLUDE_DIRECTORIES(shivavg . ilclient omx omx/include omx/vcos)


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


SET(omx_SRC
omx/common.c
omx/content_pipe_file.c
omx/queue.c
omx/content_pipe_inet.c
omx/st_static_component_loader.c
omx/omxcore.c
omx/tsemaphore.c
omx/omx_create_loaders_linux.c
omx/utils.c
omx/omx_reference_resource_manager.c
)

SET(vcos_SRC
omx/vcos/vcos_generic_event_flags.c
omx/vcos/vcos_generic_event_flags.h
omx/vcos/vcos_pthreads.c
omx/vcos/vcos_assert.h
omx/vcos/vcos_msgqueue.h
omx/vcos/vcos_platform_types.h
omx/vcos/vcos_types.h
omx/vcos/vcos.h
omx/vcos/vcos_types.h
omx/vcos/vcos_msgqueue.h
omx/vcos/vcos_msgqueue.c
omx/vcos/vcos_platform.h
omx/vcos/vcos_malloc.h
omx/vcos/vcos_logcat.c
omx/vcos/vcos.c
)



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
client/hellovg.c
#client/shapedemo.c
#video.c
${shiva_SRC}
${omx_SRC}
video.c
il/ilcore.c
il/ilclient.c
${vcos_SRC}
)

ADD_EXECUTABLE(omxregister-bellagio
omx/omxregister.c
${omx_SRC}
)


TARGET_LINK_LIBRARIES(shivavg X11 GLU glut dl GL m jpeg pthread rt)

TARGET_LINK_LIBRARIES(omxregister-bellagio dl  m pthread)


ENDIF(WIN32)



