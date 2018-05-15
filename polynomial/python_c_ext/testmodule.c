#include <Python.h>
#include "testmodule.h"

PyObject * poly(PyObject *self, PyObject *args) {
  float x;
  float mu = 10.0;
  float s;
  int j;
  float pol[100];

  if (!PyArg_ParseTuple(args, "f", &x)) {
    PyErr_SetString(PyExc_TypeError, "parameters incorrect");
    return NULL;
  }

  for (j=0; j<100; j++) {
      pol[j] = mu = (mu + 2.0) / 2.0;
    }
  s = 0.0;
  for (j=0; j<100; j++) {
      s = x*s + pol[j];
    }
  return Py_BuildValue("f", s);
}

