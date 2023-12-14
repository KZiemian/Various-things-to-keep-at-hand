struct string {
  int size;
  int capacity;
  char *data;
};

const Char *c_str() const {
  // ...
  if (data[size()] != '\0') {
    data[size()] = '\0';
  }

  return data;
}
