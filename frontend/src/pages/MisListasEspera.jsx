import { useEffect, useState } from "react";
import api from "../services/api";
import "../styles/MisListasEspera.css";

function MisListasEspera() {

  const [listas, setListas] = useState([]);
  const [error, setError] = useState("");
  const [mensaje, setMensaje] = useState("");

  const cargarListas = async () => {

    try {

      const token = localStorage.getItem("token");

      const response = await api.get(
        "/mis-listas-espera",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setListas(response.data);

    } catch {

      setError("No se pudieron cargar las listas de espera");

    }

  };

  useEffect(() => {

    let activo = true;

    const obtenerListas = async () => {

      try {

        const token = localStorage.getItem("token");

        const response = await api.get(
          "/mis-listas-espera",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        if (activo) {
          setListas(response.data);
        }

      } catch {

        if (activo) {
          setError("No se pudieron cargar las listas de espera");
        }

      }

    };

    obtenerListas();

    return () => {
      activo = false;
    };

  }, []);

  const salirLista = async (id) => {

    try {

      const token = localStorage.getItem("token");

      await api.delete(
        `/lista-espera/${id}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setMensaje("Saliste de la lista de espera");
      setError("");

      cargarListas();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo salir de la lista"
      );

    }

  };

  return (
    <div className="listas-container">

      <h1>Mis Listas de Espera</h1>

      {mensaje && (
        <p className="mensaje-exito">
          {mensaje}
        </p>
      )}

      {error && (
        <p className="mensaje-error">
          {error}
        </p>
      )}

      {listas.length === 0 ? (

        <p>No estás anotado en ninguna lista de espera.</p>

      ) : (

        <div className="listas-grid">

          {listas.map((lista) => (

            <div
              key={lista.ID}
              className="lista-card"
            >

              <h2>
                {lista.Concierto.Nombre}
              </h2>

              <p>
                📍 {lista.Concierto.Lugar}
              </p>

              <p>
                📌 Posición:
                {" "}
                {lista.PosicionCola}
              </p>

              <p>

                Estado:

                {" "}

                {lista.Estado === "esperando"
                  ? "⏳ Esperando"
                  : "✅ Entrada asignada"}

              </p>

              {lista.FechaNotificacion && (

                <p>

                  📅 Fecha de asignación:

                  {" "}

                  {new Date(
                    lista.FechaNotificacion
                  ).toLocaleString("es-AR")}

                </p>

              )}

              <button
                className="salir-btn"
                onClick={() => salirLista(lista.ID)}
              >
                Salir de Lista
              </button>

            </div>

          ))}

        </div>

      )}

    </div>
  );
}

export default MisListasEspera;