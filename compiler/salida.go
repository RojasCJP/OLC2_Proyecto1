/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_array(){
	t2=P+1;
	t3=stack[int(t2)];
	t5=heap[int(t3)];
	t3=t3+1;
	fmt.Printf("%c", int(91));
	L1:
	t7=heap[int(t3)];
	t6=t7;
	if t4 >= t5 {goto L0;}
	fmt.Printf("%d", int(t7));
	fmt.Printf("%c", int(44));
	t3=t3+1;
	t4=t4+1;
	goto L1;
	L0:
	fmt.Printf("%c", int(93));
	return;
}


func main(){
	/* compilacion de valor de variable */
	t0=H;
	heap[int(H)]=5;
	H=H+1;
	heap[int(H)]=1;
	H=H+1;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=4;
	H=H+1;
	heap[int(H)]=5;
	H=H+1;
	t0=t0+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t0;
	
	/* compilacion de accesso */
	t1=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t8=P+1;
	t8=t8+1;
	stack[int(t8)]=t1;
	P=P+1;
	print_array();
	t9=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}