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


#APPROACH: BINARY SEARCH

#binary search for missing number
def findSmallestMissingNumber(array, start, end): 
  
    if (start > end): 
        return end + 1
  
    if (start != array[start]): 
        return start; 
  
    mid = int((start + end) / 2) 
  
    # Left half has all elements 
    # from 0 to mid 
    if (array[mid] == mid): 
        return findSmallestMissingNumber(array, 
                          mid+1, end) 
  
    return findSmallestMissingNumber(array,  
                          start, mid) 

#main function definition        
if __name__ == '__main__':

    arr = [0,1,2,6,9]
    n = len(arr)
    m = 10
    print(findSmallestMissingNumber(arr, 0, n-1))


    arr = [0, 1, 2, 3]
    n = len(arr)
    m = 5
    print(findSmallestMissingNumber(arr, 0, n-1))

    arr = [0, 1, 2, 3, 4, 5, 6, 7, 10]
    n = len(arr)
    m = 11
    print(findSmallestMissingNumber(arr, 0, n-1))
    