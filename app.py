import os
from flask import Flask

app = Flask(__name__)

@app.route('/')

def hello():
    return ' <h2> Test app1 (ptyhon) <h2>'

port = int(os.environ.get('PORT', 5000))
app.run(host='0.0.0.0', port=port)
