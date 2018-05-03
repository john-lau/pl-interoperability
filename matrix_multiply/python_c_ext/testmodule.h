#ifndef __TESTMODULE_H__
#define __TESTMODULE_H__

#include <Python.h>

void mm_destroy(int n, double **m);
double **matrix_init(int n);
double **mm_mul(int n, double *const *a, double *const *b);
void convert_matrix(**double c_matrix, PyObject *python_matrix);
PyObject * matumul(PyObject *self, PyObject *args);

#endif
