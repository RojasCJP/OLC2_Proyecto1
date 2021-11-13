from optimization.Optimizador import Optimizador
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
        generator.clean_all()
        return C3D
    except Exception as e:
        print("no se puede compilar", e)
        error = {}
        error['type'] = "no contemplado"
        error['text'] = "no se puede compilar"
        Environment.errores.append(error)
    return "error"


def bloques(input):
    instructions = optimizator.parse(input)
    instructions.Bloques()
    out = instructions.get_code()
    f = open("bloques.go", 'w')
    f.write(out)
    f.close()
    return out


def mirilla(input):
    instructions = optimizator.parse(input)
    instructions.Mirilla()
    out = instructions.get_code()
    f = open("mirilla.go", 'w')
    f.write(out)
    f.close()
    return out


@app.route("/")
def hello_world():
    return {"text": "hola que tal como estas"}


@app.route("/salida", methods=['POST'])
def compilacion():
    data = request.get_json(force=True)
    codigo = compile(data["code"])
    return {"text": codigo}


@app.route("/mirilla", methods=['POST'])
def mirillaa():
    data = request.get_json(force=True)
    codigo = mirilla(data['code'])
    return {"text": codigo}


@app.route("/bloques", methods=["POST"])
def bloquess():
    data = request.get_json(force=True)
    codigo = bloques(data['code'])
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
    for key in Environment.functions.keys():
        nuevo_objeto = {}
        nuevo_objeto['id'] = key
        nuevo_objeto['value'] = 'instrucciones'
        nuevo_objeto['type'] = 'function'
        retorno.append(nuevo_objeto)
    for key in Environment.structs.keys():
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


@app.route('/errores')
def get_errores():
    retorno = {}
    retorno["value"] = Environment.errores
    Environment.errores = []
    return retorno


@app.route('/optimizacion')
def get_optimizacion():
    retorno_true = {}
    retorno = Optimizador.reglas
    retorno_true['value'] = retorno
    Optimizador.reglas = []
    return retorno_true


# compile("equis")
# f = open(
#     "/home/juanpa/Documents/Compi/OLC2_Proyecto1/salida.txt", "r")
# input = f.read()
# mirilla(input)
# bloques(input)
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=4200)
