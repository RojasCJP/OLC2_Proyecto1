/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103 float64
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func print_string(){
	t11=P+1;
	t12=stack[int(t11)];
	L27:
	t13=heap[int(t12)];
	if t13 == -1 {goto L26;}
	fmt.Printf("%c", int(t13));
	t12=t12+1;
	goto L27;
	L26:
	return;
}

/*-----FUNCS-----*/
func suma(){
	/* compilacion de accesso */
	t99=P+1;
	t98=stack[int(t99)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t101=P+2;
	t100=stack[int(t101)];
	/* fin de la compilacion de acceso */
	
	t102=t98+t100;
	stack[int(P)]=t102;
	goto L83;
	L83:
	return;
}

func main(){
	t0=5+5;
	fmt.Printf("%d", int(t0));
	fmt.Printf("%c", int(32));
	t1=1*100;
	t2=t1/2;
	fmt.Printf("%f", t2);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* inicio de expression realcional */
	if 5 > 0 {goto L0;}
	goto L1;
	/* fin de la expression relacional */
	
	L0:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L2;
	L1:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L2:
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L3;
	/* goto para evitar errores */
	goto L4;
	L3:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L5;
	L4:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L5:
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* inicio de expression realcional */
	/* inicio de expression realcional */
	if 5 == 10 {goto L6;}
	goto L7;
	/* fin de la expression relacional */
	
	L6:
	t3=1;
	goto L8;
	L7:
	t3=0;
	L8:
	goto L10;
	/* goto para evitar errores */
	goto L9;
	L9:
	t4=1;
	goto L11;
	L10:
	t4=0;
	L11:
	if t3 == t4 {goto L12;}
	goto L13;
	/* fin de la expression relacional */
	
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
	fmt.Printf("%c", int(10));
	/* inicio de expression logica */
	/* inicio de expression realcional */
	if 5 == 10 {goto L17;}
	goto L16;
	/* fin de la expression relacional */
	
	L17:
	/* inicio de expression realcional */
	if 1 != 1 {goto L15;}
	goto L16;
	/* fin de la expression relacional */
	
	/* finalizo la expression logica */
	
	L15:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L18;
	L16:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L18:
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* inicio de expression logica */
	/* inicio de expression realcional */
	if 5 == 10 {goto L19;}
	goto L21;
	/* fin de la expression relacional */
	
	L21:
	/* inicio de expression realcional */
	if 1 == 1 {goto L19;}
	goto L20;
	/* fin de la expression relacional */
	
	/* finalizo la expression logica */
	
	L19:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L22;
	L20:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L22:
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de valor de variable */
	t5=10+10;
	t6=2+2;
	t7=8/t6;
	t8=t5*t7;
	/* fin de compilacion de variable */
	stack[int(0)]=t8;
	
	/* compilacion de valor de variable */
	t9=H;
	heap[int(H)]=81;
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
	/* fin de compilacion de variable */
	stack[int(1)]=t9;
	
	/* compilacion de valor de variable */
	goto L23;
	/* goto para evitar errores */
	goto L24;
	/* fin de compilacion de variable */
	L23:
	stack[int(2)]=1;
	goto L25;
	L24:
	stack[int(2)]=0;
	L25:
	
	t10=H;
	heap[int(H)]=72;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=77;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=33;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t14=P+3;
	t14=t14+1;
	stack[int(t14)]=t10;
	P=P+3;
	print_string();
	t15=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de accesso */
	t16=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%f", t16);
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de accesso */
	t17=stack[int(1)];
	/* fin de la compilacion de acceso */
	
	t18=P+3;
	t18=t18+1;
	stack[int(t18)]=t17;
	P=P+3;
	print_string();
	t19=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de accesso */
	t20=stack[int(2)];
	if t20 == 1 {goto L28;}
	goto L29;
	/* fin de la compilacion de acceso */
	
	L28:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L30;
	L29:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L30:
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* iniciando el if */
	goto L31;
	/* goto para evitar errores */
	goto L32;
	L31:
	goto L33;
	/* goto para evitar errores */
	goto L34;
	L33:
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(101));
	goto L35;
	L34:
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(108));
	fmt.Printf("%c", int(115));
	fmt.Printf("%c", int(101));
	L35:
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L32:
	/* iniciando el if */
	/* inicio de expression realcional */
	if 5 > 1 {goto L36;}
	goto L37;
	/* fin de la expression relacional */
	
	L36:
	t21=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=78;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t22=P+3;
	t22=t22+1;
	stack[int(t22)]=t21;
	P=P+3;
	print_string();
	t23=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L37:
	/* iniciando el if */
	/* inicio de expression realcional */
	t24=5-10;
	if t24 < 0 {goto L38;}
	goto L39;
	/* fin de la expression relacional */
	
	L38:
	t25=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t26=P+3;
	t26=t26+1;
	stack[int(t26)]=t25;
	P=P+3;
	print_string();
	t27=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L40;
	L39:
	t28=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t29=P+3;
	t29=t29+1;
	stack[int(t29)]=t28;
	P=P+3;
	print_string();
	t30=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L40:
	/* iniciando el if */
	/* inicio de expression realcional */
	if 1 >= 2 {goto L41;}
	goto L42;
	/* fin de la expression relacional */
	
	L41:
	t31=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t32=P+3;
	t32=t32+1;
	stack[int(t32)]=t31;
	P=P+3;
	print_string();
	t33=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L43;
	L42:
	/* iniciando el if */
	/* inicio de expression realcional */
	if 1 <= 1 {goto L44;}
	goto L45;
	/* fin de la expression relacional */
	
	L44:
	t34=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t35=P+3;
	t35=t35+1;
	stack[int(t35)]=t34;
	P=P+3;
	print_string();
	t36=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L45:
	L43:
	/* iniciando el if */
	/* inicio de expression realcional */
	if 10 == 11 {goto L46;}
	goto L47;
	/* fin de la expression relacional */
	
	L46:
	t37=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t38=P+3;
	t38=t38+1;
	stack[int(t38)]=t37;
	P=P+3;
	print_string();
	t39=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L48;
	L47:
	/* iniciando el if */
	/* inicio de expression realcional */
	if 10 != 11 {goto L49;}
	goto L50;
	/* fin de la expression relacional */
	
	L49:
	t40=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t41=P+3;
	t41=t41+1;
	stack[int(t41)]=t40;
	P=P+3;
	print_string();
	t42=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L51;
	L50:
	t43=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t44=P+3;
	t44=t44+1;
	stack[int(t44)]=t43;
	P=P+3;
	print_string();
	t45=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L51:
	L48:
	/* iniciando el if */
	goto L53;
	/* goto para evitar errores */
	goto L52;
	L52:
	t46=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=42;
	H=H+1;
	heap[int(H)]=50;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t47=P+3;
	t47=t47+1;
	stack[int(t47)]=t46;
	P=P+3;
	print_string();
	t48=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L54;
	L53:
	/* iniciando el if */
	goto L55;
	/* goto para evitar errores */
	goto L56;
	L55:
	t49=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=42;
	H=H+1;
	heap[int(H)]=50;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=49;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t50=P+3;
	t50=t50+1;
	stack[int(t50)]=t49;
	P=P+3;
	print_string();
	t51=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L57;
	L56:
	/* iniciando el if */
	goto L59;
	/* goto para evitar errores */
	goto L58;
	L58:
	t52=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=42;
	H=H+1;
	heap[int(H)]=50;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=50;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t53=P+3;
	t53=t53+1;
	stack[int(t53)]=t52;
	P=P+3;
	print_string();
	t54=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L60;
	L59:
	t55=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=42;
	H=H+1;
	heap[int(H)]=50;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t56=P+3;
	t56=t56+1;
	stack[int(t56)]=t55;
	P=P+3;
	print_string();
	t57=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L60:
	L57:
	L54:
	/* iniciando el if */
	/* inicio de expression realcional */
	if 5 > 1 {goto L61;}
	goto L62;
	/* fin de la expression relacional */
	
	L61:
	t58=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=78;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t59=P+3;
	t59=t59+1;
	stack[int(t59)]=t58;
	P=P+3;
	print_string();
	t60=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* iniciando el if */
	goto L63;
	/* goto para evitar errores */
	goto L64;
	L63:
	t61=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t62=P+3;
	t62=t62+1;
	stack[int(t62)]=t61;
	P=P+3;
	print_string();
	t63=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* iniciando el if */
	goto L66;
	/* goto para evitar errores */
	goto L65;
	L65:
	t64=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t65=P+3;
	t65=t65+1;
	stack[int(t65)]=t64;
	P=P+3;
	print_string();
	t66=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L67;
	L66:
	t67=H;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t68=P+3;
	t68=t68+1;
	stack[int(t68)]=t67;
	P=P+3;
	print_string();
	t69=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L67:
	goto L68;
	L64:
	/* iniciando el if */
	goto L70;
	/* goto para evitar errores */
	goto L69;
	L69:
	t70=H;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t71=P+3;
	t71=t71+1;
	stack[int(t71)]=t70;
	P=P+3;
	print_string();
	t72=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L71;
	L70:
	t73=H;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
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
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=102;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t74=P+3;
	t74=t74+1;
	stack[int(t74)]=t73;
	P=P+3;
	print_string();
	t75=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	L71:
	L68:
	L62:
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	stack[int(0)]=10;
	
	L72:
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t76=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	if t76 > 0 {goto L73;}
	goto L74;
	/* fin de la expression relacional */
	
	L73:
	t77=H;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=118;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=32;
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
	t78=P+3;
	t78=t78+1;
	stack[int(t78)]=t77;
	P=P+3;
	print_string();
	t79=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	/* compilacion de accesso */
	t80=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%d", int(t80));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t81=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t82=t81-1;
	/* fin de compilacion de variable */
	stack[int(0)]=t82;
	
	goto L72;
	L74:
	L75:
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t83=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	if t83 < 5 {goto L76;}
	goto L77;
	/* fin de la expression relacional */
	
	L76:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t84=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t85=t84+1;
	/* fin de compilacion de variable */
	stack[int(0)]=t85;
	
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t86=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	if t86 == 3 {goto L78;}
	goto L79;
	/* fin de la expression relacional */
	
	L78:
	t87=H;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t88=P+3;
	t88=t88+1;
	stack[int(t88)]=t87;
	P=P+3;
	print_string();
	t89=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L75;
	goto L80;
	L79:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t90=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	if t90 == 4 {goto L81;}
	goto L82;
	/* fin de la expression relacional */
	
	L81:
	t91=H;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t92=P+3;
	t92=t92+1;
	stack[int(t92)]=t91;
	P=P+3;
	print_string();
	t93=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L77;
	L82:
	L80:
	t94=H;
	heap[int(H)]=69;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=118;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=32;
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
	t95=P+3;
	t95=t95+1;
	stack[int(t95)]=t94;
	P=P+3;
	print_string();
	t96=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	/* compilacion de accesso */
	t97=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%d", int(t97));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	goto L75;
	L77:
	t103=P+4;
	stack[int(t103)]=10;
	t103=t103+1;
	stack[int(t103)]=20;
	P=P+3;
	suma();
	t103=stack[int(P)];
	P=P-3;
	fmt.Printf("%d", int(t103));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}