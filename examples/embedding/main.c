// +build ignore

#include <stdio.h>

int main() {
	int c;
	while ((c = getc(stdin)) != EOF) {
		printf("%c", c >= 'a' && c <= 'z' ? c^' ' : c);
	}
}
