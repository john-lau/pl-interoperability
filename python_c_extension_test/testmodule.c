#include <Python.h>
#include "testmodule.h"

PyObject * test(PyObject * self) {
	return PyUnicode_FromFormat("This is a test module!");
}
