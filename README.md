# Goploit Finder
 Go Search, this is a tool made in go search for exploits.

## Compile
 Before using the Goploit Finder, we need to compile. And for that to be done you need to have GO installed on your computer.
 - https://golang.org/
 
 To compile Goploit just execute this command inside the Goploit folder:

 MacOS and Windows users:

 `go build -o goploit main.go`
 
 Linux users:

 `go build -o goploit main.go && chmod +x goploit && sudo cp goploit /usr/bin`

## How to use
 For you to use Goploit is very simple! Enough you pass the `--name` parameter that will serve to inform you the name of the exploit/failure, or the framework. And for more commands just type the parameter `-h`.

 `./goploit --name [Framework Name]`

 ![Screenshot](/screenshot/screenshot.png)

## License
GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
