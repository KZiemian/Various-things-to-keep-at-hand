#include <stdio.h>

int maxInt(int x, int y, int z) {
    if (x >= y) {
        if (x >= z) {
            return x;
        } else {
            return z;
        }
    } else {
        if (y >= z) {
            return y;
        } else {
            return z;
        }
    }
}



int main() {
    printf("maxInt(0, 1, 2) = %d\n", maxInt(0, 1, 2));
    printf("maxInt(2, 1, 0) = %d\n", maxInt(2, 1, 0));
    printf("maxInt(0, 2, 1) = %d\n", maxInt(0, 2, 1));



    return 0;
}