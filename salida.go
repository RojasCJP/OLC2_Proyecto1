/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t24=P+1;
	t25=stack[int(t24)];
	L1:
	t26=heap[int(t25)];
	if t26 == -1 {goto L0;}
	fmt.Printf("%c", int(t26));
	t25=t25+1;
	goto L1;
	L0:
	return;
}
func print_array(){
	t27=P+1;
	t28=stack[int(t27)];
	t30=heap[int(t28)];
	t28=t28+1;
	fmt.Printf("%c", int(91));
	L3:
	t32=heap[int(t28)];
	t31=t32;
	if t29 >= t30 {goto L2;}
	if t2 == t31 {goto L5;}
	if t4 == t31 {goto L5;}
	if t6 == t31 {goto L5;}
	if t8 == t31 {goto L5;}
	if t15 == t31 {goto L5;}
	if t21 == t31 {goto L5;}
	if t11 == t32 {goto L4;}
	if t13 == t32 {goto L4;}
	if t17 == t32 {goto L4;}
	if t19 == t32 {goto L4;}
	fmt.Printf("%d", int(t32));
	fmt.Printf("%c", int(44));
	t28=t28+1;
	t29=t29+1;
	goto L3;
	L5:
	t33=P;
	t34=t29;
	t35=t30;
	t36=t32;
	t37=t27;
	t38=t28;
	t29=0;
	stack[int(t27)]=t32;
	print_array();
	t29=t34+1;
	t30=t35;
	t32=t36;
	t27=t37;
	t28=t38+1;
	goto L3;
	L4:
	t39=P;
	t40=t29;
	t41=t30;
	t42=t32;
	t43=t27;
	t44=t28;
	stack[int(t27)]=t32;
	print_string();
	P=t39;
	t29=t40+1;
	t30=t41;
	t32=t42;
	t27=t43;
	t28=t44+1;
	fmt.Printf("%c", int(44));
	goto L3;
	L2:
	fmt.Printf("%c", int(93));
	t29=0;
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
	
	/* compilacion de valor de variable */
	t10=H;
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
	t10=t10+0.12837;
	t11=t10;
	t12=H;
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
	heap[int(H)]=-1;
	H=H+1;
	t12=t12+0.12837;
	t13=t12;
	t14=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=t10;
	H=H+1;
	heap[int(H)]=t12;
	H=H+1;
	t14=t14+0.12837;
	t15=t14;
	t16=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t16=t16+0.12837;
	t17=t16;
	t18=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=113;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t18=t18+0.12837;
	t19=t18;
	t20=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=t16;
	H=H+1;
	heap[int(H)]=t18;
	H=H+1;
	t20=t20+0.12837;
	t21=t20;
	t22=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=t14;
	H=H+1;
	heap[int(H)]=t20;
	H=H+1;
	t22=t22+0.12837;
	/* fin de compilacion de variable */
	stack[int(2)]=t22;
	
	/* compilacion de accesso */
	t23=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t45=P+3;
	t45=t45+1;
	stack[int(t45)]=t23;
	P=P+3;
	print_array();
	t46=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de accesso */
	t47=stack[int(1)];
	/* fin de la compilacion de acceso */
	
	t48=P+3;
	t48=t48+1;
	stack[int(t48)]=t47;
	P=P+3;
	print_array();
	t49=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t50=stack[int(1)];
	t51=1+t50;
	t50=heap[int(t51)];
	fmt.Printf("%f", t50);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t52=stack[int(0)];
	t53=3+t52;
	t52=heap[int(t53)];
	fmt.Printf("%f", t52);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t54=stack[int(2)];
	t55=1+t54;
	t54=heap[int(t55)];
	t56=1+t54;
	t54=heap[int(t56)];
	t57=P+3;
	t57=t57+1;
	stack[int(t57)]=t54;
	P=P+3;
	print_string();
	t58=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}