#include <stdlib.h>
#include <stdio.h>

int fib(int n)
{
	switch (n) {
	case 0:
		return 0;
	case 1:
		return 1;
	default:
		return fib(n - 1) + fib(n - 2);
	}
}

int main(int argc, char **argv)
{
	if (argc != 2) {
		return 2;
	}

	int n = atoi(argv[1]);
	if (n <= 0 || n > 40) {
		return 1;
	}

	printf("%i\n", fib(n));
}
