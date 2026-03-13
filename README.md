# **go-algebra**
> # **Index**
>   1. **[Description](#description)** 
>   1. **[Technologies Used](#technologies-used)**
>   1. **[Architectural Decisions](#architectural-decisions)**
>       1. **[High User Abstraction](#high-user-abstraction)**
>       1. **[Performance vs Legibility](#performance-vs-legibility)**
>       1. **[Abstract Syntax Tree](#abstract-syntax-tree)**
>   1. **[Multi-Service Communication](#multi-service-communication)**
>       1. **[Json](#json)**
>       1. **[String](#string)**
>   1. **[Setup](#setup)**
>   1. **[Usage](#usage)**
>       1. **[Equation Component](#equation-component)**
>       1. **[Expression Component](#expression-component)**
>   1. **[AST Structuring For Other Languages](#ast-structuring-for-other-languages)**
>   1. **[KNOWN ISSUES](#known-issues)**
>   1. **[TODO](#todo)**
> 
> <br>

<br>

> # **Description**
> **go-algebra** is a algebraic driven project to bring a solution to multi-service, equation relevant data handlers.
>
> The project aims to be of simple usage, and avoid language coupling the solutions.
>
> For this, a Abstract Syntax Tree (**[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)**) was introduced as core of algebraic expressions, as well as output json structures.
>
> For large scale jobs that process data, execute numeric methods, create equations, and expect to persist the result somewhere for another service to consume, this is the ideal way to reach your objective.
> 
> <br>

<br>

> # **Technologies Used**
> > ## **Golang**
> > Pure standard golang library as possible.
> 
> <br>

<br>

> # **Architectural Decisions**
> > ## **High User Abstraction**
> > Whoever consumes **go-algebra** should not be worried about the mathematical and algebraic complexities.
> >
> > **go-algebra** aims to open as little as possible of how equations works at low level, so the developer can actually focus on its business rules.
> > 
> > Simplicity is the core of this project.
>
> > ## **Performance vs Legibility**
> > Even though performance is often given away for legibility and organization, **go-algebra** have layers of legibility, where the trade-off is meaningful.
> >
> > * For user accessible methods, the performance is given away for maximum legibility since its the least impactful part of the library.
> > * For execution level, the legibility and language is more low-level driven, since is assumed that someone changing this code should actually know both, math and low-level programming, deeply.
> > * **go-algebra** do have a internal cache for pre-computed operations that works bitwise for maximum performance. Even low-level developers may have a hard time reading it.
>
> > ## **Abstract Syntax Tree**
> > **go-algebra** uses the concept of Abstract Syntax Tree (**[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)**) in a recursive object to simplify presentation as well as execution.
> >
> > In golang environment, this decision drastically simplify communication with other services, since a strict non-ambiguous pattern have been applied. More explained at **[Multi-Service Communication](#multi-service-communication)** topic.
>
> <br>

<br>

> # **Multi-Service Communication**
> ## **Json**
> > As explained at **[Abstract Syntax Tree](#abstract-syntax-tree)** topic, the internal structure of Expressions is an **[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)**, and its main advantage is a easy way to communicate equations around.
> > 
> > The Expression recursive behavior have a strict pattern already wrapped in json keys, so for a golang-golang service communication the `json.Marshal()` and `json.Unmarshal()` actually solves everything.
> > 
> > Any other language would have its own way to deal with json parsing, but the concept remains the same, language-free json structure.
> >
> > Following is a example of the json output:
> > ```json
> > {
> >     "type": "add",
> >     "args": [
> >         {
> >             "type": "integer",
> >             "value": 2
> >         },
> >         {
> >             "type": "mul",
> >             "args": [
> >                 {
> >                     "type": "float",
> >                     "value": 1.1
> >                 },
> >                 {
> >                     "type": "add",
> >                     "args": [
> >                         {
> >                             "type": "symbol",
> >                             "name": "x"
> >                         },
> >                         {
> >                             "type": "integer",
> >                             "value": 1
> >                         }
> >                     ]
> >                 }
> >             ]
> >         },
> >         {
> >             "type": "sin",
> >             "args": [
> >                 {
> >                     "type": "pow",
> >                     "args": [
> >                         {
> >                             "type": "symbol",
> >                             "name": "x"
> >                         },
> >                         {
> >                             "type": "integer",
> >                             "value": -1
> >                         }
> >                     ]
> >                 }
> >             ]
> >         },
> >         {
> >             "type": "cos",
> >             "args": [
> >                 {
> >                     "type": "symbol",
> >                     "name": "x"
> >                 }
> >             ]
> >         },
> >         {
> >             "type": "tan",
> >             "args": [
> >                 {
> >                     "type": "symbol",
> >                     "name": "x"
> >                 }
> >             ]
> >         },
> >         {
> >             "type": "exp",
> >             "args": [
> >                 {
> >                     "type": "symbol",
> >                     "name": "x"
> >                 }
> >             ]
> >         },
> >         {
> >             "type": "log",
> >             "args": [
> >                 {
> >                     "type": "symbol",
> >                     "name": "x"
> >                 }
> >             ]
> >         },
> >         {
> >             "type": "log",
> >             "args": [
> >                 {
> >                     "type": "symbol",
> >                     "name": "x"
> >                 },
> >                 {
> >                     "type": "integer",
> >                     "value": 10
> >                 }
> >             ]
> >         }
> >     ]
> > }
> > ```
> > The algebraic notation of this same Expression object is expressed at **[String](#string)** sub-topic.
>
> > ## **String**
> > **go-algebra** prepares Expression and Equation to be used as `fmt.Printf("%s")` formatter, since a `String()` method is prepared for its pointer.
> >
> > Even though `String()` is kind of heavy to process due to simplification and pre-processing, it runs only once due to caching of the result by branch, so even a further Simplify() calling would preserve cache over unchanged branches.
> > 
> > The string will be the most simplified version of your equation in the algebraic language.
> > 
> > For validation reasons, it has been enforced to be compatible with **[GeoGebra](https://www.geogebra.org/graphing)**.
> > 
> > Following a example of the string:
> > ```
> > Equation:   f(x) = 1.1(x +1) +sin(x^(-1)) +cos(x) +tan(x) +e^x +ln(x) +log(10, x) +2
> > Expression: 1.1(x +1) +sin(x^(-1)) +cos(x) +tan(x) +e^x +ln(x) +log(10, x) +2
> > ```
> > Equation differs so you can declare its signature. Other than that, is just a wrapper for Expression.
> >
> > It were decided to avoid usage of `/` and `√` symbols to prevent complex nested parenthesis handling, since it was becoming unnecessarily complex and heavy to process. You can observe it in the above example, `sin(1/x)` is rather presented as `sin(x^(-1))`.
> > 
> > So, as a reminder:
> > * Power nodes of negatives means inverse, and multiplication by a expression in this condition means division by.
> > * Power nodes of non-integer values mean root. A non-integer negative is a inverse of a root.
>
> <br>

<br>

> # **Setup**
> Open a terminal at your golang project directory and run:
> ```bash
> go get github.com/FabioLuisBoni/go-algebra@latest
> ```
> Remember to adjust to the desired version for a better stability.
>
> <br>

<br>

> # **Usage**
> **go-algebra** opens some methods for you to work with:
> > ## **Equation Component**
> > * `algebra_equation.NewEquation`
> >     - Creates a brand new `Equation` object.
> >     - Receives a `string`, representing its signature, as parameter.
> > * `.Evaluate`
> >     - Computes the result of the `Expression` it holds.
> >     - Receives a `float64`, representing the variable value, as parameter.
> > * `.String`
> >     - Produces a simplified, algebraic friendly and **[GeoGebra](https://www.geogebra.org/graphing)** compatible string.
> >     - Compatible with `fmt.Printf("%s")` callings.
>
> > ## **Expression Component**
> > * **`algebra_expression.Int`**
> >     - Creates a brand new `Expression` object that behaves like a integer constant.
> >     - Receives a `int` as parameter.
> > * **`algebra_expression.Float`**
> >     - Creates a brand new `Expression` object that behaves like a floating point constant.
> >     - Receives a `float64` as parameter.
> > * **`algebra_expression.Symbol`**
> >     - Creates a brand new `Expression` object that behaves like a symbol.
> >     - A symbol may be both, a variable, or a known constant.
> >     - List of known constants:
> >         - `e` for **[Euler's number](https://en.wikipedia.org/wiki/E_(mathematical_constant))** (`~2.718281828`).
> >         - `pi` for **[π](https://pt.wikipedia.org/wiki/Pi)** (`~3.141592654`)
> >     - Receives a `string` as parameter.
> > * **`algebra_expression.Sum`**
> >     - Creates a brand new `Expression` object that behaves like a sum of two or more `Expression` objects.
> >     - Receives multiple `Expression` objects as parameters.
> > * **`algebra_expression.Multiply`**
> >     - Creates a brand new `Expression` object that behaves like a multiplication of two or more `Expression` objects.
> >     - Receives multiple `Expression` objects as parameters.
> > * **`algebra_expression.Pow`**
> >     - Creates a brand new `Expression` object that behaves like a power of two `Expression` objects.
> >     - Receives two `Expression` objects as parameters:
> >         - First parameter is the base of the power.
> >         - Second parameter is the exponent of the power.
> > * **`algebra_expression.Sin`**
> >     - Creates a brand new `Expression` object that behaves like a sin of an `Expression` object.
> >     - Receives a single `Expression` object as parameter
> > * **`algebra_expression.Cos`**
> >     - Creates a brand new `Expression` object that behaves like a cosine of an `Expression` object.
> >     - Receives a single `Expression` object as parameter
> > * **`algebra_expression.Tan`**
> >     - Creates a brand new `Expression` object that behaves like a tangent of an `Expression` object.
> >     - Receives a single `Expression` object as parameter
> > * **`algebra_expression.Exp`**
> >     - Creates a brand new `Expression` object that behaves like a natural exponential by an `Expression` object.
> >     - Receives a single `Expression` object as parameter
> > * **`algebra_expression.Log`**
> >     - Creates a brand new `Expression` object that behaves like a logarithm of `Expression` objects.
> >     - Receives two `Expression` objects as parameters:
> >         - First parameter is the base of the logarithm.
> >         - Second parameter is the operand of the logarithm.
> > * **`algebra_expression.Ln`**
> >     - Creates a brand new `Expression` object that behaves like a natural logarithm of an `Expression` object.
> >     - Receives a single `Expression` object as parameter
> > * **`.Sum`**
> >     - Add `Expression` objects to the original `Expression` as sum operators.
> >     - Receives multiple `Expression` objects as parameters.
> >     - Make any preparation or reorganization that the original `Expression` object needs.
> > * **`.Subtract`**
> >     - Add `Expression` objects to the original `Expression` as subtraction operators.
> >     - Receives multiple `Expression` objects as parameters.
> >     - Make any preparation or reorganization that the original `Expression` object needs.
> > * **`.Multiply`**
> >     - Add `Expression` objects to the original `Expression` as multiplication operators.
> >     - Receives multiple `Expression` objects as parameters.
> >     - Make any preparation or reorganization that the original `Expression` object needs.
> > * **`.Divide`**
> >     - Add `Expression` objects to the original `Expression` as division operators.
> >     - Receives multiple `Expression` objects as parameters.
> >     - Make any preparation or reorganization that the original `Expression` object needs.
> > * **`.Execute`**
> >     - Computes the result of the whole `Expression` tree.
> >     - Receives a `float64`, representing the variable value, as parameter.
> > * **`.Equal`**
> >     - Compares two expressions and answer if they're equal.
> >     - Receives a `Expression` object as parameter.
> >     - The comparison goes really deep, caring about the behavior instead of structure, so expect to become a heavy processing for large branches.
> >     - Some expected behavior comparison includes things like:
> >         - `x^0 equals to cos(pi/2)`
> >         - `e equals to 2.718281828`
> >         - `pi equals to 3.141592654`
> > * **`.String`**
> >     - Produces a simplified, algebraic friendly and **[GeoGebra](https://www.geogebra.org/graphing)** compatible string.
> >     - Compatible with `fmt.Printf("%s")` callings.
> > 
> > <br>
> > <br>
> > 
> > The following methods are used by pre-computing reasons, and are internally cached on the `Expression` object itself. 
> > It may be helpful, so it was decided to open for user usage:
> > * `.IsConstant`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> > * `.IsSignalInvertible`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> >     - Basically tells if the overall signal of nested `Expression` objects may be simplified with a single signal inversion.
> >     - It is approximate of `.IsNegative` behavior, but no quite the same. 
> > * `.IsEvenInteger`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> > * `.IsOddInteger`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> > * `.IsAbsoluteOne`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> > * `.IsZero`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> > * `.IsEuler`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return a boolean reporting the result.
> > * `.IsNegative`
> >     - Computes the result of the whole `Expression` tree.
> >     - Return two booleans:
> >         - First reporting the result.
> >         - Second reporting is actually the result is applicable.
> >     - Unapplicable results are, for example, variable containing branches, since a variable renders the entire result dynamic.
>
> <br>

<br>

> # **AST Structuring For Other Languages**
> For compatibility, other languages may need to implements the **[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)** used by **go-algebra**.
>
> Following are the set of rules followed by each expected Expression json key:
> > ## **integer**
> > * Defines a integer constant.
> > * Must be a leaf.
> > * Must have only a int `value` field.
>
> > ## **float**
> > * Defines a floating point constant.
> > * Must be a leaf
> > * Must have only a floating point `value` field.
>
> > ## **symbol**
> > * Tricky, it may represent a variable or a known constant.
> > * List of known constants:
> >     - `e` for **[Euler's number](https://en.wikipedia.org/wiki/E_(mathematical_constant))** (`~2.718281828`).
> >     - `pi` for **[π](https://pt.wikipedia.org/wiki/Pi)** (`~3.141592654`)
> > * Must be a leaf
> > * Must have only a string `name` field.
> > 
> > Do not name any variable with same string as known constants to avoid unexpected behavior.
>
> > ## **add**
> > * Defines a sum operation (act as subtraction too, if an argument is negative).
> > * Never is a leaf
> > * Must have only the field `args` filled with at least 2 Expression json objects.
> > * No maximum limit of `args` field.
>
> > ## **mul**
> > * Defines a multiply operation (act as division too, if an argument is a power with negative exponent, also known as inverse).
> > * Never is a leaf
> > * Must have only the field `args` filled with at least 2 Expression json objects.
> > * No maximum limit of `args` field.
>
> > ## **pow**
> > * Defines a power operation (exponential included).
> > * Must have only the field `args` filled with exactly 2 Expression json object:
> >     - First argument must be the base.
> >     - Second argument must be the exponent.
>
> > ## **sin**
> > * Defines a sin operation.
> > * Must have only the field `args` filled with exactly 1 Expression json object.
>
> > ## **cos**
> > * Defines a cosine operation.
> > * Must have only the field `args` filled with exactly 1 Expression json object.
>
> > ## **tan**
> > * Defines a tangent operation.
> > * Must have only the field `args` filled with exactly 1 Expression json object.
>
> > ## **exp**
> > * Defines a natural exponential operation.
> > * Must have only the field `args` filled with exactly 1 Expression json object.
>
> > ## **log**
> > * Defines a logarithmic operation (natural logarithm included).
> > * Must have only the field `args` filled with 1 or 2 Expression json object.
> >     - Having only 1 argument, defines a ln (natural logarithm) and its argument is the operator.
> >     - Having 2 arguments:
> >         - First argument must be the operator.
> >         - Second argument must be the base.
> > 
> > It seems a little off this order, but trust me, it keep things simple at execution level since ln and log keep its operator at the same position.
>
> <br>

<br>

> # **KNOWN ISSUES**
> > ## **Expression Component**
> > > With exception of Execute method, it is well known that the library doesn't handle well multiple equal Symbol() leafs.
> > >
> > > We expect to solve this once a Simplify() method is introduced, and its helpers be able to temporarily simplify redundant, or even unnecessary, expression branches/leafs.
> 
> <br>

<br>

> # **TODO**
> 
> <br>

<br>

# **[BACK TO TOP](#top)**