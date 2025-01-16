# Quadrilateral Area and Perimeter Calculation

This program allows the user to choose between a rectangle and a square, and it calculates both the area and perimeter of the chosen shape. The calculations are done using an interface `Quadrilateral` that is implemented by both the `Rectangle` and `Square` structs.

## Problem Description

1. **Quadrilateral Interface**:  
   The interface `Quadrilateral` should define two methods: `Area()` and `Perimeter()`.
   
2. **Implementing the Interface**:  
   The `Rectangle` and `Square` structs should implement the `Quadrilateral` interface by defining their own `Area()` and `Perimeter()` methods.

3. **Print Function**:  
   A `Print()` function should be defined which takes any shape implementing the `Quadrilateral` interface and prints the area and perimeter of that shape.

### Formulae:

- **Rectangle**:
    - Area: `Area = length * width`
    - Perimeter: `Perimeter = 2 * (length + width)`
  
- **Square**:
    - Area: `Area = side * side`
    - Perimeter: `Perimeter = 4 * side`

### Example Input/Output

**Input:**
```
Enter 1 for Rectangle or 2 for Square: 1
```

**Output:**
```
Area: 200
Perimeter: 60
```
