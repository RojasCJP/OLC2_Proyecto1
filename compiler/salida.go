/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;



func main(){
	/* compilacion de valor de variable */
	t0=H;
	heap[int(H)]=6;
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
	heap[int(H)]=6;
	H=H+1;
	t0=t0+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t0;
	
	/* compilacion de valor de variable */
	t1=H;
	heap[int(H)]=6;
	H=H+1;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=0;
	H=H+1;
	t1=t1+0.12837;
	/* fin de compilacion de variable */
	stack[int(1)]=t1;
	
	/* compilacion de acceso arreglos */
	t2=stack[int(1)];
	t3=2+t2;
	t2=heap[int(t3)];
	fmt.Printf("%f", t2);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t4=1;
	L0:
	if t4 >= 7 {goto L1;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t5=P+2;
	stack[int(t5)]=t4;
	
	/* compilacion de acceso arreglos */
	t6=stack[int(1)];
	/* compilacion de accesso */
	t8=P+2;
	t7=stack[int(t8)];
	/* fin de la compilacion de acceso */
	
	t9=t7+t6;
	t6=heap[int(t9)];
	fmt.Printf("%f", t6);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t10=stack[int(0)];
	/* compilacion de accesso */
	t12=P+2;
	t11=stack[int(t12)];
	/* fin de la compilacion de acceso */
	
	t13=t11+t10;
	t10=heap[int(t13)];
	fmt.Printf("%f", t10);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t4=t4+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t14=P+2;
	stack[int(t14)]=t4;
	
	goto L0;
	L1:

}