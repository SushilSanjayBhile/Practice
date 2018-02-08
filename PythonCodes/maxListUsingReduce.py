#max in the list using reduce
def maxList(list):
    print(reduce(lambda x,y: x if x>y else y, list))



if __name__ == "__main__":

    #list comprehension
    list = [int(x) for x in raw_input("Enter elements: ").split()]

    maxList(list)
