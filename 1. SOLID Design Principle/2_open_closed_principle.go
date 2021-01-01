package main

import "fmt"

// OCP
// open for extension, closed for modification

// Color ..
type Color int

const (
	red Color = iota
	green
	blue
)

// Size ...
type Size int

const (
	small Size = iota
	medium
	large
)

// Product ...
type Product struct {
	name  string
	color Color
	size  Size
}

// Filter ...
type Filter struct {
}

// FilterByColor ...
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// FilterBySize ...
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// FilterBySizeAndColor ...
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Specification ...
type Specification interface {
	IsSatisfied(p *Product) bool
}

// ColorSpecification ...
type ColorSpecification struct {
	color Color
}

// IsSatisfied ..
func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return spec.color == p.color
}

// SizeSpecification ...
type SizeSpecification struct {
	size Size
}

// IsSatisfied ..
func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return spec.size == p.size
}

// AndSpecification ...
type AndSpecification struct {
	first, second Specification
}

// IsSatisfied ..
func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

// BetterFilter ...
type BetterFilter struct {
}

// Filter ...
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"Apple", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// ^^^ BEFORE

	// vvv AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
