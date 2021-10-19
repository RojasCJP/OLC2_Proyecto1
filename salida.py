a = 10
while a > 0:
    print(a)
    a = a - 1

while a < 5:
    a = a + 1
    if a == 3:
        print("a")
        continue
    elif a == 4:
        print("b")
        break
    print(a)
