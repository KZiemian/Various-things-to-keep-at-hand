// struct string {
//   int size;
//   int capacity;
//   char *data;
// };

// const Char *c_str() const {
//   // ...
//   if (data[size()] != '\0') {
//     data[size()] = '\0';
//   }

//   return data;
// }

#include <iostream>

int main() {
  // int&& rValueRefInt1 = 10;

  // std::cout << "rValueRefInt1 = " << rValueRefInt1 << ".\n";

  // std::string stringVar1 = "string one";
  // std::string stringVar2 = "a really long str";

  // // std::cout << "stringVar1 = " << stringVar1 << ".\n";

  // std::string&& stringVar3 = stringVar1 + " " + stringVar2;

  // std::cout << "stringVar3 = " << stringVar3 << ".\n";

  size_t n = 0;

  for (n = 0; n < 1'000'000'000; n++) {
    n += 1;

    std::cout << n << ".\n";
  }





  return 0;
}
