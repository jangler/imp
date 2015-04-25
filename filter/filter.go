// Package filter contains image filter operations and metadata.
package filter

import (
	"image"
	"sort"
)

// A Func takes an image and the current arg stack, and returns a modified
// image and arg stack.
type Func func(*image.RGBA, []string) (*image.RGBA, []string)

// A Filter is a Func and associated metadata.
type Filter struct {
	Name, Help string
	Func       Func
}

// ByName implements sort.Interface for []*Filter based on name.
type ByName []*Filter

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// List is an array of all Filters.
var List = make([]*Filter, 0)

// Map allows Filter lookup by name.
var Map = make(map[string]*Filter)

func addFilter(f *Filter) {
	List = append(List, f)
	sort.Sort(ByName(List))
	Map[f.Name] = f
}
