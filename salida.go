/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12 float64
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;


/*-----FUNCS-----*/
func factorial(){
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t1=P+1;
	t0=stack[int(t1)];
	/* fin de la compilacion de acceso */
	
	if t0 != 0 {goto L1;}
	goto L2;
	/* fin de la expression relacional */
	
	L1:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t3=P+1;
	t2=stack[int(t3)];
	/* fin de la compilacion de acceso */
	
	t4=t2-1;
	t5=P+3;
	stack[int(t5)]=t4;
	P=P+2;
	factorial();
	t5=stack[int(P)];
	P=P-2;
	/* fin de compilacion de variable */
	t6=P+2;
	stack[int(t6)]=t5;
	
	/* compilacion de accesso */
	t8=P+1;
	t7=stack[int(t8)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t10=P+2;
	t9=stack[int(t10)];
	/* fin de la compilacion de acceso */
	
	t11=t7*t9;
	stack[int(P)]=t11;
	goto L0;
	L2:
	stack[int(P)]=1;
	goto L0;
	L0:
	return;
}

func main(){
	t12=P+1;
	stack[int(t12)]=5;
	P=P+0;
	factorial();
	t12=stack[int(P)];
	P=P-0;
	fmt.Printf("%d", int(t12));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}