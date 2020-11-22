# Galaxy Merchant Trading Guide
Galaxy Merchant Trading Guide is an app that used to translate Galactic Unit to Decimal. This app also provides parser that parse a query sentence.

See Demo Here: https://pikomonde.github.io/i-view-prospace

See Requirement Documents Here: [files/Prospace_Backend_Code_Challenge.pdf](files/Prospace_Backend_Code_Challenge.pdf)

# Make Command

1. `make test`

    Gives the test coverage of the project.

2. `make run`

    Run the project.

3. `make build-run`

    Generate an executable file named `service_app` and then run the executable file.

4. `make build`

    Generate an executable file named `service_app`.

5. `make js-build`

    Generate a .wasm (Web Assembly) file for the app.

6. `make js-run`

    Run a server to server single page web app.

7. `make js-build-run`

    Generate a .wasm (Web Assembly) file and then serve it for the web app.


# Project Structure

The architecture that is used in this project is Clean Code architecture, which is a seperation between deliveries, services, and repositories. The direction of the data is: `-> delivery -> service -> repository ->`

- Delivery

    Delivery is a part where the app's expected input and output. Currently this project has 1 delivery: `Command-Line Interface`.

- Service

    Service is where the app has all logic. These logic is divided by its use cases. Currently there are 3 services for this app:

    1. transnum

        `transnum` is a service that related to "translate number" uses cases, such as translating between decimal and galactic unit, also decimal and roman numeral. This service also provide simple entry for galactic unit and roman numeral relation.

    2. resource

        `resource` is a service that related to resource entry and calculation use cases. It provides 2 main job, to entry a resource's Credits, and to calculate the resource's Credits based on previous entry.

    3. parser

        `parser` is used to parse an input sentece into a task. This task is based on the other 2 services.
        Currently, there are 4 types of command that this parser supports, there are Galactic Unit Entry, Resource Price Entry, Galactic Unit Query, Resource Price Query.

- Repository

    Repository is the part where the app usually read/store data to db, microservice, or external process. Currently this project has no repository.


# Run Project Locally

## Prerequisites

- Go 1.11+ (go module support)

- http-server

    `npm install -g http-server`

## Build and Run the Project

1. Command Line Interface

    You can run it as a command line interface by using `go build -o service_app && ./service_app` or just run using `make build-run` command above.

2. From File

    You can parse query from a file and write the solution to other file by using redirect input and redirect output.

    Example query from a file:

    `make run < files/input01.txt`

    Example query from a file then write it in another file:

    `make run < files/input01.txt > files/output01.txt`

3. Static Web Page

    You can also create a static web page by using wasm (web assembly). This web page has different `main()` function than the Cli version. The `main()` function is located at `dist/wasm_app.go`
    
    Try running `make js-build-run`, then open yout browser at `http://localhost:8432`.

# Parser Command

Right now, the parser supports 4 commands:

```
GLOSARRY

galactic_unit           : galactic's numerical system, it can be more than 1 word
galactic_unit_word      : single word of galactic_unit
roman_numeral           : a roman numeral, it can be more than 1 character
roman_numeral_character : a single character of roman numeral (I,V,X,L,C,D,M)
resource_name           : the name of resource (Iron,Silver,Gold,etc)
decimal_number          : a decimal number
```

1. Galactic Unit Entry

    This is to assign galactic unit to a roman numeral character.
    
    Format: `[galactic_unit_word] is [roman_numeral_character]`

    Example: `glob is I`

2. Resource Price Entry

    This is to assign resources and its price in Credits.
    
    Format: `[galactic_unit] [resource_name] is [decimal_number] Credits`

    Example: `glob glob Silver is 34 Credits`

3. Galactic Unit Query

    This is to translate galactic unit to decimal.
    
    Format: `how much is [galactic_unit_word] ?`

    Example: `how much is pish tegj glob glob ?`

4. Resource Price Query

    This is to get and calculate resource's price based on previous resource price entry
    
    Format: `how many Credits is [galactic_unit_word] [resource_name] ?`

    Example: `how many Credits is glob prok Silver ?`

# Future Works

There are some thoughts on future works:

- Change (or asses first) the parse logic using regular expression

- Add db repository

- Add microservice delivery
