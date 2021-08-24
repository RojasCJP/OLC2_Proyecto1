import React from 'react'
import CodeMirror from '@uiw/react-codemirror'
import 'codemirror/keymap/sublime'
import 'codemirror/theme/monokai.css'

export class Editor extends React.Component {
    render() {
        var code = "const a = 'hola que tal'"
        var console = "hola que tal"
        return (
            <div>
                <div>
                    <h1>Editor</h1>
                </div>
                <div>
                    <button class='btn btn-dark m-2'>Abrir</button>
                    <button class='btn btn-dark m-2'>Reportes</button>
                    <button class='btn btn-dark m-2'>Compilar</button>
                    <button class='btn btn-dark m-2'>Gramaticas</button>
                </div>
                <br />
                <br />
                <div class="row">
                    <div class="col">
                        <CodeMirror
                            value={code}
                            options={{
                                theme: 'monokai',
                                keyMap: 'sublime',
                                mode: 'jsx',

                            }}
                        ></CodeMirror>
                    </div>
                    <div class="col">
                        <CodeMirror
                            value={console}
                            options={{
                                theme: 'monokai',
                                keyMap: 'sublime',
                                mode: 'jsx',
                                readOnly: true,
                                lineNumbers: false,
                            }}
                        ></CodeMirror>
                    </div>
                </div>
            </div>
        )
    }
}