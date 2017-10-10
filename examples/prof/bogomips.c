#include <stdlib.h>
#include <stdio.h>

// src: https://en.wikipedia.org/wiki/BogoMips#Computation_of_BogoMIPS
static void delay_loop(long loops)
{
	long d0 = loops;
	do {
		--d0;
	}
	while (d0 >= 0);
}

int main(int argc, char **argv)
{
	if (argc != 2) {
		return 2;
	}

	int n = atoi(argv[1]);
	if (n <= 0) {
		return 1;
	}

	delay_loop(n);
}
