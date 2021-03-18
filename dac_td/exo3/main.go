package main

import "fmt"

/*
 [17,2,55,41,36,77] a = 55 , b = 80
 [2,17,36,41,55,77]
  m = len(t) / 2 = 3 t[3] = 41
  si t[m] > b : return fn(a,b , t[:m]
  sinon si t[m] < a : return fn(a,b,t[m+1:}])
  sinon return t[m]
  fn(a,b int , t [] int ) int
*/
func main() {
	t := []int{2, 17, 36, 41, 55, 77}
	a := 40
	b := 80
	fmt.Println(dac(a, b, t))
}
func dac(a, b int, t []int) int {
	if len(t) == 0 {
		return -1
	}
	m := len(t) / 2
	if t[m] > b {
		return dac(a, b, t[:m])
	} else if t[m] < a {
		return dac(a, b, t[m+1:])
	} else {
		return t[m]
	}

}
