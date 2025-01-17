# Return the Slices

## Problem Statement

Complete the program to return 3 slices of a given array based on the following conditions:

### Given Array:
```
arr = ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]
```

### Input:
Two space-separated integers representing `index1` and `index2`.

### Output:
You need to output 3 slices:
1. A slice containing all the elements from the start of the array to `index1`.
2. A slice containing all the elements from `index1` to `index2`.
3. A slice containing all the elements from `index2` to the end of the array.

### Conditions to Handle:
- If either of the input indexes is out of range of the given array, print `"Incorrect Indexes"` and exit the program.

---

### Example Test Case 1:
#### Input:
```
2 4
```
#### Output:
```
['qwe', 'wer', 'ert']
['ert', 'rty', 'tyu']
['tyu', 'yui', 'uio', 'iop']
```

### Example Test Case 2:
#### Input:
```
2 8
```
#### Output:
```
Incorrect Indexes
```
