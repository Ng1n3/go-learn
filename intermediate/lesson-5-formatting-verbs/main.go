package main

import "fmt"

func main() {

	// --- General FOrmatting Verbs
	// %v   prints the value in the default format
	// %#v  Prints the value in Go-syntax format
	// %T   Prints the type of the value
	// %%   Prints the % sign

	i := 15.5
	string := "Hello world!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	// --- Integer Formatting Verbs
	// %b   Base 2
	// %d   Base 10
	// %+d  Base 10 and always show sign
	// %o   Base 8
	// %O   Base 8, with leading 0o
	// %x   Base 16, lowercase
	// %X   Base 16, uppercase
	// %#x  Base 16, with leading 0x
	// %4d  Pad with spaces (width 4, right justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with zeroes (width 4)

	int := 18

	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%4d", int)
	fmt.Printf("%-4d", int)
	fmt.Printf("%04d", int)

	// --- String Formatting Verbs
	// %s   Prints the value as plain string
	// %q   Prints the value as a double-quoted string
	// %8s  Prints the value as Plain string (width 8, right justified)
	// %-8s Prints the value as plain string (width 8, left justified)
	// %x   Prints the value as a hex dump of byte values
	// % x  Prints the values as hex dump with spaces

	txt := "World"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("%X\n", txt)

	// --- Boolean Formatting Verbs
	// %t   Value of the boolean operator in true or false format (same as using %v)

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", f)

	// --- Float Formatting Verbs
	// %e   Scientific notation with 'e' as exponent
	// %f   Decimal point, no exponent
	// %.2f Default width, precision 2
	// %6.2f Width 6, precision 2
	// %g   Exponent as needed, only necessary digits.

	flt := 9.18

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f", flt)
	fmt.Printf("%6.2f", flt)
	fmt.Printf("%g", flt)
}
