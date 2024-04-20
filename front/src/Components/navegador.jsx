import hacker from '../iconos/hacker.gif';

import { Routes, Route, HashRouter, Link } from 'react-router-dom'
import Comandos from '../Paginas/comandos';
import Discos from '../Paginas/discos';
import Partitions from '../Paginas/partition';
import Login from '../Paginas/login';

export default function Navegador(){
    return(
        <HashRouter>
            <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
                {/*COLUMNAS*/}
                <div id="espacio">&nbsp;&nbsp;&nbsp;</div>
                
                <div className="conteiner-fluid"> 
                    <img src={hacker} alt="" width="64" height="64" className="d-inline-block align-text-top"></img>
                </div>

                <div className="conteiner"> 
                    {/*Fila 1 (titulo del proyecto, RESPALDO)*/}
                    <div className="container-fluid">
                        <a className="navbar-brand" type="submit" >
                            ARCHIVOS PROYECTO 2            
                        </a>
                        {/*Cada bloque div aqui dentro es una nueva fila hacia abajo*/}
                        {/*Fila 2 (menus)*/}
                        <div className="collapse navbar-collapse" id="navbarColor02">
                            {/*ul es una lista no ordenada*/}
                            <ul className="navbar-nav me-auto">
                                {/*LISTA DE MENUS QUE ESTARAN EN LA BARRA DE NAVEGACION*/}
                                <li className="nav-item">
                                    {/* Utiliza Link en lugar de a para navegar entre rutas */}
                                    <Link className="nav-link active" to="/Comandos">Comandos</Link>
                                </li>

                                <li className="nav-item">
                                    <Link className="nav-link" to="/Discos">Explorador</Link>
                                </li>

                                <li className="nav-item">
                                    <Link className="nav-link" to="/Logout">Logout</Link>
                                </li>

                            </ul>{/*Fin de lista de menus*/}
                        </div>{/*Fila de menus en la barra de navegacion*/}
                    </div>{/*Fila Titulo*/}
                </div>{/*Cierro tercer columna (Menu)*/}
            </nav> 

            <Routes>
                <Route path="/" element ={<Comandos/>}/> {/*home*/}
                <Route path="/Comandos" element ={<Comandos/>}/> 
                <Route path="/Discos" element ={<Discos/>}/> 
                <Route path="/Disco/:id" element ={<Partitions/>}/> 
                <Route path="/Login/:disk/:part" element ={<Login/>}/>                 
            </Routes>
        </HashRouter>
    );
}