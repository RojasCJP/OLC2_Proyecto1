/*----HEADER----*/
package main;

import (
	"fmt"
	"math"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158, t159, t160, t161, t162, t163, t164, t165, t166, t167, t168, t169, t170, t171, t172, t173, t174, t175, t176, t177, t178, t179, t180, t181, t182, t183, t184, t185, t186, t187, t188, t189, t190, t191, t192, t193, t194, t195, t196, t197, t198, t199, t200, t201, t202, t203, t204, t205, t206, t207, t208, t209, t210, t211, t212, t213, t214, t215, t216, t217, t218, t219, t220, t221, t222, t223, t224, t225, t226, t227, t228, t229, t230, t231, t232, t233, t234, t235, t236, t237, t238, t239, t240, t241, t242, t243, t244, t245, t246, t247, t248, t249, t250, t251, t252, t253, t254, t255, t256, t257, t258, t259, t260, t261, t262, t263, t264, t265, t266, t267 float64;
var P, H float64;
var stack [30101999]float64;
var heap [30101999]float64;

/*-----NATIVES-----*/
func potencia(){
	t137=P+1;
	t138=stack[int(t137)];
	t137=t137+1;
	t139=stack[int(t137)];
	t137=t138;
	L37:
	if t139 <= 1 {goto L38;}
	t138=t138*t137;
	t139=t139-1;
	goto L37;
	L38:
	stack[int(P)]=t138;
	return;
}
func print_string(){
	t149=P+1;
	t150=stack[int(t149)];
	L41:
	t151=heap[int(t150)];
	if t151 == -1 {goto L40;}
	fmt.Printf("%c", int(t151));
	t150=t150+1;
	goto L41;
	L40:
	return;
}

/*-----FUNCS-----*/
func quicksort(){
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t1=P+2;
	t0=stack[int(t1)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t2=P+4;
	stack[int(t2)]=t0;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t4=P+3;
	t3=stack[int(t4)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t5=P+5;
	stack[int(t5)]=t3;
	
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t7=P+4;
	t6=stack[int(t7)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t9=P+3;
	t8=stack[int(t9)];
	/* fin de la compilacion de acceso */
	
	if t6 >= t8 {goto L1;}
	goto L2;
	/* fin de la expression relacional */
	
	L1:
	stack[int(P)]=0;
	goto L0;
	L2:
	/* compilacion de valor de variable */
	/* compilacion de acceso arreglos */
	t11=P+1;
	t10=stack[int(t11)];
	/* compilacion de accesso */
	t13=P+4;
	t12=stack[int(t13)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t15=P+5;
	t14=stack[int(t15)];
	/* fin de la compilacion de acceso */
	
	t16=t12+t14;
	if 2 == 0 {goto L3;}
	t17=t16/2;
	goto L4;
	L3:
	fmt.Printf("%c", int(109));
	fmt.Printf("%c", int(97));
	fmt.Printf("%c", int(116));
	fmt.Printf("%c", int(104));
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(101));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(114));
	fmt.Printf("%c", int(111));
	fmt.Printf("%c", int(114));
	return;	L4:
	t18=t17;
	t19 = math.Mod(t18,1);
	t18=t18-t19;
	t21=heap[int(t10)];
	t20=t18+t10;
	if t21 < t18 {goto L5;}
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
	t10=heap[int(t20)];
	/* fin de compilacion de variable */
	t22=P+6;
	stack[int(t22)]=t10;
	
	L7:
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t24=P+4;
	t23=stack[int(t24)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t26=P+5;
	t25=stack[int(t26)];
	/* fin de la compilacion de acceso */
	
	if t23 < t25 {goto L8;}
	goto L9;
	/* fin de la expression relacional */
	
	L8:
	L10:
	/* inicio de expression logica */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t28=P+4;
	t27=stack[int(t28)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t30=P+5;
	t29=stack[int(t30)];
	/* fin de la compilacion de acceso */
	
	if t27 < t29 {goto L13;}
	goto L12;
	/* fin de la expression relacional */
	
	L13:
	/* inicio de expression realcional */
	/* compilacion de acceso arreglos */
	t32=P+1;
	t31=stack[int(t32)];
	/* compilacion de accesso */
	t34=P+4;
	t33=stack[int(t34)];
	/* fin de la compilacion de acceso */
	
	t36=heap[int(t31)];
	t35=t33+t31;
	if t36 < t33 {goto L14;}
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
	t31=heap[int(t35)];
	/* compilacion de accesso */
	t38=P+6;
	t37=stack[int(t38)];
	/* fin de la compilacion de acceso */
	
	if t31 < t37 {goto L11;}
	goto L12;
	/* fin de la expression relacional */
	
	/* finalizo la expression logica */
	
	L11:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t40=P+4;
	t39=stack[int(t40)];
	/* fin de la compilacion de acceso */
	
	t41=t39+1;
	/* fin de compilacion de variable */
	t42=P+4;
	stack[int(t42)]=t41;
	
	goto L10;
	L12:
	L16:
	/* inicio de expression logica */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t44=P+4;
	t43=stack[int(t44)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t46=P+5;
	t45=stack[int(t46)];
	/* fin de la compilacion de acceso */
	
	if t43 < t45 {goto L19;}
	goto L18;
	/* fin de la expression relacional */
	
	L19:
	/* inicio de expression realcional */
	/* compilacion de acceso arreglos */
	t48=P+1;
	t47=stack[int(t48)];
	/* compilacion de accesso */
	t50=P+5;
	t49=stack[int(t50)];
	/* fin de la compilacion de acceso */
	
	t52=heap[int(t47)];
	t51=t49+t47;
	if t52 < t49 {goto L20;}
	goto L21;
	L20:
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
	return;	L21:
	t47=heap[int(t51)];
	/* compilacion de accesso */
	t54=P+6;
	t53=stack[int(t54)];
	/* fin de la compilacion de acceso */
	
	if t47 > t53 {goto L17;}
	goto L18;
	/* fin de la expression relacional */
	
	/* finalizo la expression logica */
	
	L17:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t56=P+5;
	t55=stack[int(t56)];
	/* fin de la compilacion de acceso */
	
	t57=t55-1;
	/* fin de compilacion de variable */
	t58=P+5;
	stack[int(t58)]=t57;
	
	goto L16;
	L18:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t60=P+4;
	t59=stack[int(t60)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t62=P+5;
	t61=stack[int(t62)];
	/* fin de la compilacion de acceso */
	
	if t59 < t61 {goto L22;}
	goto L23;
	/* fin de la expression relacional */
	
	L22:
	/* compilacion de valor de variable */
	/* compilacion de acceso arreglos */
	t64=P+1;
	t63=stack[int(t64)];
	/* compilacion de accesso */
	t66=P+4;
	t65=stack[int(t66)];
	/* fin de la compilacion de acceso */
	
	t68=heap[int(t63)];
	t67=t65+t63;
	if t68 < t65 {goto L24;}
	goto L25;
	L24:
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
	return;	L25:
	t63=heap[int(t67)];
	/* fin de compilacion de variable */
	t69=P+7;
	stack[int(t69)]=t63;
	
	/* compilacion de acceso arreglos */
	t71=P+1;
	t70=stack[int(t71)];
	/* compilacion de accesso */
	t73=P+5;
	t72=stack[int(t73)];
	/* fin de la compilacion de acceso */
	
	t75=heap[int(t70)];
	t74=t72+t70;
	if t75 < t72 {goto L26;}
	goto L27;
	L26:
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
	return;	L27:
	t70=heap[int(t74)];
	/* cambiando el valor de arreglo */
	t77=P+1;
	t76=stack[int(t77)];
	/* compilacion de accesso */
	t80=P+4;
	t79=stack[int(t80)];
	/* fin de la compilacion de acceso */
	
	t82=heap[int(t76)];
	t81=t79+t76;
	if t82 < t79 {goto L28;}
	goto L29;
	L28:
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
	return;	L29:
	t76=heap[int(t81)];
	t78=t81;
	heap[int(t78)]=t70;
	/* compilacion de accesso */
	t84=P+7;
	t83=stack[int(t84)];
	/* fin de la compilacion de acceso */
	
	/* cambiando el valor de arreglo */
	t86=P+1;
	t85=stack[int(t86)];
	/* compilacion de accesso */
	t89=P+5;
	t88=stack[int(t89)];
	/* fin de la compilacion de acceso */
	
	t91=heap[int(t85)];
	t90=t88+t85;
	if t91 < t88 {goto L30;}
	goto L31;
	L30:
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
	return;	L31:
	t85=heap[int(t90)];
	t87=t90;
	heap[int(t87)]=t83;
	L23:
	goto L7;
	L9:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t93=P+5;
	t92=stack[int(t93)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t95=P+4;
	t94=stack[int(t95)];
	/* fin de la compilacion de acceso */
	
	if t92 < t94 {goto L32;}
	goto L33;
	/* fin de la expression relacional */
	
	L32:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t97=P+5;
	t96=stack[int(t97)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t98=P+7;
	stack[int(t98)]=t96;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t100=P+4;
	t99=stack[int(t100)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t101=P+5;
	stack[int(t101)]=t99;
	
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t103=P+7;
	t102=stack[int(t103)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t104=P+4;
	stack[int(t104)]=t102;
	
	L33:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t106=P+1;
	t105=stack[int(t106)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t108=P+2;
	t107=stack[int(t108)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t110=P+4;
	t109=stack[int(t110)];
	/* fin de la compilacion de acceso */
	
	t111=P+9;
	stack[int(t111)]=t105;
	t111=t111+1;
	stack[int(t111)]=t107;
	t111=t111+1;
	stack[int(t111)]=t109;
	P=P+8;
	quicksort();
	t111=stack[int(P)];
	P=P-8;
	/* fin de compilacion de variable */
	t112=P+8;
	stack[int(t112)]=t111;
	
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t113=P+9;
	stack[int(t113)]=0;
	
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t115=P+4;
	t114=stack[int(t115)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t117=P+2;
	t116=stack[int(t117)];
	/* fin de la compilacion de acceso */
	
	if t114 == t116 {goto L34;}
	goto L35;
	/* fin de la expression relacional */
	
	L34:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t119=P+4;
	t118=stack[int(t119)];
	/* fin de la compilacion de acceso */
	
	t120=t118+1;
	/* fin de compilacion de variable */
	t121=P+9;
	stack[int(t121)]=t120;
	
	goto L36;
	L35:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t123=P+4;
	t122=stack[int(t123)];
	/* fin de la compilacion de acceso */
	
	/* fin de compilacion de variable */
	t124=P+9;
	stack[int(t124)]=t122;
	
	L36:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t126=P+1;
	t125=stack[int(t126)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t128=P+9;
	t127=stack[int(t128)];
	/* fin de la compilacion de acceso */
	
	/* compilacion de accesso */
	t130=P+3;
	t129=stack[int(t130)];
	/* fin de la compilacion de acceso */
	
	t131=P+11;
	stack[int(t131)]=t125;
	t131=t131+1;
	stack[int(t131)]=t127;
	t131=t131+1;
	stack[int(t131)]=t129;
	P=P+10;
	quicksort();
	t131=stack[int(P)];
	P=P-10;
	/* fin de compilacion de variable */
	t132=P+8;
	stack[int(t132)]=t131;
	
	L0:
	return;
}

func main(){
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	stack[int(0)]=0;
	
	/* compilacion de valor de variable */
	t133=H;
	heap[int(H)]=13;
	H=H+1;
	heap[int(H)]=12;
	H=H+1;
	heap[int(H)]=9;
	H=H+1;
	heap[int(H)]=4;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=56;
	H=H+1;
	heap[int(H)]=34;
	H=H+1;
	heap[int(H)]=78;
	H=H+1;
	heap[int(H)]=22;
	H=H+1;
	heap[int(H)]=1;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	heap[int(H)]=13;
	H=H+1;
	heap[int(H)]=120;
	H=H+1;
	t133=t133+0.12837;
	t134=t133;
	t135=7*3;
	t140=P+1;
	t140=t140+1;
	stack[int(t140)]=9874;
	t140=t140+1;
	stack[int(t140)]=0;
	P=P+1;
	potencia();
	t141=stack[int(P)];
	P=P-1;
	if 0 != 0 {goto L39;}
	t141=1;
	L39:
	t142=820*10;
	t143=8*0;
	t144=t143+8;
	t145=H;
	heap[int(H)]=16;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=t135;
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
	heap[int(H)]=t141;
	H=H+1;
	heap[int(H)]=44;
	H=H+1;
	heap[int(H)]=3;
	H=H+1;
	heap[int(H)]=t142;
	H=H+1;
	heap[int(H)]=11;
	H=H+1;
	heap[int(H)]=t144;
	H=H+1;
	heap[int(H)]=10;
	H=H+1;
	t145=t145+0.12837;
	t146=t145;
	t147=H;
	heap[int(H)]=2;
	H=H+1;
	heap[int(H)]=t133;
	H=H+1;
	heap[int(H)]=t145;
	H=H+1;
	t147=t147+0.12837;
	/* fin de compilacion de variable */
	stack[int(1)]=t147;
	
	t148=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=81;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=83;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t148=t148+0.12837;
	t152=P+2;
	t152=t152+1;
	stack[int(t152)]=t148;
	P=P+2;
	print_string();
	t153=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t154=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=86;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=81;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t154=t154+0.12837;
	t155=P+2;
	t155=t155+1;
	stack[int(t155)]=t154;
	P=P+2;
	print_string();
	t156=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion length */
	/* compilacion de acceso arreglos */
	t157=stack[int(1)];
	t159=heap[int(t157)];
	t158=1+t157;
	if t159 < 1 {goto L42;}
	goto L43;
	L42:
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
	return;	L43:
	t157=heap[int(t158)];
	t160=t157;
	t160=heap[int(t160)];
	t161=1;
	L44:
	if t161 > t160 {goto L45;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t162=P+2;
	stack[int(t162)]=t161;
	
	t161=t161+1;
	/* compilacion de acceso arreglos */
	t163=stack[int(1)];
	t165=heap[int(t163)];
	t164=1+t163;
	if t165 < 1 {goto L46;}
	goto L47;
	L46:
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
	return;	L47:
	t163=heap[int(t164)];
	/* compilacion de accesso */
	t167=P+2;
	t166=stack[int(t167)];
	/* fin de la compilacion de acceso */
	
	t169=heap[int(t163)];
	t168=t166+t163;
	if t169 < t166 {goto L48;}
	goto L49;
	L48:
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
	return;	L49:
	t163=heap[int(t168)];
	fmt.Printf("%d", int(t163));
	fmt.Printf("%c", int(32));
	t170=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t170=t170+0.12837;
	t171=P+3;
	t171=t171+1;
	stack[int(t171)]=t170;
	P=P+3;
	print_string();
	t172=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t173=P+2;
	stack[int(t173)]=t161;
	
	goto L44;
	L45:
	t174=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t174=t174+0.12837;
	t175=P+2;
	t175=t175+1;
	stack[int(t175)]=t174;
	P=P+2;
	print_string();
	t176=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t177=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t177=t177+0.12837;
	t178=P+2;
	t178=t178+1;
	stack[int(t178)]=t177;
	P=P+2;
	print_string();
	t179=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t180=stack[int(1)];
	t182=heap[int(t180)];
	t181=1+t180;
	if t182 < 1 {goto L50;}
	goto L51;
	L50:
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
	return;	L51:
	t180=heap[int(t181)];
	/* compilacion length */
	/* compilacion de acceso arreglos */
	t183=stack[int(1)];
	t185=heap[int(t183)];
	t184=1+t183;
	if t185 < 1 {goto L52;}
	goto L53;
	L52:
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
	return;	L53:
	t183=heap[int(t184)];
	t186=t183;
	t186=heap[int(t186)];
	t187=P+3;
	stack[int(t187)]=t180;
	t187=t187+1;
	stack[int(t187)]=1;
	t187=t187+1;
	stack[int(t187)]=t186;
	P=P+2;
	quicksort();
	t187=stack[int(P)];
	P=P-2;
	t188=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=86;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=112;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=81;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=83;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=58;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t188=t188+0.12837;
	t189=P+2;
	t189=t189+1;
	stack[int(t189)]=t188;
	P=P+2;
	print_string();
	t190=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion length */
	/* compilacion de acceso arreglos */
	t191=stack[int(1)];
	t193=heap[int(t191)];
	t192=1+t191;
	if t193 < 1 {goto L54;}
	goto L55;
	L54:
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
	return;	L55:
	t191=heap[int(t192)];
	t194=t191;
	t194=heap[int(t194)];
	t195=1;
	L56:
	if t195 > t194 {goto L57;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t196=P+2;
	stack[int(t196)]=t195;
	
	t195=t195+1;
	/* compilacion de acceso arreglos */
	t197=stack[int(1)];
	t199=heap[int(t197)];
	t198=1+t197;
	if t199 < 1 {goto L58;}
	goto L59;
	L58:
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
	return;	L59:
	t197=heap[int(t198)];
	/* compilacion de accesso */
	t201=P+2;
	t200=stack[int(t201)];
	/* fin de la compilacion de acceso */
	
	t203=heap[int(t197)];
	t202=t200+t197;
	if t203 < t200 {goto L60;}
	goto L61;
	L60:
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
	return;	L61:
	t197=heap[int(t202)];
	fmt.Printf("%d", int(t197));
	fmt.Printf("%c", int(32));
	t204=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t204=t204+0.12837;
	t205=P+3;
	t205=t205+1;
	stack[int(t205)]=t204;
	P=P+3;
	print_string();
	t206=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t207=P+2;
	stack[int(t207)]=t195;
	
	goto L56;
	L57:
	t208=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t208=t208+0.12837;
	t209=P+2;
	t209=t209+1;
	stack[int(t209)]=t208;
	P=P+2;
	print_string();
	t210=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t211=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=86;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=110;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=81;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t211=t211+0.12837;
	t212=P+2;
	t212=t212+1;
	stack[int(t212)]=t211;
	P=P+2;
	print_string();
	t213=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion length */
	/* compilacion de acceso arreglos */
	t214=stack[int(1)];
	t216=heap[int(t214)];
	t215=2+t214;
	if t216 < 2 {goto L62;}
	goto L63;
	L62:
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
	return;	L63:
	t214=heap[int(t215)];
	t217=t214;
	t217=heap[int(t217)];
	t218=1;
	L64:
	if t218 > t217 {goto L65;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t219=P+2;
	stack[int(t219)]=t218;
	
	t218=t218+1;
	/* compilacion de acceso arreglos */
	t220=stack[int(1)];
	t222=heap[int(t220)];
	t221=2+t220;
	if t222 < 2 {goto L66;}
	goto L67;
	L66:
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
	return;	L67:
	t220=heap[int(t221)];
	/* compilacion de accesso */
	t224=P+2;
	t223=stack[int(t224)];
	/* fin de la compilacion de acceso */
	
	t226=heap[int(t220)];
	t225=t223+t220;
	if t226 < t223 {goto L68;}
	goto L69;
	L68:
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
	return;	L69:
	t220=heap[int(t225)];
	fmt.Printf("%d", int(t220));
	fmt.Printf("%c", int(32));
	t227=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t227=t227+0.12837;
	t228=P+3;
	t228=t228+1;
	stack[int(t228)]=t227;
	P=P+3;
	print_string();
	t229=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t230=P+2;
	stack[int(t230)]=t218;
	
	goto L64;
	L65:
	t231=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t231=t231+0.12837;
	t232=P+2;
	t232=t232+1;
	stack[int(t232)]=t231;
	P=P+2;
	print_string();
	t233=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	t234=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=45;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t234=t234+0.12837;
	t235=P+2;
	t235=t235+1;
	stack[int(t235)]=t234;
	P=P+2;
	print_string();
	t236=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion de acceso arreglos */
	t237=stack[int(1)];
	t239=heap[int(t237)];
	t238=2+t237;
	if t239 < 2 {goto L70;}
	goto L71;
	L70:
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
	return;	L71:
	t237=heap[int(t238)];
	/* compilacion length */
	/* compilacion de acceso arreglos */
	t240=stack[int(1)];
	t242=heap[int(t240)];
	t241=2+t240;
	if t242 < 2 {goto L72;}
	goto L73;
	L72:
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
	return;	L73:
	t240=heap[int(t241)];
	t243=t240;
	t243=heap[int(t243)];
	t244=P+3;
	stack[int(t244)]=t237;
	t244=t244+1;
	stack[int(t244)]=1;
	t244=t244+1;
	stack[int(t244)]=t243;
	P=P+2;
	quicksort();
	t244=stack[int(P)];
	P=P-2;
	t245=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=86;
	H=H+1;
	heap[int(H)]=97;
	H=H+1;
	heap[int(H)]=108;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=112;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=115;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=100;
	H=H+1;
	heap[int(H)]=101;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=81;
	H=H+1;
	heap[int(H)]=117;
	H=H+1;
	heap[int(H)]=105;
	H=H+1;
	heap[int(H)]=99;
	H=H+1;
	heap[int(H)]=107;
	H=H+1;
	heap[int(H)]=83;
	H=H+1;
	heap[int(H)]=111;
	H=H+1;
	heap[int(H)]=114;
	H=H+1;
	heap[int(H)]=116;
	H=H+1;
	heap[int(H)]=58;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t245=t245+0.12837;
	t246=P+2;
	t246=t246+1;
	stack[int(t246)]=t245;
	P=P+2;
	print_string();
	t247=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));
	/* compilacion length */
	/* compilacion de acceso arreglos */
	t248=stack[int(1)];
	t250=heap[int(t248)];
	t249=2+t248;
	if t250 < 2 {goto L74;}
	goto L75;
	L74:
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
	return;	L75:
	t248=heap[int(t249)];
	t251=t248;
	t251=heap[int(t251)];
	t252=1;
	L76:
	if t252 > t251 {goto L77;}
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t253=P+2;
	stack[int(t253)]=t252;
	
	t252=t252+1;
	/* compilacion de acceso arreglos */
	t254=stack[int(1)];
	t256=heap[int(t254)];
	t255=2+t254;
	if t256 < 2 {goto L78;}
	goto L79;
	L78:
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
	return;	L79:
	t254=heap[int(t255)];
	/* compilacion de accesso */
	t258=P+2;
	t257=stack[int(t258)];
	/* fin de la compilacion de acceso */
	
	t260=heap[int(t254)];
	t259=t257+t254;
	if t260 < t257 {goto L80;}
	goto L81;
	L80:
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
	return;	L81:
	t254=heap[int(t259)];
	fmt.Printf("%d", int(t254));
	fmt.Printf("%c", int(32));
	t261=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=32;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t261=t261+0.12837;
	t262=P+3;
	t262=t262+1;
	stack[int(t262)]=t261;
	P=P+3;
	print_string();
	t263=stack[int(P)];
	P=P-3;
	fmt.Printf("%c", int(32));
	/* compilacion de valor de variable */
	/* fin de compilacion de variable */
	t264=P+2;
	stack[int(t264)]=t252;
	
	goto L76;
	L77:
	t265=H;
	heap[int(H)]=0;
	H=H+1;
	heap[int(H)]=-1;
	H=H+1;
	t265=t265+0.12837;
	t266=P+2;
	t266=t266+1;
	stack[int(t266)]=t265;
	P=P+2;
	print_string();
	t267=stack[int(P)];
	P=P-2;
	fmt.Printf("%c", int(32));
	fmt.Printf("%c", int(10));

}