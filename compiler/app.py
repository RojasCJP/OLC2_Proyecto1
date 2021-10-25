from grammar import grammar
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


def compile():
    gen_aux = Generator()
    gen_aux.clean_all()
    generator = gen_aux.get_instance()
    new_env = Environment(None)
    ast = grammar.parse("nada gg")
    try:
        for inst in ast:
            inst.compile(new_env)
        C3D = generator.get_code()
        f = open("salida.go", 'w')
        f.write(C3D)
        f.close()
    except Exception as e:
        print("no se puede compilar", e)


@app.route("/")
def hello_world():
    return {"text": "hola que tal como estas"}


@app.route("/salida")
def compilacion():
    compile()
    return {"text": "revisa la consola para ver el resultado"}


compile()
# if __name__ == "__main__":
#     app.run(host='0.0.0.0', port=3000)
#puede que trabaje aqui alguna vez pero el deploy tiene que ser desde windows creo yo

