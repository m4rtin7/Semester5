a = 0

def inpInt():
    global a 
    a = int(input("insert int: "))

#toto riesenie funguje len ak je druhe cislo mensie ako 100.000
def multiply2():
    global a
    def codeInputes():
        global a
        if a == 0:
            inpInt()
            return
        else:
            a -= 1
            codeInputes()
            a += 100_000
    codeInputes()
    a = (a // 100_000) * (a % 100_000)


        



def main():
    global a
    inpInt()
    multiply2()
    print("product is " + str(a))

main()