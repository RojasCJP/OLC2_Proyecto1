from abstract.Expression import *


class AccessStruct(Expression):

    def __init__(self, id, attribute, line, column):
        Expression.__init__(line, column)
        self.id = id
        self.attribute = attribute
