# Game Battleship User Manual

This is a coding exercise project which from iSpot.tv. 
This is a command line, text-only game that simulates attacking battleships.
For more requirements, please refer to the document [requirements.md](docs/requirements.md)

### 1. Install Environment

This project is based on Golang 1.17, so, you may need install GoLang development environment in your local computer. 
you can go to [Golang official website](https://go.dev/) to download and install Golang development environment.
OR you can follow this instruction to install environment.

<font color=orange>**For Install Environment:**</font>

If you already install make command in your local, you can easily use the command as below to install Golang development environment.

```shell
make env
```

If you cannot use make command, you can open the file [Makefile](./Makefile) to execute the command by manually.

```yaml
# env target is used to build the environment for this Gin framework project.
env:
  go get -u github.com/gin-gonic/gin

# start target is used to build the whole project and run the program.
start:
  go run main.go

.PHONY: env start
```

you just need copy the command and paste it to the terminal tools to execute.

For example, execute this command below to build the project, then run the program.
```shell
go run main.go
```

So far, you already installed the necessary packages, and the environment should be ready.

<font color=orange>**For Unit Test:**</font>

Please use this command to run the unit test.
```shell
make test
```
OR, copy and execute this command:
```shell
go test -v -cover ./tests/
```

### 2. Directory Structure
```yaml
# Project directory structure description
Directory Structure:
⊢bootstrap            # Used to store the bootstrap program of the project. For example: Game Server, Command and so on.
⊢config               # Used to store the configuration structures.
⊢docs                 # Used to store the related docs of the project.
⊢global               # Used to store and define the global variables of the project
⊢test                 # For tests, unit tests
⊢utils                # Used to store some tool files and special function files .
```

### 3. Player Operation guide
3.1 About Game Players

There are only two players for this game. Player 1 and Player 2.

3.2 About Battleships

Each player has 4 battleships. They have different size, which are 2, 3, 4, and 5.

3.3 How To Play

This is a command line, text-only game, so players can only input the command to play the game. 
At any time in the game, player can input 'help' command to get help information.

<font color=orange>**Notice:**</font> This game is not case-sensitive. So, 'help' is the same as 'HELP'!

At the beginning of the game, each player will get a 10 X 10 grid with columns A-J and rows 1-10. 
So the command format that each player can input is between 'A1' and 'J10'. 
And the computer will automatically lay out all the battleships for both players when the game starts. 
Players only need to input the coordinate. For example: "B2", "G3" or "e5".

If command exceeds the coordinate range, the game will have corresponding prompts.

If hit a battleship, the game will have corresponding prompts.

If sank a battleship, the game will have corresponding prompts.

If admit defeat, the game will have corresponding prompts.

LOL!

3.4 How To Win

When you sink all the enemy's battleships, you will win the game!


<font color=orange>**Welcome to Battleship! Explore it! Enjoy it!**</font>
