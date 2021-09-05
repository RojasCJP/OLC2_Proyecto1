from ..abstracts.expression import *
from ..abstracts.returns import *
from enum import Enum


class ArithmeticEnum(Enum):
    PLUS = 0
    MINUS = 1
    TIMES = 2
    DIV = 3
    RAISED = 4
    MODULE = 5


class Arithmetic(Expression):
    def __init__(self, left, right, types, line, column):
        Expression.__init__(self, line, column)
        self.left = left
        self.right = right
        self.type = types

    def execute(self, environment: Environment):
        leftValue = self.left.execute(environment)
        rightValue = self.right.execute(environment)

        result = Return(0, Type.INT)

        if leftValue.type == Type.FLOAT or rightValue.type == Type.FLOAT:
            result.type = Type.FLOAT

        if self.type == ArithmeticEnum.PLUS:
            result.value = leftValue.value + rightValue.value
        elif self.type == ArithmeticEnum.MINUS:
            result.value = leftValue.value - rightValue.value
        elif self.type == ArithmeticEnum.TIMES:
            result.value = leftValue.value * rightValue.value
        elif self.type == ArithmeticEnum.DIV:
            result.value = leftValue.value / rightValue.value
            result.type = Type.FLOAT
        elif self.type == ArithmeticEnum.RAISED:
            result.value = leftValue.value ** rightValue.value
        elif self.type == ArithmeticEnum.MODULE:
            result.value = leftValue.value % rightValue.value
        return result
