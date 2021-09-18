class Node:

    def __init__(self, value):
        self.value = value
        self.child = []
        self.contador = 0
        self.grafo = ""

    def insert_child(self, temp):
        self.child.append(temp)

    def get_graph_ast(self):
        self.grafo = "digraph Tree{\n"
        self.grafo += "nodo0[label=\"" + str(self.value) + "\"];\n"
        self.contador = 1
        self.graph_ast("nodo0", self)
        self.grafo += "}"
        return self.grafo

    def graph_ast(self, padre, temp):
        for child in temp.child:
            print(child.value)
            child_name = "nodo"+str(self.contador)
            self.grafo += child_name + \
                "[label=\"" + str(child.value) + "\"];\n"
            self.grafo += padre + "->" + child_name + ";\n"
            self.contador = self.contador + 1
            self.graph_ast(child_name, child)
