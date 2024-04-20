import { useState } from "react";
import { useParams } from 'react-router-dom';
import { useNavigate } from "react-router-dom"
import partIMG from '../iconos/part.png';

export default function Partitions(){
    const { id } = useParams()
    const [ particiones, setParticiones ] = useState([]);
    const navigate = useNavigate()
    
    useState(()=>{
        fetch('http://localhost:8080/particiones', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify(id)
        })
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setParticiones(rawData);})
    }, [])

    const onClick = (particion) => {
        console.log("click",particion)
        navigate(`/login/${id}/${particion}`)
    }

    return(
        <>
            PARTICIONES EN EL DISCO {id} 
           <div style={{display:"flex", flexDirection:"row"}}>
                {particiones && particiones.length > 0 ? (
                    particiones.map((particion, index) => {
                        return (
                            <div key={index} style={{
                                display: "flex",
                                flexDirection: "column", // Alinea los elementos en columnas
                                alignItems: "center", // Centra verticalmente los elementos
                                maxWidth: "100px",
                                }}
                                onClick={() => onClick(particion)}
                            >
                                <img src={partIMG} alt="part" style={{width: "100px"}} />
                                {particion}
                            </div>
                        )
                    })
                ):(
                    <div>No hay particiones disponibles</div>
                )}
            </div> 
        </>
        
    );
}