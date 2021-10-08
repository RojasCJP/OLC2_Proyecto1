from abc import ABC, abstractmethod


class Expression(ABC):
    def __init__(self, line, column):
        self.line = line
        self.column = column
        self.true_lbl = ''
        self.false_lbl = ''

    @abstractmethod
    def compile(self, env): pass
