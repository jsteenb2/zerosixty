# Golang or Go Home

a programmer's introduction to Go (*not an intro to programming in general*).  Pull down the repo, ```cd``` into it and follow along.

## Why Go?

Is it the lightning fast compile and execution times that make Go so appealing? In short, yes, that is a definite plus for working in Go. However, that's not all there is to love about Go. To get the speed you don't have to load up on complexity in Go. 

30+ years of lessons learned have been applied to Go. The creators, the likes of which include Rob Pike & Ken Thompson (former bell labs, creators of large parts of Unix and the Plan-9 OS), Robert Griesemer(create of V8), and Russ Cox (another Plan-9 contributor and systems programming heavyweight), only add something to the language if it is agreed upon by the group. So what did the creators include in Go...? 

* native support for testing, crypto, networking.... you name it
	* built right into the std lib
* memory integrity/management
	* zero values
	* pointers but no pointer arithmetic
	* Garbage collected
* native support for multi-core computers
	* concurrency modeled after the again famous [CSP](https://en.wikipedia.org/wiki/Communicating_sequential_processes)(communicating sequential processes)
* Wealth of tools to support you in development
	* gofmt
	* benchmarking
	* memory profiling
	* code coverage analyzer

You get all that and more. The best thing about Go, above all else, is how simple it is to work in big teams. 

## Warning

Go is NOT object/class oriented

LEAVE YOUR OBJECTED ORIENTED HAT AT THE DOOR

## Go Type System

### Primitives

```go
var boolean bool

var aString string

var (
	anInt   int
	anInt8  int8
	anInt32 int32
	anInt64 int64
	uInt    uint64
)
	
var (
	aFloat32 float32
	aFloat64 float64
)

// more still
```

All your expected types and then some. Note for ints, you want to default to using ```int``` for integers unless you are absolutely sure you need something else. ```int``` changes depending on your systems architecture (address size).

Each time a var is created without declaring it to a specific value, it will assume the zero type. See the ```var_declarations_test.go``` file for an example of what that looks like. Memory is allocated for you, memory integrity comes first.

run example: ```$ go test -v var_declarations_test.go -run TestVariableDeclarationsToZeroValues```

### Var Declarations

There are a number of ways to declare variables and we'll look at a few of the most commonly used ones.

```go
// setting to zero values
var boolean bool
var aString string
var anInt   int

// long way
var maxInt uint64 = 1<<64 - 1
var falsey bool = false

// shorthand, sets type equal to the right hand assignment
anInt := 2017  // int type default
aFloat := 32.0 // float64 type default
boolean := true
```

The shorthand is very commonly used when you are setting the value of something at declaration. If you want the zero value, you should use the first set, guarantees zero value.

run example: ```$ go test -v var_declarations_test.go -run 'DelcarationAssignment|Shorthand'```

### Type Conversion

There is no type-casting in Go, but there is conversions. They look a lot like constructor functions...

```go
var anInt8 int8 = 2
var aFloat32 float32 = 32.00

var anInt16 int16 = 2
var aFloat64 float64 = 32.00
	
int16(anInt8) == anInt16 // output: true
float64(aFloat32) == aFloat64 // output: true

// anInt8 == anInt16 { // doesn't compile, cannot do operations on different types
```

Different types are not interchangeable without conversion. Compiler will crap the bed and tell you to fix it.... again an emphasis on data integrity.

run example: ```$ go test -v var_declarations_test.go -run Conversion```

## Functions

A function in Go, ```func```, are easy to work with. Use the keyword ```func``` to create a func, specify the input params and output params. There are a number of ways to set types on the input params.

```go
func odder(anInt int) int {
	return anInt % 1
}

func adder(anInt, anotherInt, thirdInt int) int {
	// input params that share types can have the type associated to the last param
	// all params without type in front of that param get same type associated
	return anInt + anotherInt + thirdInt
}
```

run example: ```$ go test -v func_test.go -run BasicFunc```

When having multiple input params that the same type, you can condense the type from the leading params and just put it on the last input param. Input params can have a primitive type, a struct type, or an interface type as we'll see in a bit.


Input param types can be mixed and matched however you see fit. Functions in Go can be variadic as well, meaning you can specify and undetermined number of inputs (of same type) with the ```...``` syntax. Your return params can be named as well. Funcs can be set to variables as well. See the example below.

```go
func variadicFunc(msg string, ints ...int) (output string) {
	// named output type used here to showcase you can indeed name an output
	var sum int
	for _, v := range ints {
		sum += v
	}
	output = fmt.Sprintf("%s: %d", msg, sum)
	return
}

fn := variadicFunc

fmt.Println(fn("what's my total", 1, 2, 3, 4)) // output: what's my total: 10
``` 

run example: ```$ go test -v func_test.go -run Variadic```

Anonymous funcs can also be created, and closure is fully available to you then as well as any other time you create a func.

```go
func TestFuncClosure(t *testing.T) {
	msg := "I'm getting closured"

	fn := func(anInt int) string {
		return fmt.Sprintf("%s\n\tinput: %d", msg, anInt)
	}

	t.Log(fn(3333))
}
```

run example: ```$ go test -v func_test.go -run Closure```

### Data Structures

#### Arrays

Nothing to far off the beaten path. The most efficient data structure from the point of view of the hardware. Go runs on the metal, take a mental note of this.

```go
var zeroArray [4]string // zero value is array with all elements at their zero value

filledArray := [4]string{"Asynchrony", "World", "Wide", "Technology"}

unFilledArray := [4]string{"Missed", "By", "1"} // all 4 entries initialized, 3 to indicated values, 1 to zero value of "" for string

unFilledArray[3] = "covfefe"

for ele := range filledArray {
	fmt.Println(ele)
}
```

Arrays are straight forward, more of the same you've probably seen before. When you initialize an array, it will be put into memory as indicated. Any empty values will assume the zero value of the type the element is. The ```range``` operator iterates over elements of the array, similar to a foreach loop.

run example: ```$ go test -v data_structures_test.go -run Array```

#### Slices

Slices are where Go makes it's money. They are basically dynamic arrays that can manipulate capacity at runtime. A slice is a 3 word data structure that looks like:

1. pointer to backing array
2. length
3. capacity

```go
strSlice := []string{"like", "arrays", "only", "better"}

for idx, value := range strSlice {
	fmt.Printf("id: %d, value: %v\n", idx, value)
}

// can extend with append
strSlice = append(strSlice, "and", "extendable") // append(dst, ...src)

fmt.Println(strSlice)
```

Adding to a slice is easy enough with append. First arg being the destination, the rest of the arguments being the source. 

run example: ```$ go test -v data_structures_test.go -run SliceTypeWithRange```

There are several ways to create a slice and reassign elements of that slice.

```go
slice := make([]string, 0, 0) // initializes slice with 0 length and 0 capacity

slice = append(slice, "with", "great", "power", "comes", "great", "responsibility")
fmt.Println(strings.Join(slice, " "))

newSlice := slice[:] // returns copy of original slice
newSlice[2] = "food"
fmt.Println(strings.Join(newSlice, " "))

newerSlice := append(slice[:len(newSlice)-1], "indigestion")
fmt.Println(strings.Join(newerSlice, " "))
```

Pretty straightforward for reassignments. Note, you can call for a range of indices by using the ```newSlice[start:end]``` syntax. The end is exclusive. Anytime you use the ```:``` it will created a copy in memory of the backing array elements that are included in your range. Don't jump into arrays right away, the performance gain isn't world changing but it is horribly restrictive.

run example: ```$ go test -v data_structures_test.go -run SliceReassign```

#### Maps

Maps in Go are akin to the object in Javascript, maps in Java/C++, the dictionary in Python, yada yada. A key can be any type that can be equated(```==```), which includes primitives, structs, you name it. Anything except an array, slice, or another map should work. Note: the hashing function is done for you, no need to roll your own. It's robust as all get out, thank you Google!

Assignments are like you saw with arrays and slices.

```go
aMap := map[string][]string{
	"marvel": {"spiderman", "hulk", "ironman"},
	"dc":     {"batman", "superman", "wonderwoman"},
}

for key, value := range aMap {
	fmt.Printf("%s", key)
	for _, hero := range value {
		fmt.Printf("\t%s", hero)
	}
}
```

You get access to range operator on maps as well. This example illustrate some Go syntax sugar. Since the type of the value was called out in the declaration of the map, you do not need to initialize the embedded slices with the full ```[]string{}``` syntax. Go's compiler fills in the gaps for you.

You can access k/v pairs using the same syntax you are use to in other languages. You also can delete k/v pairs from a map using the ```delete``` keyword. Each access also has a lookup flag that indicates if the k/v pair exists.

```go
aMap := map[string][]string{
	"marvel": {"spiderman", "hulk", "ironman"},
	"dc":     {"batman", "superman", "wonderwoman"},
}

delete(aMap, "dc") // drops k/v pair

_, ok := aMap["dc"]

fmt.Printf("dc found: %t", ok) // output: dc found: false
```

run example: ```$ go test -v data_structures_test.go -run MapType```

#### Zero Value of Reference Types

Reference types, ```map``` and ```slice```, have a reference value of ```nil```. 

```go
var zeroSlice []bool
var zeroMap map[string]bool

if zeroSlice == nil {
	fmt.Println("zeroSlice == nil")
}

if zeroMap == nil {
	fmt.Println("zeroMap == nil")
}
```

run example: ```$ go test -v data_structures_test.go -run ZeroTypes```

### Named Types & Methods

A small concerted example is something like a type phoneNumber that is a string.

```go
package company

type PhoneNumber string

func (p PhoneNumber) HumanReadable() string {
	return fmt.Sprintf("(%s) %s-%s", p[0:3], p[3:6], p[6:])
}

func (p PhoneNumber) AreaCode() int {
	ac, _ := strconv.Atoi(string(p[0:3]))
	return ac
}
```

run example: ```$ go test -v ./company/structs_named_types_test.go -run TestNamedTypeMethods```

Named types are your friend in go. Take a simple int or string, throw a name on it, and a whole new world opens up to you. Check out duration from the std lib's [time package](https://golang.org/pkg/time/#Duration) below. 

```go
package time

// other time pkg stuff...

// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
type Duration int64

const (
  	Nanosecond  Duration = 1
  	Microsecond          = 1000 * Nanosecond
  	Millisecond          = 1000 * Microsecond
  	Second               = 1000 * Millisecond
  	Minute               = 60 * Second
  	Hour                 = 60 * Minute
)

func (d Duration) String() string {
// Stringify logics here
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64 { return int64(d) }

// Seconds returns the duration as a floating point number of seconds.
func (d Duration) Seconds() float64 {
	sec := d / Second
	nsec := d % Second
	return float64(sec) + float64(nsec)/1e9
}

// Minutes returns the duration as a floating point number of minutes.
func (d Duration) Minutes() float64 {
	min := d / Minute
	nsec := d % Minute
	return float64(min) + float64(nsec)/(60*1e9)
}

// Hours returns the duration as a floating point number of hours.
func (d Duration) Hours() float64 {
	hour := d / Hour
	nsec := d % Hour
	return float64(hour) + float64(nsec)/(60*60*1e9)
}
  
// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Duration) Truncate(m Duration) Duration {
	if m <= 0 {
		return d
	}
	return d - d%m
}

// ... more time package stuff
```

The ```Duration``` type has a bunch of useful methods baked into it. With this simple named ```int64``` you can add methods that turn the ```int64``` into something with complex logic baked right in. Methods can attach to any named type or struct. 

### Structs

Structs are whatever you make them be. You can tack on as many fields as you see fit. On top of that you can have methods on your struct types. 

```go
package company

type Address struct {
	ZipCode int
	Street  string
	Number  int
	City    string
	State   string
}

func (a Address) MailFormat() string {
	return fmt.Sprintf("\n%d %s\n%s, %s %d", a.Number, a.Street, a.City, a.State, a.ZipCode)
}

func (Address) WhoAmI() string {
	return "address"
}
```

The Address struct has a number of different fields on it. The methods can access any field within a struct, both public and private. Fields can be accessed by ```.``` notation as seen in the above example.

run example: ```$ go test -v ./company/structs_named_types_test.go -run StructTypeMethods```

#### Embedding Structs

Structs can be embedded into other structs. When this happens the embedded struct fields and methods can be accessed directly. This is called field/method promotion.

```go
package company 

type Company struct {
	Name string
	Contact
}

type Contact struct {
	PointOfContact []Person
	PhoneNumber    PhoneNumber
	Address
}

type Person string

func (Company) WhoAmI() string {
	return "company"
}

func (Contact) WhoAmI() string {
	return "contact"
}

func (Person) WhoAmI() string {
	return "person"
}
```

Now when we call the structs after instantiating, we'll be able to see the field promotion. 

```go
address := company.Address{
	ZipCode: 63102,
	Street:  "Spruce Street",
	Number:  900,
}

contact := company.Contact{
	PointOfContact: []company.Person{"Pops", "RedGoatTea"},
	PhoneNumber:    "3146782200",
	Address:        address,
}

async := company.Company{"Asynchrony", contact}

fmt.Println(async.Street) // nested 2 levels deep, output: Spruce Street
fmt.Println(async.PhoneNumber) // nested 1 level deep, output: 3146782200
```

run example: ```$ go test -v ./company/structs_named_types_test.go -run StructFieldProm```

One thing to note, if you have a name collision amongst fields or methods of nested structs, then you'll have to call out the specific nested struct you are meaning to access.

```go
address := company.Address{63102, "Spruce Street", 900, "St. Louis", "MO"}

contact := company.Contact{
	PointOfContact: []company.Person{"RedGoatTea", "Trigger"},
	PhoneNumber:    "3146782200",
	Address:        address,
}

async := company.Company{"Asynchrony", contact}

fmt.Println(async.WhoAmI())   // output: company
fmt.Println(contact.WhoAmI()) // output: contact
fmt.Println(address.WhoAmI()) // output: address
```

run example: ```$ go test -v ./company/structs_named_types_test.go -run Collision```

### Interface

Interfaces are one of the places where Go should be most admired. Go is a structurally typed language and this makes all interfaces satisfied implicitly. This opens the door to polymorphic behavior in our code. You will get something akin to duck typing in Go, a strongly typed programming language. Here's a small example.

```go
type WhoAmIer interface {
	WhoAmI() string
}
```

The interface says: "to satsify me, you need to have a WhoAmI method on the type with the exact signature I say you should have". Any struct or named type with the exact method name and signature will satisfy the interface.

```go
type WhoAmIer interface {
	WhoAmI() string
}

whoamiSlice := []WhoAmIer{
	c.Contact{},
	c.Address{},
	c.PhoneNumber(5558675309),
	c.Company{},
	c.Person("ex"),
}

for _, v := range whoamiSlice {
	fmt.Println(v.WhoAmI()) // hello duck typing?
}

// the line below does not work
// the compiler knows nothing of the implementation of Address struct
// based on the WhoAmIer interface
// fmt.Println(whoamiSlice[1].MailFormat())
```

run example: ```$ go test -v ./company/interface_test.go -run Implicitly
```

All the types of the slice implement the interface, so they are all valid elements of the ```whoamiSlice```. If you attempt to access a method that isn't called out in the interface, the compiler will blow up on you.

A ```func``` can specify an interface as the type for an input parameter. When this is done, you will be able to access the specific method(s) spelled out in your interface type, but none others. You will obtain polymorphic behavior for your ```func``` params.

```go
func isCompany(who WhoAmIer) bool {
	return who.WhoAmI() == "company"
}

whoamiSlice := []WhoAmIer{
	c.Contact{},
	c.Address{},
	c.PhoneNumber(5558675309),
	c.Company{},
	c.Person("ex"),
}

for idx, v := range whoamiSlice {
	if isCompany(v) {
		fmt.Printf("company found at idx: %d", idx)
	}
}
```

run example: ```$ go test -v ./company/interface_test.go -run AsInput```

### Named Func Types

#### Middleware

### Packaging

### Concurrency in Go

#### Go routines

#### Go scheduler

### Atomic Funcs

### Mutexes

### Chans
