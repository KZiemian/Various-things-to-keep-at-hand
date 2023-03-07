void sprintf_s_exampleC3(char *szTarget, size_t bsize, const char *szSource) {
  sprintf_s(szTarget, bsize, "%s", szSource);
  puts(szTarget);
}

void itoa_exampleC1(int a, int b) {
  char outstr1[MAX_PATH];
  char outstr2[MAX_PATH];

  _itoa(a, outstr1, 16);
  _itoa(b, outstr2, 16);

  printf("%s %s\n", outstr1, outstr2);
}

void itoa_exampleC2(int a, int b) {
  char outstr1[MAX_PATH];
  char outstr2[MAX_PATH];
  int  len1 = _countof(outstr1);
  int  len2 = _countof(outstr2);


  _itoa_s(a, outstr1, len1, 16);
  _itoa_s(b, outstr2, len2, 16);

  printf("%s %s\n", outstr1, outstr2);
}

int main(int argc, char *argv[]) {
  /* ... */



  return 0;
}

int main(int argc, char **argv) {
  /* ... */



  return 0;
}
