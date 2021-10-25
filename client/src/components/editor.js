import React, { Component } from 'react';
import { UnControlled as CodeMirror } from 'react-codemirror2';
import { Redirect } from "react-router-dom";
import 'codemirror/mode/javascript/javascript';
import 'codemirror/theme/monokai.css';
import 'codemirror/lib/codemirror.css';

export class Editor extends Component {
    state = {
        value: '',
        redirect: null
    };

    onChange = (editor, data, value) => {
        this.setState({
            code: '',
            console: '',
            direcciones: '',
            stado: false,
        });
    };

    render() {
        if (this.state.redirect) {
            return <Redirect to={this.state.redirect} />;
        }
        return (
            <div>
                <div>
                    <h1>Editor</h1>
                </div>
                <div>
                    <button class='btn btn-dark m-2' onClick={() => this.entrada_datos()}>Interpretar
                    </button>
                    <button class='btn btn-dark m-2' onClick={() => this.compilador()}>Compilar</button>
                </div>
                <br />
                <br />
                <div class="row">
                    <div class="col" className="editor">
                        <CodeMirror
                            value=""
                            options={{
                                theme: 'monokai',
                                mode: 'javascript',
                                lineNumbers: true,
                            }}
                            onChange={(editor, data, value) => {
                                this.setState({ code: value });
                            }}
                        ></CodeMirror>
                    </div>
                    <div>
                        <h1>Console</h1>
                    </div>
                    <div class="col" className="editor">
                        <CodeMirror
                            value={this.state.console}
                            options={{
                                theme: 'monokai',
                                mode: 'javascript',
                                readOnly: true,
                                lineNumbers: false,
                            }}
                            onChange={(editor, data, value) => {
                                this.setState({ console: value });
                            }}
                        ></CodeMirror>
                    </div>
                </div>
                <div>
                    <h1>3 Direcciones</h1>
                </div>
                <div class="col" className="editor">
                    <CodeMirror
                        value={this.state.direcciones}
                        options={{
                            theme: 'monokai',
                            mode: 'javascript',
                            readOnly: true,
                            lineNumbers: false,
                        }}
                        onChange={(editor, data, value) => {
                            this.setState({ direcciones: value });
                        }}
                    ></CodeMirror>
                </div>
            </div>
        );
    }

    entrada_datos() {
        var codigoEnviar = { "code": this.state.code };
        console.log(codigoEnviar);
        console.log(this.state.console);
        fetch('https://jolc-server.herokuapp.com/entrada', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(codigoEnviar)
        })
            .then(async response => {
                const json = await response.json();
                this.setState({ console: json.text });
            });
    }

    compilador() {
        var codigoEnviar = { "code": this.state.code };
        console.log(codigoEnviar);
        console.log(this.state.console);
        fetch('https://jolc-compilador.herokuapp.com/salida', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(codigoEnviar)
        })
            .then(async response => {
                const json = await response.json();
                this.setState({ direcciones: json.text });
            });
    }
}