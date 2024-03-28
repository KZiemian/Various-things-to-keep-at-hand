#!/usr/bin/python3

# import math
import sympy

# print("math.sqrt(9) = ", math.sqrt(9))
# print("math.sqrt(8) = ", math.sqrt(8))

# print("sympy.sqrt(3) = ", sympy.sqrt(3))
# print("sympy.sqrt(8) = ", sympy.sqrt(8))
# x, y = sympy.symbols('x y')
# expr = x + 2*y

# print("expr = ", expr)

# print("expr + 1 = ", expr + 1)
# print("expr - x = ", expr - x)

# print("x*expr = ", x*expr)

# expanded_expr = sympy.expand(x*expr)

# print("expanded_expr = ", expanded_expr)

# factored_expr = sympy.factor(expanded_expr)

# print("factored_expr = ", factored_expr)

# x, t, z, nu = sympy.symbols('x t z nu')

# sympy.init_printing(use_unicode=True)

# print("sympy.diff(sympy.sin(x)*sympy.exp(x), x) = ",
#       sympy.diff(sympy.sin(x)*sympy.exp(x), x))

# print("sympy.integrate(sympy.exp(x)*sympy.sin(x) + sympy.exp(x)*sympy.cos(x), x) = ",
#       sympy.integrate(sympy.exp(x)*sympy.sin(x) + sympy.exp(x)*sympy.cos(x),
#                       x))

# print("sympy.integrate(sympy.sin(x**2), (x, -oo, +oo)) = ",
#       sympy.integrate(sympy.sin(x**2)))

# print("sympy.limit(sympy.sin(x)/x, x, 0) = ", sympy.limit(sympy.sin(x)/x, x,
#                                                           0))

# print("sympy.solve(x**2 - 2, x) = ", sympy.solve(x**2 - 2, x))

# y = sympy.Function('y')

# solution_var = sympy.dsolve(sympy.Eq(y(t).diff(t, t) - y(t), sympy.exp(t)),
#                            y(t))

# print("solution_var = ", solution_var)

# eigen_vals = sympy.Matrix([[1, 2], [2, 2]]).eigenvals()

# print("eigen_vals = ", eigen_vals)

# rewriten_function = sympy.besselj(nu, z).rewrite(sympy.jn)

# print("rewrite_function = ", rewriten_function)

# latex_output = sympy.latex(sympy.Integral(sympy.cos(x)**2, (x, 0, sympy.pi)))

# print("latex_output = ", latex_output)

# x + 1

# x = sympy.symbols('x')
# print("x + 1 = ", x + 1)

# x, y, z = sympy.symbols('x y z')

# print("x =", x)

# a, b = sympy.symbols('b a')

# print("a =", a)
# print("b =", b)

# crazy = sympy.symbols('unrelated')

# print("crazy + 1 =", crazy + 1)

# x = sympy.symbols('x')
# expr = x + 1
# x = 2

# print("expr =", expr)

x = sympy.symbols('x')
# expr = x + 1
# print("expr =", expr)
# print("expr.subs(x, 2) =", expr.subs(x, 2))

# print("x + 1 == 4 =>", x + 1 == 4)

# equation_var = sympy.Eq(x + 1, 4)

# print("equation_var =", equation_var)

# print("(x + 1)**2 == x**2 + 2*x + 1 ->", (x + 1)**2 == x**2 + 2*x + 1)

# a = (x + 1)**2
# b = x**2 + 2*x + 1

# c = x**2 - 2*x + 1

# print("sympy.simplify(a - b) =", sympy.simplify(a - b))

# print("sympy.simplify(a - c) =", sympy.simplify(a - c))
