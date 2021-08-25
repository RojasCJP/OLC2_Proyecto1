from flask import Flask
from flask_cors import CORS

app = Flask("__server__")
cors = CORS(app, resources={r"/api/*": {"origins": "*"}})
CORS(app)


@app.route("/")
def hello_world():
    return {"text": "hola que tal como estas"}


app.run(host='0.0.0.0', port=3000)
