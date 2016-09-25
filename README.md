# ctci6th

## 3. Stacks and Queues

3.1
Three in One: Describe how you could use a single array to implement three stacks.

3.2
Stack Min: How would you design a stack which, in addition to push and pop, has a function min which returns the minimum element? Push, pop and min should all operate in O(1) time.

3.3 Stack of Plates: Image a (literal) stack of plates. If the stack gets too high, it might topple. Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be composed of several stacks and should create a new stack once the previous one exceeds capacity. SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack).

FOLLOW UP

Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.

3.4
Queues via Stacks: Implement a MyQueue class which implements a queue using two stacks.

3.5
Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use an additional temporary stack, but you may not copy the elements into any other data structure (such as an array). The stack supports the following operations: push, pop, peek, and isEmpty.

3.6
Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a structly "first in, first out" basis. PEople must adopt either the "oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). They cannot select which specific animal they would like. Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCat. You may use the built-in LinkedList data structure.

## 4. Trees and Graphs

4.1
Route Between Nodes: Given a directed graph, design an algorithm to find out whether there is a route between two nodes.

4.2
Minimal Tree: Given a sorted (increasing order) array with unique integer elements, write an algorithm to create a binary search tree with minimal height.

4.3
List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth (e.g., if you have a tree with depth D, you'll have D linked lists).

4.4
Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of this question, a balanced tree is defined to be a stree such that the heights of the two subtrees of any node never differ by more than one.

4.5
Validate BST: Implement a function to check if a binary treee is a binary search tree.

4.6
Successor: Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a binary search tree. You may assume that each node has a link to its parent.

4.7
Build ORder: You are given a list of projects and a list of dependencies (which is a list of pairs of projects, where the second project is dependent on the first project). All of a project's dependencies must be built before the project is. Find a build order that will allow the projects to be built. If there is no valid build order, return an error.

EXAMPLE

Input:
projects: a, b, c, d, e, f
dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
Output: f, e, a, b, d, c

4.8
First Common Ancestor: Design an algorithm and write code to find the first common ancestor of two nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not necessarily a binary search tree.

4.9
BST Sequences: A binary search tree was created by traversin through an array from left to right and inserting each element. Given a binary search tree with distinct elements, print all possible arrays that could have led to this tree.

Input: 2 -> 1, 3
Output: {2, 1, 3}, {2, 3, 1}

4.10
Check Subtree: T1 and T2 are two very large binary trees, with T1 much bigger than T2. Create an algorithm to determine if T2 is a subtree of T1.

A tree T2 is a subtree of T1 if there exists a node n in T1 such that the subtree of n is identical to T2. That is, if you cut off the tree at node n, the two trees would be identical.

4.11
Random Node: You are implement a binary tree class from scratch which, in addition to insert, find, and delete, has a method getRandomNode() which returns a random node from the tree. All nodes should be equally likely to be chosen. Design and implement an algorithm for getRandomNode(), and explain how you would implement the rest of the methods.

4.12
Paths and Sums: You are given a binary tree in which each node contains an integer value (which might be positive or negative). Design an algorithm to count the number of paths that sum to a given value. The path does not need to start or end at the root or a leaf, but it must go downwards (Traveling only from parent nodes to child nodes).

## 5. Bit Manipulation

5.1
Insertion: You are given two 32-bit numbers, N and M, and two bit positions, i and j. Write a method to insert M into N such that M starts at bit j and ends at bit i. You can assume that bits j through i have enough space to fit all of M. That is, if M = 10011, you can assume that there are at least 5 bits between j and i. You would not, for example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.

EXAMPLE

```
Input:  N = 10000000000, M = 1011, i = 2, j = 6
Output: N = 10001001100
```

5.2
Binary to String: Given a real number between 0 and 1 (e.g., 0.72) that is passed in as a double, print the binary representation. If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR".

5.3
Flip Bits to Win: You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to find the length of the longest sequence of 1s you could create.

EXAMPLE

```
Input: 1775 (or: 11011101111)
Output: 8
```

5.4
Next Number: Given a positive integer, print the next smallest and the next largest number that have the same number of 1 bits in their binary representation.

5.5
Debugger: Explain what the following code does: `((n & (n-1)) == 0)`.

5.6
Conversion: Write a function to determine the number of bits you would need to flip to convert integer A to integer B.

EXAMPLE

```
Input: 29 (or: 11101), 15 (or: 01111)
Output: 2
```

5.7
Pairwise Swap: Write a program to swap odd and even bits in an integer with as few instructions as possible (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).

5.8
Draw Line: A monochrome screen is stored as a single array of bytes, allowing eight consecutive pixels to be stored in one byte. The screen has width w, where w is divisible by 8 (that is, no byte will be split across rows). The height of the screen, of course, can be derived from the length of the array and the width. Implement a function that draws a horizontal line from (x1, y) to (x2, y).

The method signature should look something like:

```
drawLine(byte[] screen, int width, int x1, int x2, int y)
```

## 10. Sorting and Searching

10.1
Sorted Merge: You are given two sorted arrays, A and B, where A has a large enough buffer at the end to hold B. Write a method to merge B into A in sorted order.

10.2
Group Anagrams: Write a method to sort an array of strings so that all the anagrams are next to each other.

10.3
Search in Rotated Array: Given a sorted array of n integers that has been rotated an unknown number of times, write code to find an element in the array. You may assume that the array was originally sorted in increasing order.

EXAMPLE

```
Input: find 5 in {15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
Output: 8 (the index of 5 in the array)
```

10.4
Sorted Search, No Size: You are given an array-like data structure Listy which lacks a size method. It does, however, have an elementAt(i) method that returns the element at index i in O(1) time. If i is beyond the bounds of the data structure, it returns -1. (For this reason, the data structure only supports positive integers.) Given a Listy which contains sorted, positive integers, find the index at which an element x occurs. If x occurs multiple times, you may return any index.

10.5
Sparse Search: Given a sorted array of strings that is interspersed with empty strings, write a method to find the location of a given string.

EXAMPLE

```
Input: ball, {"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
Output: 4
```

10.6
Sort Big File: Imagine you have a 20 GB file with one string per line. Explain how you would sort the file.

10.7
Given an input file with four billion non-negative integers, provide an algorithm to generate an integer that is not contained in the file. Assume you have 1 GB of memory available for this task.

FOLLOW UP

What if you have only 10 MB of memory? Assume that all the values are distinct and we now have no more than one billion non-negative integers.

10.8
Find Duplicates: You have an array with all the numbers from 1 to N, where N is at most 32,000. The array may have duplicate entries and you do not know what N is. With only 4 kilobytes of memory available, how would you print all duplicate elements in the array?

10.9
Sorted Matrix Search: Given an M x N matrix in which each row and each column is sorted in ascending order, write a method to find an element.

10.10
Rank from Stream: Imagine you are reading in a stream of integers. Periodically, you wish to be able to look up the rank of a number x (the number of values less than or equal to x). Implement the data structures and algorithms to support these operations. That is, implement the method track(int x), which is called when each number is generated, and the method getRankOfNumber(int x), which returns the number of values less than or equal to x (not including x itself).

EXAMPLE

```
Stream (in order of appearance): 5, 1, 4, 4, 5, 9, 7, 13, 3
getRankOfNumber(1) = 0
getRankOfNumber(3) = 1
getRankOfNumber(4) = 3
```

10.11
Peaks and Valleys: In an array of integers, a "peak" is an element which is greater than or equal to the adjacent integers and a "valley" is an element which is less than or equal to the adjacent intergers. For example, in the array {5, 8, 6, 2, 3, 4, 6}, {8, 6} are peaks and {5, 2} are valleys. Given an array of integers, sort the array into an alternating sequence of peaks and valleys.

EXAMPLE

```
Input: {5, 3, 1, 2, 3}
Output: {5, 1, 3, 2, 3}
```

## 14. Databases

Questions 1-3 use the following schema:

```sql
-- TODO: This is not the schema, look at page 173.

create table Apartments
(
	ApartmentID int not null identity,
	ApartmentAddress nvarchar(100),
	BuildingID int not null
)
go

alter table Apartments
add constraint PK_Apartments_ApartmentID primary key (ApartmentID)
go

create table Buildings
(
	BuildingID int not null identity,
	BuildingName nvarchar(100) not null,
	BuildingAddress nvarchar(500)
)
go

alter table Buildings
add constraint PK_Buildings_BuildingID primary key (BuildingID)
go

create table Apartments
(
	ApartmentID int not null identity,
	ApartmentAddress nvarchar(500),
	BuildingID int not null
)
go

alter table Apartments
add constraint PK_Apartments_ApartmentID primary key (ApartmentID)
go

alter table Apartments
add constraint FK_Apartments_BuildingID foreign key (BuildingID) references Buildings
go

create table Tenants
(
	TenantID int not null identity,
	TenantName nvarchar(100) not null,
	TenantAddress nvarchar(500)
)
go

alter table Tenants
add constraint PK_Tenants_TenantID primary key (TenantID)
go

create table TenantApartments
(
	TenantID int not null,
	ApartmentID int not null
)
go

alter table TenantApartments
add constraint PK_TenantApartments_TenantID_ApartmentID primary key (TenantID, ApartmentID)
go

alter table TenantApartments
add constraint FK_TenantApartments_TenantID foreign key (TenantID) references Tenants
go

alter table TenantApartments
add constraint FK_TenantApartments_ApartmentID foreign key (ApartmentID) references Apartments
go

```

14.1
Multiple Apartments: Write a SQL query to get a list of tenants who are renting more than one apartment.

14.2
Open Requests: Write a SQL query to get a list of all buildings, and the number of open requests (Requests in which status equals 'Open').

14.3
Close All REquests: Building #11 is undergoing a major renovation. Implement a query to close all requests from apartments in this building.

14.4
Joins: What are the different types of joins? Please explain how they differ and why certain types are better in certain situations.

14.5
Denormalization: What is denormalization? Explain the pros and cons.

14.6
Entity-Relationship Diagram: Draw an entity-relationship diagram for a database with companies, people, and professionals (people who work for companies).

14.7
Design Grade DAtabase: Imagine a simple database storing information for students' grades. Design what this database might look like and provide a SQL query to return a list of the honor roll students (top 10%), sorted by their grade point average.

## 16. Moderate

16.5
Factorial Zero: Write an algorithm which computers the number of trailing zeros in n factorial.

## 17. Hard

Sparse Similarity: The similarity of two documents (each with distinct words) is defined to be the size of the intersection divided by the size of the union. For example, if the document consists of integers, the similarity of {1, 5, 3} and {1, 7, 2, 3} is 0.4, because the intersection has size 2 and the union has size 5.

We have a long list of documents (with distinct values and each with an associated ID) where the similarity is believed to be "sparse." That is, any two arbitrarily selected documents are very likely to have similarity 0. Design an algorithm that returns a list of pairs of document IDs and the associated similarity.

Print only the pairs with similarity greater than 0. Empty documents should not be printed at all. For simplicity, you may assume each doucment is represted as an array of distinct integers.

EXAMPLE

```
Input:
13: {14, 15, 100, 9, 3}
16: {32, 1, 9, 3, 5}
19: {15, 29, 2, 6, 8, 7}
24: {7, 10}
Output:
ID1, ID2 : SIMILARITY
13, 19 : 0.1
13, 16 : 0.25
19, 24 : 0.14285714285714285
```
