package main

import (
	"fmt"
	external_dependency "github.com/marinewater/package-import-test-dep"
	subpackage_a "github.com/marinewater/package-import-test/subpackage-a"
	subpackage_b "github.com/marinewater/package-import-test/subpackage-a/subpackage-b"
)

func main() {
	fmt.Printf("Interval A: %d\n", subpackage_a.GetInterval())
	fmt.Printf("Interval B: %d\n", subpackage_b.GetInterval())
	fmt.Printf("Interval External: %d\n\n", external_dependency.GetInterval())

	subpackage_a.SetInterval(15)
	fmt.Println("Interval A set to 15")

	fmt.Printf("Interval A: %d\n", subpackage_a.GetInterval())
	fmt.Printf("Interval B: %d\n", subpackage_b.GetInterval())
	fmt.Printf("Interval External: %d\n\n", external_dependency.GetInterval())

	subpackage_b.SetInterval(20)
	fmt.Println("Interval B set to 20")

	fmt.Printf("Interval A: %d\n", subpackage_a.GetInterval())
	fmt.Printf("Interval B: %d\n", subpackage_b.GetInterval())
	fmt.Printf("Interval External: %d\n\n", external_dependency.GetInterval())

	external_dependency.SetInterval(25)
	fmt.Println("Interval External set to 25")

	fmt.Printf("Interval A: %d\n", subpackage_a.GetInterval())
	fmt.Printf("Interval B: %d\n", subpackage_b.GetInterval())
	fmt.Printf("Interval External: %d\n", external_dependency.GetInterval())
}
