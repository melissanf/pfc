import React, { useState, useEffect } from "react";
import Sidebar from "../components/Sidebar";
import "./codesPage.css";
import CodePopup from "../components/CodePopup";

const CodesPage = () => {
  const [role, setRole] = useState("");
  const [codes, setCodes] = useState([
    { code: "EMB0001", role: "enseignant" },
    { code: "PMN003", role: "staffAdministrateur" },
    { code: "DFM000", role: "chefDepartement" },
  ]);

  const [showPopup, setShowPopup] = useState(false);

  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      const user = JSON.parse(storedUser);
      setRole(user.role);
    }
  }, []);

  const handleCreateClick = () => {
    setShowPopup(true);
  };

  const handleImport = () => {
    alert("Fonction d'importation à implémenter");
  };

  const handleSaveCode = (newCodeData) => {
    setCodes((prev) => [...prev, newCodeData]);
    setShowPopup(false);
  };

  return (
    <div className="codes-page">
      <div className="codes-sidebar">
        <Sidebar />
      </div>

      <div className="codes-content">
        <h1>Codes d’Accès</h1>

        <div className="codes-buttons">
          <button onClick={handleCreateClick}>Créer un code</button>
          <button onClick={handleImport}>Importer</button>
        </div>

        {showPopup && (
          <CodePopup
            onSave={handleSaveCode}
            onClose={() => setShowPopup(false)}
          />
        )}

        <table className="codes-table">
          <thead>
            <tr>
              <th>Code</th>
              <th>Rôle</th>
            </tr>
          </thead>
          <tbody>
            {codes.map((item, index) => (
              <tr key={index}>
                <td>{item.code}</td>
                <td>{item.role}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default CodesPage;
