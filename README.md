# reggie

This is an experimental Regex engine built from scratch, supporting most Regex grammars. 

It follows the following process:
- parse input string and and store each valid regex type inside of a parser context
- build a non-DFA graph, where each character acts as a node
- construct state transitions based on node type (or episilon for trivial transitions)

## How to run the program
### Docker (easiest/recommended)
The project contains a Dockerfile to easily run the program from a local container.

Steps:
````bash
$ sudo systemctl start docker      ### run Docker daemon beforehand
$ cd ~/location-of-reggie/reggie
$ docker build -t reggie .
$ docker run -it reggie
````
### Build from source
The project could also be built from source, given that the local machine has Golang installed

Steps:
````bash
$ git clone git@github.com:AarhamH/reggie.git
$ cd ~/location-of-reggie/reggie
$ go build -v
$ go run .
````

## How to use
The program is ran through a REPL, so the Regex checker evaluates to every input into the console. When first starting the user will be asked to provide a proper Regex expression, and every input after uses the Regex to evaluate the input string.

Use the following commands to change the REPL control flow:
- `change()` - allows the user to input a new Regex formula
- `exit()` - exits the REPL

## Example

![reggie_demo](https://github.com/user-attachments/assets/9d7d5da4-c48f-4b27-a041-ce4d48edcf11)
