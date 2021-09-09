
# parsetab.py
# This file is automatically generated. Do not edit.
# pylint: disable=W,C,R
_tabversion = '3.10'

_lr_method = 'LALR'

_lr_signature = 'leftORleftANDleftIGUALQUENIGUALQUEleftMAYQUEMENQUEMAYIGUALQUEMENIGUALQUEleftMASMENOSleftPORDIVIDIDOMODULOleftPOTENCIArightUMENOSAND BOOL CADENA CHAR COMA CONCAT CORCHETEDER CORCHETEIZQ COS DECIMAL DIVIDIDO DOSP ELSE ENTERO FALSE FLOAT GLOBAL ID IF IGUAL IGUALQUE INT LLAVDER LLAVIZQ LOCAL LOG LOG10 LOWERCASE MAS MAYIGUALQUE MAYQUE MENIGUALQUE MENOS MENQUE MODULO NIGUALQUE NOT NOTHING NUMERO OR PARDER PARIZQ PARSE POR POTENCIA PRINT PRINTLN PTCOMA PUNTO SIN SQRT STRING TAN TOFLOAT TOSTRING TRUE TRUNC TYPEOF UPPERCASE WHILEinit            : instruccionesinstrucciones    : instrucciones instruccion\n                        | instruccioninstruccion      : print_instr\n                        | println_instr\n                        | definicion_instr\n                        | asignacion_instr\n                        | definicion_asignacion_instrexpression       : MENOS expression %prec UMENOS\n                        | NOT expression %prec UMENOS\n                        | expression MAS expression\n                        | expression MENOS expression\n                        | expression POR expression\n                        | expression DIVIDIDO expression\n                        | expression POTENCIA expression\n                        | expression MODULO expression\n                        | expression MAYQUE expression\n                        | expression MENQUE expression\n                        | expression MENIGUALQUE expression\n                        | expression MAYIGUALQUE expression\n                        | expression IGUALQUE expression\n                        | expression NIGUALQUE expression\n                        | expression OR expression\n                        | expression AND expression\n                        | final_expressionfinal_expression     : PARIZQ expression PARDER\n                            | DECIMAL\n                            | ENTERO\n                            | CADENA\n                            | ID\n                            | TRUE\n                            | FALSE\n                            | call_func\n                            | nativascall_func        : ID PARIZQ PARDER\n                        | ID PARIZQ exp_list PARDERnativas          : LOG PARIZQ ENTERO COMA expression PARDER\n                        | LOG10 PARIZQ expression PARDER\n                        | COS PARIZQ expression PARDER\n                        | SIN PARIZQ expression PARDER\n                        | TAN PARIZQ expression PARDER\n                        | SQRT PARIZQ expression PARDER\n                        | UPPERCASE PARIZQ expression PARDER\n                        | LOWERCASE PARIZQ expression PARDER\n                        | TOSTRING PARIZQ expression PARDER\n                        | TOFLOAT PARIZQ expression PARDER\n                        | TRUNC PARIZQ INT COMA expression PARDER\n                        | TYPEOF PARIZQ expression PARDER\n                        | PARSE PARIZQ tipo COMA expression PARDER\n                        exp_list         : exp_list COMA expression\n                        | expressionprint_instr    : PRINT PARIZQ expression PARDER PTCOMAprintln_instr  : PRINTLN PARIZQ expression PARDER PTCOMAtipo     : INT\n                | FLOAT\n                | BOOL\n                | CHAR\n                | STRING\n    definicion_instr   :  LOCAL ID PTCOMA\n                            | GLOBAL ID PTCOMAasignacion_instr   : ID IGUAL expression PTCOMA\n                            | LOCAL ID IGUAL expression PTCOMA\n                            | GLOBAL ID IGUAL expression PTCOMAdefinicion_asignacion_instr  : ID IGUAL expression DOSP DOSP tipo PTCOMA\n                                    | LOCAL ID IGUAL expression DOSP DOSP tipo PTCOMA\n                                    | GLOBAL ID IGUAL expression DOSP DOSP tipo PTCOMA'
    
