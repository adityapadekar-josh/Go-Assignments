# Panic and Recover

In the given code, the `accessSlice` function accepts a slice and an index.  
If the value is present in the slice at that index, the program should print the following:

```
item <index>, value <value at that index in slice>
```

However, if the index does not hold any value, it will lead to an "index out of range" panic in our program.

## Task

Complete the given code to recover from panic inside the `accessSlice` function when an "index out of range" panic occurs.  
Also, print the following after handling the panic:

```
internal error: <recovered panic value>
```

## Example Test Case 1:

### Input:
```
3
```

### Output:
```
item 3, value 6
```


