T *p;
T const *pc;
T volatile *pv;

pc = p;
pv = p;

int i = 37;
int const ci = 2;

void wi(int& r);
void wci(int const& rc);

wi(i);
wci(ci);
