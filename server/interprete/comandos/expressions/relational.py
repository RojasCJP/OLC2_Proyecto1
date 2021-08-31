from ..abstracts.expression import *
from ..abstracts.returns import *
from enum import Enum


class RelationalEnum(Enum):
    GREATER = 0
    MINOR = 1
    GREATER_EQUAL = 2
    MINOR_EQUAL = 3
    EQUAL_EQUAL = 4
    DIFFERENT = 5


class Relational(Expression):
    def __init__(self, left, right, type, line, column):
        Expression.__init__(self, line, column)
        self.left = left
        self.right = right
        self.type = type

    def execute(self, environment: Environment):
        leftValue = self.left.execute(environment)
        rightValue = self.right.execute(environment)

        result = Return(False, Type.BOOL)

        if self.type == RelationalEnum.GREATER:
            result.value = leftValue.value > rightValue.value
        elif self.type == RelationalEnum.MINOR:
            result.value = leftValue.value < rightValue.value
        elif self.type == RelationalEnum.GREATER_EQUAL:
            result.value == leftValue >= rightValue.value
        elif self.type == RelationalEnum.MINOR_EQUAL:
            result.value == leftValue.value <= leftValue.value
        elif self.type == RelationalEnum.EQUAL_EQUAL:
            result.value = leftValue.value == rightValue.value
        elif self.type == RelationalEnum.DIFFERENT:
            result.value = leftValue != rightValue.value

        return result
