## ltdl.mk -- includable Makefile snippet
##
##   Copyright (C) 2003-2005, 2007, 2011-2015 Free Software Foundation,
##   Inc.
##   Written by Gary V. Vaughan, 2003
##
##   NOTE: The canonical source of this file is maintained with the
##   GNU Libtool package.  Report bugs to bug-libtool@gnu.org.
##
## GNU Libltdl is free software; you can redistribute it and/or
## modify it under the terms of the GNU Lesser General Public
## License as published by the Free Software Foundation; either
## version 2 of the License, or (at your option) any later version.
##
## As a special exception to the GNU Lesser General Public License,
## if you distribute this file as part of a program or library that
## is built using GNU libtool, you may include this file under the
## same distribution terms that you use for the rest of that program.
##
## GNU Libltdl is distributed in the hope that it will be useful,
## but WITHOUT ANY WARRANTY; without even the implied warranty of
## MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
## GNU Lesser General Public License for more details.
##
## You should have received a copy of the GNU LesserGeneral Public
## License along with GNU Libltdl; see the file COPYING.LIB.  If not, a
## copy can be downloaded from http://www.gnu.org/licenses/lgpl.html,
## or obtained by writing to the Free Software Foundation, Inc.,
## 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.
#####

ACLOCAL_AMFLAGS = -I ../m4
AUTOMAKE_OPTIONS = foreign
AM_CPPFLAGS =
AM_LDFLAGS =
BUILT_SOURCES =
include_HEADERS =
noinst_LTLIBRARIES =
lib_LTLIBRARIES =
EXTRA_LTLIBRARIES =
EXTRA_DIST =
CLEANFILES =
MOSTLYCLEANFILES =

# -I$(srcdir) is needed for user that built libltdl with a sub-Automake
# (not as a sub-package!) using 'nostdinc':
AM_CPPFLAGS            += -DLT_CONFIG_H='<$(LT_CONFIG_H)>' \
                          -DLTDL -I. -I$(srcdir) -Ilibltdl \
                          -I$(srcdir)/libltdl
AM_LDFLAGS             += -no-undefined
LTDL_VERSION_INFO       = -version-info 10:1:3


if INSTALL_LTDL
ltdlincludedir          = $(includedir)/libltdl
include_HEADERS        += ltdl.h
lib_LTLIBRARIES        += libltdl.la
endif

if CONVENIENCE_LTDL
noinst_LTLIBRARIES     += libltdlc.la
endif


libltdl_la_CPPFLAGS     = -DLTDLOPEN=$(LTDLOPEN) $(AM_CPPFLAGS)
libltdl_la_LDFLAGS      = $(AM_LDFLAGS) $(LTDL_VERSION_INFO) $(LT_DLPREOPEN)
libltdl_la_LIBADD       = $(ltdl_LTLIBOBJS)
libltdl_la_DEPENDENCIES = $(ltdl_LTLIBOBJS)

libltdlc_la_CPPFLAGS    = -DLTDLOPEN=$(LTDLOPEN)c $(AM_CPPFLAGS)
libltdlc_la_LDFLAGS     = $(AM_LDFLAGS) $(LT_DLPREOPEN)
libltdlc_la_LIBADD      = $(libltdl_la_LIBADD)
libltdlc_la_DEPENDENCIES= $(libltdl_la_DEPENDENCIES)

CLEANFILES                     += libltdl.la \
                                  libltdlc.la \
                                  libdlloader.la

## Automake-1.9.6 doesn't clean subdir AC_LIBOBJ compiled objects
## automatically:
CLEANFILES             += $(ltdl_LIBOBJS) $(ltdl_LTLIBOBJS)

EXTRA_DIST             += COPYING.LIB \
                          README

## --------------------------- ##
## Gnulib Makefile.am snippets ##
## --------------------------- ##

BUILT_SOURCES   += libltdl/$(LT_ARGZ_H)
EXTRA_DIST      += libltdl/lt__argz_.h

# We need the following in order to create an <argz.h> when the system
# doesn't have one that works with the given compiler.
all-local $(lib_OBJECTS): $(LT_ARGZ_H)
libltdl/lt__argz.h: libltdl/lt__argz_.h
        $(AM_V_at)$(mkinstalldirs) . libltdl
        $(AM_V_GEN)cp $(srcdir)/libltdl/lt__argz_.h $@-t
        $(AM_V_at)mv $@-t $@
MOSTLYCLEANFILES += libltdl/lt__argz.h \
                    libltdl/lt__argz.h-t
