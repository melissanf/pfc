import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import PageWrapper from "../PageWrapper";
import { motion } from "framer-motion";
import "./login.css";
import logo from "../../assets/eduorg.logo.png";
import illustration from "../../assets/Design sans titre (1).png";

export default function LoginPage() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [code, setCode] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const handleLoginSubmit = async (e) => {
    e.preventDefault();
    setErrorMessage("");

    if (!code) {
      setErrorMessage("Veuillez entrer le code.");
      return;
    }

    try {
      const backendUrl =
        process.env.REACT_APP_BACKEND_URL || "http://localhost:8000/";

      const response = await fetch(`${backendUrl}login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
          code,
        }),
      });

      if (!response.ok) {
        const data = await response.json().catch(() => ({}));
        setErrorMessage(
          data.message || "Email, mot de passe ou code invalide."
        );
        return;
      }

      const data = await response.json();
      const token = data.token;

      if (!token) {
        setErrorMessage("Erreur lors de la connexion. Token manquant.");
        return;
      }

      localStorage.setItem("token", token);
      localStorage.setItem("user", JSON.stringify(data.user));

      navigate("/modules");
    } catch (error) {
      setErrorMessage("Erreur de connexion. Veuillez réessayer.");
      console.error("Login error:", error);
    }
  };

  return (
    <PageWrapper>
      <motion.div
        className="login-container"
        initial={{ opacity: 0, x: 50 }}
        animate={{ opacity: 1, x: 0 }}
        exit={{ opacity: 0, x: -50 }}
        transition={{ duration: 0.5 }}
      >
        <div className="login-left">
          <form className="login-form" onSubmit={handleLoginSubmit}>
            <img src={logo} alt="EduOrg Logo" className="login-logo" />
            <h2 className="login-title">Se connecter</h2>

            {errorMessage && (
              <div
                className="error-message"
                style={{ color: "red", marginBottom: "1rem" }}
              >
                {errorMessage}
              </div>
            )}

            <div className="form-group">
              <label
                htmlFor="email"
                className="form-label"
                style={{
                  fontSize: "0.8rem",
                  display: "block",
                  marginBottom: 4,
                }}
              >
                Adresse email
              </label>
              <input
                type="email"
                id="email"
                placeholder="Entrer votre adresse email"
                className="form-input"
                required
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>

            <div className="form-group">
              <label
                htmlFor="password"
                className="form-label"
                style={{
                  fontSize: "0.8rem",
                  display: "block",
                  marginBottom: 4,
                }}
              >
                Mot de passe
              </label>
              <input
                type="password"
                id="password"
                placeholder="Entrer votre mot de passe"
                className="form-input"
                required
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>

            <div className="form-group">
              <label
                htmlFor="code"
                className="form-label"
                style={{
                  fontSize: "0.8rem",
                  display: "block",
                  marginBottom: 4,
                }}
              >
                Code
              </label>
              <input
                type="text"
                id="code"
                placeholder="Entrer le code"
                className="form-input"
                required
                value={code}
                onChange={(e) => setCode(e.target.value)}
              />
            </div>

            <div className="checkbox-container custom-checkbox">
              <input type="checkbox" id="remember" />
              <label htmlFor="remember">Se souvenir de moi</label>
            </div>

            <div className="button-group">
              <button type="submit" className="login-button blue-button">
                Se connecter
              </button>
              <button
                type="button"
                className="signup-button"
                onClick={() => navigate("/signup")}
              >
                S'inscrire
              </button>
            </div>
          </form>
        </div>

        <div className="login-right">
          <img
            src={illustration}
            alt="Illustration de connexion"
            className="login-image"
          />
        </div>
      </motion.div>
    </PageWrapper>
  );
}
