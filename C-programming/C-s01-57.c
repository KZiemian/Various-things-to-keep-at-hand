#include <stdio.h>

int main() {
    double num1 = 0.0;
    double num2 = 0.0;
    char op = 'a';

    printf("Enter the number: ");
    scanf("%lf", &num1);

    printf("Enter operator: ");
    scanf(" %c", &op);

    // printf("Pierwsze tutaj\n");

    printf("Enter second number: ");
    scanf("%lf", &num2);

    // printf("Tutaj.\n");


    if (op == '+') {
        printf("%.2f + %.2f = %.2f\n", num1, num2, num1 + num2);
    } else if (op == '-') {
        printf("%.2f - %.2f = %.2f\n", num1, num2, num1 - num2);
    } else if (op == '*') {
        printf("%.2f * %.2f = %.2f\n", num1, num2, num1 * num2);
    } else if (op == '/') {
        printf("%.2f / %.2f = %.2f\n", num1, num2, num1 / num2);
    } else {
        printf("I don't know operator: %c.\n", op);
    }



    return 0;
}