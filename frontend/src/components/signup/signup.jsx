import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import PageWrapper from "../PageWrapper";
import Select from "react-select";
import { motion } from "framer-motion";
import "./signup.css";
import logo from "../../assets/eduorg.logo.png";
import illustration from "../../assets/Design sans titre (1).png";

export default function Signup() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [selectedRole, setSelectedRole] = useState(null);
  const [nom, setNom] = useState("");
  const [prenom, setPrenom] = useState("");
  const [telephone, setTelephone] = useState("");
  const [code, setCode] = useState("");

  const roleOptions = [
    { value: "chefDepartement", label: "Chef de département" },
    { value: "staffAdministrateur", label: "Staff administrateur" },
    { value: "enseignant", label: "Enseignant" },
  ];

  const customSelectStyles = {
    control: (base) => ({
      ...base,
      borderRadius: "8px",
      borderColor: "#0056b3",
      padding: "5px",
      boxShadow: "none",
      "&:hover": { borderColor: "#0056b3" },
      fontFamily: "Raleway, sans-serif",
      fontSize: "0.9rem",
    }),
    option: (base, state) => ({
      ...base,
      backgroundColor: state.isFocused ? "#0056b3" : "white",
      color: state.isFocused ? "white" : "black",
      cursor: "pointer",
      fontFamily: "Raleway, sans-serif",
    }),
  };

  const handleLoginClick = () => {
    navigate("/login");
  };

  const handleSignup = async (e) => {
    e.preventDefault();

    if (password !== confirmPassword) {
      alert("Les mots de passe ne correspondent pas !");
      return;
    }

    if (!selectedRole) {
      alert("Veuillez sélectionner un poste.");
      return;
    }

    const userData = {
      nom,
      prenom,
      email,
      password,
      numero: telephone,
      role: selectedRole.value,
      code,
    };

    try {
      const response = await fetch("http://localhost:8000/submit", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      });

      if (!response.ok) {
        const errorText = await response.text();
        alert("Échec de l'inscription : " + errorText);
        return;
      }

      const data = await response.json();
      localStorage.setItem("token", data.token);
      localStorage.setItem("user", JSON.stringify(data.user));

      navigate("/modules");
    } catch (error) {
      console.error("Erreur lors de l'inscription:", error);
      alert("Erreur serveur lors de l'inscription.");
    }
  };

  return (
    <PageWrapper>
      <motion.div
        className="signup-container"
        initial={{ opacity: 0, x: -50 }}
        animate={{ opacity: 1, x: 0 }}
        exit={{ opacity: 0, x: 50 }}
        transition={{ duration: 0.5 }}
      >
        <div className="signup-right">
          <img
            src={illustration}
            alt="Illustration d'inscription"
            className="signup-image"
          />
        </div>

        <div className="signup-left">
          <div
            className="signup-form"
            style={{ maxHeight: "80vh", overflowY: "auto" }}
          >
            <img src={logo} alt="EduOrg Logo" className="signup-logo" />
            <h2 className="signup-title">S'inscrire</h2>

            <form onSubmit={handleSignup}>
              <div className="form-group">
                <label htmlFor="nom">Nom</label>
                <input
                  type="text"
                  id="nom"
                  value={nom}
                  onChange={(e) => setNom(e.target.value)}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="prenom">Prénom</label>
                <input
                  type="text"
                  id="prenom"
                  value={prenom}
                  onChange={(e) => setPrenom(e.target.value)}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="telephone">Téléphone</label>
                <input
                  type="tel"
                  id="telephone"
                  value={telephone}
                  onChange={(e) => setTelephone(e.target.value)}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="email">Adresse email</label>
                <input
                  type="email"
                  id="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="password">Mot de passe</label>
                <input
                  type="password"
                  id="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="confirm-password">
                  Confirmer le mot de passe
                </label>
                <input
                  type="password"
                  id="confirm-password"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="role">Sélectionner un poste</label>
                <Select
                  id="role"
                  options={roleOptions}
                  styles={customSelectStyles}
                  placeholder="Sélectionner un poste"
                  value={selectedRole}
                  onChange={setSelectedRole}
                />
              </div>

              <div className="form-group">
                <label htmlFor="code">Code</label>
                <input
                  type="text"
                  id="code"
                  value={code}
                  onChange={(e) => setCode(e.target.value)}
                  required
                />
              </div>

              <div className="button-group">
                <button type="submit" className="signup-button">
                  S'inscrire
                </button>
                <button
                  type="button"
                  className="login-button"
                  onClick={handleLoginClick}
                >
                  Se connecter
                </button>
              </div>
            </form>
          </div>
        </div>
      </motion.div>
    </PageWrapper>
  );
}
