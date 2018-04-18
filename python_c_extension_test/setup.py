#!/usr/bin/env python3

from distutils.core import setup, Extension

setup(
    name = "testmodule",
    version = "1.0",
    ext_modules = [Extension("testmodule", ["connect.c", "testmodule.c"])]
);
