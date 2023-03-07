 getc(stdin);
i[17];
struct_ptr->c;
struct_value.c;
-i;
i++;
i--;
!znaleziony;
~0x7f;
*c_ptr;
&i;
sizeof(j);
(int)c;
i * j;
i / j;
i % j;
i + j;
i - j;
i << 2;
i >> 2;
i < j;
i <= j;
i > j;
i >= j;
i == 5;
i != 5;
7 & 5;
7 ^ 5;
7 | 5;
(i == 5) && (j == 6);
(i == 5) || (j == 6);
i > 4 ? i : j;
i = 7;
i *= 3;
i /= 4;
i %= 4;
i += 4;
i -= 4;
i &= 8;
i ^= 8;
i |= 8;
i <<= 8;
i >>= 8;

printf("Hello World!"), n = 7;

size_t strcat_exampleC1(const char *szStr1, const char *szStr2) {
  char szStrOut[MAX_PATH];

  strcpy(szStrOut, szStr1);
  strcat(szStrOut, " ");
  strcat(szStrOut, szStr2);
  puts(szStrOut);
}

void strcat_exampleC2(const char* szStr1, const char *szStr2) {
  char szStrOut[MAX_PATH];

  strncpy(szStrOut, szStr1, MAX_PATH);
  strncat(szStrOut, " ",    MAX_PATH);
  strncat(szStrOut, szStr2, MAX_PATH);
  puts(szStrOut);
}

void strcat_exampleC2(const char* szStr1, const char* szStr2) {
  char szStrOut[MAX_PATH];

  strncpy(szStrOut, szStr1, MAX_PATH - 1);
  strncat(szStrOut, " ",    MAX_PATH - 1);
  strncat(szStrOut, szStr2, MAX_PATH - 1);

  puts(szStrOut);
}

void strcat_exampleC3(const char *szStr1, const char *szStr2) {
  char szStrOut[MAX_PATH];

  strcpy_s(szStrOut, MAX_PATH - 1; szStr1);
  strcat_s(szStrOut, MAX_PATH - 1, " ");
  strcat_s(szStrOut, MAX_PATH - 1, szStr2);

  puts(szStrOut);
}

void strcat_exampleCPP(std::string str1, std::string str2) {
  std::string strOut = str1 + " " + str2;

  std::cout << strOut << std::endl;
}

void sprintf_exampleC1(char *szTarget, const char *szSource) {
  sprintf(szTarget, "%s", szSource);
  puts(szTarget);
}

void snprintf_exampleC2(char *szTarget, size_t bsize, const char *szSource) {
  snprintf(szTarget, bsize, "%s", szSource);
  puts(szTarget);
}
