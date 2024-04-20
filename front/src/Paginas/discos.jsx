import { useState } from "react";
import { useNavigate } from "react-router-dom"

import diskIMG from '../iconos/disk.png';

export default function Discos(){
    const [discos, setDiscos] = useState([]);
    const navigate = useNavigate()

    useState(()=>{
        fetch('http://localhost:8080/discos')
        .then(Response => Response.json())
        .then(rawData => {console.log(rawData); setDiscos(rawData);})
    }, [])

    const onClick = (disco) => {
        console.log("click",disco)
        navigate(`/Disco/${disco}`) //navegar al objeto que hice click
    }

    return(
        <>
            DISCOS 
           <div style={{display:"flex", flexDirection:"row"}}>

                {discos && discos.length > 0 ? (
                    discos.map((disco, index) => {
                        return (
                            <div key={index} style={{
                                display: "flex",
                                flexDirection: "column", // Alinea los elementos en columnas
                                alignItems: "center", // Centra verticalmente los elementos
                                maxWidth: "100px",
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
        </>
        
    );
}
