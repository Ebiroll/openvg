/*=============================================================================
Copyright (c) 2009 Broadcom Europe Limited.
All rights reserved.

Project  :  vcfw
Module   :  chip driver
File     :  $RCSfile: $
Revision :  $Revision: $

FILE DESCRIPTION
VideoCore OS Abstraction Layer - public header file
=============================================================================*/

/**
  * \mainpage OS Abstraction Layer
  *
  * \section intro Introduction
  *
  * This abstraction layer is here to allow the underlying OS to be easily changed (e.g. from
  * Nucleus to ThreadX) and to aid in porting host applications to new targets.
  *
  * \subsection error Error handling
  *
  * Wherever possible, VCOS functions assert internally and return void. The only exceptions
  * are creation functions (which might fail due to lack of resources) and functions that
  * might timeout or fail due to lack of space. Errors that might be reported by the underlying
  * OS API (e.g. invalid mutex) are treated as a programming error, and are merely asserted on.
  *
  * \section thread_synch Threads and synchronisation
  *
  * \subsection thread Threads
  *
  * The thread API is somewhat different to that found in Nucleus. In particular, threads
  * cannot just be destroyed at arbitrary times and nor can they merely exit. This is so
  * that the same API can be implemented across all interesting platforms without too much
  * difficulty. See vcos_thread.h for details. Thread attributes are configured via
  * the VCOS_THREAD_ATTR_T structure, found in vcos_thread_attr.h.
  *
  * \subsection sema Semaphores
  *
  * Counted semaphores (c.f. Nucleus NU_SEMAPHORE) are created with VCOS_SEMAPHORE_T.
  * Under ThreadX on VideoCore, semaphores are implemented using VideoCore spinlocks, and
  * so are quite a lot faster than ordinary ThreadX semaphores. See vcos_semaphore.h.
  *
  * \subsection mtx Mutexes
  *
  * Mutexes are used for locking. Attempts to take a mutex twice, or to unlock it
  * in a different thread to the one in which it was locked should be expected to fail.
  * Mutexes are not re-entrant (see vcos_reentrant_mutex.h for a slightly slower
  * re-entrant mutex).
  *
  * \subsection evflags Event flags
  *
  * Event flags (the ThreadX name - also known as event groups under Nucleus) provide
  * 32 flags which can be waited on by multiple clients, and signalled by multiple clients.
  * A timeout can be specified. See vcos_event_flags.h. An alternative to this is the
  * VCOS_EVENT_T (see vcos_event.h) which is akin to the Win32 auto-reset event, or a
  * saturating counted semaphore.
  *
  * \subsection event Events
  *
  * A VCOS_EVENT_T is a bit like a saturating semaphore. No matter how many times it
  * is signalled, the waiter will only wake up once. See vcos_event.h. You might think this
  * is useful if you suspect that the cost of reading the semaphore count (perhaps via a
  * system call) is expensive on your platform.
  *
  * \subsection tls Thread local storage
  *
  * Thread local storage is supported using vcos_tls.h. This is emulated on Nucleus
  * and ThreadX.
  *
  * \section int Interrupts
  *
  * The legacy LISR/HISR scheme found in Nucleus is supported via the legacy ISR API,
  * which is also supported on ThreadX. New code should avoid this, and old code should
  * be migrated away from it, since it is slow. See vcos_legacy_isr.h.
  *
  * Registering an interrupt handler, and disabling/restoring interrupts, is handled
  * using the functions in vcos_isr.h.
  *
  */

/**
  * \file vcos.h
  *
  * This is the top level header file. Clients include this. It pulls in the platform-specific
  * header file (vcos_platform.h) together with header files defining the expected APIs, such
  * as vcos_mutex.h, vcos_semaphore.h, etc. It is also possible to include these header files
  * directly.
  *
  */

#ifndef VCOS_H
#define VCOS_H

#include "vcos_types.h"
#include "vcos_assert.h"
#include "vcos_generic_event_flags.h"
#include "vcos_msgqueue.h"
#include "vcos_logging.h"


// We dont need all
void *vcos_malloc(VCOS_UNSIGNED size, const char *description);

void *vcos_calloc(VCOS_UNSIGNED num, VCOS_UNSIGNED size, const char *description);

void vcos_free(void *ptr);

void vcos_logging_init(void);

VCOS_STATUS_T vcos_semaphore_create(VCOS_SEMAPHORE_T *sem,
                                                        const char *name,
                                                        VCOS_UNSIGNED initial_count);

void vcos_log_set_level(VCOS_LOG_CAT_T *category, VCOS_LOG_LEVEL_T level);



VCOS_STATUS_T vcos_timer_init(void);


VCOS_STATUS_T vcos_init(void);

void vcos_deinit(void);

void vcos_semaphore_delete(VCOS_SEMAPHORE_T *sem);


VCOS_STATUS_T vcos_event_flags_create(VCOS_EVENT_FLAGS_T *flags, const char *name);

int vcos_snprintf(char *buf, size_t buflen, const char *fmt, ...);

VCOS_STATUS_T vcos_semaphore_wait(VCOS_SEMAPHORE_T *sem);

VCOS_STATUS_T vcos_semaphore_post(VCOS_SEMAPHORE_T *sem);

#if 0
#include "interface/vcos/vcos_init.h"
#include "interface/vcos/vcos_semaphore.h"
#include "interface/vcos/vcos_thread.h"
#include "interface/vcos/vcos_mutex.h"
#include "interface/vcos/vcos_event_flags.h"
#include "interface/vcos/vcos_mem.h"
#include "interface/vcos/vcos_logging.h"
#include "interface/vcos/vcos_string.h"
#include "interface/vcos/vcos_event.h"
#include "interface/vcos/vcos_thread_attr.h"
#include "interface/vcos/vcos_tls.h"
#include "interface/vcos/vcos_reentrant_mutex.h"
#include "interface/vcos/vcos_named_semaphore.h"
#include "interface/vcos/vcos_quickslow_mutex.h"
#endif
#if VCOS_HAVE_QUEUE
#include "interface/vcos/vcos_queue.h"
#endif

#if VCOS_HAVE_LEGACY_ISR
#include "interface/vcos/vcos_legacy_isr.h"
#endif

#if VCOS_HAVE_TIMER
//#include "interface/vcos/vcos_timer.h"
#endif

#if VCOS_HAVE_MEMPOOL
#include "interface/vcos/vcos_mempool.h"
#endif

#if VCOS_HAVE_ISR
#include "interface/vcos/vcos_isr.h"
#endif

#if VCOS_HAVE_ATOMIC_FLAGS
#include "interface/vcos/vcos_atomic_flags.h"
#endif

#endif /* VCOS_H */

