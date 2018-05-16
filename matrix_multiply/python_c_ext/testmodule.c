#include <Python.h>
#include "testmodule.h"

void mm_destroy(int n, double **m) {
  int i;
  for (i = 0; i < n; ++i) free(m[i]);
  free(m);
}


double **matrix_init(int n) {
  double **m;
  int i;

  m = (double**)malloc(n * sizeof(void*));
  for (i = 0; i < n; ++i)
    m[i] = calloc(n, sizeof(double));
  return m;
}


double **mm_mul(int n, double *const *a, double *const *b) {
  int i, j, k;
  double **m, **c;

  m = matrix_init(n); c = matrix_init(n);

  for (i = 0; i < n; ++i) // transpose
    for (j = 0; j < n; ++j)
      c[i][j] = b[j][i];
  for (i = 0; i < n; ++i) {
    double *p = a[i], *q = m[i];
    for (j = 0; j < n; ++j) {
      double t = 0.0, *r = c[j];
      for (k = 0; k < n; ++k)
        t += p[k] * r[k];
      q[j] = t;
    }
  }
  mm_destroy(n, c);
  return m;
}


double **convert_matrix(const int N, double **c_matrix, PyObject *python_matrix) {
  c_matrix = matrix_init(N);

  PyObject *matrix_iter = PyObject_GetIter(python_matrix);
  if (!matrix_iter) {
    PyErr_SetString(PyExc_TypeError, "matrix iterator error");
    return NULL;
  }

  int i = 0;
  PyObject *matrix_next = PyIter_Next(matrix_iter);

  while (matrix_next) { //matrix_iter: 1 for each list, we need another iter to go into each individual list

    PyObject *row_iter = PyObject_GetIter(matrix_next);

    if (!row_iter) {
      PyErr_SetString(PyExc_TypeError, "row iterator error");
      return NULL;
    }

    int j = 0;
    PyObject *row_next = PyIter_Next(row_iter);
    while (row_next) { //now we are in a specific row of the matrix.

      if (!PyFloat_Check(row_next)) {
        PyErr_SetString(PyExc_TypeError, "expected a float");
        return NULL;
      }

      double index_val = PyFloat_AsDouble(row_next);
      c_matrix[i][j] = index_val;
      row_next = PyIter_Next(row_iter);
      j++;
    }

    //move on to the next row of the matrix, increment i
    i++;
    matrix_next = PyIter_Next(matrix_iter);
  }

  return c_matrix;
}


PyObject *matmul(PyObject *self, PyObject *args) {

  /*Parameters passed in are size of matrix, matrix a, and matrix b */
  int N;
  PyObject *a;
  PyObject *b;

  if (!PyArg_ParseTuple(args, "iOO", &N, &a, &b)) {
    PyErr_SetString(PyExc_TypeError, "parameters incorrect");
    return NULL;
  }

  double **a_matrix = NULL;
  double **b_matrix = NULL;
  double **result_matrix = NULL;

  a_matrix = convert_matrix(N, a_matrix, a);
  b_matrix = convert_matrix(N, b_matrix, b);

  // we want to put the result in matrix d instead, so we dont have to return anything in our function... 
  result_matrix = mm_mul(N, a_matrix, b_matrix);
  double result = result_matrix[N/2][N/2];

  // free allocated memory
  mm_destroy(N, a_matrix);
  mm_destroy(N, b_matrix);
  mm_destroy(N, result_matrix);

  return PyFloat_FromDouble(result);

}


