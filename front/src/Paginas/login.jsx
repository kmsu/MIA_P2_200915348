import { useParams } from 'react-router-dom';
import { useState } from "react";

export default function Login(){
    const { disk, part } = useParams()
    const [ errores, setErrores ] = useState();

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

        fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        })
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setErrores(rawData);})
    }

    return(
        <>
            Login {disk} {part}
            <form onSubmit={handleSubmit}>
                <div className="container">
                    <label htmlFor="uname"><b>Username</b></label>
                    <input type="text" placeholder="Enter Username" name="uname" required/>

                    <label htmlFor="psw"><b>Password</b></label>
                    <input type="password" placeholder="Enter Password" name="psw" required/>
                    
                    <button type="submit">Login</button>
                
                </div>
            </form> 

            {errores  === -1 ? (
                <div>bienvenido </div>
                //Abrir sistema de archivos (mostrar carpeta raiz)
            ):errores === 0 ? (
                <div>Ya existe sesion activa </div>
            ):errores === 2 ?(
                <div> Particion sin formato </div>
            ):errores === 3 ?(
                <div> Contraseña incorrecta </div>
            ):errores === 4 ?(
                <div> No se encontro el usuario </div>
            ):(
                <div></div>
            )}
        </>
    )
}