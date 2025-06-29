/*
 *  dl_helper.c  --  wrapper around lt_dlopen() and dlopen()
 *
 *  Copyright (C) 2016 by Massimiliano Ghilardi
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 */

#ifndef TWIN_DL_HELPER_H
#define TWIN_DL_HELPER_H

#ifndef TWIN_H
#include "twin.h"
#endif

#include <dlfcn.h>
#undef dlinit /* dlopen() requires no initialization */
#define dlinit_once() true
#define dlhandle void *
#define dlopen(name) dlopen((name), RTLD_NOW | RTLD_GLOBAL)
#define DL_PREFIX "lib"
#ifdef __APPLE__
#define DL_EXT ".dylib"
#else
#define DL_EXT ".so"
#endif
#define DL_SUFFIX "-" TWIN_VERSION_STR DL_EXT

#endif /* TWIN_DL_HELPER_H */
