#ifndef __TESTMODULE_H__
#define __TESTMODULE_H__

#include <Python.h>

/* PyObject * test(PyObject *); */
/* PyObject * add(PyObject *, PyObject *); */
double eval_A(int, int);
void eval_A_times_u(int, const double[], double[]);
void eval_At_times_u(int, const double[], PyObject *);
/* void eval_AtA_times_u_helper(PyObject *, PyObject *); */
PyObject * eval_AtA_times_u(PyObject *, PyObject *);

#endif
