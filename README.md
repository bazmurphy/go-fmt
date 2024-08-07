# How to use the `fmt` package in Go

fmt is a Go package that is used to format basic strings, values, inputs, and outputs.
It can also be used to print and write from the terminal.
Additionally, Go has formats called verbs.  
A verb is a placeholder for a named value to be formatted.
Let’s take a look at some of them:

- `%v` represents the named value in its default format
- `%T` represents the type of the value
- `%d` expects value to be an integer type of base 10
- `%b` expects value to be an integer type of base 2
- `%s` the bytes of string or slice
- `%f` expects value to have a float type

## `fmt.Print()`

prints whatever it gets to the terminal without adding any space or newlines unless it's explicitly coded.

spaces are added when the arguments are not strings.

```go
fmt.Print("Hello World")
```

## `fmt.Printf()`

provides custom formatting of input strings using one or more verbs and then prints the formatted string to the terminal without appending any space or newlines (unless explicitly coded).

```go
year := 2024

fmt.Printf("I was born in %d", year)
```

## `fmt.Println()`

is similar to `fmt.Print()` the difference being that it adds spaces between arguments and appends a newline at the end

```go
fmt.Println("Hello", "World") // spaces are added between arguments
fmt.Println("Hello World") // newline is appended at the end
```

## `fmt.Sprint()`

formats and returns the input string without printing anything to the terminal
notice how calling the `fmt.Sprint()` function doesn't print anything
instead it returns two values that we store in variables `a` and `b`
in order to see the value we have to use a print statement

```go
a := fmt.Sprint("Hello", "World")
b := fmt.Sprint("Hello World")

fmt.Println(a)
fmt.Println(b)
```

## `fmt.Sprintln()`

is similar in function to `fmt.Sprint()` except that it automatically adds spaces between arguments

```go
a := fmt.Sprintln("Hello", "World")

fmt.Println(a)
```

## `fmt.Sprintf()`

is used to format an input string. It also works like `fmt.Printf()` the difference being that `fmt.Sprintf()` returns the value as opposed to printing it

```go
name := "Baz"
s := fmt.Sprintf("My name is %s", name)

fmt.Print(s)
```

## `fmt.Scan()`

collects user input from the standard input and stores this in successive arguments
Spaces or newlines are considered multiple values and are stored in multiple arguments

```go
var name string

fmt.Println("What's your name?")

fmt.Scan(&name) // prompts the user for input and stores it in the name variable

fmt.Println("Nice to meet you", name)
```

## `fmt.Scanf()`

reads text from the standard input and puts away progressive space-separated values into progressive arguments as specified by the format. It also expects an address of each argument to be passed.

```go
var name string

fmt.Printf("What's your name: ")

fmt.Scanf("%s", &name) // takes values from standard input

fmt.Printf("Your name is %s\n", name)
```

## `fmt.Scanln()`

functions similarly to `fmt.Scan` but it stop scanning at a newline
this means that after the last item there must be a newline or EOF

```go
var z int

fmt.Scanln(&z)

fmt.Printf("Scanln: Z = %d\n", z)
```
