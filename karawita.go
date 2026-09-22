package main

import"fmt"
	
	




func main(){


	var op int
	var num1 int
	var num2 int
	

	fmt.Println("CALCULA-DORA")
	fmt.Println("-------------")
	fmt.Println("1.SUMA")
	fmt.Println("2.RESTA")
	fmt.Println("3.MULTI")
	fmt.Println("4.DIVI")
	fmt.Println("INGRESA LA OPERACION CHAMO:")
	fmt.Scan(&op)
	fmt.Println("INGRESA EL PRIMER NUMERO TIO:")
	fmt.Scan(&num1)
	fmt.Println("INGRESA EL SEGUNDO NUMERO CHAVAL:")
	fmt.Scan(&num2)




	if op == 1{
		resul := sumar(num1,num2)
		fmt.Println(resul)
	}else if op == 2{
		resul := restar(num1,num2)
		fmt.Println(resul)
	}else if op == 3{
		resul := multi(num1,num2)
		fmt.Println(resul)
	}else if op == 4{
		resul,war := divi(num1,num2)


		if war != ""{
			fmt.Println("JAJAJAJA")
		}else {
			fmt.Println(resul)
		}

		
	}
	
	
	
	

	

}



func sumar(a int, b int) int{

	re := a + b
	return re

}

func restar(a int, b int) int{

	re := a - b
	return re

}

func multi(a int, b int) int{

	re := a * b
	return re

}

func divi(a int, b int) (int,string){
	if b == 0{
		return 0,"MALDITO MAMAWEBO :)"
	}

	re := a / b
	return re,""

}