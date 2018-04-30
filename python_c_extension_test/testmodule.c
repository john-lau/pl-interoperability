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

PyObject * modList(PyObject *self, PyObject *args) {
  PyObject *obj;
  if (!PyArg_ParseTuple(args, "O", &obj)) {
    PyErr_SetString(PyExc_TypeError, "parameter must be a list.");
    return NULL;
  }
  PyObject *iter = PyObject_GetIter(obj);
  if (!iter) {
    PyErr_SetString(PyExc_TypeError, "list iterator error");
    return NULL;
  }
  int i = 0;
  PyObject *next = PyIter_Next(iter);
  while(next) {
    /* PyObject *next = PyIter_Next(iter); */
    /* if (!next) { */
    /*   // nothing left in the iterator */
    /*   break; */
    /* } */
    if (!PyFloat_Check(next)) {
      PyErr_SetString(PyExc_TypeError, "expected a float");
      return NULL;
    }
    double foo = PyFloat_AsDouble(next);
    foo += 1.0;
    PyList_SetItem(obj, i, Py_BuildValue("f", foo));
    next = PyIter_Next(iter);
    i++;
  }
  /* PyErr_SetString(PyExc_TypeError, "successful"); */
  return Py_BuildValue("O", obj);
}
