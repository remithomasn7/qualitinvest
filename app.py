from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def hello_world():
    years = ['2019', '2020', '2021']
    revenues = [100000, 120000, 150000]
    return render_template('index.html', years=years, revenues=revenues)