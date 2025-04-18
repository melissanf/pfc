package models

import "time"

// ===================== USERS =====================
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email       string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	FullName     string `json:"full_name"`
}

// ===================== TEACHERS =====================
type Teacher struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SpecialtyID int    `json:"specialty_id"`
	MaxHours    int    `json:"max_hours"`
	CurrentHours int   `json:"current_hours"`
}

type TeacherProfile struct {
	TeacherID       int    `json:"teacher_id"`
	Degrees         string `json:"degrees"`
	YearsOfExperience int  `json:"years_of_experience"`
	DepartmentID    int    `json:"department_id"`
	Bio             string `json:"bio"`
	DomainOfInterest string `json:"domain_of_interest"`
}

type TeacherWish struct {
	TeacherID int `json:"teacher_id"`
	ModuleID  int `json:"module_id"`
	Priority  int `json:"priority"`
	WantsCours bool `json:"wants_cours"`
	WantsTD    bool `json:"wants_td"`
	WantsTP    bool `json:"wants_tp"`
}

type TeacherNote struct {
	ID        int       `json:"id"`
	TeacherID int       `json:"teacher_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// ===================== MODULES =====================
type Module struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Semester  string `json:"semester"`
	Niveau    string `json:"niveau"`
	VolumeCours int  `json:"volume_cours"`
	VolumeTD    int  `json:"volume_td"`
	VolumeTP    int  `json:"volume_tp"`
}

// ===================== ASSIGNMENTS =====================
type Assignment struct {
	TeacherID    int    `json:"teacher_id"`
	ModuleID     int    `json:"module_id"`
	Type         string `json:"type"`
	HoursAssigned int   `json:"hours_assigned"`
}

// ===================== ORGANIGRAMME =====================
type OrganigramTemplate struct {
	ID         int       `json:"id"`
	Semester   string    `json:"semester"`
	ModuleID   int       `json:"module_id"`
	TeacherID  int       `json:"teacher_id"`
	Section    string    `json:"section"`
	Type       string    `json:"type"`
	Hours      int       `json:"hours"`
	VersionID  int       `json:"version_id"`
	UpdatedBy  int       `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrganigramVersion struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	IsPublished bool      `json:"is_published"`
}

// ===================== DEPARTMENTS & SPECIALTIES =====================
type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Specialty struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ===================== WISHLIST CONSTRAINTS =====================
type WishlistConstraint struct {
	TeacherID     int    `json:"teacher_id"`
	Semester      string `json:"semester"`
	MaxTotalHours int    `json:"max_total_hours"`
	Violated      bool   `json:"violated"`
}

// ===================== COMMENTS & EXPORTS =====================
type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Context   string    `json:"context"`
	ContextID int       `json:"context_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Export struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Type     string    `json:"type"`
	FilePath string    `json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
}

// ===================== ACTION LOGS =====================
type ActionLog struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Action    string    `json:"action"`
	TableName string    `json:"table_name"`
	RecordID  int       `json:"record_id"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}
