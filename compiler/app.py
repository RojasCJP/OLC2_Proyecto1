from grammar import grammar
from grammar.optimizator import optimizator
from sym.Environment import *
from sym.Generator import *

import sys
from flask import Flask
from flask import request
from flask_cors import CORS
from jinja2 import environment


sys.setrecursionlimit(3000)

app = Flask(__name__)
app.debug = True
cors = CORS(app, resources={r"/api/*": {"origins": "*"}})
CORS(app)


def compile(entrada):
    gen_aux = Generator()
    gen_aux.clean_all()
    generator = gen_aux.get_instance()
    new_env = Environment(None)
    ast = grammar.parse(entrada)
    try:
        for inst in ast:
            inst.compile(new_env)
        C3D = generator.get_code()
        f = open("salida.go", 'w')
        f.write(C3D)
        f.close()
        return C3D
    except Exception as e:
        print("no se puede compilar", e)
    return "error"


def bloques(input):
    instructions = optimizator.parse(input)
    instructions.Bloques()
    out = instructions.get_code()
    f = open("bloques.go", 'w')
    f.write(out)
    f.close()


def mirilla(input):
    instructions = optimizator.parse(input)
    instructions.Bloques()
    out = instructions.get_code()
    f = open("mirilla.go", 'w')
    f.write(out)
    f.close()


@app.route("/")
def hello_world():
    return {"text": "hola que tal como estas"}


@app.route("/salida", methods=['POST'])
def compilacion():
    data = request.get_json(force=True)
    codigo = compile(data["code"])
    return {"text": codigo}


@app.route('/variables')
def get_variables():
    retorno_true = {}
    retorno = []
    for key in Environment.variables.keys():
        nuevo_objeto = {}
        nuevo_objeto['id'] = key
        nuevo_objeto['value'] = str(Environment.variables[key].value)
        nuevo_objeto['type'] = str(Environment.variables[key].type)
        retorno.append(nuevo_objeto)
    for key in Environment.functions.key():
        nuevo_objeto = {}
        nuevo_objeto['id'] = key
        nuevo_objeto['value'] = 'instrucciones'
        nuevo_objeto['type'] = 'function'
        retorno.append(nuevo_objeto)
    for key in Environment.structs.key():
        nuevo_objeto = {}
        nuevo_objeto['id'] = key
        nuevo_objeto['value'] = 'atributos'
        nuevo_objeto['type'] = 'struct'
        retorno.append(nuevo_objeto)
    retorno_true['value'] = retorno
    Environment.variables = {}
    Environment.functions = {}
    Environment.structs = {}
    return retorno_true


compile("equis")
f = open(
    "/home/juanpa/Documents/Compi/OLC2_Proyecto1/compiler/salida.go", "r")
input = f.read()
bloques(input)
mirilla(input)
# if __name__ == "__main__":
#     app.run(host='0.0.0.0', port=3000)
# puede que trabaje aqui alguna vez pero el deploy tiene que ser desde windows creo yo
