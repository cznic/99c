#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>

#define BUFSIZE 1<<16

int main(int argc, char **argv)
{
	char *buf = malloc(BUFSIZE);
	if (!buf) {
		return 1;
	}

	for (int i = 1; i < argc; i++) {
		int fd = open(argv[i], O_RDWR);
		if (fd < 0) {
			return 1;
		}

		ssize_t n;
		while ((n = read(fd, buf, BUFSIZE)) > 0) {
			write(0, buf, n);
		}
	}
	free(buf);
}
