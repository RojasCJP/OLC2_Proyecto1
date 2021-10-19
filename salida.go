/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2 float64
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;


/*-----FUNCS-----*/
func prueba(){
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t1=P+1;
	t0=stack[int(t1)];
	/* fin de la compilacion de acceso */
	
	if t0 == 3 {goto L1;}
	goto L2;
	/* fin de la expression relacional */
	
	L1:
	stack[int(P)]=100;
	goto L0;
	L2:
	stack[int(P)]=50;
	goto L0;
	L0:
	return;
}

func main(){
	t2=P+1;
	stack[int(t2)]=3;
	P=P+0;
	prueba();
	t2=stack[int(P)];
	P=P-0;
	fmt.Printf("%d", int(t2));

}