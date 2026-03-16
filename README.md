# **go-algebra**
> # **Index**
>   1. **[Description](#description)** 
>   1. **[Technologies Used](#technologies-used)**
>   1. **[Architectural Decisions](#architectural-decisions)**
>       1. **[High User Abstraction](#high-user-abstraction)**
>       1. **[Performance vs Legibility](#performance-vs-legibility)**
>       1. **[Abstract Syntax Tree](#abstract-syntax-tree)**
>       1. **[Angular Units: Why Radians?](#angular-units-why-radians)**
>   1. **[Multi-Service Communication](#multi-service-communication)**
>       1. **[Json](#json)**
>       1. **[String](#string)**
>   1. **[Setup](#setup)**
>   1. **[Usage](#usage)**
>       1. **[Equation Component](#equation-component)**
>           1. **[Constructor Methods](#constructor-methods)**
>           1. **[Action Methods](#action-methods)**
>       1. **[Expression Component](#expression-component)**
>           1. **[Constructor Methods](#constructor-methods-1)**
>           1. **[Action Methods](#action-methods-1)**
>           1. **[Pre-Computing Methods](#pre-computing-methods)**
>       1. **[Least Squares Component](#least-squares-component)**
>           1. **[Constructor Methods](#constructor-methods-2)**
>           1. **[Action Methods](#action-methods-2)**
>       1. **[Frame Component](#frame-component)**
>           1. **[Constructor Methods](#constructor-methods-3)**
>           1. **[Action Methods](#action-methods-3)**
>   1. **[AST Structuring](#ast-structuring)**
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
> > ## **Angular Units: Why Radians?**
> > Any trigonometric operation will work expecting its value as radian (rad). Its the formal algebraic/calculus unit since it is purely mathematical measurements over a circle.
> > 
> > Its cleaner, more reliable, and universal if compared to degree (°).
> > 
> > Keep this in mind if you're creating a constant to define an angle, don't forget conversion if you using degree (°) unit for legibility in your system:
> > * `radians = degrees * pi / 180`
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
> > Even though `String()` is kind of heavy to process due to simplification and pre-processing, it runs only once due to caching of the result by branch, so even a further `Simplify()` calling would preserve cache over unchanged branches.
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
> > > ### **Constructor Methods**
> > > * `algebra_equation.NewEquation`
> > >     - Creates a brand new `Equation` object.
> > >     - Receives a `string`, representing its signature, as parameter.
> > > * `.SetExpression`
> > >     - Takes the given pointer into the equation, so it behaves exactly as the given `Expression`.
> > >     - Returns the "self", for easiness of construction.
> > 
> > > ### **Action Methods**
> > > * `.Evaluate`
> > >     - Computes the result of the `Expression` it holds.
> > >     - Receives a `float64`, representing the variable value, as parameter.
> > > * `.String`
> > >     - Produces a simplified, algebraic friendly and **[GeoGebra](https://www.geogebra.org/graphing)** compatible string.
> > >     - Compatible with `fmt.Printf("%s")` callings.
>
> > ## **Expression Component**
> > > ### **Constructor Methods**
> > > * **`algebra_expression.Int`**
> > >     - Creates a brand new `Expression` object that behaves like a integer constant.
> > >     - Receives a `int` as parameter.
> > > * **`algebra_expression.Float`**
> > >     - Creates a brand new `Expression` object that behaves like a floating point constant.
> > >     - Receives a `float64` as parameter.
> > > * **`algebra_expression.Symbol`**
> > >     - Creates a brand new `Expression` object that behaves like a symbol.
> > >     - A symbol may be both, a variable, or a known constant.
> > >     - List of known constants:
> > >         - `e` for **[Euler's number](https://en.wikipedia.org/wiki/E_(mathematical_constant))** (`~2.718281828`).
> > >         - `pi` for **[π](https://pt.wikipedia.org/wiki/Pi)** (`~3.141592654`)
> > >     - Receives a `string` as parameter.
> > > * **`algebra_expression.Sum`**
> > >     - Creates a brand new `Expression` object that behaves like a sum of two or more `Expression` objects.
> > >     - Receives multiple `Expression` objects as parameters.
> > > * **`algebra_expression.Multiply`**
> > >     - Creates a brand new `Expression` object that behaves like a multiplication of two or more `Expression` objects.
> > >     - Receives multiple `Expression` objects as parameters.
> > > * **`algebra_expression.Pow`**
> > >     - Creates a brand new `Expression` object that behaves like a power of two `Expression` objects.
> > >     - Receives two `Expression` objects as parameters:
> > >         - First parameter is the base of the power.
> > >         - Second parameter is the exponent of the power.
> > > * **`algebra_expression.Sin`**
> > >     - Creates a brand new `Expression` object that behaves like a sin of an `Expression` object.
> > >     - Receives a single `Expression` object as parameter
> > > * **`algebra_expression.Cos`**
> > >     - Creates a brand new `Expression` object that behaves like a cosine of an `Expression` object.
> > >     - Receives a single `Expression` object as parameter
> > > * **`algebra_expression.Tan`**
> > >     - Creates a brand new `Expression` object that behaves like a tangent of an `Expression` object.
> > >     - Receives a single `Expression` object as parameter
> > > * **`algebra_expression.Exp`**
> > >     - Creates a brand new `Expression` object that behaves like a natural exponential by an `Expression` object.
> > >     - Receives a single `Expression` object as parameter
> > > * **`algebra_expression.Log`**
> > >     - Creates a brand new `Expression` object that behaves like a logarithm of `Expression` objects.
> > >     - Receives two `Expression` objects as parameters:
> > >         - First parameter is the base of the logarithm.
> > >         - Second parameter is the operand of the logarithm.
> > > * **`algebra_expression.Ln`**
> > >     - Creates a brand new `Expression` object that behaves like a natural logarithm of an `Expression` object.
> > >     - Receives a single `Expression` object as parameter
> > > * **`.Sum`**
> > >     - Add `Expression` objects to the original `Expression` as sum operators.
> > >     - Receives multiple `Expression` objects as parameters.
> > >     - Make any preparation or reorganization that the original `Expression` object needs.
> > > * **`.Subtract`**
> > >     - Add `Expression` objects to the original `Expression` as subtraction operators.
> > >     - Receives multiple `Expression` objects as parameters.
> > >     - Make any preparation or reorganization that the original `Expression` object needs.
> > > * **`.Multiply`**
> > >     - Add `Expression` objects to the original `Expression` as multiplication operators.
> > >     - Receives multiple `Expression` objects as parameters.
> > >     - Make any preparation or reorganization that the original `Expression` object needs.
> > > * **`.Divide`**
> > >     - Add `Expression` objects to the original `Expression` as division operators.
> > >     - Receives multiple `Expression` objects as parameters.
> > >     - Make any preparation or reorganization that the original `Expression` object needs.
> > 
> > > ### **Action Methods**
> > > * **`.Execute`**
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Receives a `float64`, representing the variable value, as parameter.
> > > * **`.Equal`**
> > >     - Compares two expressions and answer if they're equal.
> > >     - Receives a `Expression` object as parameter.
> > >     - The comparison goes really deep, caring about the behavior instead of structure, so expect to become a heavy processing for large branches.
> > >     - Some expected behavior comparison includes things like:
> > >         - `x^0 equals to cos(pi/2)`
> > >         - `e equals to 2.718281828`
> > >         - `pi equals to 3.141592654`
> > > * **`.Clone`**
> > >   - Gives a brand new object containing a deep copy completely detached from the original.
> > >   - The cache of the `Expression` and all its branches are cloned too.
> > > * **`.ClearCache`**
> > >   - Reset cache status, allowing all the pre-computing methods to be ran again.
> > >   - Works recursively, calling in all underlying branches.
> > > * **`.String`**
> > >     - Produces a simplified, algebraic friendly and **[GeoGebra](https://www.geogebra.org/graphing)** compatible string.
> > >     - Compatible with `fmt.Printf("%s")` callings.
> > >     - Result is cached to run processing only once.
> > 
> > > ### **Pre-Computing Methods**
> > > The following methods are used for pre-computing reasons, and are internally cached on the `Expression` object itself. 
> > > 
> > > It may be helpful, so it was decided to open for user usage:
> > > * `.IsMalformedStructure`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsIndefiniteness`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsConstant`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsZero`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsAbsoluteOne`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsEuler`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsFraction`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsInteger`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsEvenInteger`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsOddInteger`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > > * `.IsSignalInvertible`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return a boolean reporting the result.
> > >     - Basically tells if the overall signal of nested `Expression` objects may be simplified with a single signal inversion.
> > >     - It is approximate of `.IsNegative` behavior, but no quite the same. 
> > > * `.IsNegative`
> > >     - Computes the result of the whole `Expression` tree.
> > >     - Return two booleans:
> > >         - First reporting the result.
> > >         - Second reporting if result is actually applicable.
> > >     - Unapplicable results are, for example, branch containing variable leaf, since a variable renders the entire result dynamic.
> > >
> > > They potentially call each other in recursive manner thorough the whole branches of the `Expression`, but it were designed to run only once in each `Expression` and make assumptions over other conflicting methods.
> > > For example:
> > > * `.IsZero` internally calls `.IsConstant`, so both will sure to be cached after `.IsZero` calling. Even further, if `.IsZero` is true, it assumes and cache right away:
> > >     - `IsAbsoluteOne` as `false`.
> > >     - `IsEuler` as `false`.
> > >     - `IsFraction` as `false`.
> > >     - `IsInteger` as `true`.
> > >     - `IsEvenInteger` as `true`.
> > >     - `IsOddInteger` as `false`.
> > > * `.IsConstant` itself may call pre-computing over its underlying branches:
> > >     - `.IsZero`
> > >     - `.IsAbsoluteOne`
> > >     - `.IsEuler`
> > >
> > > Pre-computing methods runs only until the first deterministic information, so, its not guarantee that every branch and/or leaf have been pre-computed.
> > > 
> > > Once pre-computed, the method never runs again, it get the value from the cache right away.
> > > 
> > > The cache belongs to the `Expression` object, and if you move this branch around, the cache will remain intact, since its answer only changes if its branches changes.
> > > For this reason, is highly advisable that you don't manually modify branches of a `Expression` object, since its result would potentially be cached. Instead, try using a **go-algebra** method.
> > >
> > > If a method you need is not present, please, feel free to suggest.
> 
> > ## **Least Squares Component**
> > > ### **Constructor Methods**
> > > * `algebra_numeric_method.NewLeastSquares`
> > >     - Creates a brand new `LeastSquares` object.
> > >     - Receives a `string` as parameter, it will copy it in any equation it produces as the signature.
> > > * `.BaseOn`
> > >   - Defines the basis functions for the least squares operation using the structure of the provided equation.
> > >   - Since least squares works by adjusting the coefficients of given terms (see **[Wiki](https://en.wikipedia.org/wiki/Least_squares)**), `Expression` type determines how the basis is extracted:
> > >       - `Sum` **(Recommended)**: The primary type for multi-term regression. The library treats every immediate child of the `Sum` as an independent basis function.
> > >       - `Float`/`Integer`: Interpreted as a single constant term.
> > >       - All Other Types: The entire expression is treated as a single basis function with one coefficient to be adjusted.
> > 
> > > ### **Action Methods**
> > > * `.Solve`
> > >   - Solves the least squares problem for the given `Frame` of points and returns a new `Equation` object.
> > >   - The returned `Equation` object follows the structure defined in `LeastSquare.Base`, but with all coefficients calculated to minimize the sum of squared residuals.
> > >   - This method does not modify the existing `LeastSquares` instance, and is memory safe for the brand new `Equation` object returned.
> 
> > ## **Frame Component**
> > > ### **Constructor Methods**
> > > * `.AddPoint`
> > >   - Add a cartesian point into the frame.
> > 
> > > ### **Action Methods**
> > > * `.Sort`
> > >   - Sorts the data within the `Frame` according to the specified sorting strategy.
> > > * `.String`
> > >   - Returns a human-readable representation of the `Frame`.
> > >   - Compatible with `fmt.Printf(%s)` callings.
> > >   - Implements the `fmt.Stringer` interface for seamless integration with print functions.
> > > * `.CSV`
> > >   - Serializes the `Frame` into a text format optimized for spreadsheet compatibility.
> > >   - If `Frame.Name` is populated, the first line of the output will be the name followed by a newline.
> > >   - The data is always preceded by the fixed header line: `x;y`.
> > >   - Uses a semicolon (`;`) as delimiter to separate the x and y coordinates of each point.
>
> <br>

<br>

> # **AST Structuring**
> For compatibility, other languages may need to implements the **[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)** used by **go-algebra**.
>
> Remember that angles should be stored in radians, as explained at **[Angular Units: Why Radians?](#angular-units-why-radians)** sub-topic.
>
> Following are the set of rules followed by each expected `Expression` json key:
> > ## **integer**
> > * Defines a integer constant.
> > * Must be a leaf.
> > * Must have only a int `value` field.
>
> > ## **float**
> > * Defines a floating point constant.
> > * Must be a leaf.
> > * Must have only a floating point `value` field.
>
> > ## **symbol**
> > * Tricky, it may represent a variable or a known constant.
> > * List of known constants:
> >     - `e` for **[Euler's number](https://en.wikipedia.org/wiki/E_(mathematical_constant))** (`~2.718281828`).
> >     - `pi` for **[π](https://pt.wikipedia.org/wiki/Pi)** (`~3.141592654`)
> > * Must be a leaf.
> > * Must have only a string `name` field.
> > 
> > Do not name any variable with same string as known constants to avoid unexpected behavior.
>
> > ## **add**
> > * Defines a **[addition](https://en.wikipedia.org/wiki/Addition)** operation (act as subtraction too, if an argument is negative).
> > * Never is a leaf.
> > * Must have only the field `args` filled with at least 2 `Expression` json objects.
> > * No maximum limit of `args` field.
>
> > ## **mul**
> > * Defines a **[multiplication](https://en.wikipedia.org/wiki/Multiplication)** operation (act as division too, if an argument is a power with negative exponent, also known as inverse).
> > * Never is a leaf.
> > * Must have only the field `args` filled with at least 2 `Expression` json objects.
> > * No maximum limit of `args` field.
>
> > ## **pow**
> > * Defines a **[power](https://en.wikipedia.org/wiki/Exponentiation)** function, mathematically known as exponentiation.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 2 `Expression` json object:
> >     - First argument must be the base.
> >     - Second argument must be the exponent.
>
> > ## **exp**
> > * Defines a **[natural exponential](https://en.wikipedia.org/wiki/Exponential_function)** function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **sin**
> > * Defines a **[trigonometric](https://en.wikipedia.org/wiki/Trigonometric_functions)** sine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **cos**
> > * Defines a **[trigonometric](https://en.wikipedia.org/wiki/Trigonometric_functions)** cosine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **tan**
> > * Defines a **[trigonometric](https://en.wikipedia.org/wiki/Trigonometric_functions)** tangent function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **asin**
> > * Defines a **[trigonometric](https://en.wikipedia.org/wiki/Trigonometric_functions)** arcsine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **acos**
> > * Defines a **[trigonometric](https://en.wikipedia.org/wiki/Trigonometric_functions)** arccosine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **atan**
> > * Defines a **[trigonometric](https://en.wikipedia.org/wiki/Trigonometric_functions)** arctangent function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **sinh**
> > * Defines a **[hyperbolic](https://en.wikipedia.org/wiki/Hyperbolic_functions)** sine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **cosh**
> > * Defines a **[hyperbolic](https://en.wikipedia.org/wiki/Hyperbolic_functions)** cosine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **tanh**
> > * Defines a **[hyperbolic](https://en.wikipedia.org/wiki/Hyperbolic_functions)** tangent function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **asinh**
> > * Defines a **[hyperbolic](https://en.wikipedia.org/wiki/Hyperbolic_functions)** arcsine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **acosh**
> > * Defines a **[hyperbolic](https://en.wikipedia.org/wiki/Hyperbolic_functions)** arccosine function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **atanh**
> > * Defines a **[hyperbolic](https://en.wikipedia.org/wiki/Hyperbolic_functions)** arctangent function.
> > * Never is a leaf.
> > * Must have only the field `args` filled with exactly 1 `Expression` json object.
>
> > ## **log**
> > * Defines a **[logarithmic](https://en.wikipedia.org/wiki/Logarithm)** function (**[natural logarithm](https://en.wikipedia.org/wiki/Natural_logarithm)** included).
> > * Never is a leaf.
> > * Must have only the field `args` filled with 1 or 2 `Expression` json object.
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
> > > With exception of `Execute()` method, it is well known that the library doesn't handle well multiple equal `Symbol()` leafs, specially at `Equal()` callings.
> > >
> > > We expect to solve this once a `Simplify()` method is introduced, and its helpers be able to temporarily simplify redundant, or even unnecessary, expression branches/leafs.
> >
> > > `Execute()` method will not care for true symbology, any variable will be treated as "x" variable at a 2d equation behavior.
> > >
> > > We expect to solve this in the future introducing an object as argument for `Execute()` linking a value to a symbol, but special attention should be given for simplicity be kept.
> > 
> > > The trigonometric functions only support `Sine`, `Cosine` and `Tangent`. 
> > > 
> > > `Cosecant`, `Secant` and `Cotangent` can be achieved by:
> > > ```golang
> > > // Cosecant
> > > Pow(
> > >   Sin(
> > >     Symbol("x"),
> > >   ),
> > >   Int(-1),
> > > )
> > > 
> > > // Secant
> > > Pow(
> > >   Cos(
> > >     Symbol("x"),
> > >   ),
> > >   Int(-1),
> > > )
> > > 
> > > // Cotangent
> > > Pow(
> > >   Tan(
> > >     Symbol("x"),
> > >   ),
> > >   Int(-1),
> > > )
> > > ```
> > > And its indefiniteness are validated through power condition.
> > >
> > > That said, `Arcsine`, `Arccosine`, `Arctangent` and their counterparts `Arccosecant`, `Arcsecant` and `Arccotangent` will be introduced further in the future.
> > > 
> > > Hyperbolic functions can actually be constructed via the library itself. It is not in our map the development of dedicated methods to deal with them, unless sponsorship request it for simplification and performance. 
> 
> <br>

<br>

> # **TODO**
> 
> <br>

<br>

# **[BACK TO TOP](#top)**
