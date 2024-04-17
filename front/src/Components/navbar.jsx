import hacker from '../iconos/hacker.gif';

export default function Navbar(){
    return(
        <div>
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
                                    <a className="nav-link active" type="button" id="comandos">Comandos</a>
                                </li>

                                <li className="nav-item">
                                    <a className="nav-link" type="button" id="about">Explorador</a>
                                </li>

                                <li className="nav-item">
                                    <a className="nav-link" type="submit" id="ejecutar">Logout</a>
                                </li>

                            </ul>{/*Fin de lista de menus*/}
                        </div>{/*Fila de menus en la barra de navegacion*/}
                    </div>{/*Fila Titulo*/}
                </div>{/*Cierro tercer columna (Menu)*/}
            </nav>  
        </div>
    );
}
