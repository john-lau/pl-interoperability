#ifndef __TESTMODULE_H__
#define __TESTMODULE_H__

#include <Python.h>


double **mm_mul(int n, double *const *a, double *const *b);

PyObject * matumul(PyObject *self, PyObject *args);

#endif
