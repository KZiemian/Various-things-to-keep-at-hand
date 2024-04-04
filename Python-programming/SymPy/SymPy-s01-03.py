#!/usr/bin/python3

import math
import numpy
import sympy

x = sympy.symbols('x')

# print("sympy.pi =", sympy.pi)
# print("sympy.pi.evalf() =", sympy.pi.evalf())
# print("sympy.pi.evalf(100) =", sympy.pi.evalf(100))

# expr_var_1 = sympy.cos(2*x)

# print("expr_var_1 =", expr_var_1)
# print("expr_var_1.evalf(subs={x: 2.4}) =", expr_var_1.evalf(subs={x: 2.4}))

# one = sympy.cos(1)**2 + sympy.sin(1)**2
# print("(one - 1).evalf() =", (one - 1).evalf())
# print("(one - 1).evalf(chop=True) =", (one - 1).evalf(chop=True))

a = numpy.arange(10)
expr_var_1 = sympy.sin(x)
f_1 = sympy.lambdify(x, expr_var_1, "numpy")
print("f_1(a) =", f_1(a))

f_2 = sympy.lambdify(x, expr_var_1, "math")
print("f_2(0.1) =", f_2(0.1))

def mysin(x):
    """
    My sine. Note that this is only accurate for small x.
    """

    return x

f_3 = sympy.lambdify(x, expr_var_1, {"sin":mysin})

print("f_3(0.1) =", f_3(0.1))
