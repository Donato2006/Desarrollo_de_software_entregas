import "../styles/login.css";

function Login() {
  return (
    <div className="login-container">
      <div className="login-box">
        <h1>Iniciar Sesión</h1>

        <form>
          <input
            type="email"
            placeholder="Correo electrónico"
          />

          <input
            type="password"
            placeholder="Contraseña"
          />

          <button type="submit">
            Ingresar
          </button>
        </form>
      </div>
    </div>
  );
}

export default Login;