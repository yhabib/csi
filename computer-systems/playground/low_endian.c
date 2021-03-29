#include <stdio.h>
#include <string.h>

void show_bytes(unsigned char * start, int len) {
	int i;
	for (i = 0; i < len; i++) {
		printf("  %2.2x", start[i]);
	}
	printf("\n");
}

int main() {
	int i=0x01020304;
        show_bytes((unsigned char *) &i, sizeof(int));	
}
