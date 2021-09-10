
# parsetab.py
# This file is automatically generated. Do not edit.
# pylint: disable=W,C,R
_tabversion = '3.10'

_lr_method = 'LALR'

_lr_signature = 'leftORleftANDleftIGUALQUENIGUALQUEleftMAYQUEMENQUEMAYIGUALQUEMENIGUALQUEleftMASMENOSleftPORDIVIDIDOMODULOleftPOTENCIArightUMENOSAND BOOL BREAK CADENA CHAR COMA CONCAT CONTINUE CORCHETEDER CORCHETEIZQ COS DECIMAL DIVIDIDO DOSP ELSE END ENTERO FALSE FLOAT FUNCTION GLOBAL ID IF IGUAL IGUALQUE INT LLAVDER LLAVIZQ LOCAL LOG LOG10 LOWERCASE MAS MAYIGUALQUE MAYQUE MENIGUALQUE MENOS MENQUE MODULO NIGUALQUE NOT NOTHING NUMERO OR PARDER PARIZQ PARSE POR POTENCIA PRINT PRINTLN PTCOMA PUNTO RETURN SIN SQRT STRING TAN TOFLOAT TOSTRING TRUE TRUNC TYPEOF UPPERCASE WHILEinit            : instruccionesinstrucciones    : instrucciones instruccion\n                        | instruccioninstruccion      : print_instr PTCOMA\n                        | println_instr PTCOMA\n                        | definicion_instr PTCOMA\n                        | asignacion_instr PTCOMA\n                        | definicion_asignacion_instr PTCOMA\n                        | call_function PTCOMA\n                        | declare_function PTCOMA\n                        | return_state PTCOMA\n                        | break_state PTCOMA\n                        | continue_state PTCOMAexpression       : MENOS expression %prec UMENOS\n                        | NOT expression %prec UMENOS\n                        | expression MAS expression\n                        | expression MENOS expression\n                        | expression POR expression\n                        | expression DIVIDIDO expression\n                        | expression POTENCIA expression\n                        | expression MODULO expression\n                        | expression MAYQUE expression\n                        | expression MENQUE expression\n                        | expression MENIGUALQUE expression\n                        | expression MAYIGUALQUE expression\n                        | expression IGUALQUE expression\n                        | expression NIGUALQUE expression\n                        | expression OR expression\n                        | expression AND expression\n                        | final_expressionfinal_expression     : PARIZQ expression PARDER\n                            | DECIMAL\n                            | ENTERO\n                            | CADENA\n                            | ID\n                            | TRUE\n                            | FALSE\n                            | call_func\n                            | nativascall_func        : ID PARIZQ PARDER\n                        | ID PARIZQ exp_list PARDERnativas          : LOG PARIZQ ENTERO COMA expression PARDER\n                        | LOG10 PARIZQ expression PARDER\n                        | COS PARIZQ expression PARDER\n                        | SIN PARIZQ expression PARDER\n                        | TAN PARIZQ expression PARDER\n                        | SQRT PARIZQ expression PARDER\n                        | UPPERCASE PARIZQ expression PARDER\n                        | LOWERCASE PARIZQ expression PARDER\n                        | TOSTRING PARIZQ expression PARDER\n                        | TOFLOAT PARIZQ expression PARDER\n                        | TRUNC PARIZQ INT COMA expression PARDER\n                        | TYPEOF PARIZQ expression PARDER\n                        | PARSE PARIZQ tipo COMA expression PARDER\n                        print_instr    : PRINT PARIZQ expression PARDERprintln_instr  : PRINTLN PARIZQ expression PARDERtipo     : INT\n                | FLOAT\n                | BOOL\n                | CHAR\n                | STRING\n    definicion_instr   :  LOCAL ID\n                            | GLOBAL IDasignacion_instr   : ID IGUAL expression\n                            | LOCAL ID IGUAL expression\n                            | GLOBAL ID IGUAL expressiondefinicion_asignacion_instr  : ID IGUAL expression DOSP DOSP tipo\n                                    | LOCAL ID IGUAL expression DOSP DOSP tipo\n                                    | GLOBAL ID IGUAL expression DOSP DOSP tipocall_function    : ID PARIZQ PARDER\n                        | ID PARIZQ exp_list PARDERexp_list         : exp_list COMA expression\n                        | expressionstatement        : instruccionesdeclare_function     : FUNCTION ID PARIZQ PARDER statement END\n                            | FUNCTION ID PARIZQ dec_params PARDER statement ENDdec_params :    dec_params COMA ID\n                    | IDbreak_state      : BREAKcontinue_state      : CONTINUEreturn_state     : RETURN\n                        | RETURN expression'
    
