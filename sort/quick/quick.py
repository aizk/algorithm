import random

def wrap(part):
    def qs(a, s, e):
        if s < e:
            p = part(a, s, e)
            qs(a, s, p - 1)
            qs(a, p + 1, e)
    return qs

def partition1(a, s, e):
    t = a[s] # 第一个数
    i = s
    j = e
    while i < j:
        # find element < t from right to left
        while i < j:
            if a[j] <= t:
                a[i] = a[j] # 此时 a[i] 所指的是一个空位
                break
            j -= 1
        while i < j:
            if a[i] > t:
                a[j] = a[i] # 此时 a[j] 所指的是一个空位，因为之前已经赋值给 a[i]
                break
            i += 1
    # i == j 时指针相等，因为必然会执行一个空位
    a[i] = t
    return i

def partition2(a, s, e):
    t = a[s]
    i = s
    j = e
    while True:
        while i < j:        
            if a[j] <= t: break
            j -= 1
        while i < j:
            if a[i] >= t: break # 条件必须是 >=
            i += 1
        if i < j:
            exch(a, i, j)
        else:
            return i

def exch(a, i, j):
    t = a[i]
    a[i] = a[j]
    a[j] = t

if __name__ == "__main__":
    a = [5, 4, 3, 2, 1]

    wrap(partition1)(a, 0, len(a) - 1)
    print(a)

    random.shuffle(a)
    #print(a)

    wrap(partition2)(a, 0, len(a) - 1)
    print(a)