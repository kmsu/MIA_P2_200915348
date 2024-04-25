import { useState } from "react";
import { useNavigate } from "react-router-dom"

import diskIMG from '../iconos/disk.png';
import "../StyleSheets/fondo.css"

export default function Discos({newIp="localhost"}){
    const [discos, setDiscos] = useState([]);
    const navigate = useNavigate()

    useState(()=>{
        fetch(`http://${newIp}:8080/discos`)
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setDiscos(rawData);})
    }, [])

    const onClick = (disco) => {
        console.log("click",disco)
        navigate(`/Disco/${disco}`) //navegar al objeto que hice click
    }

    return(
        <div className='body'>
            <div>&nbsp;&nbsp;&nbsp;</div>
            <div style={{display:"flex", flexDirection:"row", justifyContent: "center"}}><h1>DISCOS</h1></div>
            <div style={{display:"flex", flexDirection:"row", justifyContent: "center"}}> 
                {discos && discos.length > 0 ? (
                    discos.map((disco, index) => {
                        return (
                            <div key={index} style={{
                                display: "flex",
                                flexDirection: "column", // Alinea los elementos en columnas
                                alignItems: "center", // Centra verticalmente los elementos
                                maxWidth: "100px",
                                margin: "10px"
                                }}
                                onClick={() => onClick(disco)}
                            >
                                <img src={diskIMG} alt="disk" style={{width: "100px"}} />
                                {disco}
                            </div>
                        )
                    })
                ):(
                    <div>No hay discos disponibles</div>
                )}
            </div> 
        </div>
    );
}
