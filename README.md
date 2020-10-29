# go Beginners' Essentials Notes

In this short text, I will provide short handy notes for go language essentials. This is meant for Go Beginners but experienced programmers. So, any one coming from C, Java, C#, Python, Javascript background, can quickly catch up. I will also mention the name of the exercise files present in the same repository for easy reference.

This will be a collection of brief notes, just enough to get you started using Go. Taking your skills beyond the basics is completely your choice. Please do not treat this as complete language reference.

## Pre-requisites
1. A Computer :)
1. Go compiler. Head to http://golang.org and download go for your OS. Steps are mentioned in the website.
1. Any Text Editor of your choice.
1. Prior experience with programming language like C, C++, Java, C#, Python. As the following notes will draw lots of analogy between the languages.

## Basics
Exercise File: 1.go

1. Similar Commenting Styles from C, C++, Java
1. Semi-colon are not mandatory.
1. Has concept of packages like java and you have to import necessary packages to run the program
1. You must need a main package definition for running on console
1. You must need a main function definition for running on console
1. `go run` runs the file
1. `go build` creates an executable based on the platform you are running
1. Strings are automatically UTF. You do not need special packages to support multi-lingual system.
1. `go fmt` command formats the source code as per its own standards thus solving the problem of a multi-developer development system where each developer follows their own formatting styles.
1. Go built application has an associated application run time. No Virtual Machine is required to run the application.
1. Go has some of the Object Oriented Features - Interfaces, Custom Types, Structs, Methods
1. Go does not have these Object Oriented Features - Inheritance, Polymorphism, Method or Operator overloading, implicit numeric conversion, Exception Handling

## Variable Declaration

Exercise Files: 2.1.go, 2.2.go, 2.3.go, 2.4.go

1. Variable declaration is a mix of javascript var and C/Java type declaration - `var x int` 
1. Go has many integer definitions - `int8`, `int16`, `int32`, `int64`, `int`. The `int` type defaults to 32 bit or 64 bit based on platform. So, be careful if you have specific use case. Oherwise, you may go out of space or may get errors supporting your desired values.
1. Go also supports unsigned integer definitions like - `uint8`, `uint16`, `uint32`, `uint64` and `uint`. Automatic size definition goes for `uint` as well.
1. Go has an alias for `uint8` and that is `byte`
1. Floats are either `float32` or `float64`. No automatic definition for floats.
1. Default value for number type variables are 0 or 0
1. If your program consists of any unused variable, the program will not even compile. So, get rid of memory issues in the compile time itself. Imagine, you have declared an array of `int64` of size 86400 * 64 and left it unused, essentially you are wasting almost 42 MB of main memory. Go, gets rid of that during compile time itself.
1. Go, even raises compiler error for unused but assigned variables
1. Like unused variables, go raises compilation error for unused imports too
1. However, go does not raise error for a blank function.

## Strong type check system

Exercise Files: 3.1.go, 3.2.go, 3.3.go, 3.4.go, 3.5.go

1. Go has a very strong type check system in place
1. Go does not allow you to assign integer variables to float variables
1. Operation between an integer and a float is not allowed
1. Operation between an integer and an unsigned integer is not allowed
1. Operation between float32 and float64 is also not allowed
1. Division between two integers results an integer
1. To get the float result, declare all associated variables as floats
1. Like python you can leave the type inference on compiler. You can use `:=` operator for that
1. You can assign multiple variables on a single line using `,`

## Branching

Exercise Files: 4.1.go (Requires an integer console argument), 4.2.go
Usage: go run 4.1.go 7
       go run 4.2.go 1

1. Supports if else ladder, however like python you do not need parenthesis to be wrapped around the condition
1. Conditional statements can have `&&`and `||`
1. If conditions can also have an optional assignment statement 
1. Also supports switch and like Java, C# it also supports string variables in the condition
1. You do not need `break` statement in `switch`
1. Multiple cases can be comma separated
1. `default` is optional
1. `switch` can work without a condition, which is an alternative to if/else
1. `switch` also can have optional assignment
1. The starting parenthesis must be placed on the same line as of the if/switch statement
