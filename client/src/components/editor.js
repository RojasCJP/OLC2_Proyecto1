import React from 'react'

export class Editor extends React.Component {
    render() {
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
                <div>
                    <textarea></textarea>
                    <textarea></textarea>
                </div>
            </div>
        )
    }
}