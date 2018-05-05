#ifndef __TESTMODULE_H__
#define __TESTMODULE_H__

#include <Python.h>

void mm_destroy(int n, double **m);
double **matrix_init(int n);
double **mm_mul(int n, double *const *a, double *const *b);
double  **convert_matrix(int N, double **c_matrix, PyObject *python_matrix);
PyObject * matmul(PyObject *self, PyObject *args);

#endif
