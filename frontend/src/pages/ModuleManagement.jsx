import React, { useState, useEffect } from 'react';
import Sidebar from '../components/Sidebar';
import ModuleTable from '../components/ModuleTable';
import Pagination from '../components/Pagination';
import ModuleModal from '../components/ModuleModal';
import ExportPopup from '../components/ExportPopup';
import './ModuleManagement.css';
import PopupCommentaire from '../components/PopupCommentaire';
import * as XLSX from 'xlsx';
import jsPDF from 'jspdf';
import 'jspdf-autotable';
import { useNavigate } from 'react-router-dom'; 

const ModuleManagement = () => {
  const [role, setRole] = useState('');
  const navigate = useNavigate();

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
  }, []);

  const [modules, setModules] = useState([
    { nom: 'Programmation web', specialite: 'Informatique', semestre: 'S3', enseignant: 'Sara Bouzid' },
    { nom: 'Probabilités', specialite: 'Mathématiques', semestre: 'S3', enseignant: 'Rami Benaissa' },
    { nom: 'Machine Learning', specialite: 'IA', semestre: 'S4', enseignant: 'Yasmine Armani' },
    { nom: 'Java Avancé', specialite: 'Informatique', semestre: 'S4', enseignant: 'Lina Hadj-Messaoud' },
    { nom: 'Réseaux 2', specialite: 'Réseaux', semestre: 'S3', enseignant: 'Karim Mansour' },
  ]);

  const [selectedModule, setSelectedModule] = useState(null);
  const [searchValue, setSearchValue] = useState('');
  const [isAdding, setIsAdding] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const [showExportPopup, setShowExportPopup] = useState(false);
  const [showCommentPopup, setShowCommentPopup] = useState(false);
  const [commentText, setCommentText] = useState('');

  const itemsPerPage = 3;

  const filteredModules = modules.filter((mod) =>
    mod.nom.toLowerCase().includes(searchValue.toLowerCase())
  );

  const totalPages = Math.ceil(filteredModules.length / itemsPerPage);
  const paginatedModules = filteredModules.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  const handleEdit = (module) => {
    if (role === 'chefDepartement') {
      setSelectedModule(module);
      setIsAdding(false);
    } else {
      alert("Vous n'avez pas la permission de modifier ce module.");
    }
  };

  const handleAdd = () => {
    if (role === 'chefDepartement') {
      setSelectedModule({ nom: '', specialite: '', semestre: '', enseignant: '' });
      setIsAdding(true);
    } else {
      alert("Vous n'avez pas la permission d'ajouter un module.");
    }
  };

  const handleClosePopup = () => {
    setSelectedModule(null);
  };

  const handleSave = (module) => {
    if (isAdding) {
      setModules([...modules, module]);
    } else {
      const updated = modules.map((mod) =>
        mod.nom === selectedModule.nom ? module : mod
      );
      setModules(updated);
    }
    setSelectedModule(null);
  };

  const handleDelete = (moduleToDelete) => {
    if (role === 'chefDepartement') {
      const confirmDelete = window.confirm(
        `Êtes-vous sûr de vouloir supprimer le module "${moduleToDelete.nom}" ?`
      );

      if (confirmDelete) {
        const updated = modules.filter((mod) => mod.nom !== moduleToDelete.nom);
        setModules(updated);
      }
    } else {
      alert("Vous n'avez pas la permission de supprimer ce module.");
    }
  };

  const exportToExcel = () => {
    try {
      // Préparer les données pour l'export
      const exportData = filteredModules.map((module, index) => ({
        'N°': index + 1,
        'Nom du Module': module.nom,
        'Spécialité': module.specialite,
        'Semestre': module.semestre,
        'Enseignant': module.enseignant
      }));

      // Créer un nouveau workbook
      const wb = XLSX.utils.book_new();
      
      // Convertir les données en worksheet
      const ws = XLSX.utils.json_to_sheet(exportData);

      // Ajuster la largeur des colonnes
      const colWidths = [
        { wch: 5 },  // N°
        { wch: 25 }, // Nom du Module
        { wch: 15 }, // Spécialité
        { wch: 10 }, // Semestre
        { wch: 20 }  // Enseignant
      ];
      ws['!cols'] = colWidths;

      // Ajouter le worksheet au workbook
      XLSX.utils.book_append_sheet(wb, ws, 'Modules');

      // Générer le nom de fichier avec la date
      const currentDate = new Date().toLocaleDateString('fr-FR').replace(/\//g, '-');
      const fileName = `modules_${currentDate}.xlsx`;

      // Télécharger le fichier
      XLSX.writeFile(wb, fileName);
      
      console.log('Export Excel réussi');
    } catch (error) {
      console.error('Erreur lors de l\'export Excel:', error);
      alert('Erreur lors de l\'export Excel');
    }
  };

  const exportToPDF = () => {
    try {
      // Créer un nouveau document PDF
      const doc = new jsPDF();

      // Titre du document
      doc.setFontSize(20);
      doc.text('Liste des Modules', 20, 20);

      // Date d'export
      const currentDate = new Date().toLocaleDateString('fr-FR');
      doc.setFontSize(12);
      doc.text(`Date d'export: ${currentDate}`, 20, 35);

      // Préparer les données pour le tableau
      const tableData = filteredModules.map((module, index) => [
        index + 1,
        module.nom,
        module.specialite,
        module.semestre,
        module.enseignant
      ]);

      // Créer le tableau
      doc.autoTable({
        head: [['N°', 'Nom du Module', 'Spécialité', 'Semestre', 'Enseignant']],
        body: tableData,
        startY: 45,
        styles: {
          fontSize: 10,
          cellPadding: 3
        },
        headStyles: {
          fillColor: [41, 128, 185],
          textColor: 255,
          fontStyle: 'bold'
        },
        alternateRowStyles: {
          fillColor: [245, 245, 245]
        },
        columnStyles: {
          0: { halign: 'center', cellWidth: 15 }, // N°
          1: { cellWidth: 50 }, // Nom du Module
          2: { cellWidth: 35 }, // Spécialité
          3: { halign: 'center', cellWidth: 20 }, // Semestre
          4: { cellWidth: 45 } // Enseignant
        }
      });

      // Ajouter un pied de page
      const pageCount = doc.internal.getNumberOfPages();
      for (let i = 1; i <= pageCount; i++) {
        doc.setPage(i);
        doc.setFontSize(8);
        doc.text(
          `Page ${i} sur ${pageCount}`,
          doc.internal.pageSize.width - 30,
          doc.internal.pageSize.height - 10
        );
      }

      // Générer le nom de fichier avec la date
      const fileName = `modules_${currentDate.replace(/\//g, '-')}.pdf`;

      // Télécharger le fichier
      doc.save(fileName);
      
      console.log('Export PDF réussi');
    } catch (error) {
      console.error('Erreur lors de l\'export PDF:', error);
      alert('Erreur lors de l\'export PDF');
    }
  };

  const handleExportClick = () => {
    setShowExportPopup(true);
  };

  const handleExport = (fileType) => {
    setShowExportPopup(false);
    
    switch (fileType.toLowerCase()) {
      case 'excel':
      case 'xlsx':
        exportToExcel();
        break;
      case 'pdf':
        exportToPDF();
        break;
      default:
        console.log(`Format d'export non supporté: ${fileType}`);
        alert(`Format d'export non supporté: ${fileType}`);
    }
  };

  const handleCommentClick = () => {
    setShowCommentPopup(true);
  };

  return (
    <div className="container">
      <Sidebar />
      <main className="main">
        <h2>Gestion des Modules</h2>

        <div className="top-bar">
          <div className="search-wrapper">
            <span className="search-icon">🔍</span>
            <input
              type="text"
              className="search-input"
              placeholder="Rechercher..."
              value={searchValue}
              onChange={(e) => {
                setSearchValue(e.target.value);
                setCurrentPage(1);
              }}
            />
          </div>

          <div className="button-group">
            {role === 'chefDepartement' && (
              <button className="button-module" onClick={handleAdd}>
                AJOUTER UN MODULE
              </button>
            )}

            <button className="button-export" onClick={handleExportClick}>
              EXPORTER LA LISTE
            </button>

            {role === 'staffAdministrateur' && (
              <button
                className="button-comment"
                onClick={handleCommentClick}
              >
                💬 COMMENTAIRES
              </button>
            )}
          </div>
        </div>

        <ModuleTable
          modules={paginatedModules}
          onEdit={handleEdit}
          onDelete={handleDelete}
          role={role}
        />

        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={setCurrentPage}
        />

        {selectedModule && (
          <ModuleModal
            module={selectedModule}
            onSave={handleSave}
            onClose={handleClosePopup}
            isAdding={isAdding}
          />
        )}

        {showExportPopup && (
          <ExportPopup
            onClose={() => setShowExportPopup(false)}
            onExport={handleExport}
          />
        )}

        {showCommentPopup && (
          <PopupCommentaire
            isOpen={showCommentPopup}
            setIsOpen={setShowCommentPopup}
            commentText={commentText}
            setCommentText={setCommentText}
            onSubmit={() => {
              console.log('Commentaire ajouté:', commentText);
              setShowCommentPopup(false);
              setCommentText('');
            }}
          />
        )}
      </main>
    </div>
  );
};

export default ModuleManagement;