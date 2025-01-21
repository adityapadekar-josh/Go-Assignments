# Problem Statement

Given a string containing conversation between Alice and Bob. In the string, if it reaches `$`, it means the end of Alice's message and if it reaches `#`, it means the end of Bob's message. If it reaches `^`, it means the end of conversation, and you should ignore the string after that.

**Note**: The given string doesn't contain any spaces. If a message contains two continuous conversations from one person, it should be printed one after another as shown in the example.

Write a program to separate out messages from Alice and Bob. Write messages from Alice and Bob on separate channels. Whenever a message from Alice or Bob appears, print it in front of their name as shown in the example below.

**Note**: There is a single space before and after the colon (`:`), and no space before and after the comma.

## Input:

A string of messages in the form of a conversation.

## Output:

A list of messages from Alice and Bob, printed on separate channels as described.

## Example

### Input:

`"helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"`

### Output:

`alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?`
