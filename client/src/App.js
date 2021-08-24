import './App.css'
import { Navbar } from './components/navbar'
import { Home } from './components/home'
import { Datos } from './components/datos'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { Reportes } from './components/reportes'
import { Editor } from './components/editor'

function App() {
  return (
    <div className="App">
      <Router>
        <Navbar></Navbar>
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact path='/datos' component={Datos} />
          <Route exact path='/reportes' component={Reportes} />
          <Route exact path='/editor' component={Editor} />
        </Switch>
      </Router>
    </div>
  )
}

export default App
