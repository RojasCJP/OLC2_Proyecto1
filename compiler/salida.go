/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t9=P+1;
	t10=stack[int(t9)];
	L3:
	t11=heap[int(t10)];
	if t11 == -1 {goto L2;}
	fmt.Printf("%c", int(t11));
	t10=t10+1;
	goto L3;
	L2:
	return;
}
func print_array(){
	t55=P+1;
	t56=stack[int(t55)];
	t58=heap[int(t56)];
	t56=t56+1;
	fmt.Printf("%c", int(91));
	L5:
	t60=heap[int(t56)];
	t59=t60;
	if t57 >= t58 {goto L4;}
	if t46 == t60 {goto L6;}
	if t48 == t60 {goto L6;}
	if t50 == t60 {goto L6;}
	fmt.Printf("%f", t60);
	fmt.Printf("%c", int(44));
	t56=t56+1;
	t57=t57+1;
	goto L5;
	L6:
	t61=P;
	t62=t57;
	t63=t58;
	t64=t60;
	t65=t55;
	t66=t56;
	stack[int(t55)]=t60;
	print_string();
	P=t61;
	t57=t62+1;
	t58=t63;
	t60=t64;
	t55=t65;
	t56=t66+1;
	fmt.Printf("%c", int(44));
	goto L5;
	L4:
	fmt.Printf("%c", int(93));
	t57=0;
	return;
}

/*-----FUNCS-----*/
func AgregarFamiliar(){
	t1=P+1;
	t0=stack[int(t1)];
	t2=t0+2;
	t3=heap[int(t2)];
	t4=t3+1;
	t6=P+1;
	t5=stack[int(t6)];
	t7=t5+2;
	heap[int(t7)]=t4;
	return;
}
func ImprimirDatosPersona(){
	t8=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=78;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=58;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t8=t8+0.12837;
	t12=P+2;
	t12=t12+1;
	stack[int(t12)]=t8;
	P=P+2;
	print_string();
	t13=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	t15=P+1;
	t14=stack[int(t15)];
	t16=t14+0;
	t17=heap[int(t16)];
	t18=P+2;
	t18=t18+1;
	stack[int(t18)]=t17;
	P=P+2;
	print_string();
	t19=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t20=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=58;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t20=t20+0.12837;
	t21=P+2;
	t21=t21+1;
	stack[int(t21)]=t20;
	P=P+2;
	print_string();
	t22=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	t24=P+1;
	t23=stack[int(t24)];
	t25=t23+1;
	t26=heap[int(t25)];
	fmt.Printf("%d", int(t26));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t27=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=78;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=58;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t27=t27+0.12837;
	t28=P+2;
	t28=t28+1;
	stack[int(t28)]=t27;
	P=P+2;
	print_string();
	t29=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	t31=P+1;
	t30=stack[int(t31)];
	t32=t30+2;
	t33=heap[int(t32)];
	fmt.Printf("%d", int(t33));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	return;
}

func main(){
	/* compilacion de valor de variable */
	t34=H;
	t35=t34;
	H=H+3;
	t36=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=77;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t36=t36+0.12837;
	heap[int(t35)]=t36;
	t35=t35+1;
	heap[int(t35)]=22;
	t35=t35+1;
	heap[int(t35)]=4;
	t35=t35+1;
	/* fin de compilacion de variable */
	stack[int(0)]=t34;
	
	/* compilacion de accesso */
	t37=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t38=P+2;
	stack[int(t38)]=t37;
	P=P+1;
	ImprimirDatosPersona();
	t38=stack[int(P)];
	P=P-1;
	/* compilacion de accesso */
	t39=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t40=P+2;
	stack[int(t40)]=t39;
	P=P+1;
	AgregarFamiliar();
	t40=stack[int(P)];
	P=P-1;
	/* compilacion de accesso */
	t41=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t42=P+2;
	stack[int(t42)]=t41;
	P=P+1;
	ImprimirDatosPersona();
	t42=stack[int(P)];
	P=P-1;
	/* compilacion de valor de variable */
	t43=H;
	t44=t43;
	H=H+1;
	t45=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=80;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t45=t45+0.12837;
	t46=t45;
	t47=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=80;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=50;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t47=t47+0.12837;
	t48=t47;
	t49=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=80;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=51;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t49=t49+0.12837;
	t50=t49;
	t51=H;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t45;
	H=H+1;
	heap[int(H)]=t47;
	H=H+1;
	heap[int(H)]=t49;
	H=H+1;
	t51=t51+0.12837;
	heap[int(t44)]=t51;
	t44=t44+1;
	/* fin de compilacion de variable */
	stack[int(1)]=t43;
	
	t52=stack[int(1)];
	t53=t52+0;
	t54=heap[int(t53)];
	P=P+2;
	P=P-2;
	t67=P+2;
	t67=t67+1;
	stack[int(t67)]=t54;
	P=P+2;
	print_array();
	t68=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}