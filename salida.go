/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t8=P+1;
	t9=stack[int(t8)];
	L3:
	t10=heap[int(t9)];
	if t10 == -1 {goto L2;}
	fmt.Printf("%c", int(t10));
	t9=t9+1;
	goto L3;
	L2:
	return;
}


func main(){
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	stack[int(0)]=3;
	
	t0=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	/* compilacion de accesso */
	t1=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t5=H;
	L0:
	t6=t6+1;
	t7=heap[int(t3)];
	L1:
	heap[int(H)]=t7;
	H=H+1;
	t3=t3+1;
	t7=heap[int(t3)];
	if t7 != -1 {goto L1;}
	t3=t3-1;
	if t1 > t6 {goto L0;}
	heap[int(H)]=-1;
	H=H+1;
	t11=P+1;
	t11=t11+1;
	stack[int(t11)]=t5;
	P=P+1;
	print_string();
	t12=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}