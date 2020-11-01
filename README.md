![alt text](Golang.png)


---  

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<a rel="license"
href="http://creativecommons.org/licenses/by-nc-nd/4.0/"><img
alt="Creative Commons Licence" style="border-width:0"
src="https://i.creativecommons.org/l/by-nc-nd/4.0/88x31.png"
/></a><br /><span xmlns:dct="http://purl.org/dc/terms/"
href="http://purl.org/dc/dcmitype/Text" property="dct:title"
rel="dct:type">Go Beginners' Essentials</span> by <span
xmlns:cc="http://creativecommons.org/ns#"
property="cc:attributionName"><a href="http://palashkantikundu.in">Palash Kanti Kundu</a></span>
is licensed under a <a rel="license"
href="http://creativecommons.org/licenses/by-nc-nd/4.0/">Creative
Commons Attribution-NonCommercial-NoDerivatives 4.0
International License</a>.
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>  

---------


<br />
<br />
<br />


## Introduction

In this short guide, I will provide short handy notes for go language essentials. This is meant for Go Beginners but experienced programmers. So, any one coming from C, Java, C#, Python, Javascript background, can quickly catch up. I will also mention the name of the exercise files present in the same repository for easy reference.

This will be a collection of brief notes, just enough to get you started using Go. Taking your skills beyond the basics is completely your choice. Please do not treat this as complete language reference.

## Pre-requisites
1. A Computer :)
1. Go compiler. Head to [http://golang.org](http://golang.org) and download go for your OS. Steps are mentioned in the website.
1. Any Text Editor of your choice.
1. Prior experience with programming language like C, C++, Java, C#, Python. As the following notes will draw lots of analogy between the languages.

## Basics
Exercise File: [1.go](ex/1.go)

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

Exercise Files: [2.1.go](ex/2.1.go), 2.2.go, 2.3.go, 2.4.go

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
1. You can get around the compilation error of unused variable by throwing away the value. You can name thhe variable as `_` to throw away its value

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
1. If conditions have an optional assignment statement 
1. Also supports switch and like Java, C# it supports string variables in the condition
1. You do not need `break` statement in `switch`
1. Multiple cases can be comma separated
1. `default` is optional
1. `switch` can work without a condition, which is an alternative to if/else
1. `switch` also can have optional assignment
1. The starting parenthesis must be placed on the same line as of the if/switch statement

## Looping

Exercise File: 5.go, fizzbuzz.go

1. go has only `for` loop
1. `for` can be used as `while` and `while true` loop as well
1. You can use `break` and `continue` within loop

## Basic Understanding of Strings

Exercise File: 6.go

1. String is an immutable byte (uint8) array in go
1. In the built in package, you get some string handling operations like, we get `strlen()` in C
1. Go string slices can be taken based on array indexes, syntax is like list slices in python
1. Like Java, C# strings in go can be concatenated
1. As already mentioned earlier, strings are by default unicode in go
1. Multi line strings can be created using backticks 
1. In the `fmt` package, we also get a funtion `Sprintf`. This function returns a `string` instead of displaying on console.

## Basic Understanding of Slices

Exercise File: 7.go

1. A `slice` is a wrapper over actual array in go
1. The declaration of `array` and `slice` are similar except you remove the length in definition of `slice` while `array` is fixed length
1. Like `ArrayList` in Java, `slice` is graowable. For all practical purposes, I would like to think `slice` as `ArrayList` in Java or `List` in C# or `List` in python. Though the concept is not exactly the same, they work similar.
1. `len` function can be used over slices, the indexing and taking out a slice from an array or another slice is pretty similar as of python or as discussed in the basic understanding of strings
1. All elements of the same slice must be of same type
1. You can use `for` loop to iterate over a slice
1. You can use `range` keyword for iteration
1. You can use `range` with two variables that gives both values and indexes

