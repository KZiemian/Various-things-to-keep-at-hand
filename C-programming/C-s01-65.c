#include <stdio.h>

int main() {
    int secretNumber = 5;
    int guess = 0;

    while (guess != secretNumber) {
        printf("Enter guess: ");
        scanf("%d", &guess);
    }
    printf("You guess secret number: %d.\n", guess);



    return 0;
}