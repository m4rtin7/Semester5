a = 0

def inpInt(msg):
    global a 
    a = int(input(msg))

def sumMultiple():
    global a
    if a == 0 :
        inpInt("insert int: ")
        if a == 0: return
        sumMultiple()
    elif a < 0: 
        a += 1
        sumMultiple()
        a -= 1
    else:
        a -= 1
        sumMultiple()
        a += 1

def main():
    global a
    sumMultiple()
    print("sum is " + str(a))

main()



