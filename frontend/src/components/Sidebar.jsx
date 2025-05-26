import React, { useState, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import "./Sidebar.css";
import logo from "../assets/eduorg.logo.png";
import {
  FiLayout,
  FiUser,
  FiBookOpen,
  FiUsers,
  FiBell,
  FiLogOut,
  FiMessageSquare,
  FiSettings,
} from "react-icons/fi";

const Sidebar = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [userRole, setUserRole] = useState("");

  useEffect(() => {
    const userString = localStorage.getItem("user");
    if (userString) {
      try {
        const user = JSON.parse(userString);
        const role = user.role;
        setUserRole(role);
      } catch (error) {
        console.error("Erreur lors du parsing des données utilisateur:", error);
        // Rediriger vers login si les données sont corrompues
        navigate("/login");
      }
    } else {
      console.warn("Aucun utilisateur trouvé dans le local storage.");
      navigate("/login");
    }
  }, [navigate]);

  const handleLogout = () => {
    localStorage.removeItem("user");
    navigate("/login");
  };

  const handleNavigate = (path) => {
    navigate(path);
  };

  const isActive = (path) => location.pathname === path;

  const getRoleDisplayName = (role) => {
    switch (role) {
      case "chefDepartement":
        return "Chef Département";
      case "staffAdministrateur":
        return "Staff Administrateur";
      case "enseignant":
        return "Enseignant";
      default:
        return "Rôle Inconnu";
    }
  };

  const renderTeacherMenu = () => (
    <>
      <div
        className={`menu-item ${isActive("/dashboardtec") ? "active" : ""}`}
        onClick={() => handleNavigate("/dashboardtec")}
      >
        <FiLayout size={18} />
        <span>Tableau de bord</span>
      </div>

      <div
        className={`menu-item ${isActive("/profil") ? "active" : ""}`}
        onClick={() => handleNavigate("/profil")}
      >
        <FiUser size={18} />
        <span>Profil</span>
      </div>

      <div
        className={`menu-item ${isActive("/modules") ? "active" : ""}`}
        onClick={() => handleNavigate("/modules")}
      >
        <FiBookOpen size={18} />
        <span>Modules</span>
      </div>

      <div
        className={`menu-item ${isActive("/alerts") ? "active" : ""}`}
        onClick={() => handleNavigate("/alerts")}
      >
        <FiBell size={18} />
        <span>Alertes</span>
      </div>
    </>
  );

  const renderAdminMenu = () => (
    <>
      <div
        className={`menu-item ${isActive("/dashboardorga") ? "active" : ""}`}
        onClick={() => handleNavigate("/dashboardorga")}
      >
        <FiLayout size={18} />
        <span>Tableau de bord</span>
      </div>

      <div
        className={`menu-item ${isActive("/enseignants") ? "active" : ""}`}
        onClick={() => handleNavigate("/enseignants")}
      >
        <FiUser size={18} />
        <span>Enseignants</span>
      </div>

      <div
        className={`menu-item ${isActive("/modules") ? "active" : ""}`}
        onClick={() => handleNavigate("/modules")}
      >
        <FiBookOpen size={18} />
        <span>Modules</span>
      </div>

      <div
        className={`menu-item ${isActive("/organigramme") ? "active" : ""}`}
        onClick={() => handleNavigate("/organigramme")}
      >
        <FiUsers size={18} />
        <span>Organigramme</span>
      </div>
    </>
  );

  const renderDepartmentHeadMenu = () => (
    <>
      <div
        className={`menu-item ${isActive("/commentaires") ? "active" : ""}`}
        onClick={() => handleNavigate("/commentaires")}
      >
        <FiMessageSquare size={18} />
        <span>Commentaires</span>
      </div>

      <div
        className={`menu-item ${isActive("/parametre") ? "active" : ""}`}
        onClick={() => handleNavigate("/parametre")}
      >
        <FiSettings size={18} />
        <span>Paramètres</span>
      </div>
    </>
  );

  return (
    <aside className="sidebar">
      <div className="logo">
        <img src={logo} alt="EduOrg Logo" />
      </div>

      <div className="role-display">
        <strong>{getRoleDisplayName(userRole)}</strong>
      </div>

      <nav className="menu">
        {/* Menu pour Enseignant */}
        {userRole === "enseignant" && renderTeacherMenu()}

        {/* Menu pour Chef de Département (inclut menu admin + spécifiques) */}
        {userRole === "chefDepartement" && (
          <>
            {renderAdminMenu()}
            {renderDepartmentHeadMenu()}
          </>
        )}

        {/* Menu pour Staff Administrateur */}
        {userRole === "staffAdministrateur" && renderAdminMenu()}
      </nav>

      <div className="logout" onClick={handleLogout}>
        <FiLogOut size={18} />
        <span>Déconnexion</span>
      </div>
    </aside>
  );
};

export default Sidebar;