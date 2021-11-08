/*----HEADER----*/
package main;

import (
	"fmt"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158, t159, t160, t161, t162, t163, t164, t165, t166 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func potencia(){
	t116=P+1;
	t117=stack[int(t116)];
	t116=t116+1;
	t118=stack[int(t116)];
	t116=t117;
	L29:
	if t118 <= 1 {goto L30;}
	t117=t117*t116;
	t118=t118-1;
	goto L29;
	L30:
	stack[int(P)]=t117;
	return;
}
func print_string(){
	t128=P+1;
	t129=stack[int(t128)];
	L32:
	t130=heap[int(t129)];
	if t130 == -1 {goto L31;}
	fmt.Printf("%c", int(t130));
	t129=t129+1;
	goto L32;
	L31:
	return;
}
func print_array(){
	t134=P+1;
	t135=stack[int(t134)];
	t137=heap[int(t135)];
	t135=t135+1;
	fmt.Printf("%c", int(91));
	L34:
	t139=heap[int(t135)];
	t138=t139;
	if t136 >= t137 {goto L33;}
	fmt.Printf("%f", t139);
	fmt.Printf("%c", int(44));
	t135=t135+1;
	t136=t136+1;
	goto L34;
	L33:
	fmt.Printf("%c", int(93));
	t136=0;
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
	
	t18=t16+t13;
	t13=heap[int(t18)];
	t15=t18;
	heap[int(t15)]=t7;
	/* compilacion de accesso */
	t20=P+4;
	t19=stack[int(t20)];
	/* fin de la compilacion de acceso */
	
	/* cambiando el valor de arreglo */
	t22=P+3;
	t21=stack[int(t22)];
	/* compilacion de accesso */
	t25=P+2;
	t24=stack[int(t25)];
	/* fin de la compilacion de acceso */
	
	t26=t24+t21;
	t21=heap[int(t26)];
	t23=t26;
	heap[int(t23)]=t19;
	return;
}
func bubbleSort(){
	/* compilacion length */
	/* compilacion de accesso */
	t28=P+1;
	t27=stack[int(t28)];
	/* fin de la compilacion de acceso */
	
	t29=t27;
	t29=heap[int(t29)];
	t30=t29-1;
	t31=0;
	L6:
	if t31 > t30 {goto L7;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t32=P+2;
	stack[int(t32)]=t31;
	
	/* compilacion length */
	/* compilacion de accesso */
	t34=P+1;
	t33=stack[int(t34)];
	/* fin de la compilacion de acceso */
	
	t35=t33;
	t35=heap[int(t35)];
	t36=t35-1;
	t37=1;
	L8:
	if t37 > t36 {goto L9;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t38=P+3;
	stack[int(t38)]=t37;
	
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de acceso arreglos */
	t40=P+1;
	t39=stack[int(t40)];
	/* compilacion de accesso */
	t42=P+3;
	t41=stack[int(t42)];
	/* fin de la compilacion de acceso */
	
	t44=heap[int(t39)];
	t43=t41+t39;
	if t44 < t41 {goto L10;}
	goto L11;
	L10:
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
	return;	L11:
	t39=heap[int(t43)];
	/* compilacion de acceso arreglos */
	t46=P+1;
	t45=stack[int(t46)];
	/* compilacion de accesso */
	t48=P+3;
	t47=stack[int(t48)];
	/* fin de la compilacion de acceso */
	
	t49=t47+1;
	t51=heap[int(t45)];
	t50=t49+t45;
	if t51 < t49 {goto L12;}
	goto L13;
	L12:
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
	return;	L13:
	t45=heap[int(t50)];
	if t39 > t45 {goto L14;}
	goto L15;
	/* fin de la expression relacional */
	
	L14:
	/* compilacion de accesso */
	t53=P+3;
	t52=stack[int(t53)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t55=P+3;
	t54=stack[int(t55)];
	/* fin de la compilacion de acceso */
	
	t56=t54+1;
	/* compilacion de accesso */
	t58=P+1;
	t57=stack[int(t58)];
	/* fin de la compilacion de acceso */
	
	t59=P+5;
	stack[int(t59)]=t52;
	t59=t59+1;
	stack[int(t59)]=t56;
	t59=t59+1;
	stack[int(t59)]=t57;
	P=P+4;
	swap();
	t59=stack[int(P)];
	P=P-4;
	L15:
	t37=t37+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t60=P+3;
	stack[int(t60)]=t37;
	
	goto L8;
	L9:
	t31=t31+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t61=P+2;
	stack[int(t61)]=t31;
	
	goto L6;
	L7:
	return;
}
func insertionSort(){
	/* compilacion length */
	/* compilacion de accesso */
	t63=P+1;
	t62=stack[int(t63)];
	/* fin de la compilacion de acceso */
	
	t64=t62;
	t64=heap[int(t64)];
	t65=2;
	L17:
	if t65 > t64 {goto L18;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t66=P+2;
	stack[int(t66)]=t65;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t68=P+2;
	t67=stack[int(t68)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t69=P+3;
	stack[int(t69)]=t67;
	
	/* compilacion de valor de variable */
	/* compilacion de acceso arreglos */
	t71=P+1;
	t70=stack[int(t71)];
	/* compilacion de accesso */
	t73=P+2;
	t72=stack[int(t73)];
	/* fin de la compilacion de acceso */
	
	t75=heap[int(t70)];
	t74=t72+t70;
	if t75 < t72 {goto L19;}
	goto L20;
	L19:
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
	return;	L20:
	t70=heap[int(t74)];
	/* fin de compilacion de variable */
	t76=P+4;
	stack[int(t76)]=t70;
	
	L21:
	/* inicio de expression logica */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t78=P+3;
	t77=stack[int(t78)];
	/* fin de la compilacion de acceso */
	
	if t77 > 1 {goto L24;}
	goto L23;
	/* fin de la expression relacional */
	
	L24:
	/* inicio de expression realcional */
	/* compilacion de acceso arreglos */
	t80=P+1;
	t79=stack[int(t80)];
	/* compilacion de accesso */
	t82=P+3;
	t81=stack[int(t82)];
	/* fin de la compilacion de acceso */
	
	t83=t81-1;
	t85=heap[int(t79)];
	t84=t83+t79;
	if t85 < t83 {goto L25;}
	goto L26;
	L25:
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
	return;	L26:
	t79=heap[int(t84)];
	/* compilacion de accesso */
	t87=P+4;
	t86=stack[int(t87)];
	/* fin de la compilacion de acceso */
	
	if t79 > t86 {goto L22;}
	goto L23;
	/* fin de la expression relacional */
	
	/* finalizo la expression logica */
	
	L22:
	/* compilacion de acceso arreglos */
	t89=P+1;
	t88=stack[int(t89)];
	/* compilacion de accesso */
	t91=P+3;
	t90=stack[int(t91)];
	/* fin de la compilacion de acceso */
	
	t92=t90-1;
	t94=heap[int(t88)];
	t93=t92+t88;
	if t94 < t92 {goto L27;}
	goto L28;
	L27:
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
	return;	L28:
	t88=heap[int(t93)];
	/* cambiando el valor de arreglo */
	t96=P+1;
	t95=stack[int(t96)];
	/* compilacion de accesso */
	t99=P+3;
	t98=stack[int(t99)];
	/* fin de la compilacion de acceso */
	
	t100=t98+t95;
	t95=heap[int(t100)];
	t97=t100;
	heap[int(t97)]=t88;
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t102=P+3;
	t101=stack[int(t102)];
	/* fin de la compilacion de acceso */
	
	t103=t101-1;
	/* fin de compilacion de variable */
	t104=P+3;
	stack[int(t104)]=t103;
	
	goto L21;
	L23:
	/* compilacion de accesso */
	t106=P+4;
	t105=stack[int(t106)];
	/* fin de la compilacion de acceso */
	
	/* cambiando el valor de arreglo */
	t108=P+1;
	t107=stack[int(t108)];
	/* compilacion de accesso */
	t111=P+3;
	t110=stack[int(t111)];
	/* fin de la compilacion de acceso */
	
	t112=t110+t107;
	t107=heap[int(t112)];
	t109=t112;
	heap[int(t109)]=t105;
	t65=t65+1;
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t113=P+2;
	stack[int(t113)]=t65;
	
	goto L17;
	L18:
	return;
}

func main(){
	/* compilacion de valor de variable */
	t114=7*3;
	t119=P+0;
	t119=t119+1;
	stack[int(t119)]=9874;
	t119=t119+1;
	stack[int(t119)]=0;
	P=P+0;
	potencia();
	t120=stack[int(P)];
	P=P-0;
	t121=820*10;
	t122=8*0;
	t123=t122+8;
	t124=H;
	heap[int(H)]=16;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=t114;
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
	heap[int(H)]=t120;
	H=H+1;
	heap[int(H)]=44;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t121;
	H=H+1;
	heap[int(H)]=11;
	H=H+1;
	heap[int(H)]=t123;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	t124=t124+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t124;
	
	/* compilacion de accesso */
	t125=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t126=P+2;
	stack[int(t126)]=t125;
	P=P+1;
	bubbleSort();
	t126=stack[int(P)];
	P=P-1;
	t127=H;
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
	t127=t127+0.12837;
	t131=P+1;
	t131=t131+1;
	stack[int(t131)]=t127;
	P=P+1;
	print_string();
	t132=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	/* compilacion de accesso */
	t133=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	P=P+1;
	P=P-1;
	t140=P+1;
	t140=t140+1;
	stack[int(t140)]=t133;
	P=P+1;
	print_array();
	t141=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de valor de variable */
	t142=7*3;
	t144=P+1;
	t144=t144+1;
	stack[int(t144)]=9874;
	t144=t144+1;
	stack[int(t144)]=1;
	P=P+1;
	potencia();
	t145=stack[int(P)];
	P=P-1;
	t146=820*10;
	t147=8*0;
	t148=t147+8;
	t149=H;
	heap[int(H)]=16;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=t142;
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
	heap[int(H)]=t145;
	H=H+1;
	heap[int(H)]=44;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t146;
	H=H+1;
	heap[int(H)]=11;
	H=H+1;
	heap[int(H)]=t148;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	t149=t149+0.12837;
	/* fin de compilacion de variable */
	stack[int(0)]=t149;
	
	/* compilacion de acceso arreglos */
	t150=stack[int(0)];
	t152=heap[int(t150)];
	t151=2+t150;
	if t152 < 2 {goto L37;}
	goto L38;
	L37:
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
	return;	L38:
	t150=heap[int(t151)];
	t154=P+1;
	t154=t154+1;
	stack[int(t154)]=t150;
	t154=t154+1;
	stack[int(t154)]=0;
	P=P+1;
	potencia();
	t155=stack[int(P)];
	P=P-1;
	/* cambiando el valor de arreglo */
	t156=stack[int(0)];
	t158=1+t156;
	t156=heap[int(t158)];
	t157=t158;
	heap[int(t157)]=t155;
	/* compilacion de accesso */
	t159=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	t160=P+2;
	stack[int(t160)]=t159;
	P=P+1;
	insertionSort();
	t160=stack[int(P)];
	P=P-1;
	t161=H;
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
	t161=t161+0.12837;
	t162=P+1;
	t162=t162+1;
	stack[int(t162)]=t161;
	P=P+1;
	print_string();
	t163=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));
	/* compilacion de accesso */
	t164=stack[int(0)];
	/* fin de la compilacion de acceso */
	
	P=P+1;
	P=P-1;
	t165=P+1;
	t165=t165+1;
	stack[int(t165)]=t164;
	P=P+1;
	print_array();
	t166=stack[int(P)];
	P=P-1;
	fmt.Printf("%c", int(32));

}