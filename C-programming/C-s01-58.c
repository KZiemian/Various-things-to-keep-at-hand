#include <stdio.h>

int main() {
    char grade = 'G';

    switch (grade) {
        case 'A':
            printf("You did great!\n");
            break;
        case 'B':
            printf("You did alright!\n");
            break;
        case 'C':
            printf("You did poorly!\n");
            break;
        case 'D':
            printf("You did very bad!\n");
            break;
        case 'F':
            printf("You failed!\n");
            break;
        default:
            printf("Invalid grade.\n");
            break;
    }


    return 0;
}