#include <stdio.h>

int main() {
#ifdef VERBOSE
	printf(GREETING);
#endif
}
