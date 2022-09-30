#include <stdio.h>

int main() {
    char line[256];
    FILE* pFile = fopen("employees.txt", "r");

    fgets(line, 256, pFile);
    printf("line: %s;\n", line);

    fgets(line, 256, pFile);
    printf("line: %s;\n", line);

    fgets(line, 256, pFile);
    printf("line: %s;\n", line);
    
    fclose(pFile);



    return 0;
}