import { useEffect, useState } from "react";
import api from "../services/api";
import "../styles/MisEntradas.css";

function MisEntradas() {
  
  const [entradaTransferir, setEntradaTransferir] = useState(null);
  const [correoDestino, setCorreoDestino] = useState("");
  const [entradas, setEntradas] = useState([]);
  const [error, setError] = useState("");
  const [mensaje, setMensaje] = useState("");

  const cargarEntradas = async () => {

    try {

      const token = localStorage.getItem("token");

      const response = await api.get(
        "/mis-entradas",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setEntradas(response.data || []);
      setError("");

    } catch (err) {

      setEntradas([]);

      setError(
        err.response?.data?.error ||
        "No se pudieron cargar las entradas"
      );

    }

  };

  useEffect(() => {

    const obtenerEntradas = async () => {
      await cargarEntradas();
    };

    obtenerEntradas();

  }, []);

  const cancelarEntrada = async (id) => {

    try {

      const token = localStorage.getItem("token");

      await api.delete(
        `/entradas/${id}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setMensaje("Entrada cancelada correctamente");
      setError("");

      await cargarEntradas();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo cancelar la entrada"
      );

    }

  }; 

  const transferirEntrada = async () => {
  try {
    const token = localStorage.getItem("token");

    await api.put(
      `/entradas/${entradaTransferir}/transferir`,
      {
        CorreoDestino: correoDestino,
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    setMensaje("Entrada transferida correctamente");
    setError("");
    setEntradaTransferir(null);
    setCorreoDestino("");

    await cargarEntradas();
  } catch (err) {
    setMensaje("");
    setError(
      err.response?.data?.error ||
        "No se pudo transferir la entrada"
    );
  }
 };

 return (
  <div className="mis-entradas-container">
    <h1>Mis Entradas</h1>

    {mensaje && <p className="mensaje-exito">{mensaje}</p>}

    {error && <p className="mensaje-error">{error}</p>}

    {entradas.length === 0 ? (
      <p>No tenés entradas compradas</p>
    ) : (
      <div className="entradas-grid">
        {entradas.map((entrada) => (
          <div key={entrada.ID} className="entrada-card">
            <h3>Entrada #{entrada.ID}</h3>

            <p>Estado: {entrada.Estado}</p>
            <p>Usuario ID: {entrada.UsuarioID}</p>
            <p>Concierto ID: {entrada.ConciertoID}</p>

            <button
              className="cancelar-btn"
              onClick={() => cancelarEntrada(entrada.ID)}
            >
              Cancelar Entrada
            </button>

            {entrada.Estado === "activa" && (
              <button
                className="transferir-btn"
                onClick={() => setEntradaTransferir(entrada.ID)}
              >
                Transferir Entrada
              </button>
            )}

            {entradaTransferir === entrada.ID && (
              <div className="transferir-box">
                <input
                  type="email"
                  placeholder="Correo del destinatario"
                  value={correoDestino}
                  onChange={(e) => setCorreoDestino(e.target.value)}
                />

                <button onClick={transferirEntrada}>
                  Confirmar transferencia
                </button>

                <button
                  onClick={() => {
                    setEntradaTransferir(null);
                    setCorreoDestino("");
                  }}
                >
                  Cancelar
                </button>
              </div>
            )}
          </div>
        ))}
      </div>
    )}
  </div>
 );
}

export default MisEntradas;