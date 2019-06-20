# Find the smallest missing number
# Given a sorted array of n distinct integers where each integer is in the range from 0 to m-1 and m > n. Find the smallest number that is missing from the array.

# Examples
# Input: {0, 1, 2, 6, 9}, n = 5, m = 10
# Output: 3

# Input: {4, 5, 10, 11}, n = 4, m = 12
# Output: 0

# Input: {0, 1, 2, 3}, n = 4, m = 5
# Output: 4

# Input: {0, 1, 2, 3, 4, 5, 6, 7, 10}, n = 9, m = 11
# Output: 8


#APPROACH: LINEAR SEARCH


#find missing number function definition
def findSmallestMissingNumber(arr, n, m):
    flag = 0

    for i in range(m-1):
        if(arr[i] != i):
            print('missing number is '+str(i))
            flag = 1
            break

    if flag == 0:
        print('missing number is '+str(m-1))
        

#main function definition
if __name__ == '__main__':
    arr = [0,1,2,6,9]
    n = len(arr)
    m = 10
    findSmallestMissingNumber(arr, n, m)


    arr = [0, 1, 2, 3]
    n = len(arr)
    m = 5
    findSmallestMissingNumber(arr, n, m)

    arr = [0, 1, 2, 3, 4, 5, 6, 7, 10]
    n = len(arr)
    m = 11
    findSmallestMissingNumber(arr, n, m)
