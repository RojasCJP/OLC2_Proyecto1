/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158, t159, t160, t161, t162, t163, t164, t165, t166, t167, t168, t169, t170, t171 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func potencia(){
	t120=P+1;
	t121=stack[int(t120)];
	t120=t120+1;
	t122=stack[int(t120)];
	t120=t121;
	L37:
	if t122 <= 1 {goto L38;}
	t121=t121*t120;
	t122=t122-1;
	goto L37;
	L38:
	stack[int(P)]=t121;
	return;
}
func print_string(){
	t132=P+1;
	t133=stack[int(t132)];
	L40:
	t134=heap[int(t133)];
	if t134 == -1 {goto L39;}
	fmt.Printf("%c", int(t134));
	t133=t133+1;
	goto L40;
	L39:
	return;
}
func print_array(){
	t138=P+1;
	t139=stack[int(t138)];
	t141=heap[int(t139)];
	t139=t139+1;
	fmt.Printf("%c", int(91));
	L42:
	t143=heap[int(t139)];
	t142=t143;
	if t140 >= t141 {goto L41;}
	fmt.Printf("%f", t143);
	fmt.Printf("%c", int(44));
	t139=t139+1;
	t140=t140+1;
	goto L42;
	L41:
	fmt.Printf("%c", int(93));
	t140=0;
	return;
}

