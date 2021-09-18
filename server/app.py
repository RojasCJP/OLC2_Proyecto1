from ply.yacc import Production
import sys
from threading import Event

from flask import Flask
from flask import request
from flask_cors import CORS
from jinja2 import environment
from sintactico import parse
from arbol import arbol
from interprete.enviroment.environment import *
from interprete.comandos.nativas.print import *

sys.setrecursionlimit(3000)

app = Flask(__name__)
cors = CORS(app, resources={r"/api/*": {"origins": "*"}})
CORS(app)

ultima_entrada = ""


def entrada(input):
    env = Environment(None)
    ast = parse(input)
    arbol(input)
    try:
        for i in ast:
            i.execute(env)
    except Exception as e:
        print(e)
    return Print.printlist


@app.route("/")
def hello_world():
    return {"text": "hola que tal como estas"}


@app.route("/variables")
def get_variables():
    retorno_true = {}
    retorno = []
    for key in Environment.variables.keys():
        nuevoObjeto = {}
        nuevoObjeto["id"] = key
        nuevoObjeto["value"] = str(Environment.variables[key].value)
        nuevoObjeto["type"] = str(Environment.variables[key].type)
        retorno.append(nuevoObjeto)
    for key in Environment.functions.keys():
        nuevoObjeto = {}
        nuevoObjeto["id"] = key
        nuevoObjeto["value"] = "instrucciones"
        nuevoObjeto["type"] = "function"
        retorno.append(nuevoObjeto)
    for key in Environment.structs.keys():
        nuevoObjeto = {}
        nuevoObjeto["id"] = key
        nuevoObjeto["value"] = str(Environment.structs[key])
        nuevoObjeto["type"] = "struct"
        retorno.append(nuevoObjeto)
    retorno_true["value"] = retorno
    print(retorno_true)
    return retorno_true


@app.route("/errores")
def get_errores():
    retorno = {}
    retorno["value"] = Environment.errores
    return retorno


@app.route("/entrada", methods=['GET', 'POST'])
def entrada_codigo():
    if request.method == "POST":
        data = request.get_json(force=True)
        consola = entrada(data["code"])
        Print.printlist = ''
        return {"text": consola}
    elif request.method == "GET":
        return {"text": "hola que tal como estas"}


@app.route("/ast")
def get_ast():
    arbol(ultima_entrada)
    return{"text": "reporte generado"}
# app.run(host='0.0.0.0', port=3000)
