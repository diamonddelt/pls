# pls

pls or _"path ls"_ is a CLI utility to make viewing the system PATH easier.

In Windows, there is already a command to view the entire list of environment variables: "SET"
However, `SET` is intended to be used primarily to manipulate or display ALL environment variables via cmd.exe, instead of the System PATH specifically. Additionally, it does not display each individual path in the combined PATH on a new line, making it difficult to find something specific.

## Getting Started

### Windows

1. Clone or download the repository: `git clone https://github.com/diamonddelt/pls.git`
2. Copy the *pls.exe* executable into your System32 directory `C:\Windows\System32`
3. In a new terminal or CMD shell, simply type `pls`

### Prerequisites

#### Windows

There are no prerequisites if you are using the provided .exe.

However, if you'd like to compile the code yourself, you will need a current version of Go installed and configured to run `go build`.

In cd into the `src/` directory of the project, and run `go build pls.go`. From there, you can copy the executable into your `%SYSTEMROOT%`, or wherever you'd like to invoke the utility from

### Usage

1. `pls` with zero arguments will invoke the "ls" or list function. This will simply list all elements on the System PATH, each on a new line.
2. `pls help` will display a help message
3. `pls version` will display the current version of the binary

## Authors

* **Ryan Rasti** - *Initial work* - [Blog](https://ryanrasti.info)
* **Brad Lugo** - *Design/Refactoring/Wisdom*


## TODO
* Add the ability to add or remove a path based on a flag "-a | --add" or "-d | --delete"