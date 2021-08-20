import React from 'react'
import { Link } from 'react-router-dom'

export class Navbar extends React.Component {
    render() {
        return (
            <div>
                <nav class="navbar navbar-expand-lg navbar-light bg-light">
                    <div class='container'>
                        <a class="navbar-brand" href='/'>JOLC</a>
                        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                            <span class="navbar-toggler-icon"></span>
                        </button>
                        <div class="collapse navbar-collapse" id="navbarSupportedContent">
                            <ul class="navbar-nav mr-auto">
                                <li class="nav-item active">
                                    <Link to='/'>
                                        <a class="nav-link">Home</a>
                                    </Link>
                                </li>
                                <li class="nav-item active">
                                    <Link to='/datos'>
                                        <a class="nav-link" >Datos</a>
                                    </Link>
                                </li>
                                <li class="nav-item active">
                                    <Link to='/editor'>
                                        <a class="nav-link" >Editor</a>
                                    </Link>
                                </li>
                                <li class="nav-item active">
                                    <Link to='/reportes'>
                                        <a class="nav-link" >Reportes</a>
                                    </Link>
                                </li>
                            </ul>
                        </div>
                    </div>
                </nav>
            </div>
        )
    }
}

