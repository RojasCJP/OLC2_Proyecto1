/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20 float64
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


func main(){
	t0=5+5;
	fmt.Printf("%d", int(t0));
	t1=1*100;
	t2=t1/2;
	fmt.Printf("%d", int(t2));
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
	fmt.Printf("%c", int(10));
	/* compilacion de accesso */
	t16=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	fmt.Printf("%d", int(t16));
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
	fmt.Printf("%c", int(10));

}