_lr_action_items = {'PRINT':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[14,14,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,14,14,14,]),'PRINTLN':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[15,15,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,15,15,15,]),'LOCAL':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[16,16,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,16,16,16,]),'GLOBAL':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[18,18,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,18,18,18,]),'ID':([0,2,3,16,18,19,20,23,24,25,26,27,28,29,30,31,32,33,34,35,37,38,42,43,45,69,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,115,157,158,159,161,171,173,],[17,17,-3,36,39,40,49,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,49,49,49,49,49,49,49,49,49,114,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,17,17,17,179,49,49,49,]),'FUNCTION':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[19,19,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,19,19,19,]),'RETURN':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[20,20,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,20,20,20,]),'BREAK':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[21,21,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,21,21,21,]),'CONTINUE':([0,2,3,23,24,25,26,27,28,29,30,31,32,33,115,157,158,],[22,22,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,22,22,22,]),'$end':([1,2,3,23,24,25,26,27,28,29,30,31,32,33,],[0,-1,-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,]),'END':([3,23,24,25,26,27,28,29,30,31,32,33,156,157,178,],[-3,-2,-4,-5,-6,-7,-8,-9,-10,-11,-12,-13,177,-74,185,]),'PTCOMA':([4,5,6,7,8,9,10,11,12,13,20,21,22,36,39,41,44,46,47,48,49,50,51,52,53,70,71,90,91,107,108,109,111,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,147,148,149,150,151,160,162,163,164,165,166,167,168,169,170,172,175,177,183,184,185,186,187,188,],[24,25,26,27,28,29,30,31,32,33,-81,-79,-80,-62,-63,-82,-30,-32,-33,-34,-35,-36,-37,-38,-39,-64,-70,-14,-15,-55,-56,-65,-71,-66,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,-28,-29,-31,-40,-57,-58,-59,-60,-61,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,-67,-75,-68,-69,-76,-42,-52,-54,]),'PARIZQ':([14,15,17,20,34,35,37,38,40,42,43,45,49,54,55,56,57,58,59,60,61,62,63,64,65,66,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[34,35,38,45,45,45,45,45,75,45,45,45,93,94,95,96,97,98,99,100,101,102,103,104,105,106,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,]),'IGUAL':([17,36,39,],[37,69,74,]),'MENOS':([20,34,35,37,38,41,42,43,44,45,46,47,48,49,50,51,52,53,67,68,69,70,73,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,95,96,97,98,99,100,101,102,103,105,109,112,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,161,162,163,164,165,166,167,168,169,170,171,172,173,180,181,182,186,187,188,],[42,42,42,42,42,77,42,42,-30,42,-32,-33,-34,-35,-36,-37,-38,-39,77,77,42,77,77,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,-14,-15,77,42,42,42,42,42,42,42,42,42,42,42,77,42,77,-16,-17,-18,-19,-20,-21,77,77,77,77,77,77,77,77,-31,-40,77,77,77,77,77,77,77,77,77,77,77,-41,42,-43,-44,-45,-46,-47,-48,-49,-50,-51,42,-53,42,77,77,77,-42,-52,-54,]),'NOT':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,]),'DECIMAL':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,46,]),'ENTERO':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,94,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,134,47,47,47,47,47,47,47,47,47,47,47,47,47,47,]),'CADENA':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,48,]),'TRUE':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,]),'FALSE':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,]),'LOG':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,54,]),'LOG10':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,55,]),'COS':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,56,]),'SIN':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,]),'TAN':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,58,]),'SQRT':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,59,]),'UPPERCASE':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,60,]),'LOWERCASE':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,61,]),'TOSTRING':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,62,]),'TOFLOAT':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,63,]),'TRUNC':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,]),'TYPEOF':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,65,]),'PARSE':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,]),'PARDER':([38,44,46,47,48,49,50,51,52,53,67,68,72,73,75,90,91,92,93,114,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,179,180,181,182,186,187,188,],[71,-30,-32,-33,-34,-35,-36,-37,-38,-39,107,108,111,-73,115,-14,-15,131,132,-78,158,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,-28,-29,-31,-40,160,162,163,164,165,166,167,168,169,170,172,-72,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,-77,186,187,188,-42,-52,-54,]),'MAS':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[76,-30,-32,-33,-34,-35,-36,-37,-38,-39,76,76,76,76,-14,-15,76,76,76,-16,-17,-18,-19,-20,-21,76,76,76,76,76,76,76,76,-31,-40,76,76,76,76,76,76,76,76,76,76,76,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,76,76,76,-42,-52,-54,]),'POR':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[78,-30,-32,-33,-34,-35,-36,-37,-38,-39,78,78,78,78,-14,-15,78,78,78,78,78,-18,-19,-20,-21,78,78,78,78,78,78,78,78,-31,-40,78,78,78,78,78,78,78,78,78,78,78,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,78,78,78,-42,-52,-54,]),'DIVIDIDO':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[79,-30,-32,-33,-34,-35,-36,-37,-38,-39,79,79,79,79,-14,-15,79,79,79,79,79,-18,-19,-20,-21,79,79,79,79,79,79,79,79,-31,-40,79,79,79,79,79,79,79,79,79,79,79,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,79,79,79,-42,-52,-54,]),'POTENCIA':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[80,-30,-32,-33,-34,-35,-36,-37,-38,-39,80,80,80,80,-14,-15,80,80,80,80,80,80,80,-20,80,80,80,80,80,80,80,80,80,-31,-40,80,80,80,80,80,80,80,80,80,80,80,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,80,80,80,-42,-52,-54,]),'MODULO':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[81,-30,-32,-33,-34,-35,-36,-37,-38,-39,81,81,81,81,-14,-15,81,81,81,81,81,-18,-19,-20,-21,81,81,81,81,81,81,81,81,-31,-40,81,81,81,81,81,81,81,81,81,81,81,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,81,81,81,-42,-52,-54,]),'MAYQUE':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[82,-30,-32,-33,-34,-35,-36,-37,-38,-39,82,82,82,82,-14,-15,82,82,82,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,82,82,82,82,-31,-40,82,82,82,82,82,82,82,82,82,82,82,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,82,82,82,-42,-52,-54,]),'MENQUE':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[83,-30,-32,-33,-34,-35,-36,-37,-38,-39,83,83,83,83,-14,-15,83,83,83,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,83,83,83,83,-31,-40,83,83,83,83,83,83,83,83,83,83,83,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,83,83,83,-42,-52,-54,]),'MENIGUALQUE':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[84,-30,-32,-33,-34,-35,-36,-37,-38,-39,84,84,84,84,-14,-15,84,84,84,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,84,84,84,84,-31,-40,84,84,84,84,84,84,84,84,84,84,84,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,84,84,84,-42,-52,-54,]),'MAYIGUALQUE':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[85,-30,-32,-33,-34,-35,-36,-37,-38,-39,85,85,85,85,-14,-15,85,85,85,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,85,85,85,85,-31,-40,85,85,85,85,85,85,85,85,85,85,85,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,85,85,85,-42,-52,-54,]),'IGUALQUE':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[86,-30,-32,-33,-34,-35,-36,-37,-38,-39,86,86,86,86,-14,-15,86,86,86,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,86,86,-31,-40,86,86,86,86,86,86,86,86,86,86,86,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,86,86,86,-42,-52,-54,]),'NIGUALQUE':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[87,-30,-32,-33,-34,-35,-36,-37,-38,-39,87,87,87,87,-14,-15,87,87,87,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,87,87,-31,-40,87,87,87,87,87,87,87,87,87,87,87,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,87,87,87,-42,-52,-54,]),'OR':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[88,-30,-32,-33,-34,-35,-36,-37,-38,-39,88,88,88,88,-14,-15,88,88,88,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,-28,-29,-31,-40,88,88,88,88,88,88,88,88,88,88,88,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,88,88,88,-42,-52,-54,]),'AND':([41,44,46,47,48,49,50,51,52,53,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,135,136,137,138,139,140,141,142,143,145,154,160,162,163,164,165,166,167,168,169,170,172,180,181,182,186,187,188,],[89,-30,-32,-33,-34,-35,-36,-37,-38,-39,89,89,89,89,-14,-15,89,89,89,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,89,-29,-31,-40,89,89,89,89,89,89,89,89,89,89,89,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,89,89,89,-42,-52,-54,]),'DOSP':([44,46,47,48,49,50,51,52,53,70,90,91,109,110,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,152,155,160,162,163,164,165,166,167,168,169,170,172,186,187,188,],[-30,-32,-33,-34,-35,-36,-37,-38,-39,110,-14,-15,152,153,155,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,-28,-29,-31,-40,174,176,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,-42,-52,-54,]),'COMA':([44,46,47,48,49,50,51,52,53,72,73,90,91,114,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,134,144,146,147,148,149,150,151,154,160,162,163,164,165,166,167,168,169,170,172,179,186,187,188,],[-30,-32,-33,-34,-35,-36,-37,-38,-39,112,-73,-14,-15,-78,159,-16,-17,-18,-19,-20,-21,-22,-23,-24,-25,-26,-27,-28,-29,-31,-40,112,161,171,173,-57,-58,-59,-60,-61,-72,-41,-43,-44,-45,-46,-47,-48,-49,-50,-51,-53,-77,-42,-52,-54,]),'INT':([104,106,153,174,176,],[144,147,147,147,147,]),'FLOAT':([106,153,174,176,],[148,148,148,148,]),'BOOL':([106,153,174,176,],[149,149,149,149,]),'CHAR':([106,153,174,176,],[150,150,150,150,]),'STRING':([106,153,174,176,],[151,151,151,151,]),}

