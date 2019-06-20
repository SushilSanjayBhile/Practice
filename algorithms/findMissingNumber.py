# You are given a list of n-1 integers and these integers are in the range of 1 to n.
# There are no duplicates in list. 
# One of the integers is missing in the list. 
# Write an efficient code to find the missing integer.

# Example :
# I/P    [1, 2, 4, ,6, 3, 7, 8]
# O/P    5


#find missing number function definition
def findMissingNumber(InputArray):
    n = len(InputArray)+1
    total = n*(n+1)/2
    sum = 0

    for i in InputArray:
        sum += i
    
    print('list is:- ',InputArray)
    print('Missing number is:- '+ str(total-sum))


#main function definition
if __name__ == '__main__':
    ip = [1,2,4,6,3,7,8]

    findMissingNumber(ip)