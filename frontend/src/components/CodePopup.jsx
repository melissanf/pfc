import React, { useState, useEffect } from "react";

const CodePopup = ({ onSave, onClose }) => {
  const [codeData, setCodeData] = useState({ code: "", role: "enseignant" });
  const [isFormValid, setIsFormValid] = useState(true);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCodeData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const isValid = codeData.code.trim() !== "" && codeData.role.trim() !== "";
    setIsFormValid(isValid);

    if (isValid) {
      onSave(codeData);
    } else {
      alert("Tous les champs doivent être remplis.");
    }
  };

  const styles = {
    overlay: {
      position: "fixed",
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      backgroundColor: "rgba(0, 0, 0, 0.5)",
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
      zIndex: 1000,
    },
    container: {
      backgroundColor: "#fff",
      padding: "20px",
      borderRadius: "8px",
      width: "350px",
      boxShadow: "0 4px 8px rgba(0, 0, 0, 0.2)",
    },
    title: {
      marginBottom: "20px",
      textAlign: "center",
    },
    formField: {
      marginBottom: "15px",
    },
    label: {
      display: "block",
      marginBottom: "5px",
      fontWeight: "bold",
    },
    input: {
      width: "100%",
      padding: "8px",
      borderRadius: "4px",
      border: "1px solid #ccc",
    },
    select: {
      width: "100%",
      padding: "8px",
      borderRadius: "4px",
      border: "1px solid #ccc",
    },
    inputError: {
      border: "1px solid red",
    },
    buttons: {
      display: "flex",
      justifyContent: "space-between",
    },
    button: {
      padding: "10px 15px",
      border: "none",
      borderRadius: "4px",
      cursor: "pointer",
    },
    saveButton: {
      backgroundColor: "#4CAF50",
      color: "#fff",
    },
    closeButton: {
      backgroundColor: "#f44336",
      color: "#fff",
    },
    errorMessage: {
      color: "red",
      fontSize: "14px",
      marginBottom: "10px",
    },
  };

  return (
    <div style={styles.overlay}>
      <div style={styles.container}>
        <h3 style={styles.title}>Créer un nouveau code</h3>
        <form onSubmit={handleSubmit}>
          <div style={styles.formField}>
            <label style={styles.label}>Code:</label>
            <input
              type="text"
              name="code"
              value={codeData.code}
              onChange={handleChange}
              style={{
                ...styles.input,
                ...(isFormValid || codeData.code ? {} : styles.inputError),
              }}
              required
            />
          </div>

          <div style={styles.formField}>
            <label style={styles.label}>Rôle:</label>
            <select
              name="role"
              value={codeData.role}
              onChange={handleChange}
              style={styles.select}
            >
              <option value="enseignant">Enseignant</option>
              <option value="staffAdministrateur">Staff Administrateur</option>
              <option value="chefDepartement">Chef de Département</option>
            </select>
          </div>

          {!isFormValid && (
            <p style={styles.errorMessage}>
              Tous les champs sont obligatoires.
            </p>
          )}

          <div style={styles.buttons}>
            <button
              type="submit"
              style={{ ...styles.button, ...styles.saveButton }}
            >
              Sauvegarder
            </button>
            <button
              type="button"
              onClick={onClose}
              style={{ ...styles.button, ...styles.closeButton }}
            >
              Fermer
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CodePopup;
