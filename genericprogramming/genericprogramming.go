package genericprogramming
import "golang.org/x/exp/constraints"
//import "fmt"

func AnyOf[T comparable] (value []T, f func (T) bool) bool {
	for _, v := range value {
		if f(v) == true {
			return true
		}
	} 
	return false
}

func FindIf[T comparable] (value []T , f func (T) bool) int {
	for i, v := range value {
		if (f(v) == true) {
			return i
		}
	}
	return -1
}

func AdjacentFind[T comparable] (slice []T, f func (T,T) bool) int {
	var previousValue T
	for i, v := range slice {
		if (i == 0) {
			previousValue = v
			continue
		}
		if (f(v,previousValue) == true) {
				return i
		}
		previousValue = v
	}
	return -1
}

func Equal[T comparable] (a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i,_ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ReplaceIf[T comparable] (slice []T,value T, f func (T) bool) int {
	var replacements int
	for i, v := range slice {
		if (f(v) == true) {
			replacements = replacements + 1
			slice[i] = value
		}
	}
	return replacements
}

func RemoveIf[T comparable] (slice *[]T, f func (T) bool) int {
	var removals int
	for i, v := range *slice {
		if (f(v) == true) {
			removals = removals + 1
			(*slice)[i] = (*slice)[len(*slice) - removals]
		}
	}
	*slice = (*slice)[:len(*slice) - removals]
	return removals
}

func IsSorted[T constraints.Ordered] (slice []T) bool {
	var previousValue T
	for i, v := range slice {
		if (i == 0) {
			previousValue = v
			continue
		}
		if (v < previousValue) {
				return false
		}
		previousValue = v
	}
	return true
}

//Junta slices pre-ordenados en un nuevo slice ordenado
func Merge[T constraints.Ordered] (a, b []T) []T {
	var result []T
	var i,j int
	for ; i < len(a); i++{
		for ; j < len(b); j++ {
			if (a[i] < b[j]) {
				result = append(result, a[i])
				break
			} else {
				result = append(result, b[j])
			}
		}
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}
