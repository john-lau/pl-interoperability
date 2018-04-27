#include <Python.h>
#include "testmodule.h"

PyObject * test(PyObject * self) {
	return PyUnicode_FromFormat("This is a test module!");
}

PyObject * add(PyObject *self, PyObject *args) {
  int a;
  int b;

  if(!PyArg_ParseTuple(args, "ii", &a, &b)) {
    return NULL;
  }

  return Py_BuildValue("i", a + b);
}
