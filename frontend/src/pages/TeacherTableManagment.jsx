import React, { useState, useEffect, useRef } from 'react';
import Sidebar from '../components/Sidebar';
import TeacherTable from '../components/TeacherTable';
import Pagination from '../components/Pagination';
import Popup from '../components/Popup';
import ExportPopup from '../components/ExportPopup';
import './TeacherTableManagment.css';
import PopupCommentaire from '../components/PopupCommentaire';
import * as XLSX from 'xlsx';

const TeacherTableManagment = () => {
  const [role, setRole] = useState(''); // état initial vide
  const fileInputRef = useRef(null);

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

  const [teachers, setTeachers] = useState([
    { nom: 'Sara Bouzid', specialite: 'Informatique', semestre: 'S3', module: 'Programmation web' },
    { nom: 'Rami Benaissa', specialite: 'Mathématiques', semestre: 'S3', module: 'Probabilités' },
    { nom: 'Yasmine Armani', specialite: 'IA', semestre: 'S4', module: 'Machine Learning' },
    { nom: 'Lina Hadj-Messaoud', specialite: 'Informatique', semestre: 'S4', module: 'Java Avancé' },
    { nom: 'Karim Mansour', specialite: 'Réseaux', semestre: 'S3', module: 'Réseaux 2' },
  ]);

  const [selectedTeacher, setSelectedTeacher] = useState(null);
  const [searchValue, setSearchValue] = useState('');
  const [isAdding, setIsAdding] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const [showExportPopup, setShowExportPopup] = useState(false);
  const [showCommentPopup, setShowCommentPopup] = useState(false);
  const [commentText, setCommentText] = useState('');
  const [isImporting, setIsImporting] = useState(false);

  const itemsPerPage = 3;

  const filteredTeachers = teachers.filter((t) =>
    t.nom.toLowerCase().includes(searchValue.toLowerCase())
  );

  const totalPages = Math.ceil(filteredTeachers.length / itemsPerPage);
  const paginatedTeachers = filteredTeachers.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  const handleEdit = (teacher) => {
    if (role === 'chefDepartement') {
      setSelectedTeacher(teacher);
      setIsAdding(false);
    } else {
      alert("Vous n'avez pas la permission de modifier cet enseignant.");
    }
  };

  const handleAdd = () => {
    if (role === 'chefDepartement') {
      setSelectedTeacher({ nom: '', specialite: '', semestre: '', module: '' });
      setIsAdding(true);
    } else {
      alert("Vous n'avez pas la permission d'ajouter un enseignant.");
    }
  };

  const handleClosePopup = () => {
    setSelectedTeacher(null);
  };

  const handleSave = (teacher) => {
    if (isAdding) {
      setTeachers([...teachers, teacher]);
    } else {
      const updated = teachers.map((t) =>
        t.nom === selectedTeacher.nom ? teacher : t
      );
      setTeachers(updated);
    }
    setSelectedTeacher(null);
  };

  const handleDelete = (teacherToDelete) => {
    if (role === 'chefDepartement') {
      const confirmDelete = window.confirm(
        `Êtes-vous sûr de vouloir supprimer l'enseignant "${teacherToDelete.nom}" ?`
      );

      if (confirmDelete) {
        const updated = teachers.filter((t) => t.nom !== teacherToDelete.nom);
        setTeachers(updated);
      }
    } else {
      alert("Vous n'avez pas la permission de supprimer cet enseignant.");
    }
  };
  const handleFileChange = (event) => { 
    
  }
  const handleCommentClick = () => {
    setShowCommentPopup(true);
  };

  // Fonction pour déclencher l'import de fichiers
  const handleImportClick = () => {
    if (role === 'chefDepartement') {
      fileInputRef.current?.click();
    } else {
      alert("Vous n'avez pas la permission d'importer des enseignants.");
    }
  };

  // Fonction pour traiter les fichiers Excel
  const exportToExcel = () => {
    try {
      const exportData = filteredTeachers.map((teacher, index) => ({
        'N°': index + 1,
        'Nom enseignant': teacher.nom,
        'Spécialité': teacher.specialite,
        'Semestre': teacher.semestre,
        'Module': teacher.module
      }));
      const wb = XLSX.utils.book_new();
      const ws = XLSX.utils.json_to_sheet(exportData);

      const colWidths = [
        { wch: 5 },  // N°
        { wch: 25 }, // Nom du Module
        { wch: 15 }, // Spécialité
        { wch: 10 }, // Semestre
        { wch: 20 }  // Enseignant
      ];
      ws['!cols'] = colWidths;

      // Ajouter le worksheet au workbook
      XLSX.utils.book_append_sheet(wb, ws, 'teachers');

      // Générer le nom de fichier avec la date
      const currentDate = new Date().toLocaleDateString('fr-FR').replace(/\//g, '-');
      const fileName = `teachers_${currentDate}.xlsx`;

      // Télécharger le fichier
      XLSX.writeFile(wb, fileName);
      
      console.log('Export Excel réussi');
    } catch (error) {
      console.error('Erreur lors de l\'export Excel:', error);
      alert('Erreur lors de l\'export Excel');
    }
  };

  // Fonction pour traiter les fichiers PDF (extraction de texte basique)
  const exportToPDF = () => {
    try {
      // Créer un nouveau document PDF
      const doc = new jsPDF();

      // Titre du document
      doc.setFontSize(20);
      doc.text('Liste de teacher', 20, 20);

      // Date d'export
      const currentDate = new Date().toLocaleDateString('fr-FR');
      doc.setFontSize(12);
      doc.text(`Date d'export: ${currentDate}`, 20, 35);

      // Préparer les données pour le tableau
      const tableData = filteredTeachers.map((teacher, index) => ({
        'N°': index + 1,
        'Nom enseignant': teacher.nom,
        'Spécialité': teacher.specialite,
        'Semestre': teacher.semestre,
        'Module': teacher.module
      }));

      // Créer le tableau
      doc.autoTable({
        head: [['N°', 'Nom enseignant', 'Spécialité', 'Semestre', 'Module']],
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
      const fileName = `teachers_${currentDate.replace(/\//g, '-')}.pdf`;

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

  return (
    <div className="container">
      <Sidebar />
      <main className="main">
        <h2>Gestion des Enseignants</h2>

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
              <>
                <button className="button-module" onClick={handleAdd}>
                  AJOUTER UN ENSEIGNANT
                </button>
                
                <button 
                  className="button-import" 
                  onClick={handleImportClick}
                  disabled={isImporting}
                >
                  {isImporting ? 'IMPORTATION...' : '📁 IMPORTER'}
                </button>
              </>
            )}

            <button className="button-export" onClick={handleExportClick}>
              EXPORTER LA LISTE
            </button>

            {role === 'staffAdministrateur' && (
              <button className="button-comment" onClick={handleCommentClick}>
                💬 COMMENTAIRES
              </button>
            )}
          </div>
        </div>

        {/* Input file caché pour l'import */}
        <input
          type="file"
          ref={fileInputRef}
          onChange={handleFileChange}
          accept=".xlsx,.xls,.pdf"
          style={{ display: 'none' }}
        />

        <TeacherTable
          teachers={paginatedTeachers}
          onEdit={handleEdit}
          onDelete={handleDelete}
          role={role}
        />

        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={setCurrentPage}
        />

        {selectedTeacher && (
          <Popup
            module={selectedTeacher}
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
              onSubmit={async () => {
                try {
                  const response = await fetch('/staff/commentaire', {
                    method: 'POST',
                    headers: {
                      'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ message: commentText }),
                  });

                  if (!response.ok) {
                    throw new Error('Erreur lors de l’envoi du commentaire');
                  }

                  alert('Commentaire envoyé avec succès.');
                  setShowCommentPopup(false);
                  setCommentText('');
                } catch (error) {
                  console.error('Erreur:', error);
                  alert('Échec de l’envoi du commentaire');
                }
            }}
          />
          )}

      </main>
    </div>
  );
};

export default TeacherTableManagment;