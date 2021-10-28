/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t15=P+1;
	t16=stack[int(t15)];
	L1:
	t17=heap[int(t16)];
	if t17 == -1 {goto L0;}
	fmt.Printf("%c", int(t17));
	t16=t16+1;
	goto L1;
	L0:
	return;
}
func print_array(){
	t18=P+1;
	t19=stack[int(t18)];
	t21=heap[int(t19)];
	t19=t19+1;
	fmt.Printf("%c", int(91));
	L3:
	t23=heap[int(t19)];
	t22=t23;
	if t20 >= t21 {goto L2;}
	if t1 == t22 {goto L5;}
	if t3 == t22 {goto L5;}
	if t5 == t22 {goto L5;}
	if t7 == t22 {goto L5;}
	if t9 == t22 {goto L5;}
	fmt.Printf("%d", int(t23));
	fmt.Printf("%c", int(44));
	t19=t19+1;
	t20=t20+1;
	goto L3;
	L5:
	t24=P;
	t25=t20;
	t26=t21;
	t27=t23;
	t28=t18;
	t29=t19;
	t20=0;
	stack[int(t18)]=t23;
	print_array();
	t20=t25+1;
	t21=t26;
	t23=t27;
	t18=t28;
	t19=t29+1;
	goto L3;
	L2:
	fmt.Printf("%c", int(93));
	t20=0;
	return;
}


func main(){
	/* compilacion de valor de variable */
	t0=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=1;
	H=H+1;
	heap[int(H)]=2;
	H=H+1;
	t0=t0+0.12837;
	t1=t0;
	t2=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=4;
	H=H+1;
	t2=t2+0.12837;
	t3=t2;
	t4=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=5;
	H=H+1;
	heap[int(H)]=6;
	H=H+1;
	t4=t4+0.12837;
	t5=t4;
	t6=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=7;
	H=H+1;
	heap[int(H)]=8;
	H=H+1;
	t6=t6+0.12837;
	t7=t6;
	t8=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=9;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	t8=t8+0.12837;
	t9=t8;
	t10=H;
	heap[int(H)]=5;
	H=H+1;
	heap[int(H)]=t0;
	H=H+1;
	heap[int(H)]=t2;
	H=H+1;
	heap[int(H)]=t4;
	H=H+1;
	heap[int(H)]=t6;
	H=H+1;
	heap[int(H)]=t8;
	H=H+1;
	t10=t10+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t10;
	
	/* compilacion de valor de variable */
	t11=H;
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
	t11=t11+0.12837;
	/* fin de compilacion de variable */
	stack[int(1)]=t11;
	
	/* compilacion de acceso arreglos */
	t12=stack[int(0)];
	t13=2+t12;
	t12=heap[int(t13)];
	t14=2+t12;
	t12=heap[int(t14)];
	t30=P+2;
	t30=t30+1;
	stack[int(t30)]=t12;
	P=P+2;
	print_array();
	t31=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t32=1;
	L6:
	if t32 >= 5 {goto L7;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t33=P+2;
	stack[int(t33)]=t32;
	
	/* compilacion de acceso arreglos */
	t34=stack[int(1)];
	/* compilacion de accesso */
	t36=P+2;
	t35=stack[int(t36)];
	/* fin de la compilacion de acceso */
	
	t37=t35+t34;
	t34=heap[int(t37)];
	fmt.Printf("%f", t34);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t32=t32+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t38=P+2;
	stack[int(t38)]=t32;
	
	goto L6;
	L7:

}