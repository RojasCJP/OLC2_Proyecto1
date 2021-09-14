from ..abstracts.expression import *
from ..abstracts.returns import *


class ReturnST(Expression):

    def __init__(self, expression, line, column):
        Expression.__init__(self, line, column)
        self.expression = expression

    def execute(self, environment):
        res = Return(None, Type.RETURN_ST)
        try:
            if self.expression is not None:
                value = self.expression.execute(environment)
                res.value = value
            return res
        except Exception as e:
            print("Erorr en Return", e)
