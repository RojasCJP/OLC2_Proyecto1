/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t37=P+1;
	t38=stack[int(t37)];
	L11:
	t39=heap[int(t38)];
	if t39 == -1 {goto L10;}
	fmt.Printf("%c", int(t39));
	t38=t38+1;
	goto L11;
	L10:
	return;
}


func main(){
	t0=0;
	L0:
	if t0 >= 9 {goto L1;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t1=P+0;
	stack[int(t1)]=t0;
	
	/* compilacion de valor de variable */
	t2=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	/* fin de compilacion de variable */
	t3=P+1;
	stack[int(t3)]=t2;
	
	/* compilacion de accesso */
	t5=P+0;
	t4=stack[int(t5)];
	/* fin de la compilacion de acceso */
	
	t6=10-t4;
	t7=0;
	L2:
	if t7 >= t6 {goto L3;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t8=P+2;
	stack[int(t8)]=t7;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t10=P+1;
	t9=stack[int(t10)];
	/* fin de la compilacion de acceso */
	
	t11=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t15=H;
	t13=t9;
	t14=t11;
	t16=heap[int(t13)];
	t17=heap[int(t14)];
	L4:
	heap[int(H)]=t16;
	H=H+1;
	t13=t13+1;
	t16=heap[int(t13)];
	if t16 != -1 {goto L4;}
	L5:
	heap[int(H)]=t17;
	H=H+1;
	t14=t14+1;
	t17=heap[int(t14)];
	if t17 != -1 {goto L5;}
	heap[int(H)]=-1;
	H=H+1;
	/* fin de compilacion de variable */
	t18=P+1;
	stack[int(t18)]=t15;
	
	t7=t7+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t19=P+2;
	stack[int(t19)]=t7;
	
	goto L2;
	L3:
	/* compilacion de accesso */
	t21=P+0;
	t20=stack[int(t21)];
	/* fin de la compilacion de acceso */
	
	t22=0;
	L6:
	if t22 >= t20 {goto L7;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t23=P+2;
	stack[int(t23)]=t22;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t25=P+1;
	t24=stack[int(t25)];
	/* fin de la compilacion de acceso */
	
	t26=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=42;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t30=H;
	t28=t24;
	t29=t26;
	t31=heap[int(t28)];
	t32=heap[int(t29)];
	L8:
	heap[int(H)]=t31;
	H=H+1;
	t28=t28+1;
	t31=heap[int(t28)];
	if t31 != -1 {goto L8;}
	L9:
	heap[int(H)]=t32;
	H=H+1;
	t29=t29+1;
	t32=heap[int(t29)];
	if t32 != -1 {goto L9;}
	heap[int(H)]=-1;
	H=H+1;
	/* fin de compilacion de variable */
	t33=P+1;
	stack[int(t33)]=t30;
	
	t22=t22+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t34=P+2;
	stack[int(t34)]=t22;
	
	goto L6;
	L7:
	/* compilacion de accesso */
	t36=P+1;
	t35=stack[int(t36)];
	/* fin de la compilacion de acceso */
	
	t40=P+2;
	t40=t40+1;
	stack[int(t40)]=t35;
	P=P+2;
	print_string();
	t41=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t0=t0+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t42=P+0;
	stack[int(t42)]=t0;
	
	goto L0;
	L1:
	t43=H;
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
	t44=P+0;
	t44=t44+1;
	stack[int(t44)]=t43;
	P=P+0;
	print_string();
	t45=stack[int(P)];
	P=P-0;
	fmt.Printf("%c", int(32));
	fmt.Printf("%d", int(3));
	fmt.Printf("%c", int(32));
	goto L12;
	/* goto para evitar errores */
	goto L13;
	L12:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L14;
	L13:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L14:
	fmt.Printf("%c", int(32));
	fmt.Printf("%f", 3.123);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}