import { useEffect, useState } from "react";
import api from "../services/api";
import Navbar from "../components/Navbar";
import CardConcierto from "../components/CardConcierto";
import "../styles/Home.css";

function Home() {
  const [conciertos, setConciertos] = useState([]);
  const [busqueda, setBusqueda] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    const cargarConciertos = async () => {
      try {
        const response = await api.get("/conciertos");
        setConciertos(response.data);
      } catch {
        setError("No se pudieron cargar los conciertos");
      }
    };

    cargarConciertos();
  }, []);

  const conciertosFiltrados = conciertos.filter((concierto) =>
    concierto.Nombre.toLowerCase().includes(busqueda.toLowerCase())
  );

  return (
    <div className="home-container">
      <Navbar
        busqueda={busqueda}
        setBusqueda={setBusqueda}
      />

      <div className="home-content">
        <h1 className="home-title">
          Próximos Conciertos
        </h1>

        {error && <p>{error}</p>}

        {conciertosFiltrados.length === 0 ? (
          <p>No se encontraron conciertos</p>
        ) : (
          <div className="conciertos-grid">
            {conciertosFiltrados.map((concierto) => (
              <CardConcierto
                key={concierto.ID}
                concierto={concierto}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

export default Home;