_lr_action = {}
for _k, _v in _lr_action_items.items():
   for _x,_y in zip(_v[0],_v[1]):
      if not _x in _lr_action:  _lr_action[_x] = {}
      _lr_action[_x][_k] = _y
del _lr_action_items

_lr_goto_items = {'init':([0,],[1,]),'instrucciones':([0,115,158,],[2,157,157,]),'instruccion':([0,2,115,157,158,],[3,23,3,23,3,]),'print_instr':([0,2,115,157,158,],[4,4,4,4,4,]),'println_instr':([0,2,115,157,158,],[5,5,5,5,5,]),'definicion_instr':([0,2,115,157,158,],[6,6,6,6,6,]),'asignacion_instr':([0,2,115,157,158,],[7,7,7,7,7,]),'definicion_asignacion_instr':([0,2,115,157,158,],[8,8,8,8,8,]),'call_function':([0,2,115,157,158,],[9,9,9,9,9,]),'declare_function':([0,2,115,157,158,],[10,10,10,10,10,]),'return_state':([0,2,115,157,158,],[11,11,11,11,11,]),'break_state':([0,2,115,157,158,],[12,12,12,12,12,]),'continue_state':([0,2,115,157,158,],[13,13,13,13,13,]),'expression':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[41,67,68,70,73,90,91,92,109,113,117,118,119,120,121,122,123,124,125,126,127,128,129,130,73,135,136,137,138,139,140,141,142,143,145,154,180,181,182,]),'final_expression':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,]),'call_func':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,52,]),'nativas':([20,34,35,37,38,42,43,45,69,74,76,77,78,79,80,81,82,83,84,85,86,87,88,89,93,95,96,97,98,99,100,101,102,103,105,112,161,171,173,],[53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,53,]),'exp_list':([38,93,],[72,133,]),'dec_params':([75,],[116,]),'tipo':([106,153,174,176,],[146,175,183,184,]),'statement':([115,158,],[156,178,]),}

