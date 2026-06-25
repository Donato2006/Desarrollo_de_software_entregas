import { useNavigate } from "react-router-dom";
import "../styles/CardConcierto.css";

function CardConcierto({ concierto }) {

  const navigate = useNavigate();

  const fecha = new Date(concierto.Fecha).toLocaleDateString("es-AR", {
    day: "2-digit",
    month: "long",
    year: "numeric",
  });

  const hora = new Date(concierto.Fecha).toLocaleTimeString("es-AR", {
    hour: "2-digit",
    minute: "2-digit",
  });

  const manejarDetalle = () => {

    navigate(`/conciertos/${concierto.ID}`);

  };

  return (
    <div className="card-concierto">

      <div className="card-image">
        <span>{concierto.Nombre}</span>
      </div>

      <div className="card-body">

        <h2>{concierto.Nombre}</h2>

        <p>📍 {concierto.Lugar}</p>
        <p>📅 {fecha}</p>
        <p>🕘 {hora}</p>
        <p>🎫 Cupos disponibles: {concierto.CuposDisponibles}</p>

        <button
          className="card-button"
          onClick={manejarDetalle}
        >
          Ver detalle
        </button>

      </div>

    </div>
  );
}

export default CardConcierto;