#!/usr/bin/python3

import sympy

x = sympy.symbols('x')
y = sympy.symbols('y')
z = sympy.symbols('z')

# a = sympy.cos(x)**2 - sympy.sin(x)**2
# b = sympy.cos(2*x)

# print("a.equals(b) ->", a.equals(b))

# print("True ^ False ->", True ^ False)

# print("sympy.Xor(x, y) ->", sympy.Xor(x, y))

# print("type(sympy.Integer(1) + 1) =", type(sympy.Integer(1) + 1))
# print("type(1 + 1) =", type(1 + 1))

# sympy_var_1 = sympy.Integer(1)/sympy.Integer(3)
# print("sympy_var_1 =", sympy_var_1)
# print("type(sympy_var_1) =", type(sympy_var_1))

# sympy_var_2 = sympy.Rational(1, 2)

# print("sympy_var_2 =", sympy_var_2)
# print("type(sympy_var_2) =", type(sympy_var_2))

# expr_var_1 = x + 1/2
# print("expr_var_1 =", expr_var_1)

# print("type(expr_var_1) =", type(expr_var_1))

# expr_var_1 = sympy.cos(x) + 1
# value_var_1 = expr_var_1.subs(x, y)

# print("expr_var_1 =", expr_var_1)
# print("type(expr_var_1) =", type(expr_var_1))

# print("value_var_1 =", value_var_1)
# print("type(value_var_1) =", type(value_var_1))

# value_var_2 = expr_var_1.subs(x, 0)

# print("value_var_2 =", value_var_2)
# print("type(value_var_2) =", type(value_var_2))

# expr_var_1 = x**y
# print("expr_var_1 =", expr_var_1)
# print("type(expr_var_1) =", type(expr_var_1))

# expr_var_1 = expr_var_1.subs(y, x**y)
# print("expr_var_1 =", expr_var_1)

# expr_var_1 = expr_var_1.subs(y, x**x)
# print("expr_var_1 =", expr_var_1)

# print("type(expr_var_1) =", type(expr_var_1))

# expr_var_1 = sympy.sin(2*x) + sympy.cos(2*x)
# expr_var_2 = sympy.expand_trig(expr_var_1)

# print("expr_var_1 =", expr_var_1)
# print("type(expr_var_1) =", type(expr_var_1))

# print("expr_var_2 =", expr_var_2)
# print("type(expr_var_2) =", type(expr_var_2))

# expr_var_3 = expr_var_1.subs(sympy.sin(2*x), 2*sympy.sin(x)*sympy.cos(x))
# print("expr_var_3 =", expr_var_3)
# print("type(expr_var_3) =", type(expr_var_3))

# expr_var_1 = sympy.cos(x)
# print("expr_var_1.subs(x, 0) =", expr_var_1.subs(x, 0))
# print("expr_var_1 =", expr_var_1)
# print("x =", x)

# expr_var_1 = x**3 + 4*x*y - z

# print("expr_var_1 =", expr_var_1)
# print("expr_var_1.subs([...]) =", expr_var_1.subs([(x, 2), (y, 4), (z, 0)]))

# expr_var_1 = x**4 - 4*x**3 + 4*x**2 - 2*x + 3
# replacements = [(x**i, y**i) for i in range(5) if i % 2 == 0]

# print("expr_var_1 =", expr_var_1)
# print("expr_var_1.subs(replacements) =", expr_var_1.subs(replacements))

# str_expr_1 = "x**2 + 3*x - 1/2"
# expr_var_1 = sympy.sympify(str_expr_1)

# print("str_expr_1 =", str_expr_1)
# print("type(str_expr_1) =", type(str_expr_1))

# print("expr_var_1 =", expr_var_1)
# print("type(expr_var_1) =", type(expr_var_1))

expr_var_1 = sympy.sqrt(8)
print("expr_var_1 =", expr_var_1)
print("expr_var_1.evalf() =", expr_var_1.evalf())
