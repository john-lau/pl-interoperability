#include "lib.h"

#include <stdlib.h>
#include <stdio.h>
#include <math.h>

void swap(int * a, int * b) {
	int tmp = * a; 
	* a = * b; 
	* b = tmp;
}

void bubble_sort(int array[], int len) {
	for (int i = 0; i < len; i++) {
		for (int j = i; j < len; j++) {
			if (array[i] > array[j]) {
				swap( & array[i], & array[j]);
			}
		}
	}
}