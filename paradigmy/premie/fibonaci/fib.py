from math import sqrt

def fib(n): 
    return 1/sqrt(5) * (((1+sqrt(5))/2)**n -((1-sqrt(5))/2)**n)



print(fib(1_000))