import React from 'react';
import { UnControlled as CodeMirror } from 'react-codemirror2';
import 'codemirror/mode/javascript/javascript';
import 'codemirror/theme/monokai.css';
import 'codemirror/lib/codemirror.css';

export class Reportes extends React.Component {
    state = {
        errores: [],
        symbolos: [],
        optimizacion: [],
        console: "reporte ast"
    };

    componentDidMount = () => {
        this.actualizarErrores();
        this.actualizarVariables();
        this.reporteOptimizacion();
        this.getAST();
    };

    onChange = (editor, data, value) => {
        this.setState({
            code: '',
            console: '',
            stado: false,
        });
    };

    render() {
        return (
            <div>
                <h1>Errores</h1>
                <table className="table table-striped table-hover table-light">
                    <thead>
                        <tr>
                            <th scope="col">NO.</th>
                            <th scope="col">Tipo</th>
                            <th scope="col">Descripcion</th>
                        </tr>
                    </thead>
                    <tbody style={{ textAlign: 'center' }}>
                        {this.state.errores.map(
                            element =>
                                <tr key={element.num}>
                                    <td >{element.num}</td>
                                    <td >{element.type}</td>
                                    <td >{element.value}</td>
                                </tr>
                        )}
                    </tbody>
                </table>
                <h1>Variables</h1>
                <table className="table table-striped table-hover table-light">
                    <thead>
                        <tr>
                            <th scope="col">NO.</th>
                            <th scope="col">Id</th>
                            <th scope="col">Type</th>
                            <th scope="col">Value</th>
                        </tr>
                    </thead>
                    <tbody style={{ textAlign: 'center' }}>
                        {this.state.symbolos.map(
                            element =>
                                <tr key={element.num}>
                                    <td >{element.num}</td>
                                    <td >{element.id}</td>
                                    <td >{element.type}</td>
                                    <td >{element.value}</td>
                                </tr>
                        )}
                    </tbody>
                </table>
                <h1>Optimizacion</h1>
                <table className="table table-striped table-hover table-light">
                    <thead>
                        <tr>
                            <th scope="col">NO.</th>
                            <th scope="col">Operacion</th>
                        </tr>
                    </thead>
                    <tbody style={{ textAlign: 'center' }}>
                        {this.state.optimizacion.map(
                            element =>
                                <tr key={element.num}>
                                    <td >{element.num}</td>
                                    <td >{element.value}</td>
                                </tr>
                        )}
                    </tbody>
                </table>
                <h1>Codigo Graphviz</h1>
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
        );
    }
    actualizarErrores() {
        var ruta = 'https://jolc-compilador.herokuapp.com/errores';
        fetch(ruta, {
            method: 'GET',
            headers: { "Content-Type": "application/json" }
        }).then(async response => {
            var i = 1;
            const jsonInicial = await response.json();
            const json = jsonInicial.value;
            if (json != null) {
                json.forEach(element => {
                    element["num"] = i;
                    i++;
                });
                this.setState({ errores: json });
            }
        });
    }

    actualizarVariables() {
        var ruta = 'https://jolc-compilador.herokuapp.com/variables';
        fetch(ruta, {
            method: 'GET',
            headers: { "Content-Type": "application/json" }
        }).then(async response => {
            var i = 1;
            const jsonInicial = await response.json();
            const json = jsonInicial.value;
            if (json != null) {
                json.forEach(element => {
                    element["num"] = i;
                    i++;
                });
                this.setState({ symbolos: json });
            }
        });
    }
    reporteOptimizacion() {
        var ruta = 'https://jolc-compilador.herokuapp.com/optimizacion';
        fetch(ruta, {
            method: 'GET',
            headers: { "Content-Type": "application/json" }
        }).then(async response => {
            var i = 1;
            const jsonInicial = await response.json();
            const json = jsonInicial.value;
            console.log(json);
            if (json != null) {
                var auxiliar = [];
                json.forEach(element => {
                    var miniAux = {
                        num: i,
                        value: element
                    };
                    auxiliar.push(miniAux);
                    i++;
                });
                this.setState({ optimizacion: auxiliar });
                console.log(this.state.optimizacion);
            }
        });
    }
    getAST() {
        var ruta = 'https://jolc-server.herokuapp.com/ast';
        fetch(ruta, {
            method: 'GET',
            headers: { "Content-Type": "application/json" }
        }).then(async response => {
            var i = 1;
            const jsonInicial = await response.json();
            const json = jsonInicial.text;
            this.setState({ console: json });

        });
    }
}