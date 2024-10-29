package models

type User struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	Email         string        `gorm:"uniqueIndex" json:"email"`
	Password      string        `json:"-"`
	Profile       UserProfile   `gorm:"foreignKey:UserID" json:"profile"`
	Workouts      []Workout     `gorm:"foreignKey:UserID" json:"workouts"`
	NutritionLogs []Nutrition    `gorm:"foreignKey:UserID" json:"nutritionLogs"`
	Progress      []Progress     `gorm:"foreignKey:UserID" json:"progress"`
}

type UserProfile struct {
	ID              uint     `gorm:"primaryKey" json:"id"`
	UserID          uint     `gorm:"not null" json:"userId"`
	Name            string   `json:"name"`
	FitnessGoals    []Goal   `gorm:"foreignKey:UserProfileID" json:"fitnessGoals"`
	ActivityHistory []Activity `gorm:"foreignKey:UserProfileID" json:"activityHistory"`
}

type Goal struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	UserProfileID   uint   `json:"userProfileId"`
	Goal            string `json:"goal"`
}

type Activity struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	UserProfileID   uint   `json:"userProfileId"`
	Activity        string `json:"activity"`
}

type Workout struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	UserID         uint   `gorm:"not null" json:"userId"`
	Type           string `json:"type"`
	Duration       int    `json:"duration"`
	CaloriesBurned int    `json:"caloriesBurned"`
	Date           string `json:"date"`
}

type Nutrition struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `gorm:"not null" json:"userId"`
	Meal     string `json:"meal"`
	Calories int    `json:"calories"`
	Date     string `json:"date"`
}

type Progress struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	UserID       uint    `gorm:"not null" json:"userId"`
	Weight       float64 `json:"weight"`
	Measurements string   `json:"measurements"`
	Date         string   `json:"date"`
}
