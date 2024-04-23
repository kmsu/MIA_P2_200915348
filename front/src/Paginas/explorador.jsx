import { useState } from "react";

import texto from '../iconos/txt1.png';
import carpeta from '../iconos/carpeta.png';
import volver from '../iconos/volver.png';
import "../StyleSheets/explorer.css"

export default function Explorer(){
    const [ archivos, setArchivos ] = useState([]);
    const [ path, setPath ] = useState("path: /");

    useState(()=>{
        fetch('http://localhost:8080/explorer')
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setArchivos(rawData);})
    }, [])

    const onClick = (archivo) => {
        console.log("buscar",archivo)
        let tmp = path+archivo+"/"
        setPath(tmp)
        fetch('http://localhost:8080/contenido', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify(archivo)
        })
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setArchivos(rawData);})
        //navigate(`/login/${id}/${particion}`)
    }

    const back = () =>{
        let tmp2 = path.split("/")
        if (tmp2.length > 2) {
            let newPath = "path: /"
            for (let i=1; i < tmp2.length-2; i++){
                newPath += tmp2[i]+"/"
            }
            console.log("back ", newPath)
            setPath(newPath)
            fetch('http://localhost:8080/back')
            .then(Response => Response.json())
            .then(rawData => {console.log(rawData); setArchivos(rawData);})
        }
    }

    return(
        <>
            <div className="container">
                <div className="d-flex justify-content-center">
                    
                    <div className="explorer-card">
                        <div className="explorer-card-header">
                            <img onClick={back} src={volver} alt="volver" style={{width: "20px", margin: "5px"}} />
                            {path}
                        </div>
                        <div className="container-with-scroll" style={{display:"flex", flexDirection:"row"}}>
                            {archivos && archivos.length > 0 ? (
                                archivos.map((archivo, index) => {
                                    return (
                                        <div key={index} style={{
                                            display: "flex",
                                            flexDirection: "column", // Alinea los elementos en columnas
                                            alignItems: "center", // Centra verticalmente los elementos
                                            maxWidth: "100px",
                                            margin: "10px"
                                            }}
                                        >
                                            {archivo.endsWith('.txt')? (
                                                <img src={texto} alt="archivo" style={{width: "100px"}} />    
                                            ):(
                                                <img onClick={() => onClick(archivo)} src={carpeta} alt="archivo" style={{width: "100px"}} />
                                            )}
                                            {archivo}
                                        </div>
                                    )
                                })
                            ):(
                                <div>No hay archivos disponibles</div>
                            )}
                            
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}