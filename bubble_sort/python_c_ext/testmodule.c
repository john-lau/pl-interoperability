#include <Python.h>
#include "testmodule.h"


void swap(int * a, int * b) {
	int tmp = * a; 
	* a = * b; 
	* b = tmp;
}

void sort(int array[], int len) {
	for (int i = 0; i < len; i++) {
		for (int j = i; j < len; j++) {
			if (array[i] > array[j]) {
				swap( & array[i], & array[j]);
			}
		}
	}
}

PyObject * c_sort(PyObject *self, PyObject *args) {

	int n;
	PyObject *list;

	if (!PyArg_ParseTuple(args, "iO", &n, &list)) {
		PyErr_SetString(PyExc_TypeError, "parameters incorrect");
		return NULL;
	}

	int c_list[n];
	PyObject *iter = PyObject_GetIter(list);

	if (!iter) {
		PyErr_SetString(PyExc_TypeError, "list iterator error");
		return NULL;
	}

	int i = 0;
	PyObject *next = PyIter_Next(iter);

	while(next) {
		// if (!PyFloat_Check(next)) {
		// 	PyErr_SetString(PyExc_TypeError, "expected an int");
		// 	return NULL;
		// }
		int foo = (int) next;
		c_list[i] = foo;
		next = PyIter_Next(iter);
		i++;
	}

	return Py_BuildValue("i", 1);

}