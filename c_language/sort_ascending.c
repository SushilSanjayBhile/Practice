//Program to sort inputs {only '0' & '1'} in ascending order.
//RESTRICTION: you've to go through whole list just for ONE TIME, throughout the program.


#include<stdio.h>

void main() {
	int arr[50],n,i,temp,k,j,valid;

	printf("Enter no of elements\n");
	scanf("%d",&n);

	for(i=0;i<n;i++) {
		valid=-1;

		printf("\nEnter values: %d",i);

		do{							//filter to check if I/P is other than {0,1}
			scanf("%d",&valid);
			printf("valid=%d",valid);

			if((valid!=1) && (valid!=0)) {
				printf("\nWrong valuedd, Enter again: ");
			}

		}while((valid!=1)&&(valid!=0));
		
		arr[i]=valid;
	}

	printf("\nBefore sorting: ");

	for(i=0;i<n;i++) {
		printf("%d ",arr[i]);
	}

		i=0;
		j=n-1;
//sorting
	for(k=0;k<n-3;k++) {


		if((arr[i]==1)&&(arr[j]==0)){
			temp=arr[i];
			arr[i]=arr[j];
			arr[j]=temp;
			i++;
			j--;
		}
		else {
			if(arr[i]==0){
				i++;}
			if(arr[j]==1){
				j--;}
		 }
		}

	/*	
		if(arr[i]>arr[i+1]) {
			temp=arr[i+1];
			arr[i+1]=arr[i];
			arr[i]=temp;
		}
	}
*/
	printf("\nAfter sorting: ");

	for(i=0;i<n;i++) {
		printf("%d ",arr[i]);
	}
}
