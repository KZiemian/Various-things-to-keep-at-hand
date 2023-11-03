// #include <iostream>

// int main() {
//   // int x = 3;
//   // int const *const p = &x;

//   // std::cout << "*p: " << *p << '\n';

//   // double costOfCar1 = 35000;
//   // double costOfCar2 = 25000;
//   double costOfCar1 = 100000;
//   double costOfCar2 = 80000;

//   double costOfFule = 4;

//   double carUseOfFule1 = 5;
//   double carUseOfFule2 = 8;

//   double yearOfUse = 10;
//   double kmPerYear = 25000;

//   costOfCar1 += costOfFule * (carUseOfFule1 / 100) * yearOfUse * kmPerYear;
//   costOfCar2 += costOfFule * (carUseOfFule2 / 100) * yearOfUse * kmPerYear;

//   std::cout << "costOfCar1 = " << costOfCar1 << ".\n";
//   std::cout << "costOfCar2 = " << costOfCar2 << ".\n";


//   if (costOfCar1 < costOfCar2) {
//     std::cout << "Wybierz samochod 1.\n";
//   } else if (costOfCar1 > costOfCar2) {
//     std::cout << "Wybierz samochod 2.\n";
//   } else {
//     std::cout << "Nie ma znaczenia ktory samochod wybierzesz.\n";
//   }

//   // std::cout << "carUseOfFule1 / 100 = " << carUseOfFule1 / 100 << ".\n";
//   // std::cout << costOfFule * (carUseOfFule1 / 100) * yearOfUse * kmPerYear
//   // 	    << ".\n";





//   return 0;
// }

#include <iostream>

int main() {
  // double var = 1E3;
  // int var = 1E3;
  // auto var = 1E3;
  // auto var = 1e3;

  // std::cout << "var = " << var << ".\n";
  // std::cout << "typeid(var) = " << typeid(var).name() << ".\n";

  // int integerVar = 0;

  // std::cout << "integerVar: " << integerVar << ".\n";

  // const int BOTTLE_NUMBER = 2;

  // std::cout << "BOTTLE_NUMBER type: " << typeid(BOTTLE_NUMBER).name()
  // 	    << ".\n";

  // // BOTTLE_NUMBER = 3;

  // std::cout << "BOTTLE_NUMBER = " << BOTTLE_NUMBER << ".\n";

  // const double CAN_VOLUME = 0.355;

  // std::cout << "CAN_VOLUME = " << CAN_VOLUME << ".\n";

  // const int CANS_PER_PACK = 6;
  // const double CAN_VOLUME = 0.355;
  // const double BOTTLE_VOLUME = 2.0;



  // double totalVolume = CANS_PER_PACK * CAN_VOLUME;


  // std::cout << "Sześciopak puszek po 12 uncji ma pojemność "
  // 	    << totalVolume << " litra.\n";

  // totalVolume = totalVolume + BOTTLE_VOLUME;

  // std::cout << "Sześciopak i dwulitrowa butelka mają pojemność "
  // 	    << totalVolume << " litra.\n";

  const int INT_VAR = 0;

  std::cout << "sizeof(INT_VAR): " << sizeof(INT_VAR) << ".\n";
  std::cout << "typeid(INT_VAR).name(): " << typeid(INT_VAR).name()
	    << ".\n";





  return 0;
}
