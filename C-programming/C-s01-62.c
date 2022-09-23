#include <stdio.h>
#include <string.h>

struct Student{
    char name[50];
    char major[50];
    int age;
    double gpa;
};

int main() {
    struct Student student1;
    student1.age = 22;
    student1.gpa = 3.2;
    strcpy(student1.name, "Jim");
    strcpy(student1.major, "Bussiness");

    printf("student1.name --> %s\n", student1.name);
    printf("student1.major --> %s\n", student1.major);
    printf("student1.age --> %d\n", student1.age);
    printf("student1.gpa --> %.2f\n", student1.gpa);



    return 0;
}