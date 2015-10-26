#include "vcos.h"
#include "vcos_types.h"
#include <stdlib.h>




void *vcos_malloc(VCOS_UNSIGNED size, const char *description) {
   return malloc(size);
}

void *vcos_calloc(VCOS_UNSIGNED num, VCOS_UNSIGNED size, const char *description) {
   return calloc(num, size);
}

void vcos_free(void *ptr) {
   free(ptr);
}

void * vcos_malloc_aligned(VCOS_UNSIGNED size, VCOS_UNSIGNED align, const char *description) {
  // return vcos_generic_mem_alloc_aligned(size, align, description);
    return malloc(size);
}

void vcos_abort(void)
{
    abort();
}
