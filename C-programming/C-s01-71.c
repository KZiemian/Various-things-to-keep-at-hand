#include <stdio.h>

int main() {
    int intVar = 30;
    int* pInt = &intVar;

    printf("*pInt = %d\n", *pInt);

    *pInt = 40;

    printf("Po zmianie: *pInt = %d\n", *pInt);



    return 0;
}