_lr_action_items = {'PRINT':([0,2,3,4,5,6,7,8,14,47,50,86,90,126,127,130,155,160,161,],[9,9,-3,-4,-5,-6,-7,-8,-2,-59,-60,-61,-52,-53,-62,-63,-64,-65,-66,]),'PRINTLN':([0,2,3,4,5,6,7,8,14,47,50,86,90,126,127,130,155,160,161,],[10,10,-3,-4,-5,-6,-7,-8,-2,-59,-60,-61,-52,-53,-62,-63,-64,-65,-66,]),'LOCAL':([0,2,3,4,5,6,7,8,14,47,50,86,90,126,127,130,155,160,161,],[11,11,-3,-4,-5,-6,-7,-8,-2,-59,-60,-61,-52,-53,-62,-63,-64,-65,-66,]),'GLOBAL':([0,2,3,4,5,6,7,8,14,47,50,86,90,126,127,130,155,160,161,],[13,13,-3,-4,-5,-6,-7,-8,-2,-59,-60,-61,-52,-53,-62,-63,-64,-65,-66,]),'ID':([0,2,3,4,5,6,7,8,11,13,14,15,16,18,20,22,23,47,48,50,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,86,90,126,127,130,133,134,144,146,155,160,161,],[12,12,-3,-4,-5,-6,-7,-8,17,19,-2,28,28,28,28,28,28,-59,28,-60,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,-61,-52,-53,-62,-63,28,28,28,28,-64,-65,-66,]),'$end':([1,2,3,4,5,6,7,8,14,47,50,86,90,126,127,130,155,160,161,],[0,-1,-3,-4,-5,-6,-7,-8,-2,-59,-60,-61,-52,-53,-62,-63,-64,-65,-66,]),'PARIZQ':([9,10,15,16,18,20,22,23,28,33,34,35,36,37,38,39,40,41,42,43,44,45,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[15,16,20,20,20,20,20,20,70,71,72,73,74,75,76,77,78,79,80,81,82,83,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,20,]),'IGUAL':([12,17,19,],[18,48,51,]),'MENOS':([15,16,18,20,21,22,23,24,25,26,27,28,29,30,31,32,46,48,49,51,52,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,72,73,74,75,76,77,78,79,80,82,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,150,151,152,153,157,158,159,],[22,22,22,22,55,22,22,-25,-27,-28,-29,-30,-31,-32,-33,-34,55,22,55,22,55,22,22,22,22,22,22,22,22,22,22,22,22,22,22,-9,-10,22,22,22,22,22,22,22,22,22,22,22,55,55,-26,-11,-12,-13,-14,-15,-16,55,55,55,55,55,55,55,55,-35,55,55,55,55,55,55,55,55,55,55,55,-36,22,22,-38,-39,-40,-41,-42,-43,-44,-45,-46,22,-48,22,55,55,55,55,-37,-47,-49,]),'NOT':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,23,]),'DECIMAL':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,25,]),'ENTERO':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,71,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,108,26,26,26,26,26,26,26,26,26,26,26,26,26,26,]),'CADENA':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,27,]),'TRUE':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,29,]),'FALSE':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,]),'LOG':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,33,]),'LOG10':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,34,]),'COS':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,35,]),'SIN':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,36,]),'TAN':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,37,]),'SQRT':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,38,]),'UPPERCASE':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,39,]),'LOWERCASE':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,40,]),'TOSTRING':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,41,]),'TOFLOAT':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,42,]),'TRUNC':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,43,]),'TYPEOF':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,44,]),'PARSE':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,45,]),'PTCOMA':([17,19,24,25,26,27,28,29,30,31,32,49,53,68,69,84,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,121,122,123,124,125,132,135,136,137,138,139,140,141,142,143,145,148,154,156,157,158,159,],[47,50,-25,-27,-28,-29,-30,-31,-32,-33,-34,86,90,-9,-10,126,127,130,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24,-35,-54,-55,-56,-57,-58,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,155,160,161,-37,-47,-49,]),'PARDER':([21,24,25,26,27,28,29,30,31,32,46,52,68,69,70,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[53,-25,-27,-28,-29,-30,-31,-32,-33,-34,84,89,-9,-10,105,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24,-35,132,-51,135,136,137,138,139,140,141,142,143,145,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,-50,157,158,159,-37,-47,-49,]),'MAS':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[54,-25,-27,-28,-29,-30,-31,-32,-33,-34,54,54,54,-9,-10,54,54,-26,-11,-12,-13,-14,-15,-16,54,54,54,54,54,54,54,54,-35,54,54,54,54,54,54,54,54,54,54,54,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,54,54,54,54,-37,-47,-49,]),'POR':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[56,-25,-27,-28,-29,-30,-31,-32,-33,-34,56,56,56,-9,-10,56,56,-26,56,56,-13,-14,-15,-16,56,56,56,56,56,56,56,56,-35,56,56,56,56,56,56,56,56,56,56,56,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,56,56,56,56,-37,-47,-49,]),'DIVIDIDO':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[57,-25,-27,-28,-29,-30,-31,-32,-33,-34,57,57,57,-9,-10,57,57,-26,57,57,-13,-14,-15,-16,57,57,57,57,57,57,57,57,-35,57,57,57,57,57,57,57,57,57,57,57,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,57,57,57,57,-37,-47,-49,]),'POTENCIA':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[58,-25,-27,-28,-29,-30,-31,-32,-33,-34,58,58,58,-9,-10,58,58,-26,58,58,58,58,-15,58,58,58,58,58,58,58,58,58,-35,58,58,58,58,58,58,58,58,58,58,58,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,58,58,58,58,-37,-47,-49,]),'MODULO':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[59,-25,-27,-28,-29,-30,-31,-32,-33,-34,59,59,59,-9,-10,59,59,-26,59,59,-13,-14,-15,-16,59,59,59,59,59,59,59,59,-35,59,59,59,59,59,59,59,59,59,59,59,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,59,59,59,59,-37,-47,-49,]),'MAYQUE':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[60,-25,-27,-28,-29,-30,-31,-32,-33,-34,60,60,60,-9,-10,60,60,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,60,60,60,60,-35,60,60,60,60,60,60,60,60,60,60,60,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,60,60,60,60,-37,-47,-49,]),'MENQUE':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[61,-25,-27,-28,-29,-30,-31,-32,-33,-34,61,61,61,-9,-10,61,61,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,61,61,61,61,-35,61,61,61,61,61,61,61,61,61,61,61,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,61,61,61,61,-37,-47,-49,]),'MENIGUALQUE':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[62,-25,-27,-28,-29,-30,-31,-32,-33,-34,62,62,62,-9,-10,62,62,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,62,62,62,62,-35,62,62,62,62,62,62,62,62,62,62,62,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,62,62,62,62,-37,-47,-49,]),'MAYIGUALQUE':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[63,-25,-27,-28,-29,-30,-31,-32,-33,-34,63,63,63,-9,-10,63,63,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,63,63,63,63,-35,63,63,63,63,63,63,63,63,63,63,63,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,63,63,63,63,-37,-47,-49,]),'IGUALQUE':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[64,-25,-27,-28,-29,-30,-31,-32,-33,-34,64,64,64,-9,-10,64,64,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,64,64,-35,64,64,64,64,64,64,64,64,64,64,64,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,64,64,64,64,-37,-47,-49,]),'NIGUALQUE':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[65,-25,-27,-28,-29,-30,-31,-32,-33,-34,65,65,65,-9,-10,65,65,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,65,65,-35,65,65,65,65,65,65,65,65,65,65,65,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,65,65,65,65,-37,-47,-49,]),'OR':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[66,-25,-27,-28,-29,-30,-31,-32,-33,-34,66,66,66,-9,-10,66,66,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24,-35,66,66,66,66,66,66,66,66,66,66,66,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,66,66,66,66,-37,-47,-49,]),'AND':([21,24,25,26,27,28,29,30,31,32,46,49,52,68,69,85,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,107,109,110,111,112,113,114,115,116,117,119,132,135,136,137,138,139,140,141,142,143,145,150,151,152,153,157,158,159,],[67,-25,-27,-28,-29,-30,-31,-32,-33,-34,67,67,67,-9,-10,67,67,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,67,-24,-35,67,67,67,67,67,67,67,67,67,67,67,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,67,67,67,67,-37,-47,-49,]),'DOSP':([24,25,26,27,28,29,30,31,32,49,68,69,85,87,88,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,128,131,132,135,136,137,138,139,140,141,142,143,145,157,158,159,],[-25,-27,-28,-29,-30,-31,-32,-33,-34,87,-9,-10,128,129,131,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24,-35,147,149,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,-37,-47,-49,]),'COMA':([24,25,26,27,28,29,30,31,32,68,69,89,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,118,120,121,122,123,124,125,132,135,136,137,138,139,140,141,142,143,145,150,157,158,159,],[-25,-27,-28,-29,-30,-31,-32,-33,-34,-9,-10,-26,-11,-12,-13,-14,-15,-16,-17,-18,-19,-20,-21,-22,-23,-24,-35,133,-51,134,144,146,-54,-55,-56,-57,-58,-36,-38,-39,-40,-41,-42,-43,-44,-45,-46,-48,-50,-37,-47,-49,]),'INT':([81,83,129,147,149,],[118,121,121,121,121,]),'FLOAT':([83,129,147,149,],[122,122,122,122,]),'BOOL':([83,129,147,149,],[123,123,123,123,]),'CHAR':([83,129,147,149,],[124,124,124,124,]),'STRING':([83,129,147,149,],[125,125,125,125,]),}

