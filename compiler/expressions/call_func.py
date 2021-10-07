from abstract.Expression import *
from abstract.Return import *
from symbol.Environment import *


class CallFunc(Expression):

    def __init__(self, id, params, line, column):
        Expression.__init__(line, column)
        self.id = id
        self.params = params
