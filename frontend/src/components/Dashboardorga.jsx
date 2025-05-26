import React, { useEffect, useState } from "react";
import "./Dashboardorga.css";
import { useNavigate } from "react-router-dom";
import Sidebar from "../components/Sidebar";
import organigrammeS1Data from "../data/OrganigrammeS1.json";
import {
  FaChalkboardTeacher,
  FaBookOpen,
  FaGraduationCap,
  FaExclamationTriangle,
} from "react-icons/fa";

import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const Dashboardorga = ({ userRole }) => {
  const navigate = useNavigate();

  const [stats, setStats] = useState({
    enseignants: 0,
    modules: 0,
    specialites: 0,
  });

  const [alertMessage, setAlertMessage] = useState("");
  const [miniTableData, setMiniTableData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const initializeDashboard = () => {
      try {
        // Initialisation des statistiques
        const defaultStats = {
          enseignants: 130,
          modules: 12,
          specialites: 3,
        };
        
        const storedStats = localStorage.getItem("stats");
        if (storedStats) {
          const parsedStats = JSON.parse(storedStats);
          setStats(parsedStats);
        } else {
          localStorage.setItem("stats", JSON.stringify(defaultStats));
          setStats(defaultStats);
        }

        // Initialisation du message d'alerte
        const storedAlert = localStorage.getItem("alertMessage");
        if (storedAlert) {
          setAlertMessage(storedAlert);
        } else {
          const defaultAlert = "CERTAINES HEURES D'ENSEIGNEMENT EXCÉDENTAIRES ONT ÉTÉ DÉTECTÉES";
          localStorage.setItem("alertMessage", defaultAlert);
          setAlertMessage(defaultAlert);
        }

        // Données du mini tableau
        if (organigrammeS1Data && organigrammeS1Data.length > 0) {
          setMiniTableData(organigrammeS1Data.slice(0, 5));
        }
      } catch (error) {
        console.error("Erreur lors de l'initialisation du dashboard:", error);
      } finally {
        setIsLoading(false);
      }
    };

    initializeDashboard();
  }, []);

  // Données pour le graphique à barres
  const graphData = [
    { name: "Enseignants", value: Number(stats.enseignants) || 0 },
    { name: "Modules", value: Number(stats.modules) || 0 },
    { name: "Spécialités", value: Number(stats.specialites) || 0 },
  ];

  const navigationHandlers = {
    organigramme: () => navigate("/organigramme"),
    modules: () => navigate("/modules"),
    enseignants: () => navigate("/enseignants"),
    alerts: () => navigate("/alerts"),
  };

  const renderMiniTable = () => {
    if (isLoading) {
      return (
        <div className="loading-state">
          <p>Chargement des données...</p>
        </div>
      );
    }

    if (miniTableData.length === 0) {
      return (
        <div className="empty-state">
          <p>Aucune donnée disponible</p>
        </div>
      );
    }

    return (
      <div className="mini-table-container">
        <table className="mini-table">
          <thead>
            <tr>
              <th>Section</th>
              <th>Module</th>
              <th>Cours</th>
              <th>TD1</th>
            </tr>
          </thead>
          <tbody>
            {miniTableData.map((ligne, index) => (
              <tr key={index} className="table-row">
                <td className="table-cell">{ligne.Section || "-"}</td>
                <td className="table-cell">{ligne.Module || "-"}</td>
                <td className="table-cell">{ligne.Cours || "-"}</td>
                <td className="table-cell">{ligne.TD1 || "-"}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };

  const renderAlertSection = () => {
    if (!alertMessage || userRole !== "enseignant") return null;

    return (
      <div className="alert-section">
        <div className="alert-message">
          <div className="alert-content">
            <FaExclamationTriangle className="alert-icon" />
            <div className="alert-text">
              <strong>Problème détecté :</strong>
              <span>{alertMessage}</span>
            </div>
          </div>
          <button 
            className="orga-action-btn"
            onClick={navigationHandlers.alerts}
            aria-label="Voir toutes les alertes"
          >
            Voir les alertes
          </button>
        </div>
      </div>
    );
  };

  const renderStatsCards = () => {
    const cardsData = [
      {
        title: "Total d'Enseignants",
        value: stats.enseignants,
        icon: FaChalkboardTeacher,
        onClick: navigationHandlers.enseignants,
        clickable: true,
      },
      {
        title: "Modules",
        value: stats.modules,
        icon: FaBookOpen,
        onClick: navigationHandlers.modules,
        clickable: true,
      },
      {
        title: "Spécialités",
        value: stats.specialites,
        icon: FaGraduationCap,
        clickable: false,
      },
    ];

    return (
      <div className="cards-wrapper">
        {cardsData.map((card, index) => (
          <div
            key={index}
            className={`card ${card.clickable ? 'clickable' : ''}`}
            onClick={card.clickable ? card.onClick : undefined}
            role={card.clickable ? "button" : "presentation"}
            tabIndex={card.clickable ? 0 : -1}
            onKeyDown={card.clickable ? (e) => {
              if (e.key === 'Enter' || e.key === ' ') {
                e.preventDefault();
                card.onClick();
              }
            } : undefined}
          >
            <h2 className="card-title">
              <card.icon className="card-icon" />
              {card.title}
            </h2>
            <p className="card-value">{card.value}</p>
          </div>
        ))}
        
        {/* Carte Statistiques avec graphique */}
        <div className="card stats-chart-card">
          <h2 className="card-title">Statistiques Visuelles</h2>
          <div className="chart-container">
            <ResponsiveContainer width="100%" height={150}>
              <BarChart
                data={graphData}
                margin={{ top: 10, right: 20, left: 0, bottom: 5 }}
                barSize={20}
              >
                <XAxis 
                  dataKey="name" 
                  stroke="#1e3a8a"
                  fontSize={12}
                />
                <YAxis fontSize={12} />
                <Tooltip 
                  contentStyle={{
                    backgroundColor: '#f8fafc',
                    border: '1px solid #e2e8f0',
                    borderRadius: '6px'
                  }}
                />
                <Bar 
                  dataKey="value" 
                  fill="#2563eb" 
                  radius={[4, 4, 0, 0]}
                />
              </BarChart>
            </ResponsiveContainer>
          </div>
        </div>
      </div>
    );
  };

  if (isLoading) {
    return (
      <div className="dashboardorga-container">
        <Sidebar />
        <main className="orga-main-content">
          <div className="loading-dashboard">
            <p>Chargement du tableau de bord...</p>
          </div>
        </main>
      </div>
    );
  }

  return (
    <div className="dashboardorga-container">
      <Sidebar />
      <main className="orga-main-content">
        <header className="orga-header">
          <h1 className="welcome">Bienvenue sur votre tableau de bord</h1>
        </header>

        {renderAlertSection()}

        <div className="dashboard-content">
          <div className="top-section">
            <div 
              className="mini-organigramme-box"
              onClick={navigationHandlers.organigramme}
              role="button"
              tabIndex={0}
              onKeyDown={(e) => {
                if (e.key === 'Enter' || e.key === ' ') {
                  e.preventDefault();
                  navigationHandlers.organigramme();
                }
              }}
              aria-label="Voir l'organigramme complet"
            >
              <h3 className="mini-title">Aperçu Organigramme S1</h3>
              
              {renderMiniTable()}

              {organigrammeS1Data.length > 5 && (
                <p className="mini-note">+ {organigrammeS1Data.length - 5} autres modules...</p>
              )}
            </div>

            {renderStatsCards()}
          </div>
        </div>
      </main>
    </div>
  );
};

export default Dashboardorga;