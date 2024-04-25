import { useParams, useNavigate } from 'react-router-dom';
import { useState } from "react";

import "../StyleSheets/login.css"
import user from '../iconos/profile.png';
import key from '../iconos/key.png';

export default function Login({newIp="localhost"}){
    const { disk, part } = useParams()
    const [ estado, setEstado ] = useState();
    const navigate = useNavigate()

    const handleSubmit = (e) => {
        e.preventDefault()
        console.log("submit", disk, part)
  
        const user = e.target.uname.value
        const pass = e.target.psw.value
  
        console.log("user", user, pass)

        const data = {
            usuario: user,
            password: pass,
            id: part
        };

        fetch(`http://${newIp}:8080/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        })
        .then(Response => Response.json())
        .then(rawData => {
            console.log(rawData); 
            setEstado(rawData);
            if (rawData === -1){
                onClick(part)
            }
        })
    }

    const onClick = (particion) => {
        console.log("nueva pagina",particion)
        navigate(`/explorador/${particion}`)
    }

    return(
        <>
            <div className="container">
                <div className="d-flex justify-content-center">
                    <div className="card ">
                        <div className="card-header">
                            <h3>Sign In</h3>
                        </div>
                        <div className="card-body">
                            <form onSubmit={handleSubmit}>
                                <div className="input-group form-group">
                                    <div className="input-group-prepend">
                                        <span className="input-group-text" style={{padding: '0', marginRight:'10px'}}>
                                            <img src={user} alt="user" style={{width: '100%', height: '100%'}} />
                                        </span>
                                    </div>
                                    <input type="text" className="form-control" placeholder="username" name="uname" required/>
                                </div>
                                
                                <div>&nbsp;&nbsp;&nbsp;</div>
                                
                                <div className="input-group form-group">
                                    <div className="input-group-prepend">
                                        <span className="input-group-text" style={{padding: '0', marginRight:'10px'}}>
                                            <img src={key} alt="user" style={{width: '100%', height: '100%'}} />
                                        </span>
                                    </div>
                                    <input type="password" className="form-control" placeholder="password" name="psw" required/>
                                </div>

                                <div>&nbsp;&nbsp;&nbsp;</div>
                                
                                <div style={{textAlign:'center'}}>
                                    <button type="submit" className="btn btn-primary login_btn">Login</button>
                                </div>
                            </form>

                            <div>&nbsp;&nbsp;&nbsp;</div>

                            <div className='estadoLogin'>
                                {estado === 0 ? (
                                    <div>Ya existe sesion activa</div>
                                ):estado === 2 ?(
                                    <div>Particion sin formato</div>
                                ):estado === 3 ?(
                                    <div>Contraseña incorrecta</div>
                                ):estado === 4 ?(
                                    <div>No se encontro el usuario</div>
                                ):(
                                    <div></div>
                                )}
                            </div>
                            
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}
