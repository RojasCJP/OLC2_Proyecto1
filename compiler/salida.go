/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;



func main(){
	/* compilacion de valor de variable */
	t0=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=104;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t0=t0+0.12837;
	t1=t0;
	t2=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t2=t2+0.12837;
	t3=t2;
	t4=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t4=t4+0.12837;
	t5=t4;
	t6=H;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t0;
	H=H+1;
	heap[int(H)]=t2;
	H=H+1;
	heap[int(H)]=t4;
	H=H+1;
	t6=t6+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t6;
	
	t7=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=104;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t7=t7+0.12837;
	t9=t7+1;
	t8=heap[int(t9)];
	L0:
	if t8 == -1 {goto L1;}
	t8=heap[int(t9)];
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t10=P+1;
	stack[int(t10)]=t8;
	
	/* compilacion de accesso */
	t12=P+1;
	t11=stack[int(t12)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%c", int(t11));
	fmt.Printf("%c", int(32));
	t9=t9+1;
	goto L0;
	L1:

}