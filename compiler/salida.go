/*----HEADER----*/
package main

import (
	"fmt"
	"math"
)

var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43 float64
var P, H float64
var stack [30101999]float64
var heap [30101999]float64

/*-----FUNCS-----*/
func fibonacci() {
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t1 = P + 1
	t0 = stack[int(t1)]
	/* fin de la compilacion de acceso */

	if t0 > 1 {
		goto L1
	}
	goto L2
	/* fin de la expression relacional */

L1:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t3 = P + 1
	t2 = stack[int(t3)]
	/* fin de la compilacion de acceso */

	t4 = t2 - 1
	t5 = P + 3
	stack[int(t5)] = t4
	P = P + 2
	fibonacci()
	t5 = stack[int(P)]
	P = P - 2
	/* fin de compilacion de variable */
	t6 = P + 2
	stack[int(t6)] = t5

	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t8 = P + 1
	t7 = stack[int(t8)]
	/* fin de la compilacion de acceso */

	t9 = t7 - 2
	t10 = P + 4
	stack[int(t10)] = t9
	P = P + 3
	fibonacci()
	t10 = stack[int(P)]
	P = P - 3
	/* fin de compilacion de variable */
	t11 = P + 3
	stack[int(t11)] = t10

	/* compilacion de accesso */
	t13 = P + 2
	t12 = stack[int(t13)]
	/* fin de la compilacion de acceso */

	/* compilacion de accesso */
	t15 = P + 3
	t14 = stack[int(t15)]
	/* fin de la compilacion de acceso */

	t16 = t12 + t14
	stack[int(P)] = t16
	goto L0
	goto L3
L2:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t18 = P + 1
	t17 = stack[int(t18)]
	/* fin de la compilacion de acceso */

	if t17 == 1 {
		goto L4
	}
	goto L5
	/* fin de la expression relacional */

L4:
	stack[int(P)] = 1
	goto L0
	goto L6
L5:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t20 = P + 1
	t19 = stack[int(t20)]
	/* fin de la compilacion de acceso */

	if t19 == 0 {
		goto L7
	}
	goto L8
	/* fin de la expression relacional */

L7:
	stack[int(P)] = 0
	goto L0
L8:
L6:
L3:
L0:
	return
}
func impares() {
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t22 = P + 1
	t21 = stack[int(t22)]
	/* fin de la compilacion de acceso */

	if t21 == 1 {
		goto L10
	}
	goto L11
	/* fin de la expression relacional */

L10:
	stack[int(P)] = 1
	goto L9
	goto L12
L11:
	/* iniciando el if */
	/* inicio de expression realcional */
	/* compilacion de accesso */
	t24 = P + 1
	t23 = stack[int(t24)]
	/* fin de la compilacion de acceso */

	t25 = math.Mod(t23, 2)
	if t25 == 1 {
		goto L13
	}
	goto L14
	/* fin de la expression relacional */

L13:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t27 = P + 1
	t26 = stack[int(t27)]
	/* fin de la compilacion de acceso */

	t28 = t26 - 2
	t29 = P + 3
	stack[int(t29)] = t28
	P = P + 2
	impares()
	t29 = stack[int(P)]
	P = P - 2
	/* fin de compilacion de variable */
	t30 = P + 2
	stack[int(t30)] = t29

	/* compilacion de accesso */
	t32 = P + 1
	t31 = stack[int(t32)]
	/* fin de la compilacion de acceso */

	/* compilacion de accesso */
	t34 = P + 2
	t33 = stack[int(t34)]
	/* fin de la compilacion de acceso */

	t35 = t31 * t33
	stack[int(P)] = t35
	goto L9
L14:
L12:
	/* compilacion de valor de variable */
	/* compilacion de accesso */
	t37 = P + 1
	t36 = stack[int(t37)]
	/* fin de la compilacion de acceso */

	t38 = t36 - 1
	t39 = P + 4
	stack[int(t39)] = t38
	P = P + 3
	impares()
	t39 = stack[int(P)]
	P = P - 3
	/* fin de compilacion de variable */
	t40 = P + 3
	stack[int(t40)] = t39

	/* compilacion de accesso */
	t42 = P + 3
	t41 = stack[int(t42)]
	/* fin de la compilacion de acceso */

	stack[int(P)] = t41
	goto L9
L9:
	return
}

func main() {
	t43 = P + 1
	stack[int(t43)] = 7
	P = P + 0
	impares()
	t43 = stack[int(P)]
	P = P - 0
	fmt.Printf("%d", int(t43))
	fmt.Printf("%c", int(32))
	fmt.Printf("%c", int(10))

}
