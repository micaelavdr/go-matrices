
Start server using
```
go run .
```
In root folder, send request using
```
curl -F 'file=@matrix.csv' "localhost:8080/echo"
```

Run tests with
```
go test
```

# Matrix operations

Given the uploaded file:
```
1,2,3
4,5,6
7,8,9
```

1. /echo
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. /invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. /flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. /sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. /multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

