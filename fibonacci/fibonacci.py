def get_fibonacci_series(num):
     series = [0, 1]
     while len(series) < num:
          series.append(series[len(series)-1] + series[len(series)-2])
     return series

print(f"Fibonacci series length - 8: {get_fibonacci_series(8)}")