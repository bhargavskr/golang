package main

// Certifications details of the user
type Certifications struct {
	Name           string `json:"name"`
	YearIssued     int    `json:"year_issued"`
	Urlcertificate string `json:"url_certificate"`
}

// Education details of the user
type Education struct {
	GraduationYear int    `json:"graduation_year"`
	University     string `json:"university"`
	Field          string `json:"field"`
	Degree         string `json:"degree"`
	Grade          string `json:"grade"`
	Website        string `json:"website"`
}

// Employment details of the user
type Employment struct {
	JobTitle       string   `json:"job_title"`
	JobDescription []string `json:"job_description"`
	Institution    string   `json:"institution"`
	LocationCity   string   `json:"location_city"`
	LocationState  string   `json:"location_state"`
	From           string   `json:"from"`
	To             string   `json:"to"`
	Website        string   `json:"website"`
}

// SocialNetworks details of the user
type SocialNetworks struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

// PersonalInfo details of the user
type PersonalInfo struct {
	UserFirstname     string         `json:"user_firstname"`
	UserLastname      string         `json:"user_lastname"`
	Phone             string         `json:"phone"`
	Email             string         `json:"email"`
	AddressStreet     string         `json:"address_street"`
	AddressApt        string         `json:"address_apt"`
	AddressCity       string         `json:"address_city"`
	AddressState      string         `json:"address_state"`
	AddressCode       int            `json:"address_code"`
	ProfilePictureURL string         `json:"profilePictureURL"`
	LinkedinInfo      SocialNetworks `json:"linkedinInfo"`
	GithubInfo        SocialNetworks `json:"githubInfo"`
}

// Feedback provided to the user
type Feedback struct {
	Email string `json:"email"`
	Notes string `json:"notes"`
}

// Projects done by the user
type Projects struct {
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
}

// Profile of the user
type Profile struct {
	UserID                string                  `json:"userId"`
	EmailVerified         bool                    `json:"emailVerified"`
	Objective             string                  `json:"objective"`
	PersonalInfo          PersonalInfo            `json:"personal_info"`
	Education             []Education             `json:"education"`
	Certifications        []Certifications        `json:"certifications"`
	Employement           []Employment            `json:"employement"`
	Projects              []Projects              `json:"projects"`
	TechnicalCompetencies []TechnicalCompetencies `json:"technical_competencies"`
	Feedback              map[string]Feedback     `json:"feedback"`
}

// TechnicalCompetencies of the user
type TechnicalCompetencies struct {
	Area         string   `json:"area"`
	Technologies []string `json:"technologies"`
}
