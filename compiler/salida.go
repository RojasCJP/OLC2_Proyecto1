/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t9=P+1;
	t10=stack[int(t9)];
	L1:
	t11=heap[int(t10)];
	if t11 == -1 {goto L0;}
	fmt.Printf("%c", int(t11));
	t10=t10+1;
	goto L1;
	L0:
	return;
}


func main(){
	/* compilacion de valor de variable */
	t0=H;
	t1=t0;
	H=H+2;
	t2=H;
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
	heap[int(H)]=32;
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
	heap[int(H)]=-1;
	H=H+1;
	t2=t2+0.12837;
	heap[int(t1)]=t2;
	t1=t1+1;
	heap[int(t1)]=412;
	t1=t1+1;
	/* fin de compilacion de variable */
	stack[int(0)]=t0;
	
	t3=stack[int(0)];
	t4=t3+1;
	t5=heap[int(t4)];
	fmt.Printf("%d", int(t5));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t6=stack[int(0)];
	t7=t6+0;
	t8=heap[int(t7)];
	t12=P+1;
	t12=t12+1;
	stack[int(t12)]=t8;
	P=P+1;
	print_string();
	t13=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}