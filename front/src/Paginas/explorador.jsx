import { useState } from "react";
import { useParams } from 'react-router-dom';
import texto from '../iconos/txt1.png';
import carpeta from '../iconos/carpeta.png';
import "../StyleSheets/explorer.css"

export default function Explorer(){
    const [ archivos, setArchivos ] = useState([]);

    useState(()=>{
        fetch('http://localhost:8080/explorer')
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setArchivos(rawData);})
    }, [])

    return(
        <>
            <div className="container">
                <div className="d-flex justify-content-center">
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
                                            <img src={carpeta} alt="archivo" style={{width: "100px"}} />
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
        </>
    );
}