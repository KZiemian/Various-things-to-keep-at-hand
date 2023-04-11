int *p = nullptr;

int x = *p;

double weight(double impact){
  double weight_var;
  if (10.9 < impact && impact < 11.1){weight_var=1./11000. ;}
    else if (11.9<impact && impact<12.1){weight_var=1./12000. ;}
    else if (12.9<impact && impact<13.1){weight_var=1./13000. ;}
    else if (13.9<impact && impact<14.1){weight_var=1./14000. ;}
    else if (14.9<impact && impact<15.1){weight_var=1./15000. ;}
    else if (15.9<impact && impact<16.1){weight_var=1./16000. ;}
    else if (16.9<impact && impact<17.1){weight_var=1./17000. ;}
    else {
      cout << "Nieznana wartosc parametru impact:" << impact << '\n';

      exit(0);
    }

  return weight_var;
}
