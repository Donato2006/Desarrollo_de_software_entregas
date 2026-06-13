import { Navigate } from "react-router-dom";

function ProtectedAdminRoute({ children }) {

  const token = localStorage.getItem("token");
  const rol = localStorage.getItem("rol");

  if (!token) {
    return <Navigate to="/login" />;
  }

  if (rol !== "admin") {
    return <Navigate to="/" />;
  }

  return children;
}

export default ProtectedAdminRoute;