void f(int n) {
	if (n) {
		f(n-1);
		return;
	}

	*(char *)n;
}

int main() {
	f(4);
}
