// +build ignore

static int answer;

int main()
{
	// Any library initialization comes here.
	answer = 42;
}

// Use the -99lib option to prevent the linker from eliminating this function.
int f42(int arg)
{
	return arg * answer;
}
