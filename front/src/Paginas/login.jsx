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
            <form onSubmit={handleSubmit}>
                <section className="vh-100" style={{backgroundColor: '#508bfc'}}>
                <div className="container py-5 h-100">
                    <div className="row d-flex justify-content-center align-items-center h-100">
                    <div className="col-12 col-md-8 col-lg-6 col-xl-5">
                        <div className="card shadow-2-strong" style={{borderRadius: "1rem"}}>
                        <div className="card-body p-5 text-center">

                            <h3 className="mb-5">Sign in</h3>

                            <div data-mdb-input-init className="form-outline mb-4">
                            <input type="text" className="form-control form-control-lg" placeholder="Enter Username" name="uname" required/>
                            <label className="form-label">User</label>
                            </div>

                            <div data-mdb-input-init className="form-outline mb-4">
                            <input type="password" placeholder="Enter Password" className="form-control form-control-lg" name="psw" required/>
                            <label className="form-label" >Password</label>
                            </div>

                            <button data-mdb-button-init data-mdb-ripple-init className="btn btn-primary btn-lg btn-block" type="submit">Login</button>

                        </div>
                        </div>
                    </div>
                    </div>
                </div>
                </section>
            </form>

            {errores  === -1 ? (
                <div>bienvenido </div>
                //Abrir sistema de archivos (mostrar carpeta raiz)
            ):errores === 0 ? (
                alert('Ya existe sesion activa')
            ):errores === 2 ?(
                alert('Particion sin formato')
            ):errores === 3 ?(
                alert('Contraseña incorrecta')
            ):errores === 4 ?(
                alert('No se encontro el usuario')
            ):(
                <div></div>
            )}
           
        </>
    )
}