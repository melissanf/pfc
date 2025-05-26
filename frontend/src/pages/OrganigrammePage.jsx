import React, { useState, useEffect } from "react";
import Sidebar from "../components/Sidebar";
import OrganigrammeTable from "../components/OrganigrammeTable";
import PopupCommentaire from "../components/PopupCommentaire";
import ExportPopup from "../components/ExportPopup";
import PopupEchange from "../components/PopupEchange";
import organigrammeS1 from "../data/OrganigrammeS1.json";
import organigrammeS2 from "../data/OrganigrammeS2.json";
import "./OrganigrammePage.css";

import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import jsPDF from "jspdf";
import autoTable from "jspdf-autotable";

const OrganigrammePage = () => {
  const [active, setActive] = useState("");
  const [role, setRole] = useState("");
  const [showCommentPopup, setShowCommentPopup] = useState(false);
  const [showExportPopup, setShowExportPopup] = useState(false);
  const [showEchangePopup, setShowEchangePopup] = useState(false);
  const [commentText, setCommentText] = useState("");
  const [ligneToEdit, setLigneToEdit] = useState(null);
  const [dataS1, setDataS1] = useState([]);
  const [dataS2, setDataS2] = useState([]);

  const [isGeneratedS1, setIsGeneratedS1] = useState(false);
  const [isGeneratedS2, setIsGeneratedS2] = useState(false);
  const [isValidatedS1, setIsValidatedS1] = useState(false);
  const [isValidatedS2, setIsValidatedS2] = useState(false);

  useEffect(() => {
    const storedRole = localStorage.getItem("user") || "chefDepartement";
    if (storedRole) {
      try {
        const user = JSON.parse(storedRole);
        setRole(user.role);
      } catch (error) {
        console.error("Erreur parsing user:", error);
        setRole("chefDepartement"); // fallback
      }
    }
    setActive("S1"); // Default active semester
  }, []);

  // Nouvelle fonction pour envoyer des notifications à TOUS les utilisateurs
  const sendNotificationToAllUsers = async (semester) => {
    try {
      console.log(`Envoi de notification à tous les utilisateurs pour ${semester}`);

      // Obtenir le token d'authentification
      const token = localStorage.getItem('authToken') || 
                    localStorage.getItem('token') || 
                    localStorage.getItem('accessToken');

      if (!token) {
        console.error("Token d'authentification manquant");
        alert("Erreur: Token d'authentification manquant. Veuillez vous reconnecter.");
        return false;
      }

      // Envoyer une notification générale à tous les utilisateurs
      const response = await fetch("http://localhost:8000/Enseignant/Notif", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({
          notify_all_users: true,
          semestre: semester,
          type: "organigramme_valide",
          titre: `Organigramme ${semester} validé`,
          message: `L'organigramme du semestre ${semester} a été validé par le chef de département. Consultez la nouvelle planification des cours.`,
        }),
      });

      if (!response.ok) {
        const errorData = await response.text();
        console.error("Erreur lors de l'envoi des notifications:", errorData);
        alert("Erreur lors de l'envoi des notifications aux utilisateurs.");
        return false;
      }

      const result = await response.json();
      console.log("Notifications envoyées avec succès:", result);
      alert(`Notification envoyée avec succès à ${result.created_count} utilisateurs.`);
      return true;

    } catch (error) {
      console.error("Erreur réseau lors de l'envoi des notifications:", error);
      alert("Erreur réseau lors de l'envoi des notifications.");
      return false;
    }
  };

  // Fonction pour envoyer des notifications spécifiques aux professeurs (optionnelle)
  const sendNotificationsToTeachers = async (semester, organigrammeData) => {
    try {
      // Extraire les professeurs uniques de l'organigramme
      const professeurs = [...new Set(
        organigrammeData
          .map(ligne => ligne.enseignant || ligne.professeur || ligne.teacher)
          .filter(prof => prof && prof.trim() !== "")
      )];

      console.log(`Envoi de notifications spécifiques à ${professeurs.length} professeurs pour ${semester}`);

      // Préparer les notifications spécifiques
      const notifications = professeurs.map(professeur => ({
        destinataire: professeur,
        type: "organigramme_attribution",
        titre: `Vos attributions ${semester}`,
        message: `Vous avez des cours attribués dans l'organigramme du semestre ${semester}. Consultez vos nouvelles attributions d'enseignement.`,
        semestre: semester,
        date_creation: new Date().toISOString(),
        lu: false
      }));

      // Obtenir le token d'authentification
      const token = localStorage.getItem('authToken') || 
                    localStorage.getItem('token') || 
                    localStorage.getItem('accessToken');

      if (!token) {
        console.error("Token d'authentification manquant");
        return false;
      }

      // Envoyer les notifications via l'API
      const response = await fetch("http://localhost:8000/Enseignant/Notif", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({
          notify_all_users: false,
          notifications: notifications
        }),
      });

      if (!response.ok) {
        const errorData = await response.text();
        console.error("Erreur lors de l'envoi des notifications spécifiques:", errorData);
        return false;
      }

      const result = await response.json();
      console.log("Notifications spécifiques envoyées avec succès:", result);
      return true;

    } catch (error) {
      console.error("Erreur réseau lors de l'envoi des notifications spécifiques:", error);
      return false;
    }
  };

  const handleGenerate = (semester) => {
    if (semester === "S1") {
      setDataS1(organigrammeS1);
      setIsGeneratedS1(true);
      setActive("S1");
    } else {
      setDataS2(organigrammeS2);
      setIsGeneratedS2(true);
      setActive("S2");
    }
  };

  const handleValidate = async (semester) => {
    const currentData = semester === "S1" ? dataS1 : dataS2;
    
    // Confirmer la validation
    const confirmValidation = window.confirm(
      `Êtes-vous sûr de vouloir valider l'organigramme ${semester} ? Cette action enverra une notification à TOUS les utilisateurs de l'application.`
    );
    
    if (!confirmValidation) {
      return;
    }

    // Marquer comme validé
    if (semester === "S1") {
      setIsValidatedS1(true);
    } else {
      setIsValidatedS2(true);
    }

    // Envoyer la notification générale à tous les utilisateurs
    const generalNotificationSent = await sendNotificationToAllUsers(semester);
    
    // Optionnel : Envoyer aussi des notifications spécifiques aux professeurs concernés
    if (generalNotificationSent) {
      await sendNotificationsToTeachers(semester, currentData);
    }
  };

  const handleCancelValidation = (semester) => {
    if (semester === "S1") {
      setDataS1([]);
      setIsGeneratedS1(false);
      setIsValidatedS1(false);
    } else {
      setDataS2([]);
      setIsGeneratedS2(false);
      setIsValidatedS2(false);
    }
  };

  // Nouvelle fonction pour regénérer un semestre (réinitialise la validation)
  const handleRegenerate = (semester) => {
    if (semester === "S1") {
      setDataS1(organigrammeS1);
      setIsGeneratedS1(true);
      setIsValidatedS1(false); // Réinitialise la validation lors de la regénération
      setActive("S1");
    } else {
      setDataS2(organigrammeS2);
      setIsGeneratedS2(true);
      setIsValidatedS2(false); // Réinitialise la validation lors de la regénération
      setActive("S2");
    }
  };

  const handleExportClick = () => {
    setShowExportPopup(true);
  };

  const exportToExcel = (data, fileName = "organigramme") => {
    const worksheet = XLSX.utils.json_to_sheet(data);
    const workbook = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(workbook, worksheet, "Sheet1");
    const excelBuffer = XLSX.write(workbook, {
      bookType: "xlsx",
      type: "array",
    });
    const blob = new Blob([excelBuffer], {
      type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8",
    });
    saveAs(blob, `${fileName}.xlsx`);
  };

  const exportToPDF = (data, fileName = "organigramme") => {
    const doc = new jsPDF();
    const headers = Object.keys(data[0] || {});
    const rows = data.map((item) => headers.map((key) => item[key]));
    autoTable(doc, {
      head: [headers],
      body: rows,
    });
    doc.save(`${fileName}.pdf`);
  };

  const handleDelete = (ligneToDelete) => {
    if (active === "S1") {
      setDataS1((prev) => prev.filter((ligne) => ligne !== ligneToDelete));
    } else {
      setDataS2((prev) => prev.filter((ligne) => ligne !== ligneToDelete));
    }
  };

  const handleEdit = (ligne) => {
    setLigneToEdit(ligne);
    setShowEchangePopup(true);
  };

  const handleEchangeSubmit = (nouvelleLigne) => {
    if (active === "S1") {
      setDataS1((prevData) =>
        prevData.map((ligne) => (ligne === ligneToEdit ? nouvelleLigne : ligne))
      );
    } else {
      setDataS2((prevData) =>
        prevData.map((ligne) => (ligne === ligneToEdit ? nouvelleLigne : ligne))
      );
    }
    setShowEchangePopup(false);
    setLigneToEdit(null);
  };

  const isGenerated = active === "S1" ? isGeneratedS1 : isGeneratedS2;
  const isValidated = active === "S1" ? isValidatedS1 : isValidatedS2;
  const currentData = active === "S1" ? dataS1 : dataS2;

  return (
    <div className="organigramme-page">
      <div className="organigramme-sidebar">
        <Sidebar />
      </div>

      <div className="organigramme-content">
        <h1 className="h1-organigramme">Organigramme</h1>

        <div className="organigramme-top-bar">
          {role === "chefDepartement" && (
            <div className="organigramme-buttons">
              <button 
                onClick={() => handleGenerate("S1")}
                disabled={isGeneratedS1 && isValidatedS1}
              >
                {isGeneratedS1 ? (isValidatedS1 ? "S1 Validé" : "Afficher S1") : "Générer S1"}
              </button>
              <button 
                onClick={() => handleGenerate("S2")}
                disabled={isGeneratedS2 && isValidatedS2}
              >
                {isGeneratedS2 ? (isValidatedS2 ? "S2 Validé" : "Afficher S2") : "Générer S2"}
              </button>
              
              {/* Boutons pour regénérer si déjà validé */}
              {isValidatedS1 && (
                <button 
                  onClick={() => handleRegenerate("S1")}
                  className="btn-regenerate"
                >
                  Regénérer S1
                </button>
              )}
              {isValidatedS2 && (
                <button 
                  onClick={() => handleRegenerate("S2")}
                  className="btn-regenerate"
                >
                  Regénérer S2
                </button>
              )}
            </div>
          )}

          {isGenerated && !isValidated && (
            <div className="organigramme-validate-buttons">
              <button
                onClick={() => handleValidate(active)}
                style={{ marginRight: "10px" }}
                title={`Valider et notifier tous les utilisateurs de la validation de ${active}`}
              >
                Valider {active} (Notifier tous)
              </button>
              <button onClick={() => handleCancelValidation(active)}>
                Ne pas valider
              </button>
            </div>
          )}

          <div className="organigramme-tools-buttons">
            {role === "chefDepartement" && isGenerated && (
              <button className="btn-exporter" onClick={handleExportClick}>
                Exporter
              </button>
            )}
            {role === "staffAdministrateur" && (
              <button
                className="btn-commenter"
                onClick={() => setShowCommentPopup(true)}
              >
                💬 Commenter
              </button>
            )}
          </div>
        </div>

        {isGenerated && (
          <OrganigrammeTable
            data={currentData}
            title={`Organigramme ${active}`}
            role={role}
            onEdit={isValidated ? null : handleEdit}
            onDelete={isValidated ? null : handleDelete}
            isEditable={!isValidated}
          />
        )}

        {showCommentPopup && (
          <PopupCommentaire
            isOpen={showCommentPopup}
            setIsOpen={setShowCommentPopup}
            commentText={commentText}
            setCommentText={setCommentText}
            onSubmit={() => {
              console.log("Commentaire ajouté:", commentText);
              setShowCommentPopup(false);
              setCommentText("");
            }}
          />
        )}

        {showExportPopup && (
          <ExportPopup
            onClose={() => setShowExportPopup(false)}
            onExport={(type) => {
              if (type === "excel")
                exportToExcel(currentData, `Organigramme_${active}`);
              else if (type === "pdf")
                exportToPDF(currentData, `Organigramme_${active}`);
              setShowExportPopup(false);
            }}
          />
        )}

        {showEchangePopup &&
          ligneToEdit &&
          role === "chefDepartement" &&
          !isValidated && (
            <PopupEchange
              ligne={ligneToEdit}
              onClose={() => setShowEchangePopup(false)}
              onSubmit={handleEchangeSubmit}
            />
          )}
      </div>
    </div>
  );
};

export default OrganigrammePage;