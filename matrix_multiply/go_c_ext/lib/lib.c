#include "lib.h"

#include <stdlib.h>
#include <stdio.h>
#include <math.h>

void mm_destroy(int n, float **m) {
	int i;
	for (i = 0; i < n; ++i) free(m[i]);
		free(m);
}


float **matrix_init(int n) {
	float **m;
	int i;

	m = (float**)malloc(n * sizeof(void*));
	for (i = 0; i < n; ++i)
		m[i] = calloc(n, sizeof(float));
	return m;
}


void matmul(int n, float *const *a, float *const *b) {
	int i, j, k;
	float **m, **c;


	m = matrix_init(n); c = matrix_init(n);

	for (i = 0; i < n; ++i) 	// transpose
		for (j = 0; j < n; ++j)
			c[i][j] = b[j][i];

	for (i = 0; i < n; ++i) {
		float *p = a[i], *q = m[i];
		for (j = 0; j < n; ++j) {
			float t = 0.0, *r = c[j];
			for (k = 0; k < n; ++k)
				t += p[k] * r[k];
			q[j] = t;
		}
	}
	mm_destroy(n, c);
	
	float result = m[n/2][n/2];
	// return result;
}

