#include <iostream>

int main() {
  int const number = 8675309;

  std::cout << "number: " << number << '\n';

  int n = 2 * number;

  std::cout << "n: " << n << '\n';

  // number = 2;

  // std::cout << "number: " << number << '\n';

  int i = 3;
  int const sizeArray = 3;

  i = 4;

  int x[i];
  int y[sizeArray];

  x[0] = 0;
  y[0] = 0;


  std::cout << "x[0]: " << x[0] << '\n';
  std::cout << "y[0]: " << y[0] << '\n';



  return 0;
}
