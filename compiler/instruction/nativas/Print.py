import uuid
from abstract.Instruction import *
from abstract.Return import *
from symbol.Generator import *


class Print(Instruction):
    def __init__(self, value, line, column, new_line=False):
        Instruction.__init__(line, column)
        self.value = value
        self.new_line = new_line

    def compile(self, env):
        val = self.value.compile(env)
        gen_aux = Generator()
        generator = gen_aux.get_instance()

        if val.type == Type.INT:
            generator.add_print("d", val.value)
        elif val.type == Type.BOOL:
            temp_lbl = generator.new_label()
            generator.put_label(val.true_lbl)
            generator.print_true()
            generator.add_goto(temp_lbl)
            generator.put_label(val.false_lbl)
            generator.print_false()
            generator.put_label(temp_lbl)
        elif val.type == Type.STRING:
            generator.fprint_string()
            param_temp = generator.add_temp()
            generator.add_expression(param_temp, 'P', env.size, '+')
            generator.add_expression(param_temp, param_temp, '1', '+')
            generator.set_stack(param_temp, val.value)
            generator.new_env(env.size)
            generator.call_fun('printString')
            temp = generator.add_temp()
            generator.get_stack(temp, 'P')
            generator.ret_env(env.size)
        else:
            print("falta")
        if self.new_line:
            generator.add_print("c", 10)