_lr_action = {}
for _k, _v in _lr_action_items.items():
   for _x,_y in zip(_v[0],_v[1]):
      if not _x in _lr_action:  _lr_action[_x] = {}
      _lr_action[_x][_k] = _y
del _lr_action_items

_lr_goto_items = {'init':([0,],[1,]),'instrucciones':([0,],[2,]),'instruccion':([0,2,],[3,14,]),'print_instr':([0,2,],[4,4,]),'println_instr':([0,2,],[5,5,]),'definicion_instr':([0,2,],[6,6,]),'asignacion_instr':([0,2,],[7,7,]),'definicion_asignacion_instr':([0,2,],[8,8,]),'expression':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[21,46,49,52,68,69,85,88,91,92,93,94,95,96,97,98,99,100,101,102,103,104,107,109,110,111,112,113,114,115,116,117,119,150,151,152,153,]),'final_expression':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,]),'call_func':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,31,]),'nativas':([15,16,18,20,22,23,48,51,54,55,56,57,58,59,60,61,62,63,64,65,66,67,70,72,73,74,75,76,77,78,79,80,82,133,134,144,146,],[32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,32,]),'exp_list':([70,],[106,]),'tipo':([83,129,147,149,],[120,148,154,156,]),}

_lr_goto = {}
for _k, _v in _lr_goto_items.items():
   for _x, _y in zip(_v[0], _v[1]):
       if not _x in _lr_goto: _lr_goto[_x] = {}
       _lr_goto[_x][_k] = _y
