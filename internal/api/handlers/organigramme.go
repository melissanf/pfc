package handlers

import (
	"Devenir_dev/internal/api/models"
	utils "Devenir_dev/pkg"
	"encoding/json"
	"net/http"
)

// Input structure for organigramme processing
type OrganigramInput struct {
	TimeSlots        []models.OrganigramTemplate `json:"time_slots"`
	TeacherWishes    []models.TeacherWish        `json:"teacher_wishes"`
	AvailableModules []models.Module             `json:"available_modules"`
	Teachers         []models.Teacher            `json:"teachers"`
}

// Output structure for assignments
type OrganigramOutput struct {
	Assignments []models.Assignment `json:"assignments"`
}

func Orga(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var input OrganigramInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		http.Error(res, "Entrée invalide", http.StatusBadRequest)
		return
	}

	// Map to track teacher hours
	teacherHours := make(map[int]int)
	var output OrganigramOutput

	for _, slot := range input.TimeSlots {
		teacher := utils.FindTeacher(slot.TeacherID, input.Teachers)
		if teacher == nil {
			continue
		}

		// Find matching module based on teacher wishes and availability
		module := utils.FindModuleForTeacher(teacher.ID, slot.Type, input.TeacherWishes, input.AvailableModules, teacherHours[teacher.ID])
		if module != nil {
			hours := utils.GetHoursForType(module, slot.Type)
			if hours > 0 {
				teacherHours[teacher.ID] += hours
				output.Assignments = append(output.Assignments, models.Assignment{
					TeacherID:     teacher.ID,
					ModuleID:      module.ID,
					Type:          slot.Type,
					HoursAssigned: hours,
				})
				// Remove assigned module from available modules
				input.AvailableModules = utils.RemoveModule(module.ID, input.AvailableModules)
			}
		}
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(output)
}
