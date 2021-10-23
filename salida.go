/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t15=P+1;
	t16=stack[int(t15)];
	L5:
	t17=heap[int(t16)];
	if t17 == -1 {goto L4;}
	fmt.Printf("%c", int(t17));
	t16=t16+1;
	goto L5;
	L4:
	return;
}


func main(){
	t0=1;
	L0:
	if t0 >= 5 {goto L1;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t1=P+0;
	stack[int(t1)]=t0;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t3=P+0;
	t2=stack[int(t3)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t4=P+1;
	stack[int(t4)]=t2;
	
	t5=H;
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
	/* compilacion de accesso */
	t7=P+1;
	t6=stack[int(t7)];
	/* fin de la compilacion de acceso */
	
	t11=H;
	t9=t5;
	t10=t6;
	t12=t5;
	L2:
	t9=t12;
	t13=t13+1;
	t14=heap[int(t9)];
	L3:
	heap[int(H)]=t14;
	H=H+1;
	t9=t9+1;
	t14=heap[int(t9)];
	if t14 != -1 {goto L3;}
	t9=t9-1;
	if t6 > t13 {goto L2;}
	heap[int(H)]=-1;
	H=H+1;
	t18=P+2;
	t18=t18+1;
	stack[int(t18)]=t11;
	P=P+2;
	print_string();
	t19=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t0=t0+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t20=P+0;
	stack[int(t20)]=t0;
	
	goto L0;
	L1:
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	stack[int(0)]=4;
	
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	stack[int(0)]=12;
	
	t21=H;
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
	/* compilacion de accesso */
	t22=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t26=H;
	t24=t21;
	t25=t22;
	t27=t21;
	L6:
	t24=t27;
	t28=t28+1;
	t29=heap[int(t24)];
	L7:
	heap[int(H)]=t29;
	H=H+1;
	t24=t24+1;
	t29=heap[int(t24)];
	if t29 != -1 {goto L7;}
	t24=t24-1;
	if t22 > t28 {goto L6;}
	heap[int(H)]=-1;
	H=H+1;
	t30=P+1;
	t30=t30+1;
	stack[int(t30)]=t26;
	P=P+1;
	print_string();
	t31=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}