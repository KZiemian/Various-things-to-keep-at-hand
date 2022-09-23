#include <stdio.h>

int maxInt(int num1, int num2) {
    if (num1 >= num2) {
    return num1;
    } else {
        return num2;
    }
}

int main() {
    printf("maxInt(0, 1) = %d\n", maxInt(0, 1));
    printf("maxInt(1, 0) = %d\n", maxInt(1, 0));
    printf("maxInt(2, 0) = %d\n", maxInt(2, 0));
    printf("maxInt(0, 0) = %d\n", maxInt(0, 0));



    return 0;
}