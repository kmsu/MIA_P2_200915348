import React, { useState } from 'react';
import "../StyleSheets/comandos.css"

export default function Comandos(){
    const [textValue, setTextValue] = useState('');

    const handleTextChange = (event) => {
        setTextValue(event.target.value);
    };

    const sendData = async (e) => {
        e.preventDefault();
        const data = {
            text: textValue
        };
        
        try {
            const response = await fetch('http://localhost:8080/analizar', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });
    
            if (!response.ok) {
                throw new Error('Error al enviar datos');
            }
    
            const responseData = await response.json();
            console.log('Respuesta del servidor:', responseData);
        } catch (error) {
            console.error('Error:', error);
        }

    }

    return(
        <div className='contenedorEjecutar'>
            <div id="espacio">&nbsp;&nbsp;&nbsp;</div>

            <table>
                <tbody>
                    <tr>
                        <td>
                            <textarea
                                className='entrada'
                                value={textValue}
                                onChange={handleTextChange}
                                cols="80"
                                rows="50"
                                placeholder='Ingrese comandos'
                            />
                        </td>
                    </tr>

                    <tr>
                        <td style={{textAlign:'center'}}>
                            <button type="button" className="btn btn-primary" onClick={(e) => sendData(e)}>Ejecutar</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    );
}