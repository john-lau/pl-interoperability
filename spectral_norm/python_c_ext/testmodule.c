#include <Python.h>
#include "testmodule.h"

double eval_A(int i, int j) {
  return 1.0/((i+j)*(i+j+1)/2+i+1);
}

void eval_A_times_u(int N, const double u[], double Au[]) {
  int i,j;
  for(i = 0; i < N; i++) {
    Au[i] = 0;
    for(j=0;j<N;j++) {
      Au[i]+=eval_A(i,j)*u[j];
    }
  }
}


void eval_At_times_u(int N, const double u[], PyObject *Au) {
  int i = 0;
  int j;
  PyObject *iter = PyObject_GetIter(Au);
  if (!iter) {
    PyErr_SetString(PyExc_TypeError, "list iterator error");
    return;
  }
  PyObject *next = PyIter_Next(iter);
  while(next) {
    if (!PyFloat_Check(next)) {
      PyErr_SetString(PyExc_TypeError, "expected a float");
      return;
    }
    PyList_SetItem(Au, i, Py_BuildValue("f", 0.0));
    for(j = 0; j < N; j++) {
      PyObject *item = PyList_GetItem(Au, i);
      double item_val = PyFloat_AsDouble(item);
      double fin = item_val + eval_A(j,i)*u[j];
      PyList_SetItem(Au, i, Py_BuildValue("f", fin));
    }
    next = PyIter_Next(iter);
    i++;
  }
}
/* void eval_AtA_times_u_helper(PyObject *, PyObject *); */
PyObject * eval_AtA_times_u(PyObject *self, PyObject *args) {
  int N;
  PyObject *u_obj;
  PyObject *AtAu_obj;
  if (!PyArg_ParseTuple(args, "iOO", &N, &u_obj, &AtAu_obj)) {
    PyErr_SetString(PyExc_TypeError, "parameters incorrect");
    return NULL;
  }
  double u[N];

  PyObject *iter = PyObject_GetIter(u_obj);
  if (!iter) {
    PyErr_SetString(PyExc_TypeError, "list iterator error");
    return NULL;
  }
  int i = 0;
  PyObject *next = PyIter_Next(iter);
  while(next) {
    if (!PyFloat_Check(next)) {
      PyErr_SetString(PyExc_TypeError, "expected a float");
      return NULL;
    }
    double foo = PyFloat_AsDouble(next);
    u[i] = foo;
    next = PyIter_Next(iter);
    i++;
  }

  double v[N];
  eval_A_times_u(N, u, v);
  eval_At_times_u(N, v, AtAu_obj);

  return Py_BuildValue("i", 1);
}

