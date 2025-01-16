# Program to Accept Different Data Types Based on User Input

## Problem Statement

The program allows the user to input an integer value between 1 and 4. Each integer corresponds to an option that determines the type of object passed to a method named `AcceptAnything`. The `AcceptAnything` method will then print a message based on the type of object received.

### Types of Input:
1. **Integer**: If the user selects the integer option, an integer value will be passed.
2. **String**: If the user selects the string option, a string will be passed.
3. **Boolean**: If the user selects the boolean option, a boolean value will be passed.
4. **Custom Data Type Hello**: If the user selects the Hello option, an instance of a custom type `Hello` will be passed.

### `AcceptAnything` Method Behavior:
- For each type, the method should print a message in the following format:
  - **Integer**: `"This is a value of type Integer, <value>"`
  - **String**: `"This is a value of type String, <value>"`
  - **Boolean**: `"This is a value of type Boolean, <value>"`
  - **Custom Data Type `Hello`**: `"This is a value of type Hello, <value>"`

