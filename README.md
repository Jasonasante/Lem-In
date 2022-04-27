1) Authors: Abdul Raheem Khan and Jason Asante-Twumasi

2)Description: This is a program that creates n number of rooms and connects them via paths based on the data transmitted by the user and distributes the ants along the paths so that all ants reaches from start to end in a minimum of steps. 

3) Example Data given:

4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

3)How to run:

This program only takes 1 argument.

- go run . "filename.txt"

Example Output:

4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

L1-2 
L1-3 L2-2 
L1-1 L2-3 L3-2 
L2-1 L3-3 L4-2 
L3-1 L4-3 
L4-1 