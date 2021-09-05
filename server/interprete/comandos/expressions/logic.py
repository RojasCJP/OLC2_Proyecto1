from ..abstracts.expression import *
from ..abstracts.returns import *
from enum import Enum


class LogicEnum(Enum):
    NOT = 0
    AND = 1
    OR = 2


class Logic(Expression):

    def __init__(self, left, right, types, line, column):
        Expression.__init__(self, line, column)
        self.left = left
        self.right = right
        self.type = types

    def execute(self, environment: Environment):
        leftValue = self.left.execute(environment)
        rightValue = self.right.execute(environment)

        result = Return(0, Type.BOOL)

        if self.type == LogicEnum.NOT:
            if leftValue.value == rightValue.value:
                result.value = not leftValue.value
            else:
                print("su not no funciona")
        elif self.type == LogicEnum.AND:
            result.value = leftValue.value and rightValue.value
        elif self.type == LogicEnum.OR:
            result.value = leftValue.value or rightValue.value

        return result
