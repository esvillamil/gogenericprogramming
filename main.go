// Ejercicio de Programacion Generica
package main

import (
"fmt"

"github.com/esvillamil/gogenericprogramming/genericprogramming"
)

func EsNumeroPar(i int) bool {
	return i % 2 == 0
}

func EsMultiplo(i,j int) bool {
	if (j == 0) {
		return false
	}
	return (i % j == 0)
}


func main() {
	var resultB bool
	var resultI int
	slice1 := []int{3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{-1, 3, 9}

	resultB = genericprogramming.AnyOf(slice1,EsNumeroPar)
	fmt.Println("AnyOf: ¿hay pares en slice= ",slice1, "? Result: ",resultB)
	resultB = genericprogramming.AnyOf(slice2,EsNumeroPar)
	fmt.Println("AnyOf: ¿hay pares en slice= ",slice2, "? Result: ",resultB)

	resultI = genericprogramming.FindIf(slice1,EsNumeroPar)
	fmt.Println("FindIf: Encuentra un par en slice= ",slice1, ". Result= ",resultI)
	resultI = genericprogramming.FindIf(slice2,EsNumeroPar)
	fmt.Println("FindIf: Encuentra un par en slice= ",slice2, ". Result= ",resultI)

	resultI = genericprogramming.AdjacentFind(slice1,EsMultiplo)
	fmt.Println("AdjacentFind: Encuentra multiplos adyacentes en slice= ",slice1, ". Result= ",resultI)
	resultI = genericprogramming.AdjacentFind(slice2,EsMultiplo)
	fmt.Println("AdjacentFind: Encuentra multiplos adyacentes en slice= ",slice2, ". Result= ",resultI)

	resultB = genericprogramming.Equal(slice1,slice2)
	fmt.Println("Equal: ¿Son iguales slice1= ",slice1," y slice2= ",slice2, "? Result= ",resultB)
	resultB = genericprogramming.Equal(slice2,slice2)
	fmt.Println("Equal: ¿Son iguales slice1= ",slice2," y slice2= ",slice2, "? Result= ",resultB)

	fmt.Print("ReplaceIf: Reemplaza los pares por -1 en slice = ",slice1)
	resultI = genericprogramming.ReplaceIf(slice1,-1,EsNumeroPar)
	fmt.Println(". Result= ", resultI, " slice= ",slice1)

	slice3 := []int{3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("RemoveIf: Borra los pares en slice = ",slice3)
	resultI = genericprogramming.RemoveIf(&slice3,EsNumeroPar)
	fmt.Println(". Result= ", resultI, " slice= ",slice3)

	resultB = genericprogramming.IsSorted(slice3)
	fmt.Println("IsSorted: ¿Está ordenada la slice = ",slice3, "? Result= ",resultB)
	resultB = genericprogramming.IsSorted(slice2)
	fmt.Println("IsSorted: ¿Está ordenada la slice = ",slice2, "? Result= ",resultB)

	slice4 := []int{ 5, 6, 7, 8, 9, 10}
	slice5 := genericprogramming.Merge(slice2,slice4)
	fmt.Println("Merge: Mergea slice1= ",slice2," slice2= ",slice4, ". Result= ",slice5,"")
}
