# How to find the sum of digits of a positive number using recursion?

def getSumOfDigitsByRecursion(num):
     if num < 10:
          return num
     return num%10 + getSumOfDigitsByRecursion(round(num/10))


def getSumOfDigitsByIteration(num):
     sum = 0
     while num >= 10:
          sum += num % 10
          num = round(num / 10)
     sum += num
     return sum

print(f"Sum of digits of 123 is {getSumOfDigitsByRecursion(123)}")
print(f"Sum of digits of 125 is {getSumOfDigitsByRecursion(125)}")
print(f"Sum of digits of 1 is {getSumOfDigitsByRecursion(1)}")
print("------------------------------------------")
print(f"Sum of digits of 123 is {getSumOfDigitsByIteration(123)}")
print(f"Sum of digits of 125 is {getSumOfDigitsByIteration(125)}")
print(f"Sum of digits of 1 is {getSumOfDigitsByIteration(1)}")
