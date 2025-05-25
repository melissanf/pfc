import React, { useState, useEffect } from "react";
import "./Profil.css";
import Sidebar from "../components/Sidebar";

const Profil = () => {
  const [user, setUser] = useState({
    nom: "",
    prenom: "",
    email: "",
    numero: "",
    departement: "",
    statut: "",
    specialite: "",
    formation: "",
  });

  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      const parsedUser = JSON.parse(storedUser);
      setUser((prev) => ({ ...prev, ...parsedUser }));
    }
  }, []);

  const labelStyle = {
    display: "block",
    fontSize: "0.85rem",
    marginBottom: "4px",
    fontWeight: 600,
    color: "#333",
  };

  const inputStyle = {
    width: "100%",
    marginBottom: "16px",
    padding: "8px",
    boxSizing: "border-box",
  };

  return (
    <div className="app-container">
      <Sidebar />

      <main className="profile-form">
        <h1>Mon Profil</h1>
        <form>
          <div className="form-row" style={{ display: "flex", gap: "16px" }}>
            <div style={{ flex: 1 }}>
              <label style={labelStyle} htmlFor="nom">
                Nom
              </label>
              <input
                id="nom"
                type="text"
                placeholder="Votre nom"
                value={user.nom}
                style={inputStyle}
                readOnly
              />
            </div>
            <div style={{ flex: 1 }}>
              <label style={labelStyle} htmlFor="prenom">
                Prénom
              </label>
              <input
                id="prenom"
                type="text"
                placeholder="Votre prénom"
                value={user.prenom}
                style={inputStyle}
                readOnly
              />
            </div>
          </div>

          <div className="form-row" style={{ display: "flex", gap: "16px" }}>
            <div style={{ flex: 1 }}>
              <label style={labelStyle} htmlFor="email">
                Email
              </label>
              <input
                id="email"
                type="email"
                placeholder="votre.email@exemple.com"
                value={user.email}
                style={inputStyle}
                readOnly
              />
            </div>
            <div style={{ flex: 1 }}>
              <label style={labelStyle} htmlFor="numero">
                Téléphone
              </label>
              <input
                id="numero"
                type="tel"
                placeholder="+213 6 XX XX XX XX"
                value={user.numero}
                style={inputStyle}
                readOnly
              />
            </div>
          </div>

          <div className="form-row" style={{ display: "flex", gap: "16px" }}>
            <div style={{ flex: 1 }}>
              <label style={labelStyle} htmlFor="departement">
                Département
              </label>
              <select
                id="departement"
                value={user.departement}
                style={inputStyle}
                onChange={(e) =>
                  setUser((prev) => ({ ...prev, departement: e.target.value }))
                }
              >
                <option value="">Sélectionnez un département</option>
                <option value="SIQ">
                  Département des Systèmes Informatiques (SIQ)
                </option>
                <option value="AI">
                  Département de AI/ Science des Données
                </option>
              </select>
            </div>
            <div style={{ flex: 1 }}>
              <label style={labelStyle} htmlFor="departement">
                grade
              </label>
              <select
                id="statut"
                value={user.statut}
                style={inputStyle}
                onChange={(e) =>
                  setUser((prev) => ({ ...prev, statut: e.target.value }))
                }
              >
                <option value="">Sélectionnez un statut</option>
                <option value="maitreA">Maître Assistant Classe B</option>
                <option value="maitreB">Maître Assistant Classe A</option>
                <option value="confB">Maître de Conférences Classe B</option>
                <option value="confA">Maître de Conférences Classe A</option>
                <option value="prof">Professeur (PES)</option>
              </select>
            </div>
          </div>

          <div style={{ marginBottom: "16px" }}>
            <label style={labelStyle} htmlFor="specialite">
              Domaine de spécialité
            </label>
            <input
              id="specialite"
              type="text"
              placeholder="Votre domaine de spécialité"
              value={user.specialite}
              style={inputStyle}
            />
          </div>

          <div style={{ marginBottom: "16px" }}>
            <label style={labelStyle} htmlFor="formation">
              Formation
            </label>
            <textarea
              id="formation"
              placeholder="Formation"
              value={user.formation}
              style={{ ...inputStyle, height: "80px" }}
              readOnly
            />
          </div>

          <button
            style={{ width: "fit-content", padding: "10px" }}
            type="submit"
            disabled
          >
            Sauvegarder
          </button>
        </form>
      </main>
    </div>
  );
};

export default Profil;
