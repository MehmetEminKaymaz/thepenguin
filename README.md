# thepenguin
Helpers for making the use of reflection easier

![image](https://user-images.githubusercontent.com/50118591/127775539-47b11c0c-7dac-40a7-b2d6-9dea47ba86e6.png)

This project is dedicated to the penguin walking towards the mountains...

For more information about the project name (penguin) : https://youtu.be/zWH_9VRWn8Y

### Installation
```
go get github.com/MehmetEminKaymaz/thepenguin
```

## Examples

#### Get property names of struct

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.PropertyNames())
  //output: [Name Surname Age]


```

#### Get property names of struct filtered by kind

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.PropertyNamesByKind(reflect.Int))
  //output: [Age]


```

#### Get property values

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.PropertyValues())
  //output: [Harry Brown 80]
```
#### Get property names and values

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.PropertyNamesWithValues())
  //output: map[Age:80 Name:Harry Surname:Brown]
```
#### Get property type names

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.PropertyTypeNames())
  //output: [string string int]
```
#### Set property value by index

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.SetPropertyValueByIndex(0,"X"))
  //output: {X Brown 80}
```
#### Set property value by Name

```Go
type Test struct {
	Name string
	Surname string
	Age int
}

refl:= thepenguin.Analyze(Test{
		Name:    "Harry",
		Surname: "Brown",
		Age:     80,
	})
	fmt.Print(refl.SetPropertyValueByName("Name","Y"))
  //output: {Y Brown 80}
```
### Deal with slices

#### Append

```Go
var example []int = []int{1,2,3,4}

	refl:=thepenguin.Analyze(example)
	fmt.Print(refl.Append(1))
  //output: [1 2 3 4 1]
```
#### DeleteAt

```Go
var example []int = []int{1,2,3,4}

	refl:=thepenguin.Analyze(example)
	fmt.Print(refl.DeleteAt(3))
  //output: [1 2 3]
```
#### InsertAt

```Go
var example []int = []int{1,2,3,4}

	refl:=thepenguin.Analyze(example)
	fmt.Print(refl.InsertAt(2,2))
  //output: [1 2 2 3 4]
```
#### Create
```Go
var example []int = []int{1,2,3,4}

	refl:=thepenguin.Analyze(example)
	fmt.Print(refl.Create()) //it creates empty slice
  //output: []
```
#### Copy
```Go
var example []int = []int{1,2,3,4}

	refl:=thepenguin.Analyze(example)
	fmt.Print(refl.Copy()) //create a copy of the slice
  //output: [1 2 3 4]
```



