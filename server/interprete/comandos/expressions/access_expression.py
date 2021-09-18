from ..abstracts.expression import *


class AccessStruct(Expression):
    def __init__(self, idd, attribute, line, column):
        Expression.__init__(self, line, column)
        self.id = idd
        self.attribute = attribute

    def execute(self, environment):
        var = None
        if isinstance(self.id, AccessStruct):
            var = self.id.execute(environment)
        else:
            var = environment.get_var(self.id)
            return var.attributes[self.attribute]
        if var.attributes[self.attribute] is not None:
            return var.attributes[self.attribute]
        else:
            print("no existe este atributo en esta struct")
