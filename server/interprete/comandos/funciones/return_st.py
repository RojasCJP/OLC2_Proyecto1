from ..abstracts.expression import *
from ..abstracts.returns import *


class ReturnST(Expression):

    def __init__(self, expression, line, column):
        Expression.__init__(self, line, column)
        self.expression = expression

    def execute(self, environment: Environment):
        try:
            value = self.expression.execute(environment)
            return {"value": value, "type": Type.RETURN_ST}
        except Exception as e:
            print(e)
            print("error al retornar un valor")
