def gen(n, counter_open, counter_close, ans):
    if counter_open + counter_close == 2 * n:
        print(ans)
        return ans
    if counter_open < n:
        gen(n, counter_open + 1, counter_close, ans + '(')
    if counter_open > counter_close:
        gen(n, counter_open, counter_close + 1, ans + ')')
gen(2, 0, 0, "")