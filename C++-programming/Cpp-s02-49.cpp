#include <iostream>

int main() {
  int const n = 3;

  constexpr int max = n;

  constexpr int min = 1;


  std::cout << "max: " << max << '\n';

  std::cout << "min: " << min << '\n';



  return 0;
}
