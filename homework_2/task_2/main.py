def h_func(mas, x):
    h.append(0)
    deg.append(1)
    for i in range(0, len(mas)):
        h.append((h[-1] * x + mas[i]) % (10 ** 9 + 9))
        deg.append((deg[-1] * x) % (10 ** 9 + 9))
def check(f1, f2, k):
    first = h[f1 + k - 1] - h[f1 - 1] * deg[k]
    second = h[f2 + k - 1] - h[f2 - 1] * deg[k]
    if first % (10 ** 9 + 9) == second % (10 ** 9 + 9):
        return 1
    else:
        return 0
str = "zzz"
n = len(str)
mas = [ord(i) for i in str]
x = 27
deg = []
h = []
h_func(mas, x)
for k in range(1, n + 1):
    print("i: ", k)
    b = 1
    for i in range(1, n + 2 - k, k):
        print("j: ",  i)
        if check(1, i, k) == 0:
            b = 0
            break
    if check(1, k * (n // k) + 1, n % k) == 0:
        b = 0
    if b:
        print(k)
        raise SystemExit