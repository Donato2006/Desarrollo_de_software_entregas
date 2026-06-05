import "../styles/Navbar.css";

function Navbar({ busqueda, setBusqueda }) {
  return (
    <nav className="navbar">
      <h2 className="navbar-logo">Ticket Conciertos</h2>

      <input
        className="navbar-search"
        type="text"
        placeholder="Buscar concierto..."
        value={busqueda}
        onChange={(e) => setBusqueda(e.target.value)}
      />

      <button className="navbar-button">
        Iniciar sesión
      </button>
    </nav>
  );
}

export default Navbar;