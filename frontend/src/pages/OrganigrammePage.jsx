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
      const user = JSON.parse(storedRole);
      setRole(user.role);
    }
    setActive("S1"); // Default active semester
  }, []);

  const handleGenerate = (semester) => {
    if (semester === "S1") {
      setDataS1(organigrammeS1);
      setIsGeneratedS1(true);
      setIsValidatedS1(false);
      setActive("S1");
    } else {
      setDataS2(organigrammeS2);
      setIsGeneratedS2(true);
      setIsValidatedS2(false);
      setActive("S2");
    }
  };

  const handleValidate = (semester) => {
    if (semester === "S1") {
      setIsValidatedS1(true);
    } else {
      setIsValidatedS2(true);
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
              <button onClick={() => handleGenerate("S1")}>Générer S1</button>
              <button onClick={() => handleGenerate("S2")}>Générer S2</button>
            </div>
          )}

          {isGenerated && !isValidated && (
            <div className="organigramme-validate-buttons">
              <button
                onClick={() => handleValidate(active)}
                style={{ marginRight: "10px" }}
              >
                Valider
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
