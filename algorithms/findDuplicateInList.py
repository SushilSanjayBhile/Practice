# Find duplicates in a given array when elements are not limited to a range
# Given an array of n integers. The task is to print the duplicates in the given array. If there are no duplicates then print -1.

# Examples:

# Input : {2, 10, 100, 2, 10, 11}
# Output : 2 10

# Input : {5, 40, 1, 40, 100000, 1, 5, 1}
# Output : 5 40 1



def findDuplicate(arr):
    dict = {}

    for ele in arr:
        try:
            dict[ele] += 1
        except:
            dict[ele] = 1

    print("List:- ")
    print(arr)
    print("\nDuplicates:-")

    for item in dict:
        if(dict[item] > 1):
            print(item)

    print("\n")
    
if __name__ == "__main__":
    list = [2, 10, 100, 2, 10, 11]
    findDuplicate(list)

    list = [5, 40, 1, 40, 100000, 1, 5, 1]
    findDuplicate(list)
