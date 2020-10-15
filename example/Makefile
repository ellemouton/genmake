
CC=g++
CCFLAGS = -std=c++11

prog: fib.o main.o fib.h 
	$(CC) $(CCFLAGS) fib.o main.o -o prog


fib.o: fib.cpp fib.h 
	$(CC) $(CCFLAGS) fib.cpp -c

main.o: main.cpp fib.h 
	$(CC) $(CCFLAGS) main.cpp -c


clean:
	@rm *.o
	@rm prog
