// int n=0
int turn=1
bool end 

proctype P(){
    int t;
    int i=0;

    do
    ::i<10 -> 
        // t=n
        // n=t+1

        printf("P Msg 1:No critical section\n")
        printf("P Msg 2:No critical section\n")
        printf("P Msg 3:No critical section\n")

        do
        :: turn!=2-> printf("Algo")
        od
        
        printf("P Msg 1:No critical section X\n")
        printf("P Msg 2:No critical section X\n")
        printf("P Msg 3:No critical section X\n")
        turn=2
        end=true
        i++
    ::else -> break
    od
}

proctype Q(){
    // int t;
    int i=0;

    do
    ::i<10 -> 
        // t=n
        // n=t+1
        printf("P Msg 1:No critical section\n")
        printf("P Msg 2:No critical section\n")
        printf("P Msg 3:No critical section\n")

        do
        :: turn!=1-> printf("Algo")
        od

        printf("P Msg 1:No critical section X\n")
        printf("P Msg 2:No critical section X\n")
        printf("P Msg 3:No critical section X\n")
        turn=1
        end=true
        i++
    ::else -> break
    od
}

init{
    atomic{
        run P()
        run Q()
    }
    // (_nr_pr==1) -> printf("n=%d\n",n)
}