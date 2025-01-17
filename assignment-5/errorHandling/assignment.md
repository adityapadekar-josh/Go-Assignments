# Error Handling Example

In the given code, the `accessSlice` function accepts a slice and an index.  
If the value is present in the slice at that index, the program should print the following:

```
item <index>, value <value present at the index>
```

However, if the index does not hold any value, it will lead to an "index out of range" panic in our program.

## Task

To safeguard the program from panicking, add a condition to handle the case where:

```
index > lengthOfSlice - 1
```

If this condition is met, return an error from the `accessSlice` function with the following error message:

```
length of the slice should be more than index
```

Complete the given program to return an error from the `accessSlice` function and handle the returned error inside the `main` function to print the message.

## Example Test Case 1:

### Input:
```
3
```

### Output:
```
item 3, value 6
```

