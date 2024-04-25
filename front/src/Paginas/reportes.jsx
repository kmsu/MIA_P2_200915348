import { useState } from "react";
import { useNavigate } from "react-router-dom"

import diskIMG from '../iconos/png.png';
import "../StyleSheets/fondo.css"

export default function Reportes({newIp="localhost"}){
    const [reportes, setReportes] = useState([]);

    useState(()=>{
        fetch(`http://${newIp}:8080/reportes`)
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setReportes(rawData);})
    }, [])

    const onClick = (repo) => {
        console.log("click",repo)
        //setEstado(true)
        fetch(`http://${newIp}:8080/descargar`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify(repo)
        })
        .then(response => {
            // Verificar si la respuesta fue exitosa (código 200)
            if (response.ok) {
                // Crear un objeto Blob desde los datos de la respuesta
                return response.blob();
            } else {
                // Si la respuesta no fue exitosa, lanzar un error
                throw new Error('Error en la respuesta del servidor');
            }
        })
        .then(blob => {
            // Crear un enlace temporal para descargar el archivo
            const url = window.URL.createObjectURL(blob);
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', repo); // Nombre de archivo predeterminado
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
            window.URL.revokeObjectURL(url); // Liberar recursos
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }

    return(
        <div className='body'>
            <div>&nbsp;&nbsp;&nbsp;</div>
            <div style={{display:"flex", flexDirection:"row", justifyContent: "center"}}><h1>REPORTES</h1></div>
            <div style={{display:"flex", flexDirection:"row", justifyContent: "center"}}> 
                {reportes && reportes.length > 0 ? (
                    reportes.map((reporte, index) => {
                        return (
                            <div key={index} style={{
                                display: "flex",
                                flexDirection: "column", // Alinea los elementos en columnas
                                alignItems: "center", // Centra verticalmente los elementos
                                maxWidth: "100px",
                                margin: "30px"
                                }} 
                                onClick={() => onClick(reporte)}
                            >
                                <img src={diskIMG} alt="disk" style={{width: "100px"}} />
                                {reporte}
                            </div>
                        )
                    })
                ):(
                    <div>No hay reportes disponibles</div>
                )}

            </div> 
        </div>
        
    );
}
