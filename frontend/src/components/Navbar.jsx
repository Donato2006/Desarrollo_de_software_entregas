import { useNavigate } from "react-router-dom";
import "../styles/Navbar.css";

function Navbar({ busqueda, setBusqueda }) {

  const navigate = useNavigate();

  const token = localStorage.getItem("token");

  const cerrarSesion = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("rol");

    navigate("/");
    window.location.reload();
  };

  return (
    <nav className="navbar">

      <h2 className="navbar-logo">
        Ticket Conciertos
      </h2>

      <input
        className="navbar-search"
        type="text"
        placeholder="Buscar concierto..."
        value={busqueda}
        onChange={(e) => setBusqueda(e.target.value)}
      />

      {!token ? (
        <button
          className="navbar-button"
          onClick={() => navigate("/login")}
        >
          Iniciar sesión
        </button>
      ) : (
        <button
          className="navbar-button"
          onClick={cerrarSesion}
        >
          Cerrar sesión
        </button>
      )}

    </nav>
  );
}

export default Navbar;