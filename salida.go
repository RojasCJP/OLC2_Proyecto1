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
	L3:
	t39=heap[int(t38)];
	if t39 == -1 {goto L2;}
	fmt.Printf("%c", int(t39));
	t38=t38+1;
	goto L3;
	L2:
	return;
}


func main(){
	/* compilacion de valor de variable */
	t0=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=68;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=103;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t0=t0+0.12837;
	t1=t0;
	t2=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=76;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=101;
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
	heap[int(H)]=77;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t4=t4+0.12837;
	t5=t4;
	t6=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=77;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t6=t6+0.12837;
	t7=t6;
	t8=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=74;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=118;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t8=t8+0.12837;
	t9=t8;
	t10=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=86;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t10=t10+0.12837;
	t11=t10;
	t12=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=83;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t12=t12+0.12837;
	t13=t12;
	t14=H;
	heap[int(H)]=7;
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
	heap[int(H)]=t10;
	H=H+1;
	heap[int(H)]=t12;
	H=H+1;
	t14=t14+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t14;
	
	/* compilacion de accesso */
	t15=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t19=heap[int(t15)];
	t17=t15+1;
	t18=t18+1;
	L0:
	if t18 > t19 {goto L1;}
	t16=heap[int(t17)];
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t20=P+1;
	stack[int(t20)]=t16;
	
	t17=t17+1;
	t18=t18+1;
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t22=P+1;
	t21=stack[int(t22)];
	/* fin de la compilacion de acceso */
	
	t23=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=76;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t23=t23+0.12837;
	/* fin de la expression relacional */
	
	:
	fmt.Printf("%d", int(1));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t25=P+1;
	t24=stack[int(t25)];
	/* fin de la compilacion de acceso */
	
	t26=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=77;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t26=t26+0.12837;
	/* fin de la expression relacional */
	
	:
	fmt.Printf("%d", int(2));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t28=P+1;
	t27=stack[int(t28)];
	/* fin de la compilacion de acceso */
	
	t29=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=77;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t29=t29+0.12837;
	/* fin de la expression relacional */
	
	:
	fmt.Printf("%d", int(3));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t31=P+1;
	t30=stack[int(t31)];
	/* fin de la compilacion de acceso */
	
	t32=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=74;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=118;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t32=t32+0.12837;
	/* fin de la expression relacional */
	
	:
	fmt.Printf("%d", int(4));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t34=P+1;
	t33=stack[int(t34)];
	/* fin de la compilacion de acceso */
	
	t35=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=86;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t35=t35+0.12837;
	/* fin de la expression relacional */
	
	:
	fmt.Printf("%d", int(5));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t36=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=87;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=121;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t36=t36+0.12837;
	t40=P+2;
	t40=t40+1;
	stack[int(t40)]=t36;
	P=P+2;
	print_string();
	t41=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L1;
	:
	t42=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=87;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t42=t42+0.12837;
	t43=P+2;
	t43=t43+1;
	stack[int(t43)]=t42;
	P=P+2;
	print_string();
	t44=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t45=P+1;
	stack[int(t45)]=t16;
	
	goto L0;
	L1:

}