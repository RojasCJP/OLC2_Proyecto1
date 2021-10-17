/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5 float64
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;



func main(){
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	stack[int(0)]=0;
	
	/* compilacion de accesso */
	t0=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%d", int(t0));
	fmt.Printf("%c", int(10));
	L0:
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t1=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	if t1 < 10 {goto L1;}
	goto L2;
	/* fin de la expression relacional */
	
	L1:
	/* compilacion de accesso */
	t2=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%d", int(t2));
	fmt.Printf("%c", int(10));
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t3=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t4=t3+1;
	/* fin de compilacion de variable */
	stack[int(0)]=t4;
	
	goto L0;
	L2:
	/* compilacion de accesso */
	t5=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%d", int(t5));
	fmt.Printf("%c", int(10));

}