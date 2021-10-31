from abstract.Instruction import *
from abstract.Return import *
from instruction.variables.Declaration import *
from expressions.Literal import *
from sym.Generator import *


class For(Instruction):
    def __init__(self, variable, value1, instructions, line, column, value2=None):
        Instruction.__init__(self, line, column)
        self.variable = variable
        self.value1 = value1
        self.value2 = value2
        self.instructions = instructions

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        left_val = self.value1.compile(env)
        if self.value2 is not None:
            right_val = self.value2.compile(env)
            temp1 = generator.add_temp()
            lit_temp1 = Literal(temp1, Type.INT, self.line, self.column)
            generator.add_expression(temp1, left_val.value, '', '')
            continue_lbl = generator.new_label()
            generator.put_label(continue_lbl)
            end_lbl = generator.new_label()
            generator.add_if(temp1, right_val.value, ">", end_lbl)
            new_env = Environment(env)
            new_env.break_lbl = end_lbl
            new_env.continue_lbl = continue_lbl
            declaration = Declaration(
                self.variable, lit_temp1, self.line, self.column)
            declaration.compile(new_env)
            self.instructions.compile(new_env)
            generator.add_expression(temp1, temp1, '1', '+')
            declaration.compile(new_env)
            generator.add_goto(continue_lbl)
            generator.put_label(end_lbl)
