#addition of elements in list using reduce
def addList(list):
    print(reduce(lambda x,y: x+y, list))


#main function
if __name__ == "__main__":

    #list comprehension
    list = [int(x) for x in raw_input("Enter elementes: ").split()]

    addList(list)
