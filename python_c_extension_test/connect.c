#include "testmodule.h"

char testfunc_docs[] = "This are test functions";

PyMethodDef test_funcs[] = {
	{	"test",
		(PyCFunction)test,
		METH_NOARGS,
		testfunc_docs},
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
