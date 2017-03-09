#include "replication.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int main(int argc, char **argv) {
	char *c = malloc(5);
	strncpy(c, "hi", 3);
    printf("Should print 'hi' twice below\n");
	PrintStringFromGo(c, 3);
	free(c);
	return 0;
}
