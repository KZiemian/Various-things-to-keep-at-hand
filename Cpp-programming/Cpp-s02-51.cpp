#include <iostream>
// #include <cmath>

// int main() {
//   int x = 0;
//   // int y = 1;

//   int *const p = &x;


//   std::cout << " x: " << x << '\n';

//   std::cout << "*p: " << *p << '\n';


//   *p = 5;


//   std::cout << "*p: " << *p << '\n';

//   std::cout << " x: " << x << '\n';


//   // p = &y;


//   // std::cout << "*p: " << *p << '\n';




//   return 0;
// }

// int main() {
//   // double sqrtFiveTest = 2.23;

//   // std::cout << "2.23 * 2.23 = " << sqrtFiveTest * sqrtFiveTest << ".\n";

//   // sqrtFiveTest = 2.236;

//   // std::cout <<"2.236 * 2.236 = " << sqrtFiveTest * sqrtFiveTest << ".\n";

//   double piApprox = 3.1415;

//   std::cout << "cos(piApprox) = " << cos(piApprox) << ".\n";




//   return 0;
// }

// int main() {
//   std::cout << "1.41 * 1.41 = " << 1.41 * 1.41 << ".\n";
//   std::cout << "1.73 * 1.73 = " << 1.73 * 1.73 << ".\n";
//   std::cout << "2.23 * 2.23 = " << 2.23 * 2.23 << ".\n";
//   std::cout << "1.414 * 1.414 = " << 1.414 * 1.414 << ".\n";
//   std::cout << "2.236 * 2.236 = " << 2.236 * 2.236 << ".\n";




//   return 0;
// }




int main() {
  // std::cout << "Hello, World!\n";
  // std::cout << "Helo, \"World!\"\n";
  // std::cout << "Hello, \\ World!\n";

  // double x = 1;

  // x /= 0;

  // std::cout << "x = " << x << ".\n";

  int i = 0;

  double stateOfTheAccount = 10000;

  while (stateOfTheAccount < 20000) {
    stateOfTheAccount *= 1.05;

    i++;
  }

  std::cout << "Stan konta: " << stateOfTheAccount << ".\n";

  std::cout << "Zajelo to " << i << " lat.\n";





  return 0;
}
