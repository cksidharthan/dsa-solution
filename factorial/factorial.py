def get_factorial(number) -> int:
     assert number >= 0
     if number < 2:
          return 1
     else:
          return get_factorial(number - 1) * number


print(f"Factorial of 5: {get_factorial(5)}")
print(f"Factorial of 10: {get_factorial(10)}")
print(f"Factorial of 0: {get_factorial(0)}")
