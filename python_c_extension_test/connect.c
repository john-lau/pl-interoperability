#include "testmodule.h"

char testfunc_docs[] = "Prints out test";
char addfunc_docs[] = "This is a simple add function";

PyMethodDef test_funcs[] = {
	{	"test",
		(PyCFunction)test,
		METH_NOARGS,
		testfunc_docs},
	{	"add",
		(PyCFunction)add,
		METH_VARARGS,
		addfunc_docs},
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
