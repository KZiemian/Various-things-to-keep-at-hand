#include <stdio.h>

int main() {
    FILE* pFile = fopen("employees.txt", "w");

    fprintf(pFile, "Jim, Salesman\nPam, Receptionit\nOscar, Accounting");

    fclose(pFile);



    return 0;
}