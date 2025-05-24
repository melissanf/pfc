package utils

import (
	"strings"

	"github.com/melissanf/pfc/backend/internal/api/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"errors"
	"sort"
	"time"
)
type SessionType string

const (
	Cours SessionType = "cours"
	TD    SessionType = "TD"
	TP    SessionType = "TP"
)

type ModuleInfo struct {
	Module string
	Niveau string
	Types  map[SessionType]bool
}

type TeacherModules map[string][]ModuleInfo

func loadOrga(filePath string) (TeacherModules, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	// Suppose la feuille s'appelle "Feuille1"
	sheetName := f.GetSheetName(0)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	// Indices fixés pour les colonnes (selon ta structure)
	const (
		colNiveau = 0
		colModule = 1
		colCours  = 2
		colTD1    = 3
		colTD4    = 6
		colTP1    = 7
		colTP4    = 10
	)

	teacherMap := make(TeacherModules)

	for i, row := range rows {
		if i == 0 || len(row) < 3 { // Ignorer l'en-tête ou lignes incomplètes
			continue
		}

		niveau := row[colNiveau]
		module := row[colModule]

		// Cours
		if len(row) > colCours {
			name := strings.TrimSpace(row[colCours])
			if name != "" {
				addModule(teacherMap, name, niveau, module, Cours)
			}
		}

		// TD1 à TD4
		for i := colTD1; i <= colTD4 && i < len(row); i++ {
			name := strings.TrimSpace(row[i])
			if name != "" {
				addModule(teacherMap, name, niveau, module, TD)
			}
		}

		// TP1 à TP4
		for i := colTP1; i <= colTP4 && i < len(row); i++ {
			name := strings.TrimSpace(row[i])
			if name != "" {
				addModule(teacherMap, name, niveau, module, TP)
			}
		}
	}

	return teacherMap, nil
}

func addModule(m TeacherModules, teacher, niveau, module string, t SessionType) {
	if _, ok := m[teacher]; !ok {
		m[teacher] = []ModuleInfo{}
	}

	// Vérifier si le module existe déjà
	for i, mod := range m[teacher] {
		if mod.Module == module && mod.Niveau == niveau {
			m[teacher][i].Types[t] = true
			return
		}
	}

	// Sinon, ajouter nouveau module
	m[teacher] = append(m[teacher], ModuleInfo{
		Module: module,
		Niveau: niveau,
		Types:  map[SessionType]bool{t: true},
	})
}
func loadVoeuxFromDB(db *gorm.DB) ([]models.Voeux, error) {
	var voeux []models.Voeux
	err := db.Preload("Module").Preload("Niveau").Find(&voeux).Error
	if err != nil {
		return nil, err
	}
	return voeux, nil
}

func GetTeacherInfo(db *gorm.DB, teacherID uint) (*models.Teacher, error) {
	var t models.Teacher
	err := db.First(&t, teacherID).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ResolveConflit(db *gorm.DB, candidats []models.Voeux, moduleSpeciality string) (*models.Voeux, error) {
	if len(candidats) == 0 {
		return nil, errors.New("aucun candidat")
	}
	type ScoredVoeu struct {
		Voeu  models.Voeux
		Note int
	}

	var scored []ScoredVoeu
	for _, v := range candidats {
		teacher, err := GetTeacherInfo(db, v.TeacherID)
		if err != nil {
			continue
		}
		note := 0
		i:=0
		if ok, err :=TeacherHasSpeciality(db, *teacher, moduleSpeciality) ; err == nil && ok {
			note += 10
		} else 
		{
			for _, sp := range teacher.Specialities {
				if isSpecialityClose(sp.Nom, moduleSpeciality) {	
					i++;
			    }
		    }
			if i > 0 {
			   note += 5
			}
	    }
		switch teacher.Grade {
		case "Professeur":
			note += 5
		case "Maitre de conf":
			note += 3
		case "Doctorant":
			note += 1
		}
		years := time.Now().Year() - teacher.Year_entrance
		note += years

		scored = append(scored, ScoredVoeu{Voeu: v, Note: note})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Note > scored[j].Note
	})

	if len(scored) == 0 {
		return nil, errors.New("aucun prof valide")
	}

	return &scored[0].Voeu, nil
}
func TeacherHasSpeciality(db *gorm.DB, teacher models.Teacher, specialityName string) (bool, error) {
	var count int64

	err := db.Model(&models.TeacherSpeciality{}).
		Joins("JOIN specialities ON specialities.id = teacher_specialities.speciality_id").
		Where("teacher_specialities.teacher_id = ? AND specialities.name = ?", teacher.ID, specialityName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func isSpecialityClose(sp1, sp2 string) bool {
     return true
}

