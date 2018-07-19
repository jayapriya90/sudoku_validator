Requirements
------------
 This program requires Go version >= 1.10 

Assumptions
-----------
 1. Input is comma seperated string of rows

 Steps to run the solution
 -------------------------
 1. cd into the directory which contains sudoku_validator.go
 2. Run this command: go run sudoku_validator.go

 Tests handled (in the main function)
 -----------------------------------
 1. Empty grid
 2. Invalid number of rows
 3. Invalid number of columns
 4. Duplicate number in a row
 5. Duplicate number in a column
 6. Duplicate number in a block
 7. Valid sudoku

 Sample result
 -------------
[jayapriya@stux-2 sudoku]$ go run sudoku_validator.go
input1 is ''
false. Invalid grid. Input is empty

input2 is '759462813,463518297,128973645,697184532,584236971,231795486,345821769,912647358,876359124,123456789'
false. Invalid input grid. Number of rows is not equal to 9

input3 is '7594628135,463518297,128973645,697184532,584236971,231795486,345821769,912647358,876359124'
false. Invalid input grid. Number of columns is not equal to 9

input4 is '835416927,196857431,417293658,569134782,123678549,748529163,652781394,981345276,374962815'
false. Invalid input grid. Duplicate number in row

input5 is '835416927,296857431,817293658,569134782,123678549,748529163,652781394,981345276,374962815'
false. Invalid input grid. Duplicate number in column

input6 is '835416927,296857431,419293658,569134782,123678549,748529163,652781394,981345276,374962815'
false. Invalid input grid. Duplicate number in row

input7 is '835416927,296857431,417293658,569134782,123678549,748529163,652781394,981345276,374962815'
true.

Time Complexity: O(n^2)
--------------- 
 