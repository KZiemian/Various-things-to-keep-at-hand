#include <stdio.h>

int main() {
    int secretNumber = 5;
    int guess = secretNumber + 1;
    int numberOfGuesses = 3;

    while ((guess != secretNumber) && (numberOfGuesses > 0)) {
        printf("Enter a number: ");
        scanf("%d", &guess);

        numberOfGuesses--;
    }

    if (numberOfGuesses >= 0) {
        printf("You guess secret number: %d.\n", secretNumber);
    } else {
        printf("You didn't guess secret number: %d.\n", secretNumber);
    }



    return 0;
}