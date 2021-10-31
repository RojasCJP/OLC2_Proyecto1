/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t11=P+1;
	t12=stack[int(t11)];
	L1:
	t13=heap[int(t12)];
	if t13 == -1 {goto L0;}
	fmt.Printf("%c", int(t13));
	t12=t12+1;
	goto L1;
	L0:
	return;
}
func print_array(){
	t14=P+1;
	t15=stack[int(t14)];
	t17=heap[int(t15)];
	t15=t15+1;
	fmt.Printf("%c", int(91));
	L3:
	t19=heap[int(t15)];
	t18=t19;
	if t16 >= t17 {goto L2;}
	if t2 == t18 {goto L5;}
	if t4 == t18 {goto L5;}
	if t6 == t18 {goto L5;}
	if t8 == t18 {goto L5;}
	fmt.Printf("%d", int(t19));
	fmt.Printf("%c", int(44));
	t15=t15+1;
	t16=t16+1;
	goto L3;
	L5:
	t20=P;
	t21=t16;
	t22=t17;
	t23=t19;
	t24=t14;
	t25=t15;
	t16=0;
	stack[int(t14)]=t19;
	print_array();
	t16=t21+1;
	t17=t22;
	t19=t23;
	t14=t24;
	t15=t25+1;
	goto L3;
	L2:
	fmt.Printf("%c", int(93));
	t16=0;
	return;
}


func main(){
	/* compilacion de valor de variable */
	t0=H;
	heap[int(H)]=4;
	H=H+1;
	heap[int(H)]=1;
	H=H+1;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=4;
	H=H+1;
	t0=t0+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t0;
	
	/* compilacion de valor de variable */
	t1=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=1;
	H=H+1;
	heap[int(H)]=2;
	H=H+1;
	t1=t1+0.12837;
	t2=t1;
	t3=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=4;
	H=H+1;
	t3=t3+0.12837;
	t4=t3;
	t5=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=5;
	H=H+1;
	heap[int(H)]=6;
	H=H+1;
	t5=t5+0.12837;
	t6=t5;
	t7=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=7;
	H=H+1;
	heap[int(H)]=8;
	H=H+1;
	t7=t7+0.12837;
	t8=t7;
	t9=H;
	heap[int(H)]=4;
	H=H+1;
	heap[int(H)]=t1;
	H=H+1;
	heap[int(H)]=t3;
	H=H+1;
	heap[int(H)]=t5;
	H=H+1;
	heap[int(H)]=t7;
	H=H+1;
	t9=t9+0.12837;
	/* fin de compilacion de variable */
	stack[int(1)]=t9;
	
	/* compilacion de accesso */
	t10=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t26=P+2;
	t26=t26+1;
	stack[int(t26)]=t10;
	P=P+2;
	print_array();
	t27=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de accesso */
	t28=stack[int(1)];
	/* fin de la compilacion de acceso */
	
	t29=P+2;
	t29=t29+1;
	stack[int(t29)]=t28;
	P=P+2;
	print_array();
	t30=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t31=stack[int(0)];
	t32=4+t31;
	t31=heap[int(t32)];
	fmt.Printf("%f", t31);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t33=stack[int(1)];
	t34=4+t33;
	t33=heap[int(t34)];
	fmt.Printf("%f", t33);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}