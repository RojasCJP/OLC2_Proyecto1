from abc import ABC, abstractmethod


class Expression(ABC):
    def __init__(self, line, column):
        self.line = line
        self.column = column

    @abstractmethod
    def execute(self, environment): pass
