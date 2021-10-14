from grammar import grammar
from sym.Environment import *
from sym.Generator import *


def compile():
    gen_aux = Generator()
    gen_aux.clean_all()
    generator = gen_aux.get_instance()
    new_env = Environment(None)
    ast = grammar.parse("nada gg")
    try:
        for inst in ast:
            inst.compile(new_env)
        print(generator.get_code())
    except:
        print("no se puede compilar")


compile()
