#include <stdio.h>

int main() {
    int age = 30;
    int* pAge = &age;

    double gpa = 3.4;
    double* pGpa = &gpa;

    char grade = 'A';
    char* pGrade = &grade;

    printf("&age = %p\n", &age);
    printf("pAge = %p\n\n", pAge);

    printf("&gpa = %p\n", &gpa);
    printf("pGpa = %p\n\n", pGpa);

    printf("&grade = %p\n", &grade);
    printf("pGrade = %pN\n", pGrade);



    return 0;
}