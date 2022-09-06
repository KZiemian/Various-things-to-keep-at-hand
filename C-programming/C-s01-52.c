#include <stdio.h>

void sayHi(char name[]) {
    printf("Hello %s.\n", name);
}

int main() {
    sayHi("Mike");
    sayHi("Tom");
    sayHi("Hank");



    return 0;
}