_lr_goto = {}
for _k, _v in _lr_goto_items.items():
   for _x, _y in zip(_v[0], _v[1]):
       if not _x in _lr_goto: _lr_goto[_x] = {}
       _lr_goto[_x][_k] = _y
del _lr_goto_items
_lr_productions = [
  ("S' -> init","S'",1,None,None,None),
  ('init -> instrucciones','init',1,'p_init','sintactico.py',221),
  ('instrucciones -> instrucciones instruccion','instrucciones',2,'p_instrucciones_lista','sintactico.py',226),
  ('instrucciones -> instruccion','instrucciones',1,'p_instrucciones_lista','sintactico.py',227),
  ('instruccion -> print_instr PTCOMA','instruccion',2,'p_instruccion','sintactico.py',236),
  ('instruccion -> println_instr PTCOMA','instruccion',2,'p_instruccion','sintactico.py',237),
  ('instruccion -> definicion_instr PTCOMA','instruccion',2,'p_instruccion','sintactico.py',238),
  ('instruccion -> asignacion_instr PTCOMA','instruccion',2,'p_instruccion','sintactico.py',239),
  ('instruccion -> definicion_asignacion_instr PTCOMA','instruccion',2,'p_instruccion','sintactico.py',240),
  ('instruccion -> call_function PTCOMA','instruccion',2,'p_instruccion','sintactico.py',241),
  ('instruccion -> declare_function PTCOMA','instruccion',2,'p_instruccion','sintactico.py',242),
  ('instruccion -> return_state PTCOMA','instruccion',2,'p_instruccion','sintactico.py',243),
  ('instruccion -> break_state PTCOMA','instruccion',2,'p_instruccion','sintactico.py',244),
  ('instruccion -> continue_state PTCOMA','instruccion',2,'p_instruccion','sintactico.py',245),
  ('expression -> MENOS expression','expression',2,'p_expression','sintactico.py',250),
  ('expression -> NOT expression','expression',2,'p_expression','sintactico.py',251),
  ('expression -> expression MAS expression','expression',3,'p_expression','sintactico.py',252),
  ('expression -> expression MENOS expression','expression',3,'p_expression','sintactico.py',253),
  ('expression -> expression POR expression','expression',3,'p_expression','sintactico.py',254),
  ('expression -> expression DIVIDIDO expression','expression',3,'p_expression','sintactico.py',255),
  ('expression -> expression POTENCIA expression','expression',3,'p_expression','sintactico.py',256),
  ('expression -> expression MODULO expression','expression',3,'p_expression','sintactico.py',257),
  ('expression -> expression MAYQUE expression','expression',3,'p_expression','sintactico.py',258),
  ('expression -> expression MENQUE expression','expression',3,'p_expression','sintactico.py',259),
  ('expression -> expression MENIGUALQUE expression','expression',3,'p_expression','sintactico.py',260),
  ('expression -> expression MAYIGUALQUE expression','expression',3,'p_expression','sintactico.py',261),
  ('expression -> expression IGUALQUE expression','expression',3,'p_expression','sintactico.py',262),
  ('expression -> expression NIGUALQUE expression','expression',3,'p_expression','sintactico.py',263),
  ('expression -> expression OR expression','expression',3,'p_expression','sintactico.py',264),
  ('expression -> expression AND expression','expression',3,'p_expression','sintactico.py',265),
  ('expression -> final_expression','expression',1,'p_expression','sintactico.py',266),
  ('final_expression -> PARIZQ expression PARDER','final_expression',3,'p_final_expression','sintactico.py',323),
  ('final_expression -> DECIMAL','final_expression',1,'p_final_expression','sintactico.py',324),
  ('final_expression -> ENTERO','final_expression',1,'p_final_expression','sintactico.py',325),
  ('final_expression -> CADENA','final_expression',1,'p_final_expression','sintactico.py',326),
  ('final_expression -> ID','final_expression',1,'p_final_expression','sintactico.py',327),
  ('final_expression -> TRUE','final_expression',1,'p_final_expression','sintactico.py',328),
  ('final_expression -> FALSE','final_expression',1,'p_final_expression','sintactico.py',329),
  ('final_expression -> call_func','final_expression',1,'p_final_expression','sintactico.py',330),
  ('final_expression -> nativas','final_expression',1,'p_final_expression','sintactico.py',331),
  ('call_func -> ID PARIZQ PARDER','call_func',3,'p_call_func','sintactico.py',353),
  ('call_func -> ID PARIZQ exp_list PARDER','call_func',4,'p_call_func','sintactico.py',354),
  ('nativas -> LOG PARIZQ ENTERO COMA expression PARDER','nativas',6,'p_nativas','sintactico.py',362),
  ('nativas -> LOG10 PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',363),
  ('nativas -> COS PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',364),
  ('nativas -> SIN PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',365),
  ('nativas -> TAN PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',366),
  ('nativas -> SQRT PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',367),
  ('nativas -> UPPERCASE PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',368),
  ('nativas -> LOWERCASE PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',369),
  ('nativas -> TOSTRING PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',370),
  ('nativas -> TOFLOAT PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',371),
  ('nativas -> TRUNC PARIZQ INT COMA expression PARDER','nativas',6,'p_nativas','sintactico.py',372),
  ('nativas -> TYPEOF PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',373),
  ('nativas -> PARSE PARIZQ tipo COMA expression PARDER','nativas',6,'p_nativas','sintactico.py',374),
  ('print_instr -> PRINT PARIZQ expression PARDER','print_instr',4,'p_print_instr','sintactico.py',405),
  ('println_instr -> PRINTLN PARIZQ expression PARDER','println_instr',4,'p_println_instr','sintactico.py',410),
  ('tipo -> INT','tipo',1,'p_tipo','sintactico.py',415),
  ('tipo -> FLOAT','tipo',1,'p_tipo','sintactico.py',416),
  ('tipo -> BOOL','tipo',1,'p_tipo','sintactico.py',417),
  ('tipo -> CHAR','tipo',1,'p_tipo','sintactico.py',418),
  ('tipo -> STRING','tipo',1,'p_tipo','sintactico.py',419),
  ('definicion_instr -> LOCAL ID','definicion_instr',2,'p_definicion_instr','sintactico.py',434),
  ('definicion_instr -> GLOBAL ID','definicion_instr',2,'p_definicion_instr','sintactico.py',435),
  ('asignacion_instr -> ID IGUAL expression','asignacion_instr',3,'p_asignacion_instr','sintactico.py',443),
  ('asignacion_instr -> LOCAL ID IGUAL expression','asignacion_instr',4,'p_asignacion_instr','sintactico.py',444),
  ('asignacion_instr -> GLOBAL ID IGUAL expression','asignacion_instr',4,'p_asignacion_instr','sintactico.py',445),
  ('definicion_asignacion_instr -> ID IGUAL expression DOSP DOSP tipo','definicion_asignacion_instr',6,'p_definicion_asignacion_instr','sintactico.py',456),
  ('definicion_asignacion_instr -> LOCAL ID IGUAL expression DOSP DOSP tipo','definicion_asignacion_instr',7,'p_definicion_asignacion_instr','sintactico.py',457),
  ('definicion_asignacion_instr -> GLOBAL ID IGUAL expression DOSP DOSP tipo','definicion_asignacion_instr',7,'p_definicion_asignacion_instr','sintactico.py',458),
  ('call_function -> ID PARIZQ PARDER','call_function',3,'p_call_function_instr','sintactico.py',469),
  ('call_function -> ID PARIZQ exp_list PARDER','call_function',4,'p_call_function_instr','sintactico.py',470),
  ('exp_list -> exp_list COMA expression','exp_list',3,'p_exp_list_instr','sintactico.py',478),
  ('exp_list -> expression','exp_list',1,'p_exp_list_instr','sintactico.py',479),
  ('statement -> instrucciones','statement',1,'p_statement','sintactico.py',488),
  ('declare_function -> FUNCTION ID PARIZQ PARDER statement END','declare_function',6,'p_declare_function','sintactico.py',493),
  ('declare_function -> FUNCTION ID PARIZQ dec_params PARDER statement END','declare_function',7,'p_declare_function','sintactico.py',494),
  ('dec_params -> dec_params COMA ID','dec_params',3,'p_dec_params','sintactico.py',502),
  ('dec_params -> ID','dec_params',1,'p_dec_params','sintactico.py',503),
  ('break_state -> BREAK','break_state',1,'p_break','sintactico.py',512),
  ('continue_state -> CONTINUE','continue_state',1,'p_continue','sintactico.py',517),
  ('return_state -> RETURN','return_state',1,'p_return','sintactico.py',522),
  ('return_state -> RETURN expression','return_state',2,'p_return','sintactico.py',523),
]
