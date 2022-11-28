#include <iostream>

int main() {
  int lengthVar = 3;

  int const constLength = lengthVar + 4;

  int x[lengthVar];
  int y[constLength];

  lengthVar++;

  for (int i = 0; i < lengthVar; ++i) {
    x[i] = 10 + i;
  }

  for (int i = 0; i < lengthVar; ++i) {
    std::cout << "x[" << i << "]: " << x[i] << '\n';
  }

  for (int i = 0; i < constLength; ++i) {
    y[i] = -10 + i;
  }

  for (int i = 0; i < constLength; ++i) {
    std::cout << "y[" << i << "]: " << y[i] << '\n';
  }



  return 0;
}
