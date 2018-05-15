#include "testmodule.h"

char poly_func_docs[] = "Prints out poly";

PyMethodDef test_funcs[] = {
	{	"poly",
		(PyCFunction)poly,
		METH_VARARGS,
		poly_func_docs},
	{	NULL}
};

char testmod_docs[] = "This is a test module";

PyModuleDef test_mod = {
	PyModuleDef_HEAD_INIT,
	"testmodule",
	testmod_docs,
	-1,
	test_funcs,
	NULL,
	NULL,
	NULL,
	NULL
};

PyMODINIT_FUNC PyInit_testmodule(void) {
	return PyModule_Create(&test_mod);
}
