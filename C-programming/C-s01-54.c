#include <stdio.h>

int cube(int number) {
    return number * number * number;
}

int main() {
    printf("cube(-1) = %d\n", cube(-1));
    printf("cube(0) = %d\n", cube(0));
    printf("cube(1) = %d\n", cube(1));
    printf("cube(2) = %d\n", cube(2));
    printf("cube(3) = %d\n", cube(3));
    printf("cube(4) = %d\n", cube(4));
    printf("cube(5) = %d\n", cube(5));
    printf("cube(6) = %d\n", cube(6));
    printf("cube(7) = %d\n", cube(7));
    printf("cube(8) = %d\n", cube(8));
    printf("cube(9) = %d\n", cube(9));
    printf("cube(10) = %d\n", cube(10));

    

    return 0;
}