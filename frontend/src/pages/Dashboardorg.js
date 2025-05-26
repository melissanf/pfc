import React from 'react';
import './Dashboardorg.css';
import { Link } from 'react-router-dom';

const Dashboardorg = ({ userRole }) => {
  return (
    <div className="dashboard-container">
      <header className="dashboard-header">
        <h1 className="dashboard-title">Tableau de bord</h1>
      </header>

      {userRole === 'chef' ? (
        <div className="chef-dashboard">
          <div className="organigramme-section">
            <Link to="/organigramme" className="organigramme-link">
              <span className="organigramme-text">ORGANIGRAMME</span>
              <span className="arrow">➝</span>
            </Link>
          </div>

          <div className="alert-section">
            <div className="alert-box">
              <div className="alert-header">
                <span className="alert-icon">⚠️</span>
                <span className="alert-title">PROBLÈME DÉTECTÉ</span>
              </div>
              <p className="alert-message">
                CERTAINES HEURES D'ENSEIGNEMENT EXCÉDENTAIRES ONT ÉTÉ DÉTECTÉES
              </p>
            </div>
          </div>

          <div className="action-buttons">
            <button className="btn btn-primary">
              MODIFIER L'ORGANIGRAMME
            </button>
            <button className="btn btn-secondary">
              VOIR LES STATISTIQUES GÉNÉRALES
            </button>
          </div>

          <div className="stats-grid">
            <div className="stat-card">
              <div className="stat-number">25 à 30</div>
              <div className="stat-label">TOTAL D'ENSEIGNANTS</div>
            </div>
            <div className="stat-card">
              <div className="stat-number">10 à 12</div>
              <div className="stat-label">Modules</div>
            </div>
            <div className="stat-card">
              <div className="stat-number">3</div>
              <div className="stat-label">Spécialités</div>
            </div>
          </div>
        </div>
      ) : (
        <div className="regular-dashboard">
          <div className="overview-section">
            <h2 className="overview-title">Vue d'ensemble de l'organigramme</h2>
            <div className="overview-controls">
              <Link to="/organigramme" className="organigramme-link">
                <span className="organigramme-text">ORGANIGRAMME</span>
                <span className="arrow">➝</span>
              </Link>
              <div className="filter-dropdown">
                <button className="filter-button">
                  FILTRER PAR DÉPARTEMENT
                  <span className="dropdown-arrow">⌄</span>
                </button>
              </div>
            </div>

            <div className="stats-grid">
              <div className="stat-card">
                <div className="stat-number">58</div>
                <div className="stat-label">Enseignants</div>
              </div>
              <div className="stat-card">
                <div className="stat-number">6</div>
                <div className="stat-label">Spécialités</div>
              </div>
              <div className="stat-card">
                <div className="stat-number">20</div>
                <div className="stat-label">Modules</div>
              </div>
            </div>
          </div>

          <div className="action-buttons">
            <Link to="/enseignants" className="btn btn-icon">
              <span className="btn-icon">📚</span>
              <span className="btn-text">VOIR ENSEIGNANTS</span>
            </Link>
            <Link to="/modules" className="btn btn-icon">
              <span className="btn-icon">📘</span>
              <span className="btn-text">VOIR MODULES</span>
            </Link>
            <Link to="/organigramme" className="btn btn-icon">
              <span className="btn-icon">🕘</span>
              <span className="btn-text">VOIR ORGANIGRAMME</span>
            </Link>
            <button className="btn btn-icon">
              <span className="btn-icon">💬</span>
              <span className="btn-text">AJOUTER UN COMMENTAIRE</span>
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default Dashboardorg;