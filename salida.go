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
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t4=H;
	t2=t0;
	t3=4;
	t5=t0;
	L0:
	t2=t5;
	t6=t6+1;
	t7=heap[int(t2)];
	L1:
	heap[int(H)]=t7;
	H=H+1;
	t2=t2+1;
	t7=heap[int(t2)];
	if t7 != -1 {goto L1;}
	t2=t2-1;
	if 4 > t6 {goto L0;}
	heap[int(H)]=-1;
	H=H+1;
	t11=P+0;
	t11=t11+1;
	stack[int(t11)]=t4;
	P=P+0;
	print_string();
	t12=stack[int(P)];
	P=P-0;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}