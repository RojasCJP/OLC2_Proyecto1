from interprete.enviroment.environment import Environment
from ..abstracts.expression import *
from ..abstracts.returns import *


class CallFunc(Expression):

    def __init__(self, idd, params, line, column):
        Expression.__init__(self, line, column)
        self.id = idd
        self.params = params

    def execute(self, environment):
        try:
            func = environment.get_function(self.id)
            if func is not None:
                new_env = Environment(environment.get_global_env())
                for i, param in enumerate(self.params):
                    value = self.params[i].execute(environment)
                    new_env.save_var(
                        func.params[i].id, value.value, value.type)
                res = func.instructions.execute(new_env)
                if res is not None:
                    if res.type == Type.RETURN_ST:
                        return res.value
                    else:
                        return res
            else:
                struct = environment.get_struct(self.id)
                if struct is not None:
                    attributes = {}
                    for i, param in enumerate(self.params):
                        value = self.params[i].execute(environment)
                        attributes.update({
                            struct[i]: value
                        })
                        return Return(attributes, Type.STRUCT, self.id)
        except Exception as e:
            print("error en la llamada de la funcion por esta exception ", e)
