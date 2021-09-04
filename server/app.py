from flask import Flask
from flask_cors import CORS
from server.analizadores.sintactico import parse
from server.interprete.enviroment.environment import *

app = Flask("__server__")
cors = CORS(app, resources={r"/api/*": {"origins": "*"}})
CORS(app)


# @app.route("/")
# def hello_world():
#     return {"text": "hola que tal como estas"}
#
#
# app.run(host='0.0.0.0', port=3000)

def main():
    env = Environment(None)
    ast = parse()

    try:
        for i in ast:
            i.execute(env)
    except Exception as e:
        print(e)


main()