/*-----FUNCS-----*/
func swap(){
	/* compilacion de valor de variable */
	/* compilacion de acceso arreglos */
	t1=P+3;
	t0=stack[int(t1)];
	/* compilacion de accesso */
	t3=P+1;
	t2=stack[int(t3)];
	/* fin de la compilacion de acceso */
	
	t5=heap[int(t0)];
	t4=t2+t0;
	if t5 < t2 {goto L1;}
	goto L2;
	L1:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L2:
	t0=heap[int(t4)];
	/* fin de compilacion de variable */
	t6=P+4;
	stack[int(t6)]=t0;
	
	/* compilacion de acceso arreglos */
	t8=P+3;
	t7=stack[int(t8)];
	/* compilacion de accesso */
	t10=P+2;
	t9=stack[int(t10)];
	/* fin de la compilacion de acceso */
	
	t12=heap[int(t7)];
	t11=t9+t7;
	if t12 < t9 {goto L3;}
	goto L4;
	L3:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L4:
	t7=heap[int(t11)];
	/* cambiando el valor de arreglo */
	t14=P+3;
	t13=stack[int(t14)];
	/* compilacion de accesso */
	t17=P+1;
	t16=stack[int(t17)];
	/* fin de la compilacion de acceso */
	
	t19=heap[int(t13)];
	t18=t16+t13;
	if t19 < t16 {goto L5;}
	goto L6;
	L5:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L6:
	t13=heap[int(t18)];
	t15=t18;
	heap[int(t15)]=t7;
	/* compilacion de accesso */
	t21=P+4;
	t20=stack[int(t21)];
	/* fin de la compilacion de acceso */
	
	/* cambiando el valor de arreglo */
	t23=P+3;
	t22=stack[int(t23)];
	/* compilacion de accesso */
	t26=P+2;
	t25=stack[int(t26)];
	/* fin de la compilacion de acceso */
	
	t28=heap[int(t22)];
	t27=t25+t22;
	if t28 < t25 {goto L7;}
	goto L8;
	L7:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L8:
	t22=heap[int(t27)];
	t24=t27;
	heap[int(t24)]=t20;
	return;
}
func bubbleSort(){
	/* compilacion length */
	/* compilacion de accesso */
	t30=P+1;
	t29=stack[int(t30)];
	/* fin de la compilacion de acceso */
	
	t31=t29;
	t31=heap[int(t31)];
	t32=t31-1;
	t33=0;
	L10:
	if t33 > t32 {goto L11;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t34=P+2;
	stack[int(t34)]=t33;
	
	/* compilacion length */
	/* compilacion de accesso */
	t36=P+1;
	t35=stack[int(t36)];
	/* fin de la compilacion de acceso */
	
	t37=t35;
	t37=heap[int(t37)];
	t38=t37-1;
	t39=1;
	L12:
	if t39 > t38 {goto L13;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t40=P+3;
	stack[int(t40)]=t39;
	
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de acceso arreglos */
	t42=P+1;
	t41=stack[int(t42)];
	/* compilacion de accesso */
	t44=P+3;
	t43=stack[int(t44)];
	/* fin de la compilacion de acceso */
	
	t46=heap[int(t41)];
	t45=t43+t41;
	if t46 < t43 {goto L14;}
	goto L15;
	L14:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L15:
	t41=heap[int(t45)];
	/* compilacion de acceso arreglos */
	t48=P+1;
	t47=stack[int(t48)];
	/* compilacion de accesso */
	t50=P+3;
	t49=stack[int(t50)];
	/* fin de la compilacion de acceso */
	
	t51=t49+1;
	t53=heap[int(t47)];
	t52=t51+t47;
	if t53 < t51 {goto L16;}
	goto L17;
	L16:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L17:
	t47=heap[int(t52)];
	if t41 > t47 {goto L18;}
	goto L19;
	/* fin de la expression relacional */
	
	L18:
	/* compilacion de accesso */
	t55=P+3;
	t54=stack[int(t55)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t57=P+3;
	t56=stack[int(t57)];
	/* fin de la compilacion de acceso */
	
	t58=t56+1;
	/* compilacion de accesso */
	t60=P+1;
	t59=stack[int(t60)];
	/* fin de la compilacion de acceso */
	
	t61=P+5;
	stack[int(t61)]=t54;
	t61=t61+1;
	stack[int(t61)]=t58;
	t61=t61+1;
	stack[int(t61)]=t59;
	P=P+4;
	swap();
	t61=stack[int(P)];
	P=P-4;
	L19:
	t39=t39+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t62=P+3;
	stack[int(t62)]=t39;
	
	goto L12;
	L13:
	t33=t33+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t63=P+2;
	stack[int(t63)]=t33;
	
	goto L10;
	L11:
	return;
}
func insertionSort(){
	/* compilacion length */
	/* compilacion de accesso */
	t65=P+1;
	t64=stack[int(t65)];
	/* fin de la compilacion de acceso */
	
	t66=t64;
	t66=heap[int(t66)];
	t67=2;
	L21:
	if t67 > t66 {goto L22;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t68=P+2;
	stack[int(t68)]=t67;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t70=P+2;
	t69=stack[int(t70)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t71=P+3;
	stack[int(t71)]=t69;
	
	/* compilacion de valor de variable */
	/* compilacion de acceso arreglos */
	t73=P+1;
	t72=stack[int(t73)];
	/* compilacion de accesso */
	t75=P+2;
	t74=stack[int(t75)];
	/* fin de la compilacion de acceso */
	
	t77=heap[int(t72)];
	t76=t74+t72;
	if t77 < t74 {goto L23;}
	goto L24;
	L23:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L24:
	t72=heap[int(t76)];
	/* fin de compilacion de variable */
	t78=P+4;
	stack[int(t78)]=t72;
	
	L25:
	/* inicio de expression logica */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t80=P+3;
	t79=stack[int(t80)];
	/* fin de la compilacion de acceso */
	
	if t79 > 1 {goto L28;}
	goto L27;
	/* fin de la expression relacional */
	
	L28:
	/* inicio de expression realcional */
	/* compilacion de acceso arreglos */
	t82=P+1;
	t81=stack[int(t82)];
	/* compilacion de accesso */
	t84=P+3;
	t83=stack[int(t84)];
	/* fin de la compilacion de acceso */
	
	t85=t83-1;
	t87=heap[int(t81)];
	t86=t85+t81;
	if t87 < t85 {goto L29;}
	goto L30;
	L29:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L30:
	t81=heap[int(t86)];
	/* compilacion de accesso */
	t89=P+4;
	t88=stack[int(t89)];
	/* fin de la compilacion de acceso */
	
	if t81 > t88 {goto L26;}
	goto L27;
	/* fin de la expression relacional */
	
	/* finalizo la expression logica */
	
	L26:
	/* compilacion de acceso arreglos */
	t91=P+1;
	t90=stack[int(t91)];
	/* compilacion de accesso */
	t93=P+3;
	t92=stack[int(t93)];
	/* fin de la compilacion de acceso */
	
	t94=t92-1;
	t96=heap[int(t90)];
	t95=t94+t90;
	if t96 < t94 {goto L31;}
	goto L32;
	L31:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L32:
	t90=heap[int(t95)];
	/* cambiando el valor de arreglo */
	t98=P+1;
	t97=stack[int(t98)];
	/* compilacion de accesso */
	t101=P+3;
	t100=stack[int(t101)];
	/* fin de la compilacion de acceso */
	
	t103=heap[int(t97)];
	t102=t100+t97;
	if t103 < t100 {goto L33;}
	goto L34;
	L33:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L34:
	t97=heap[int(t102)];
	t99=t102;
	heap[int(t99)]=t90;
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t105=P+3;
	t104=stack[int(t105)];
	/* fin de la compilacion de acceso */
	
	t106=t104-1;
	/* fin de compilacion de variable */
	t107=P+3;
	stack[int(t107)]=t106;
	
	goto L25;
	L27:
	/* compilacion de accesso */
	t109=P+4;
	t108=stack[int(t109)];
	/* fin de la compilacion de acceso */
	
	/* cambiando el valor de arreglo */
	t111=P+1;
	t110=stack[int(t111)];
	/* compilacion de accesso */
	t114=P+3;
	t113=stack[int(t114)];
	/* fin de la compilacion de acceso */
	
	t116=heap[int(t110)];
	t115=t113+t110;
	if t116 < t113 {goto L35;}
	goto L36;
	L35:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L36:
	t110=heap[int(t115)];
	t112=t115;
	heap[int(t112)]=t108;
	t67=t67+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t117=P+2;
	stack[int(t117)]=t67;
	
	goto L21;
	L22:
	return;
}

func main(){
	/* compilacion de valor de variable */
	t118=7*3;
	t123=P+0;
	t123=t123+1;
	stack[int(t123)]=9874;
	t123=t123+1;
	stack[int(t123)]=0;
	P=P+0;
	potencia();
	t124=stack[int(P)];
	P=P-0;
	t125=820*10;
	t126=8*0;
	t127=t126+8;
	t128=H;
	heap[int(H)]=16;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=t118;
	H=H+1;
	heap[int(H)]=7;
	H=H+1;
	heap[int(H)]=89;
	H=H+1;
	heap[int(H)]=56;
	H=H+1;
	heap[int(H)]=909;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=9;
	H=H+1;
	heap[int(H)]=t124;
	H=H+1;
	heap[int(H)]=44;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t125;
	H=H+1;
	heap[int(H)]=11;
	H=H+1;
	heap[int(H)]=t127;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	t128=t128+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t128;
	
	/* compilacion de accesso */
	t129=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t130=P+2;
	stack[int(t130)]=t129;
	P=P+1;
	bubbleSort();
	t130=stack[int(P)];
	P=P-1;
	t131=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=66;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=98;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=83;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=61;
	H=H+1;
	heap[int(H)]=62;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t131=t131+0.12837;
	t135=P+1;
	t135=t135+1;
	stack[int(t135)]=t131;
	P=P+1;
	print_string();
	t136=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	/* compilacion de accesso */
	t137=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	P=P+1;
	P=P-1;
	t144=P+1;
	t144=t144+1;
	stack[int(t144)]=t137;
	P=P+1;
	print_array();
	t145=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de valor de variable */
	t146=7*3;
	t148=P+1;
	t148=t148+1;
	stack[int(t148)]=9874;
	t148=t148+1;
	stack[int(t148)]=1;
	P=P+1;
	potencia();
	t149=stack[int(P)];
	P=P-1;
	t150=820*10;
	t151=8*0;
	t152=t151+8;
	t153=H;
	heap[int(H)]=16;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=t146;
	H=H+1;
	heap[int(H)]=7;
	H=H+1;
	heap[int(H)]=89;
	H=H+1;
	heap[int(H)]=56;
	H=H+1;
	heap[int(H)]=909;
	H=H+1;
	heap[int(H)]=109;
	H=H+1;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=9;
	H=H+1;
	heap[int(H)]=t149;
	H=H+1;
	heap[int(H)]=44;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t150;
	H=H+1;
	heap[int(H)]=11;
	H=H+1;
	heap[int(H)]=t152;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	t153=t153+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t153;
	
	/* compilacion de acceso arreglos */
	t154=stack[int(0)];
	t156=heap[int(t154)];
	t155=2+t154;
	if t156 < 2 {goto L45;}
	goto L46;
	L45:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L46:
	t154=heap[int(t155)];
	t158=P+1;
	t158=t158+1;
	stack[int(t158)]=t154;
	t158=t158+1;
	stack[int(t158)]=0;
	P=P+1;
	potencia();
	t159=stack[int(P)];
	P=P-1;
	/* cambiando el valor de arreglo */
	t160=stack[int(0)];
	t163=heap[int(t160)];
	t162=1+t160;
	if t163 < 1 {goto L47;}
	goto L48;
	L47:
	fmt.Printf("%c", int(105));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(100));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(120));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(117));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(102));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(110));
	fmt.Printf("%c", int(103));
	fmt.Printf("%c", int(101));
	return;	L48:
	t160=heap[int(t162)];
	t161=t162;
	heap[int(t161)]=t159;
	/* compilacion de accesso */
	t164=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t165=P+2;
	stack[int(t165)]=t164;
	P=P+1;
	insertionSort();
	t165=stack[int(P)];
	P=P-1;
	t166=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=73;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=83;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=61;
	H=H+1;
	heap[int(H)]=62;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t166=t166+0.12837;
	t167=P+1;
	t167=t167+1;
	stack[int(t167)]=t166;
	P=P+1;
	print_string();
	t168=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	/* compilacion de accesso */
	t169=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	P=P+1;
	P=P-1;
	t170=P+1;
	t170=t170+1;
	stack[int(t170)]=t169;
	P=P+1;
	print_array();
	t171=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));

}