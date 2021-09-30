a = 0

def inpInt():
    global a 
    a = int(input("insert int: "))

def multiply2():
    global a
    if a == 0:
        inpInt()
        return
    if a % 2 == 1:
        a -= 1
        a /= 2
        multiply2()
        a *= 2
        a += 1
    else:
        a /= 2
        multiply2()
        a *= 2



def main():
    global a
    inpInt()
    multiply2()
    print("product is " + str(a))

main()