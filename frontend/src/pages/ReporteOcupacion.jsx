import { useEffect, useState } from "react";
import api from "../services/api";
import "../styles/ReporteOcupacion.css";

function ReporteOcupacion() {

  const [reporte, setReporte] = useState([]);

  useEffect(() => {

    const cargarReporte = async () => {

      try {

        const token = localStorage.getItem("token");

        const response = await api.get(
          "/reporte-ocupacion",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        setReporte(response.data);

      } catch (error) {

        console.log(error);

      }

    };

    cargarReporte();

  }, []);

  return (
    <div className="reporte-container">

      <h1>Reporte de Ocupación</h1>

      {reporte.map((item) => (

        <div
          key={item.ID}
          className="reporte-card"
        >

          <h2>{item.Nombre}</h2>

          <p>
            Cupo total: {item.CupoTotal}
          </p>

          <p>
            Entradas vendidas: {item.EntradasVendidas}
          </p>

          <p>
            Disponibles: {item.CuposDisponibles}
          </p>

          <p>
            Ocupación:
            {" "}
            {item.Porcentaje.toFixed(2)}%
          </p>

        </div>

      ))}

    </div>
  );
}

export default ReporteOcupacion;