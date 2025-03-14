# fibonacci series
#  0, 1, 1, 2, 3, 5, 8, 13, 21

final_series = [0, 1]
n = 9


for  i, j in enumerate(final_series):

    if len(final_series) == n:
        break
    a, b = final_series[i], final_series[i+1]
    # print(i)
    final_series.append(a+b)
    # print(final_series)


print(final_series)

