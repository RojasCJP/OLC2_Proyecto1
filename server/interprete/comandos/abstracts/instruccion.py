from abc import ABC, abstractmethod
from interprete.enviroment.environment import Environment


class Instruction(ABC):
    def __init__(self, line, column):
        self.line = line
        self.column = column

    @abstractmethod
    def execute(self, env: Environment): pass