del _lr_goto_items
_lr_productions = [
  ("S' -> init","S'",1,None,None,None),
  ('init -> instrucciones','init',1,'p_init','sintactico.py',207),
  ('instrucciones -> instrucciones instruccion','instrucciones',2,'p_instrucciones_lista','sintactico.py',212),
  ('instrucciones -> instruccion','instrucciones',1,'p_instrucciones_lista','sintactico.py',213),
  ('instruccion -> print_instr','instruccion',1,'p_instruccion','sintactico.py',222),
  ('instruccion -> println_instr','instruccion',1,'p_instruccion','sintactico.py',223),
  ('instruccion -> definicion_instr','instruccion',1,'p_instruccion','sintactico.py',224),
  ('instruccion -> asignacion_instr','instruccion',1,'p_instruccion','sintactico.py',225),
  ('instruccion -> definicion_asignacion_instr','instruccion',1,'p_instruccion','sintactico.py',226),
  ('expression -> MENOS expression','expression',2,'p_expression','sintactico.py',231),
  ('expression -> NOT expression','expression',2,'p_expression','sintactico.py',232),
  ('expression -> expression MAS expression','expression',3,'p_expression','sintactico.py',233),
  ('expression -> expression MENOS expression','expression',3,'p_expression','sintactico.py',234),
  ('expression -> expression POR expression','expression',3,'p_expression','sintactico.py',235),
  ('expression -> expression DIVIDIDO expression','expression',3,'p_expression','sintactico.py',236),
  ('expression -> expression POTENCIA expression','expression',3,'p_expression','sintactico.py',237),
  ('expression -> expression MODULO expression','expression',3,'p_expression','sintactico.py',238),
  ('expression -> expression MAYQUE expression','expression',3,'p_expression','sintactico.py',239),
  ('expression -> expression MENQUE expression','expression',3,'p_expression','sintactico.py',240),
  ('expression -> expression MENIGUALQUE expression','expression',3,'p_expression','sintactico.py',241),
  ('expression -> expression MAYIGUALQUE expression','expression',3,'p_expression','sintactico.py',242),
  ('expression -> expression IGUALQUE expression','expression',3,'p_expression','sintactico.py',243),
  ('expression -> expression NIGUALQUE expression','expression',3,'p_expression','sintactico.py',244),
  ('expression -> expression OR expression','expression',3,'p_expression','sintactico.py',245),
  ('expression -> expression AND expression','expression',3,'p_expression','sintactico.py',246),
  ('expression -> final_expression','expression',1,'p_expression','sintactico.py',247),
  ('final_expression -> PARIZQ expression PARDER','final_expression',3,'p_final_expression','sintactico.py',304),
  ('final_expression -> DECIMAL','final_expression',1,'p_final_expression','sintactico.py',305),
  ('final_expression -> ENTERO','final_expression',1,'p_final_expression','sintactico.py',306),
  ('final_expression -> CADENA','final_expression',1,'p_final_expression','sintactico.py',307),
  ('final_expression -> ID','final_expression',1,'p_final_expression','sintactico.py',308),
  ('final_expression -> TRUE','final_expression',1,'p_final_expression','sintactico.py',309),
  ('final_expression -> FALSE','final_expression',1,'p_final_expression','sintactico.py',310),
  ('final_expression -> call_func','final_expression',1,'p_final_expression','sintactico.py',311),
  ('final_expression -> nativas','final_expression',1,'p_final_expression','sintactico.py',312),
  ('call_func -> ID PARIZQ PARDER','call_func',3,'p_call_func','sintactico.py',334),
  ('call_func -> ID PARIZQ exp_list PARDER','call_func',4,'p_call_func','sintactico.py',335),
  ('nativas -> LOG PARIZQ ENTERO COMA expression PARDER','nativas',6,'p_nativas','sintactico.py',343),
  ('nativas -> LOG10 PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',344),
  ('nativas -> COS PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',345),
  ('nativas -> SIN PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',346),
  ('nativas -> TAN PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',347),
  ('nativas -> SQRT PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',348),
  ('nativas -> UPPERCASE PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',349),
  ('nativas -> LOWERCASE PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',350),
  ('nativas -> TOSTRING PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',351),
  ('nativas -> TOFLOAT PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',352),
  ('nativas -> TRUNC PARIZQ INT COMA expression PARDER','nativas',6,'p_nativas','sintactico.py',353),
  ('nativas -> TYPEOF PARIZQ expression PARDER','nativas',4,'p_nativas','sintactico.py',354),
  ('nativas -> PARSE PARIZQ tipo COMA expression PARDER','nativas',6,'p_nativas','sintactico.py',355),
  ('exp_list -> exp_list COMA expression','exp_list',3,'p_exp_list','sintactico.py',382),
  ('exp_list -> expression','exp_list',1,'p_exp_list','sintactico.py',383),
  ('print_instr -> PRINT PARIZQ expression PARDER PTCOMA','print_instr',5,'p_print_instr','sintactico.py',387),
  ('println_instr -> PRINTLN PARIZQ expression PARDER PTCOMA','println_instr',5,'p_println_instr','sintactico.py',392),
  ('tipo -> INT','tipo',1,'p_tipo','sintactico.py',397),
  ('tipo -> FLOAT','tipo',1,'p_tipo','sintactico.py',398),
  ('tipo -> BOOL','tipo',1,'p_tipo','sintactico.py',399),
  ('tipo -> CHAR','tipo',1,'p_tipo','sintactico.py',400),
  ('tipo -> STRING','tipo',1,'p_tipo','sintactico.py',401),
  ('definicion_instr -> LOCAL ID PTCOMA','definicion_instr',3,'p_definicion_instr','sintactico.py',416),
  ('definicion_instr -> GLOBAL ID PTCOMA','definicion_instr',3,'p_definicion_instr','sintactico.py',417),
  ('asignacion_instr -> ID IGUAL expression PTCOMA','asignacion_instr',4,'p_asignacion_instr','sintactico.py',425),
  ('asignacion_instr -> LOCAL ID IGUAL expression PTCOMA','asignacion_instr',5,'p_asignacion_instr','sintactico.py',426),
  ('asignacion_instr -> GLOBAL ID IGUAL expression PTCOMA','asignacion_instr',5,'p_asignacion_instr','sintactico.py',427),
  ('definicion_asignacion_instr -> ID IGUAL expression DOSP DOSP tipo PTCOMA','definicion_asignacion_instr',7,'p_definicion_asignacion_instr','sintactico.py',438),
  ('definicion_asignacion_instr -> LOCAL ID IGUAL expression DOSP DOSP tipo PTCOMA','definicion_asignacion_instr',8,'p_definicion_asignacion_instr','sintactico.py',439),
  ('definicion_asignacion_instr -> GLOBAL ID IGUAL expression DOSP DOSP tipo PTCOMA','definicion_asignacion_instr',8,'p_definicion_asignacion_instr','sintactico.py',440),
]
