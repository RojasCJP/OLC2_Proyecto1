from abc import ABC, abstractmethod
from server.interprete.enviroment.environment import *


class Expression(ABC):
    def __init__(self, line, column):
        self.line = line
        self.column = column

    @abstractmethod
    def execute(self, environment: Environment): pass
