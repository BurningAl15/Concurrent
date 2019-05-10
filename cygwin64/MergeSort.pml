#define wait(s) atomic { s> 0 -> s--}
#define signal(s) s++

#define NUM_PROCESS 3

byte s1=0
byte s2=0

byte arr[N]={5,3,6,12,7,4,1,9}

proctype BubbleSort(int _arr[],int n)
{
    int i,j;
    int count1=0;
    int count2=0;
    do 
    :: count1 < n-count1-1 -> 
            do
            ::  count2<n-count2-1 -> 
                if
                :: arr[count2]> arr[count2+1] -> Swap(arr[count2],arr[count2+1])
                fi
                
            count2++
            od
        count1++
    od
}

proctype Swap(int a, int b)
{
    int temp=a;
    a=b;
    b=temp;
}

active proctype MergeSort(int _arr[])
{
    int tempArr1[4],tempArr2[4];
    int count=0;
    do
    :: count > 
    count++
    od
}


active[NUM_PROCESS] proctype P(){
    do
    ::  printf("P%d NCS 1\n", _pid )
        printf("P%d NCS 2\n", _pid )
        printf("P%d NCS 3\n", _pid )
        wait(s)
        printf("P%d CRITICAL 1\n", _pid )
        printf("P%d CRITICAL 2\n", _pid )
        printf("P%d CRITICAL 3\n", _pid )
        signal(s)
    od
}