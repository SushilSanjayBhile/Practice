/*
program to count the occurance of character in given string.
capital and small are differentiated.
*/

#include<stdio.h>
void print_table(char[]);


int main(){
    char a[50];

    printf("Enter a string without spaces: ");
    scanf("%s",a);
    print_table(a);

    print_table("SushIlSanJayBHiLE");
    print_table("aAAbAAaBbBaBa");
}


void print_table(char a[]){
    printf("%s\n",a);

    char cTemp;
    int i,iTemp;
    int res[52];

    for(i=0;i<52;i++){
        res[i] = 0;
    }
    for(i=0; a[i]!='\0';i++){
        if((a[i]>='a') && (a[i]<='z')){
            iTemp = a[i]-'a';
            res[iTemp] = res[iTemp] + 1;
        }
        else if((a[i]>='A') && (a[i]<='Z')){
            iTemp = a[i]- 39;
            res[iTemp] = res[iTemp] + 1;
        }
    }

    for(i=0; i<52; i++){
        if((i>25) && (res[i]>0)){
            printf("%c\t",(i+39));
        }
        else if(res[i]>0){
            printf("%c\t",(i+97));
        }
    }

    printf("\n");

    for(i=0;i<52;i++){
        if(res[i]>0){
            printf("%d\t",res[i]);
        }
    }
    
    printf("\n\n");
        
}
