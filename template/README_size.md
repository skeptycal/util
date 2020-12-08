# size - memory consumption at runtime

>Measures the size of an object in your Go program at runtime based on `binary.Size()` from the Go standard library.

Adapted from [DmitriyVTitov](https://github.com/DmitriyVTitov)'s [size](https://github.com/DmitriyVTitov/size) ([MIT License](LICENSE))

Features:
- supports non-fixed size variables and struct fields: `struct`, `int`, `slice`, `string`, `map`;
- supports complex types including structs with non-fixed size fields;
- supports all basic types (numbers, `bool`);
- supports `chan` and `interface`;
- supports pointers;
- implements infinite recursion detection (i.e. `pointer` inside `struct field` references to parent `struct`).

### Usage example

```go
package main

import (
	"fmt"

	"github.com/skeptycal/size"
)

func main() {
	a := struct {
		a int
		b string
		c bool
		d int32
		e []byte
		f [3]int64
	}{
		a: 10,                    // 8 bytes
		b: "Text",                // 4 bytes
		c: true,                  // 1 byte
		d: 25,                    // 4 bytes
		e: []byte{'c', 'd', 'e'}, // 3 bytes
		f: [3]int64{1, 2, 3},     // 24 bytes
	}

	fmt.Println(size.Of(a))
}

// Output: 44
```
