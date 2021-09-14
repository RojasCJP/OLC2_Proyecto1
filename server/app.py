import sys

from flask import Flask
from flask import request
from flask_cors import CORS
from server.analizadores.sintactico import parse
from server.interprete.enviroment.environment import *

sys.setrecursionlimit(3000)

app = Flask("__server__")
cors = CORS(app, resources={r"/api/*": {"origins": "*"}})
CORS(app)


def entrada(input):
    env = Environment(None)
    ast = parse(input)

    try:
        for i in ast:
            i.execute(env)
    except Exception as e:
        print(e)


entrada("x")

@app.route("/")
def hello_world():
    return {"text": "hola que tal como estas"}


@app.route("/entrada", methods=['GET', 'POST'])
def entrada_codigo():
    if request.method == "POST":
        data = request.get_json(force=True)
        entrada(data["code"])
        return {"text": "se realizo exitosamente el parse"}
    elif request.method == "GET":
        return {"text": "hola que tal como estas"}

# app.run(host='0.0.0.0', port=3000)
